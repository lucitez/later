package request

import (
	"later/pkg/model"
	"later/pkg/util/wrappers"
)

// UserSignUpRequestBody Binding from json
type UserSignUpRequestBody struct {
	Username    wrappers.NullString `form:"username" json:"username"`
	Email       wrappers.NullString `form:"email" json:"email"`
	PhoneNumber string              `form:"phone_number" json:"phone_number" binding:"required"`
}

func (body *UserSignUpRequestBody) ToUser() model.User {
	return model.NewUserFromSignUp(
		body.Username,
		body.Email,
		body.PhoneNumber,
	)
}
