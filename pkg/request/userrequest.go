package request

import "later.co/pkg/util/wrappers"

// UserSignUpRequestBody Binding from json
type UserSignUpRequestBody struct {
	Username    string              `form:"username" json:"username" binding:"required"`
	Email       wrappers.NullString `form:"email" json:"email"`
	PhoneNumber wrappers.NullString `form:"phone_number" json:"phone_number,string"`
}
