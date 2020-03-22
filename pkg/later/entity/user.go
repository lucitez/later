package entity

import (
	"database/sql"

	"github.com/google/uuid"

	"time"

	"later.co/pkg/util/wrappers"
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

// NewUser constructor for creating a new user
// TODO validate email, phone number
func NewUser(
	username wrappers.NullString,
	email wrappers.NullString,
	phoneNumber string,
	signingUp bool) (*User, error) {

	newUUID, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	now := time.Now()

	var signedUpAt wrappers.NullTime

	if signingUp == true {
		signedUpAt = *wrappers.NewNullTime(&now)
	}

	user := User{
		ID:          newUUID,
		Username:    username,
		Email:       email,
		PhoneNumber: phoneNumber,
		SignedUpAt:  signedUpAt,

		CreatedAt: now,
		UpdatedAt: now}

	return &user, nil
}

// ScanRows ...
func (user *User) ScanRows(rows *sql.Rows) error {
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

	return err
}

// ScanRow ...
func (user *User) ScanRow(row *sql.Row) error {
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

	return err
}
