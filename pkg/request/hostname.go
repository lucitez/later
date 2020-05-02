package request

import (
	"github.com/lucitez/later/pkg/service/body"
)

// HostnameCreateRequestBody Binding from json
type HostnameCreateRequestBody struct {
	Hostname    string `form:"hostname" json:"hostname" binding:"required"`
	ContentType string `form:"content_type" json:"content_type" binding:"required"`
}

func (requestBody *HostnameCreateRequestBody) ToHostnameCreateBody() body.HostnameCreateBody {
	return body.HostnameCreateBody{
		Hostname:    requestBody.Hostname,
		ContentType: requestBody.ContentType,
	}
}
