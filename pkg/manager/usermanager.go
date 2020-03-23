package manager

import (
	"github.com/google/uuid"
	"later.co/pkg/later/entity"
	"later.co/pkg/repository"
	"later.co/pkg/request"
	"later.co/pkg/util/wrappers"
)

type UserManager interface {
	SignUp(body request.UserSignUpRequestBody) (*entity.User, error)
	ByID(id uuid.UUID) (*entity.User, error)
	ByIDs(ids []uuid.UUID) ([]entity.User, error)
	ByPhoneNumber(phoneNumber string) (*entity.User, error)
	All(limit int) ([]entity.User, error)
	NewUserFromPhoneNumber(phoneNumber string) (*entity.User, error)
}

type UserManagerImpl struct {
	Repository repository.UserRepository
}

// NewUserFromPhoneNumber inserts a new user using just phone number
func (manager *UserManagerImpl) NewUserFromPhoneNumber(phoneNumber string) (*entity.User, error) {
	newUser, err := entity.NewUserFromShare(
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

func (manager *UserManagerImpl) SignUp(body request.UserSignUpRequestBody) (*entity.User, error) {
	user, err := entity.NewUserFromSignUp(
		body.Username,
		body.Email,
		body.PhoneNumber)

	if err != nil {
		return nil, err
	}

	return manager.Repository.Insert(user)
}

func (manager *UserManagerImpl) ByID(id uuid.UUID) (*entity.User, error) {
	return manager.Repository.ByID(id)
}

func (manager *UserManagerImpl) ByPhoneNumber(phoneNumber string) (*entity.User, error) {
	return manager.Repository.ByPhoneNumber(phoneNumber)
}

func (manager *UserManagerImpl) All(limit int) ([]entity.User, error) {
	return manager.Repository.All(limit)
}
