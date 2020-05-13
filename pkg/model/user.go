package model

import (
	"github.com/google/uuid"

	"time"

	"github.com/lucitez/later/pkg/util/wrappers"
)

// User object
type User struct {
	ID          uuid.UUID           `json:"id"`
	Name        string              `json:"name"`
	Username    string              `json:"username"`
	Email       wrappers.NullString `json:"email"`
	PhoneNumber string              `json:"phone_number"`
	Password    string              `json:"password"`
	ExpoToken   wrappers.NullString `json:"expo_token"`

	CreatedAt  time.Time         `json:"created_at"`
	SignedUpAt wrappers.NullTime `json:"signed_up_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
	DeletedAt  wrappers.NullTime `json:"deleted_at"`
}

// NewUserFromSignUp constructor for creating a new user
// TODO validate email, phone number
func NewUserFromSignUp(
	username string,
	name string,
	email wrappers.NullString,
	phoneNumber string,
	password string,
) User {

	newUUID, _ := uuid.NewRandom()

	now := time.Now().UTC()

	return User{
		ID:          newUUID,
		Username:    username,
		Name:        name,
		Email:       email,
		PhoneNumber: phoneNumber,
		Password:    password,

		SignedUpAt: wrappers.NewNullTime(&now),
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}
