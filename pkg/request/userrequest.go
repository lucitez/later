package request

import "later.co/pkg/util/wrappers"

// UserSignUpRequestBody Binding from json
type UserSignUpRequestBody struct {
	Username    wrappers.NullString `form:"username" json:"username"`
	Email       wrappers.NullString `form:"email" json:"email"`
	PhoneNumber string              `form:"phone_number" json:"phone_number" binding:"required"`
}
