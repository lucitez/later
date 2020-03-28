package manager

import (
	"github.com/google/uuid"
	"later.co/pkg/later/entity"
	"later.co/pkg/repository"
	"later.co/pkg/response"
)

type UserContentManager interface {
	Create(userContent *entity.UserContent) (*entity.UserContent, error)
	ByID(id uuid.UUID) (*entity.UserContent, error)
	All(limit int) ([]entity.UserContent, error)
	Feed(
		userID uuid.UUID,
		senderType *string,
		contentType *string,
		archived *bool) ([]response.WireUserContent, error)
}

type UserContentManagerImpl struct {
	Repository repository.UserContentRepository
}

func (manager *UserContentManagerImpl) Create(userContent *entity.UserContent) (*entity.UserContent, error) {
	return manager.Repository.Insert(userContent)
}

func (manager *UserContentManagerImpl) ByID(id uuid.UUID) (*entity.UserContent, error) {
	return manager.Repository.ByID(id)
}

func (manager *UserContentManagerImpl) All(limit int) ([]entity.UserContent, error) {
	return manager.Repository.All(limit)
}

func (manager *UserContentManagerImpl) Feed(
	userID uuid.UUID,
	senderType *string,
	contentType *string,
	archived *bool) ([]response.WireUserContent, error) {

	return manager.Repository.Feed(
		userID,
		senderType,
		contentType,
		archived)
}
