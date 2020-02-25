package request

// ContentCreateRequestBody Binding from json
type ContentCreateRequestBody struct {
	URL string `form:"url" json:"url" binding:"required"`
}
