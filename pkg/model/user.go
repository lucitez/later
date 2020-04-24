package model

import (
	"database/sql"
	"log"

	"github.com/google/uuid"

	"time"

	"later/pkg/util/wrappers"
)

// User object
type User struct {
	ID          uuid.UUID           `json:"id"`
	Name        wrappers.NullString `json:"name"`
	Username    wrappers.NullString `json:"username"`
	Email       wrappers.NullString `json:"email"`
	PhoneNumber string              `json:"phone_number"`
	Password    string              `json:"password"`

	CreatedAt  time.Time         `json:"created_at"`
	SignedUpAt wrappers.NullTime `json:"signed_up_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
	DeletedAt  wrappers.NullTime `json:"deleted_at"`
}

// NewUserFromSignUp constructor for creating a new user
// TODO validate email, phone number
func NewUserFromSignUp(
	username string,
	Name string,
	email wrappers.NullString,
	phoneNumber string,
	password string,
) User {

	newUUID, _ := uuid.NewRandom()

	now := time.Now().UTC()

	return User{
		ID:          newUUID,
		Username:    wrappers.NewNullStringFromString(username),
		Name:        wrappers.NewNullStringFromString(Name),
		Email:       email,
		PhoneNumber: phoneNumber,
		Password:    password,

		SignedUpAt: wrappers.NewNullTime(&now),
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}

// NewUserFromShare constructor for creating a new user
// TODO validate email, phone number
func NewUserFromShare(
	username wrappers.NullString,
	email wrappers.NullString,
	phoneNumber string,
) User {

	newUUID, _ := uuid.NewRandom()

	now := time.Now()

	return User{
		ID:          newUUID,
		Username:    username,
		Email:       email,
		PhoneNumber: phoneNumber,

		CreatedAt: now,
		UpdatedAt: now,
	}
}

// ScanRows ...
func (user *User) ScanRows(rows *sql.Rows) {
	err := rows.Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&user.Email,
		&user.PhoneNumber,
		&user.Password,
		&user.CreatedAt,
		&user.SignedUpAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)

	if err != nil {
		log.Fatal(err)
	}
}

// ScanRow ...
func (user *User) ScanRow(row *sql.Row) *User {
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&user.Email,
		&user.PhoneNumber,
		&user.Password,
		&user.CreatedAt,
		&user.SignedUpAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}

	return user
}
