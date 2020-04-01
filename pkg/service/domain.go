package service

import (
	"later/pkg/repository"
	"later/pkg/model"
)

// DomainManager ...
type DomainManager struct {
	Repository repository.Domain
}

// NewDomainManager for wire generation
func NewDomainManager(repository repository.Domain) DomainManager {
	return DomainManager{repository}
}

// Create a new domain
func (manager *DomainManager) Create(domain *model.Domain) (*model.Domain, error) {
	return manager.Repository.Insert(domain)
}

// ByDomain returns a domain by its domain name
func (manager *DomainManager) ByDomain(domainName string) (*model.Domain, error) {
	return manager.Repository.ByDomain(domainName)
}

// All returns all domains given a limit
func (manager *DomainManager) All(limit int) ([]model.Domain, error) {
	return manager.Repository.All(limit)
}
