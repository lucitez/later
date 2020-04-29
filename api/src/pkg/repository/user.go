package repository

import (
	"database/sql"
	"later/pkg/service/body"
	"log"
	"strconv"

	"github.com/google/uuid"
	"github.com/lib/pq"

	"later/pkg/model"
	"later/pkg/repository/util"
)

/*
User is the struct that implements the User interface and provides the database connection
*/
type User struct {
	DB *sql.DB
}

// NewUser ...
func NewUser(db *sql.DB) User {
	return User{db}
}

var userSelectStatement = util.GenerateSelectStatement(model.User{}, "users")

// Insert inserts a new user
func (repository *User) Insert(user model.User) error {

	statement := `
	INSERT INTO USERS (
		id,
		name,
		username,
		email,
		phone_number,
		password,
		created_at,
		signed_up_at,
		updated_at,
		deleted_at
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		crypt($6, gen_salt('bf')),
		$7,
		$8,
		$9,
		$10
	);
	`

	_, err := repository.DB.Exec(
		statement,
		user.ID,
		user.Name,
		user.Username,
		user.Email,
		user.PhoneNumber,
		user.Password,
		user.CreatedAt,
		user.SignedUpAt,
		user.UpdatedAt,
		user.DeletedAt,
	)

	return err
}

// ByID gets a user by id
func (repository *User) ByID(id uuid.UUID) *model.User {
	var user model.User

	statement := userSelectStatement + ` WHERE id = $1;`

	row := repository.DB.QueryRow(statement, id)

	return user.ScanRow(row)
}

// ByIdentifierAndPassword gets a user by id
func (repository *User) ByIdentifierAndPassword(identifier string, password string) *model.User {
	var user model.User

	statement := userSelectStatement + `
	WHERE (
		email = $1
		OR username = $1
		OR phone_number = $1
	) AND password = crypt($2, password)
	AND deleted_at IS NULL;
	`

	row := repository.DB.QueryRow(statement, identifier, password)

	return user.ScanRow(row)
}

// ByIDs ...
func (repository *User) ByIDs(ids []uuid.UUID) []model.User {
	statement := userSelectStatement + `
	WHERE id = ANY($1)
	AND deleted_at IS NULL;
	`

	rows, err := repository.DB.Query(statement, pq.Array(ids))

	if err != nil {
		log.Fatal(err)
	}

	return repository.scanRows(rows)
}

// ByPhoneNumber gets a user by their phone number
func (repository *User) ByPhoneNumber(phoneNumber string) *model.User {
	var user model.User

	statement := userSelectStatement + `
	WHERE phone_number = $1
	AND deleted_at IS NULL;
	`

	row := repository.DB.QueryRow(statement, phoneNumber)

	return user.ScanRow(row)
}

// AddFriendFilter ...
func (repository *User) AddFriendFilter(
	userID uuid.UUID,
	search *string,
) []model.User {
	withStatement := `
	WITH user_friends AS (
		SELECT friend_user_id
		FROM friends
		WHERE user_id = $1
		AND deleted_at IS NULL
	)
	`
	statement := withStatement + userSelectStatement + `
	WHERE id != $1
	AND id NOT IN (SELECT * FROM user_friends)
	`

	var fuzzySearch *string = nil

	if search != nil {
		statement = statement + `
		AND (
			username ILIKE $2
			OR email ILIKE $2
			OR name ILIKE $2
		
		)
		AND deleted_at IS NULL
		`
		fuzzySearch = search
		*fuzzySearch = "%" + *fuzzySearch + "%"
	} else {
		statement = statement + `
		AND deleted_at IS NULL
		`
	}

	args := util.GenerateArguments(
		userID,
		fuzzySearch,
	)

	rows, err := repository.DB.Query(statement, args...)

	if err != nil {
		log.Fatal(err)
	}

	return repository.scanRows(rows)

}

// Filter returns all users with a limit
func (repository *User) Filter(
	search *string,
	limit int,
	offset int,
) []model.User {
	statement := userSelectStatement
	counter := 1
	var fuzzySearch *string = nil

	if search != nil {
		statement = statement + `
		WHERE (
			username ILIKE $1
			OR email ILIKE $1
			OR name ILIKE $1
		
		)
		AND deleted_at IS NULL
		`
		counter++
		fuzzySearch = search
		*fuzzySearch = "%" + *fuzzySearch + "%"
	} else {
		statement = statement + `
		WHERE deleted_at IS NULL
		`
	}

	statement += `
	LIMIT $` + strconv.Itoa(counter)

	counter++

	statement += `
	OFFSET $` + strconv.Itoa(counter)

	args := util.GenerateArguments(
		fuzzySearch,
		limit,
		offset,
	)

	rows, err := repository.DB.Query(statement, args...)

	if err != nil {
		log.Fatal(err)
	}

	return repository.scanRows(rows)
}

func (repository *User) ByIdentifiers(
	phoneNumber string,
	username string,
) *model.User {
	var user model.User

	statement := util.GenerateSelectStatement(user, "users")

	statement += `
	WHERE (
		phone_number = $1
		OR username = $2
	) AND deleted_at IS NULL;
	`

	row := repository.DB.QueryRow(statement, phoneNumber, username)

	return user.ScanRow(row)
}

func (repository *User) Update(body body.UserUpdate) error {
	statement, arguments := util.GenerateUpdateStatement("users", body)

	_, err := repository.DB.Exec(statement, arguments...)

	return err
}

func (repository *User) scanRows(rows *sql.Rows) []model.User {
	users := []model.User{}

	defer rows.Close()

	for rows.Next() {
		var user model.User
		user.ScanRows(rows)

		users = append(users, user)
	}

	err := rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return users
}