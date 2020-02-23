package user

import (
	"github.com/google/uuid"

	"time"

	"later.co/pkg/util/wrappers"
)

// User object
type User struct {
	ID          uuid.UUID
	Username    string
	Email       wrappers.NullString
	PhoneNumber wrappers.NullString

	CreatedAt  time.Time
	SignedUpAt wrappers.NullTime
	UpdatedAt  time.Time
	DeletedAt  wrappers.NullTime
}

// New constructor for creating a new user
// TODO validate email, phone number
func New(
	username string,
	email wrappers.NullString,
	phoneNumber wrappers.NullString,
	signingUp bool) (*User, error) {

	newUUID, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	now := time.Now()

	var signedUpAt wrappers.NullTime

	if signingUp == true {
		signedUpAt = *wrappers.NewNullTime(now)
	}

	f := User{
		ID:          newUUID,
		Username:    username,
		Email:       email,
		PhoneNumber: phoneNumber,
		SignedUpAt:  signedUpAt,

		CreatedAt: now,
		UpdatedAt: now}

	return &f, nil
}
