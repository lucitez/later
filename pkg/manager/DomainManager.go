package manager

import (
	"later.co/pkg/later/entity"
	"later.co/pkg/repository"
)

// DomainManager ...
type DomainManager struct {
	Repository repository.DomainRepository
}

// NewDomainManager for wire generation
func NewDomainManager(repository repository.DomainRepository) DomainManager {
	return DomainManager{repository}
}

// Create a new domain
func (manager *DomainManager) Create(domain *entity.Domain) (*entity.Domain, error) {
	return manager.Repository.Insert(domain)
}

// ByDomain returns a domain by its domain name
func (manager *DomainManager) ByDomain(domainName string) (*entity.Domain, error) {
	return manager.Repository.ByDomain(domainName)
}

// All returns all domains given a limit
func (manager *DomainManager) All(limit int) ([]entity.Domain, error) {
	return manager.Repository.All(limit)
}
