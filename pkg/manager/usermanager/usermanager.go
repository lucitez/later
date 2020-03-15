package usermanager

import (
	"later.co/pkg/later/entity"
	"later.co/pkg/repository/userrepo"
	"later.co/pkg/util/wrappers"
)

// NewUserFromPhoneNumber inserts a new user using just phone number
func NewUserFromPhoneNumber(phoneNumber string) (*entity.User, error) {
	newUser, err := entity.NewUser(
		wrappers.NewNullString(nil), // user_name
		wrappers.NewNullString(nil), // email
		phoneNumber,
		false) // signing_up?

	if err != nil {
		return nil, err
	}

	user, err := userrepo.Insert(newUser)

	if err != nil {
		return nil, err
	}

	return user, nil
}
