package body

import (
	"later/pkg/model"
	"later/pkg/util/wrappers"

	"github.com/google/uuid"
)

// UserUpdate ...
type UserUpdate struct {
	ID          uuid.UUID
	FirstName   wrappers.NullString
	LastName    wrappers.NullString
	Email       wrappers.NullString
	PhoneNumber wrappers.NullString
}

type UserSignUp struct {
	Username    string
	FirstName   string
	LastName    wrappers.NullString
	Email       wrappers.NullString
	PhoneNumber string
	Password    string
}

func (b *UserSignUp) ToUser() model.User {
	return model.NewUserFromSignUp(
		b.Username,
		b.FirstName,
		b.LastName,
		b.Email,
		b.PhoneNumber,
		b.Password,
	)
}
