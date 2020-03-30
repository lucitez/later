package repository

import (
	"database/sql"

	"github.com/google/uuid"

	"github.com/lib/pq"
	"later.co/pkg/later/entity"
)

/*
UserRepository is the struct that implements the UserRepository interface and provides the database connection
*/
type UserRepository struct {
	DB *sql.DB
}

// NewUserRepository ...
func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{db}
}

// Insert inserts a new user
func (repository *UserRepository) Insert(user *entity.User) (*entity.User, error) {

	statement := `
	INSERT INTO users (
		id,
		first_name,
		last_name,
		username,
		email,
		phone_number,
		created_at,
		signed_up_at,
		updated_at,
		deleted_at
	)
	VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8,
		$9,
		$10
	)
	`

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
func (repository *UserRepository) ByID(id uuid.UUID) (*entity.User, error) {
	var user entity.User

	statement := entity.UserSelectStatement() + `
	FROM users 
	WHERE id = $1;
	`

	row := repository.DB.QueryRow(statement, id)

	err := user.ScanRow(row)

	return &user, err
}

// ByIDs ...
func (repository *UserRepository) ByIDs(ids []uuid.UUID) ([]entity.User, error) {
	statement := entity.UserSelectStatement() + `
	FROM users
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
func (repository *UserRepository) ByPhoneNumber(phoneNumber string) (*entity.User, error) {
	var user entity.User

	statement := entity.UserSelectStatement() + `
	FROM users
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
func (repository *UserRepository) All(limit int) ([]entity.User, error) {
	statement := entity.UserSelectStatement() + `
	FROM users
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

func (repository *UserRepository) scanRows(rows *sql.Rows) ([]entity.User, error) {
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
