package userrepo

import (
	"database/sql"

	"github.com/google/uuid"

	// Postgres driver
	_ "github.com/lib/pq"
	"later.co/pkg/later/user"
)

// DB is this repository's database connection
var DB *sql.DB

// Insert inserts a new user
func Insert(user *user.User) (*user.User, error) {

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
func ByID(id uuid.UUID) (*user.User, error) {
	var user user.User

	statement := `
	SELECT * FROM users 
	WHERE id = $1
	`

	row := DB.QueryRow(statement, id)

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PhoneNumber,
		&user.CreatedAt,
		&user.SignedUpAt,
		&user.UpdatedAt,
		&user.DeletedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// All returns all users with a limit
func All(limit int) ([]user.User, error) {
	users := []user.User{}

	rows, err := DB.Query(`SELECT * FROM users LIMIT $1`, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user user.User
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.PhoneNumber,
			&user.CreatedAt,
			&user.SignedUpAt,
			&user.UpdatedAt,
			&user.DeletedAt)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}
