package manager

import (
	"github.com/google/uuid"
	"later.co/pkg/later/entity"
	"later.co/pkg/repository"
)

type ContentManager interface {
	Create(content *entity.Content) (*entity.Content, error)
	ByID(id uuid.UUID) (*entity.Content, error)
	All(limit int) ([]entity.Content, error)
}

type ContentManagerImpl struct {
	Repository repository.ContentRepository
}

func (manager *ContentManagerImpl) Create(content *entity.Content) (*entity.Content, error) {
	return manager.Repository.Insert(content)
}

func (manager *ContentManagerImpl) ByID(id uuid.UUID) (*entity.Content, error) {
	return manager.Repository.ByID(id)
}

func (manager *ContentManagerImpl) All(limit int) ([]entity.Content, error) {
	return manager.Repository.All(limit)
}
