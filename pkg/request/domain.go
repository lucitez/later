package request

import (
	"later/pkg/service/body"
)

// DomainCreateRequestBody Binding from json
type DomainCreateRequestBody struct {
	Domain      string `form:"user_name" json:"user_name" binding:"required"`
	ContentType string `form:"email" json:"email" binding:"required"`
}

func (requestBody *DomainCreateRequestBody) ToDomainCreateBody() body.DomainCreateBody {
	return body.DomainCreateBody{
		Domain:      requestBody.Domain,
		ContentType: requestBody.ContentType,
	}
}
