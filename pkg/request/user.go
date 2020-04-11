package request

import (
	"later/pkg/model"
	"later/pkg/util/wrappers"
)

// UserSignUpRequestBody Binding from json
type UserSignUpRequestBody struct {
	Username    string              `form:"username" json:"username" binding:"required"`
	FirstName   string              `form:"first_name" json:"first_name" binding:"required"`
	LastName    wrappers.NullString `form:"last_name" json:"last_name"`
	Email       wrappers.NullString `form:"email" json:"email"`
	PhoneNumber string              `form:"phone_number" json:"phone_number" binding:"required"`
}

func (body *UserSignUpRequestBody) ToUser() model.User {
	return model.NewUserFromSignUp(
		body.Username,
		body.FirstName,
		body.LastName,
		body.Email,
		body.PhoneNumber,
	)
}
