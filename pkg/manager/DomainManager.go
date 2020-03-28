package manager

import (
	"later.co/pkg/later/entity"
	"later.co/pkg/repository"
)

func NewDomainManager(repo repository.DomainRepositoryImpl) DomainManagerImpl {
	return DomainManagerImpl{repo}
}

// DomainManager ...
type DomainManager interface {
	Create(domain *entity.Domain) (*entity.Domain, error)
	ByDomain(domainName string) (*entity.Domain, error)
	All(limit int) ([]entity.Domain, error)
}

// DomainManagerImpl ...
type DomainManagerImpl struct {
	Repository repository.DomainRepositoryImpl
}

// Create a new domain
func (manager *DomainManagerImpl) Create(domain *entity.Domain) (*entity.Domain, error) {
	return manager.Repository.Insert(domain)
}

// ByDomain returns a domain by its domain name
func (manager *DomainManagerImpl) ByDomain(domainName string) (*entity.Domain, error) {
	return manager.Repository.ByDomain(domainName)
}

// All returns all domains given a limit
func (manager *DomainManagerImpl) All(limit int) ([]entity.Domain, error) {
	return manager.Repository.All(limit)
}
