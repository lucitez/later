package repository

import (
	"database/sql"
	"log"

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
	statement := util.GenerateInsertStatement(user, "users")

	_, err := repository.DB.Exec(
		statement,
		user.ID,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		user.PhoneNumber,
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

	user.ScanRow(row)

	return &user
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

	user.ScanRow(row)

	return &user
}

// All returns all users with a limit
func (repository *User) All(limit int) []model.User {
	statement := userSelectStatement + `
	WHERE deleted_at IS NULL
	ORDER BY created_at desc
	LIMIT $1;
	`

	rows, err := repository.DB.Query(statement, limit)

	if err != nil {
		log.Fatal(err)
	}

	return repository.scanRows(rows)
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
