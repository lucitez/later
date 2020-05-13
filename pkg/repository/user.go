package repository

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/lucitez/later/pkg/service/body"

	"github.com/google/uuid"
	"github.com/lib/pq"

	"github.com/lucitez/later/pkg/model"
	"github.com/lucitez/later/pkg/repository/util"
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
	_, err := repository.DB.Exec(
		util.GenerateInsertStatement(user, "users"),
		util.GenerateInsertArguments(user),
	)

	return err
}

// UpdateExpoToken updates this user's expo token for push notifications
func (repository *User) UpdateExpoToken(token string, id uuid.UUID) error {
	statement := `
	UPDATE users SET expo_token = $1
	WHERE id = $2;
	`

	_, err := repository.DB.Exec(statement, token, id)
	return err
}

// ByID gets a user by id
func (repository *User) ByID(id uuid.UUID) (*model.User, error) {
	var user model.User

	statement := userSelectStatement + ` WHERE id = $1;`

	row := repository.DB.QueryRow(statement, id)

	err := util.ScanRowInto(row, &user)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &user, err
}

// ByIdentifierAndPassword gets a user by id
func (repository *User) ByIdentifierAndPassword(identifier string, password string) (*model.User, error) {
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

	err := util.ScanRowInto(row, &user)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &user, err
}

// ByIDs ...
func (repository *User) ByIDs(ids []uuid.UUID) ([]model.User, error) {
	statement := userSelectStatement + `
	WHERE id = ANY($1)
	AND deleted_at IS NULL;
	`

	rows, err := repository.DB.Query(statement, pq.Array(ids))

	if err != nil {
		return nil, err
	}

	return repository.scanRows(rows)
}

// ByPhoneNumber gets a user by their phone number
func (repository *User) ByPhoneNumber(phoneNumber string) (*model.User, error) {
	var user model.User

	statement := userSelectStatement + `
	WHERE phone_number = $1
	AND deleted_at IS NULL;
	`

	row := repository.DB.QueryRow(statement, phoneNumber)

	err := util.ScanRowInto(row, &user)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &user, err
}

// AddFriendFilter ...
func (repository *User) AddFriendFilter(
	userID uuid.UUID,
	search *string,
) ([]model.User, error) {
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
) ([]model.User, error) {
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
) (*model.User, error) {
	var user model.User

	statement := util.GenerateSelectStatement(user, "users")

	statement += `
	WHERE (
		phone_number = $1
		OR username = $2
	) AND deleted_at IS NULL;
	`

	row := repository.DB.QueryRow(statement, phoneNumber, username)

	err := util.ScanRowInto(row, &user)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &user, err
}

func (repository *User) Update(body body.UserUpdate) error {
	statement, arguments := util.GenerateUpdateStatement("users", body)

	_, err := repository.DB.Exec(statement, arguments...)

	return err
}

func (repository *User) scanRows(rows *sql.Rows) ([]model.User, error) {
	users := []model.User{}

	defer rows.Close()

	for rows.Next() {
		var user model.User
		err := util.ScanRowsInto(rows, &user)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
