package body

import (
	"github.com/lucitez/later/pkg/model"
)

type DomainCreateBody struct {
	Domain      string
	ContentType string
}

func (body *DomainCreateBody) ToDomain() model.Domain {
	return model.NewDomain(
		body.Domain,
		body.ContentType,
	)
}