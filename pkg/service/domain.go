package service

import (
	"github.com/lucitez/later/pkg/model"
	"github.com/lucitez/later/pkg/repository"
	"github.com/lucitez/later/pkg/service/body"
)

// Domain ...
type Domain struct {
	Repository repository.Domain
}

// NewDomain for wire generation
func NewDomain(repository repository.Domain) Domain {
	return Domain{repository}
}

// Create a new domain
func (manager *Domain) Create(body body.DomainCreateBody) (*model.Domain, error) {
	domain := body.ToDomain()
	if err := manager.Repository.Insert(domain); err != nil {
		return nil, err
	}

	return &domain, nil
}

// ByDomain returns a domain by its domain name
func (manager *Domain) ByDomain(domainName string) *model.Domain {
	return manager.Repository.ByDomain(domainName)
}

// All returns all domains given a limit
func (manager *Domain) All(limit int) []model.Domain {
	return manager.Repository.All(limit)
}
