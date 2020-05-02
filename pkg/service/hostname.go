package service

import (
	"github.com/lucitez/later/pkg/model"
	"github.com/lucitez/later/pkg/repository"
	"github.com/lucitez/later/pkg/service/body"
)

// Hostname ...
type Hostname struct {
	Repository repository.Hostname
}

// NewHostname for wire generation
func NewHostname(repository repository.Hostname) Hostname {
	return Hostname{repository}
}

// Create a new hostname
func (manager *Hostname) Create(body body.HostnameCreateBody) (*model.Hostname, error) {
	hostname := body.ToHostname()
	if err := manager.Repository.Insert(hostname); err != nil {
		return nil, err
	}

	return &hostname, nil
}

// ByHostname returns a hostname by its hostname name
func (manager *Hostname) ByHostname(hn string) *model.Hostname {
	return manager.Repository.ByHostname(hn)
}

// All returns all hostnames given a limit
func (manager *Hostname) All(limit int) []model.Hostname {
	return manager.Repository.All(limit)
}
