package service

import (
	"github.com/lucitez/later/pkg/model"
	"github.com/lucitez/later/pkg/repository"
	"github.com/lucitez/later/pkg/service/body"

	"github.com/google/uuid"
)

// User ...
type User struct {
	Repository repository.User
}

// NewUser ...
func NewUser(repository repository.User) User {
	return User{repository}
}

// SignUp ...
func (service *User) SignUp(body body.UserSignUp) (*model.User, error) {
	user := body.ToUser()

	if err := service.Repository.Insert(user); err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateExpoToken ...
func (service *User) UpdateExpoToken(token string, id uuid.UUID) error {
	return service.Repository.UpdateExpoToken(token, id)
}

// ByID ...
func (service *User) ByID(id uuid.UUID) (*model.User, error) {
	return service.Repository.ByID(id)
}

// ByIdentifierAndPassword ...
func (service *User) ByIdentifierAndPassword(identifier string, password string) (*model.User, error) {
	return service.Repository.ByIdentifierAndPassword(identifier, password)
}

// ByPhoneNumber ...
func (service *User) ByPhoneNumber(phoneNumber string) (*model.User, error) {
	return service.Repository.ByPhoneNumber(phoneNumber)
}

// Filter ...
func (service *User) Filter(
	search *string,
	limit int,
	offset int,
) ([]model.User, error) {
	return service.Repository.Filter(
		search,
		limit,
		offset,
	)
}

// AddFriendFilter ...
func (service *User) AddFriendFilter(
	userID uuid.UUID,
	search *string,
) ([]model.User, error) {
	return service.Repository.AddFriendFilter(
		userID,
		search,
	)
}

// Update ...
func (service *User) Update(body body.UserUpdate) error {
	return service.Repository.Update(body)
}
