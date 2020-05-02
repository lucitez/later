package body

import (
	"github.com/lucitez/later/pkg/model"
)

type HostnameCreateBody struct {
	Hostname    string
	ContentType string
}

func (body *HostnameCreateBody) ToHostname() model.Hostname {
	return model.NewHostname(
		body.Hostname,
		body.ContentType,
	)
}
