package request

type SMSConfirmationRequestBody struct {
	PhoneNumber string `form:"phone_number" json:"phone_number" binding:"required"`
}
