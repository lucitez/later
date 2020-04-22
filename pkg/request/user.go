package request

import (
	"later/pkg/service/body"
	"later/pkg/util/wrappers"

	"github.com/google/uuid"
)

// UserSignUpRequestBody Binding from json
type UserSignUpRequestBody struct {
	Username  string              `form:"username" json:"username" binding:"required"`
	FirstName string              `form:"first_name" json:"first_name" binding:"required"`
	LastName  wrappers.NullString `form:"last_name" json:"last_name"`
	Email     wrappers.NullString `form:"email" json:"email"`
}

func (b *UserSignUpRequestBody) ToUserSignUpBody(phoneNumber string, password string) body.UserSignUp {
	return body.UserSignUp{
		Username:    b.Username,
		FirstName:   b.FirstName,
		LastName:    b.LastName,
		Email:       b.Email,
		PhoneNumber: phoneNumber,
		Password:    password,
	}
}

// UserUpdate Binding from json
type UserUpdate struct {
	ID          uuid.UUID           `form:"id" json:"id" binding:"required"`
	FirstName   wrappers.NullString `form:"first_name" json:"first_name" binding:"required"`
	LastName    wrappers.NullString `form:"last_name" json:"last_name"`
	Email       wrappers.NullString `form:"email" json:"email"`
	PhoneNumber wrappers.NullString `form:"phone_number" json:"phone_number" binding:"required"`
}

func (requestBody *UserUpdate) ToUserUpdateBody() body.UserUpdate {
	return body.UserUpdate{
		requestBody.ID,
		requestBody.FirstName,
		requestBody.LastName,
		requestBody.Email,
		requestBody.PhoneNumber,
	}
}
