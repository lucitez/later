package service

import (
	"github.com/google/uuid"
	"later/pkg/model"
	"later/pkg/repository"
	"later/pkg/request"
	"later/pkg/util/wrappers"
)

// UserManager ...
type UserManager struct {
	Repository repository.UserRepository
}

// NewUserManager ...
func NewUserManager(repository repository.UserRepository) UserManager {
	return UserManager{repository}
}

// NewUserFromPhoneNumber inserts a new user using just phone number
func (manager *UserManager) NewUserFromPhoneNumber(phoneNumber string) (*model.User, error) {
	newUser, err := model.NewUserFromShare(
		wrappers.NewNullString(nil), // user_name
		wrappers.NewNullString(nil), // email
		phoneNumber)

	if err != nil {
		return nil, err
	}

	user, err := manager.Repository.Insert(newUser)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// SignUp ...
func (manager *UserManager) SignUp(body request.UserSignUpRequestBody) (*model.User, error) {
	user, err := model.NewUserFromSignUp(
		body.Username,
		body.Email,
		body.PhoneNumber)

	if err != nil {
		return nil, err
	}

	return manager.Repository.Insert(user)
}

// ByID ...
func (manager *UserManager) ByID(id uuid.UUID) (*model.User, error) {
	return manager.Repository.ByID(id)
}

// ByPhoneNumber ...
func (manager *UserManager) ByPhoneNumber(phoneNumber string) (*model.User, error) {
	return manager.Repository.ByPhoneNumber(phoneNumber)
}

// All ...
func (manager *UserManager) All(limit int) ([]model.User, error) {
	return manager.Repository.All(limit)
}
