package request

import "later.co/pkg/util/wrappers"

// UserSignUpRequestBody Binding from json
type UserSignUpRequestBody struct {
	Username    string              `form:"user_name" json:"user_name" binding:"required"`
	Email       wrappers.NullString `form:"email" json:"email"`
	PhoneNumber wrappers.NullString `form:"phone_number" json:"phone_number,string"`
}
