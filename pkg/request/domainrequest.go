package request

// DomainCreateRequestBody Binding from json
type DomainCreateRequestBody struct {
	Domain      string `form:"user_name" json:"user_name" binding:"required"`
	ContentType string `form:"email" json:"email" binding:"required"`
}
