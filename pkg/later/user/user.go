package user

import (
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

// New constructor for creating a new user
// TODO validate email, phone number
func New(
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
