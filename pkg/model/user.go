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
	FirstName   wrappers.NullString `json:"first_name"`
	LastName    wrappers.NullString `json:"last_name"`
	Username    wrappers.NullString `json:"username"`
	Email       wrappers.NullString `json:"email"`
	PhoneNumber string              `json:"phone_number"`

	CreatedAt  time.Time         `json:"created_at"`
	SignedUpAt wrappers.NullTime `json:"signed_up_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
	DeletedAt  wrappers.NullTime `json:"deleted_at"`
}

// NewUserFromSignUp constructor for creating a new user
// TODO validate email, phone number
func NewUserFromSignUp(
	username wrappers.NullString,
	email wrappers.NullString,
	phoneNumber string,
) User {

	newUUID, _ := uuid.NewRandom()

	now := time.Now().UTC()

	return User{
		ID:          newUUID,
		Username:    username,
		Email:       email,
		PhoneNumber: phoneNumber,
		SignedUpAt:  wrappers.NewNullTime(&now),

		CreatedAt: now,
		UpdatedAt: now,
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
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.PhoneNumber,
		&user.CreatedAt,
		&user.SignedUpAt,
		&user.UpdatedAt,
		&user.DeletedAt)

	if err != nil {
		log.Fatal(err)
	}
}

// ScanRow ...
func (user *User) ScanRow(row *sql.Row) {
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.PhoneNumber,
		&user.CreatedAt,
		&user.SignedUpAt,
		&user.UpdatedAt,
		&user.DeletedAt)

	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
}
