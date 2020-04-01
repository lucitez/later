package repository

import (
	"database/sql"

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
func (repository *User) Insert(user *model.User) (*model.User, error) {

	statement := util.GenerateInsertStatement(*user, "users")

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
		user.DeletedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// ByID gets a user by id
func (repository *User) ByID(id uuid.UUID) (*model.User, error) {
	var user model.User

	statement := userSelectStatement + ` WHERE id = $1;`

	row := repository.DB.QueryRow(statement, id)

	err := user.ScanRow(row)

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

	users, err := repository.scanRows(rows)

	if err != nil {
		return nil, err
	}

	return users, nil
}

// ByPhoneNumber gets a user by their phone number
func (repository *User) ByPhoneNumber(phoneNumber string) (*model.User, error) {
	var user model.User

	statement := userSelectStatement + `
	WHERE phone_number = $1;
	`

	row := repository.DB.QueryRow(statement, phoneNumber)

	err := user.ScanRow(row)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// All returns all users with a limit
func (repository *User) All(limit int) ([]model.User, error) {
	statement := userSelectStatement + `
	WHERE deleted_at IS NULL
	ORDER BY created_at desc
	LIMIT $1;
	`

	rows, err := repository.DB.Query(statement, limit)

	if err != nil {
		return nil, err
	}

	users, err := repository.scanRows(rows)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repository *User) scanRows(rows *sql.Rows) ([]model.User, error) {
	users := []model.User{}

	defer rows.Close()

	for rows.Next() {
		var user model.User
		err := user.ScanRows(rows)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	err := rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}
