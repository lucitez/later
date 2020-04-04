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
func (manager *User) NewUserFromPhoneNumber(phoneNumber string) model.User {
	newUser := model.NewUserFromShare(
		wrappers.NewNullString(nil), // user_name
		wrappers.NewNullString(nil), // email
		phoneNumber,
	)

	user := manager.Repository.Insert(newUser)

	return user
}

// SignUp ...
func (manager *User) SignUp(body request.UserSignUpRequestBody) model.User {
	user := model.NewUserFromSignUp(
		body.Username,
		body.Email,
		body.PhoneNumber,
	)

	return manager.Repository.Insert(user)
}

// ByID ...
func (manager *User) ByID(id uuid.UUID) *model.User {
	return manager.Repository.ByID(id)
}

// ByPhoneNumber ...
func (manager *User) ByPhoneNumber(phoneNumber string) (*model.User, error) {
	return manager.Repository.ByPhoneNumber(phoneNumber)
}

// All ...
func (manager *User) All(limit int) []model.User {
	return manager.Repository.All(limit)
}
