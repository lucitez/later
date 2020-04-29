package request

import (
	"github.com/lucitez/later/pkg/service/body"
	"github.com/lucitez/later/pkg/util/wrappers"

	"github.com/google/uuid"
)

// UserSignUpRequestBody Binding from json
type UserSignUpRequestBody struct {
	Username string              `form:"username" json:"username" binding:"required"`
	Name     string              `form:"name" json:"name" binding:"required"`
	Email    wrappers.NullString `form:"email" json:"email"`
}

func (b *UserSignUpRequestBody) ToUserSignUpBody(phoneNumber string, password string) body.UserSignUp {
	return body.UserSignUp{
		Username:    b.Username,
		Name:        b.Name,
		Email:       b.Email,
		PhoneNumber: phoneNumber,
		Password:    password,
	}
}

// UserUpdate Binding from json
type UserUpdate struct {
	Name        wrappers.NullString `form:"name" json:"name"`
	Email       wrappers.NullString `form:"email" json:"email"`
	PhoneNumber wrappers.NullString `form:"phone_number" json:"phone_number"`
}

func (requestBody *UserUpdate) ToUserUpdateBody(userID uuid.UUID) body.UserUpdate {
	return body.UserUpdate{
		ID:          userID,
		Name:        requestBody.Name,
		Email:       requestBody.Email,
		PhoneNumber: requestBody.PhoneNumber,
	}
}
