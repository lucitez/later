package service

import (
	"later/pkg/model"
	"later/pkg/repository"
	"later/pkg/service/body"
	"later/pkg/util/wrappers"

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

// NewUserFromPhoneNumber inserts a new user using just phone number
func (manager *User) NewUserFromPhoneNumber(phoneNumber string) (*model.User, error) {
	user := model.NewUserFromShare(
		wrappers.NewNullString(nil), // username
		wrappers.NewNullString(nil), // email
		phoneNumber,
	)

	if err := manager.Repository.Insert(user); err != nil {
		return nil, err
	}

	return &user, nil
}

// SignUp ...
func (manager *User) SignUp(body body.UserSignUp) (*model.User, error) {
	user := body.ToUser()

	if err := manager.Repository.Insert(user); err != nil {
		return nil, err
	}

	return &user, nil
}

// ByID ...
func (manager *User) ByID(id uuid.UUID) *model.User {
	return manager.Repository.ByID(id)
}

// ByIdentifierAndPassword ...
func (manager *User) ByIdentifierAndPassword(identifier string, password string) *model.User {
	return manager.Repository.ByIdentifierAndPassword(identifier, password)
}

// ByPhoneNumber ...
func (manager *User) ByPhoneNumber(phoneNumber string) *model.User {
	return manager.Repository.ByPhoneNumber(phoneNumber)
}

// Filter ...
func (manager *User) Filter(
	search *string,
	limit int,
	offset int,
) []model.User {
	return manager.Repository.Filter(
		search,
		limit,
		offset,
	)
}

// AddFriendFilter ...
func (manager *User) AddFriendFilter(
	userID uuid.UUID,
	search *string,
) []model.User {
	return manager.Repository.AddFriendFilter(
		userID,
		search,
	)
}

// Update ...
func (manager *User) Update(body body.UserUpdate) error {
	return manager.Repository.Update(body)
}
