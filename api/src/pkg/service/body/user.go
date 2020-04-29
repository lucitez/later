package body

import (
	"later/pkg/model"
	"later/pkg/util/wrappers"

	"github.com/google/uuid"
)

// UserUpdate ...
type UserUpdate struct {
	ID          uuid.UUID
	Name        wrappers.NullString
	Email       wrappers.NullString
	PhoneNumber wrappers.NullString
}

type UserSignUp struct {
	Username    string
	Name        string
	Email       wrappers.NullString
	PhoneNumber string
	Password    string
}

func (b *UserSignUp) ToUser() model.User {
	return model.NewUserFromSignUp(
		b.Username,
		b.Name,
		b.Email,
		b.PhoneNumber,
		b.Password,
	)
}
