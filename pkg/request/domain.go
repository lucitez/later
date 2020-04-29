package request

import (
	"github.com/lucitez/later/pkg/service/body"
)

// DomainCreateRequestBody Binding from json
type DomainCreateRequestBody struct {
	Domain      string `form:"domain" json:"domain" binding:"required"`
	ContentType string `form:"content_type" json:"content_type" binding:"required"`
}

func (requestBody *DomainCreateRequestBody) ToDomainCreateBody() body.DomainCreateBody {
	return body.DomainCreateBody{
		Domain:      requestBody.Domain,
		ContentType: requestBody.ContentType,
	}
}
