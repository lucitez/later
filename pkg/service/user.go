package service

import (
	"later/pkg/model"
	"later/pkg/repository"
	"later/pkg/request"
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
		wrappers.NewNullString(nil), // user_name
		wrappers.NewNullString(nil), // email
		phoneNumber,
	)

	if err := manager.Repository.Insert(user); err != nil {
		return nil, err
	}

	return &user, nil
}

// SignUp ...
func (manager *User) SignUp(body request.UserSignUpRequestBody) (*model.User, error) {
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

// ByPhoneNumber ...
func (manager *User) ByPhoneNumber(phoneNumber string) *model.User {
	return manager.Repository.ByPhoneNumber(phoneNumber)
}

// All ...
func (manager *User) All(limit int) []model.User {
	return manager.Repository.All(limit)
}
