package user

import (
	"database/sql"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"later.co/pkg/later"
)

// DB is this repository's database connection
var DB *sql.DB

// ByID gets a user by id
func ByID(id uuid.UUID) (*later.EntityUser, error) {
	var user later.EntityUser

	row := DB.QueryRow(`SELECT id, username, email, created_at, updated_at FROM users WHERE ID = $1`, id)

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
