package userrepo

import (
	"database/sql"

	"github.com/google/uuid"

	// Postgres driver
	_ "github.com/lib/pq"
	"later.co/pkg/later/entity"
)

// DB is this repository's database connection
var DB *sql.DB

// Insert inserts a new user
func Insert(user *entity.User) (*entity.User, error) {

	statement := `
	INSERT INTO users (username, email, phone_number)
	VALUES (
		$1,
		$2,
		$3
	)
	`

	_, err := DB.Exec(
		statement,
		user.Username,
		user.Email,
		user.PhoneNumber)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// ByID gets a user by id
func ByID(id uuid.UUID) (*entity.User, error) {
	var user entity.User

	statement := `
	SELECT * FROM users 
	WHERE id = $1;
	`

	row := DB.QueryRow(statement, id)

	err := user.ScanRow(row)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// ByIDs ...
func ByIDs(ids []uuid.UUID) ([]entity.User, error) {
	statement := `
	SELECT * FROM users
	WHERE id in $1
	AND deleted_at IS NULL;
	`

	rows, err := DB.Query(statement, ids)

	if err != nil {
		return nil, err
	}

	users, err := scanRows(rows)

	if err != nil {
		return nil, err
	}

	return users, nil
}

// ByPhoneNumber gets a user by their phone number
func ByPhoneNumber(phoneNumber string) (*entity.User, error) {
	var user entity.User

	statement := `
	SELECT * FROM users 
	WHERE phone_number = $1;
	`

	row := DB.QueryRow(statement, phoneNumber)

	err := user.ScanRow(row)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// All returns all users with a limit
func All(limit int) ([]entity.User, error) {
	statement := `
	SELECT * FROM users
	WHERE deleted_at IS NULL
	LIMIT $1;
	`

	rows, err := DB.Query(statement, limit)

	if err != nil {
		return nil, err
	}

	users, err := scanRows(rows)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func scanRows(rows *sql.Rows) ([]entity.User, error) {
	users := []entity.User{}

	defer rows.Close()

	for rows.Next() {
		var user entity.User
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
