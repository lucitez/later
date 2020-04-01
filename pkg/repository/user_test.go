package repository_test

import (
	"testing"

	// Postgres driver
	"github.com/google/uuid"
	_ "github.com/lib/pq"

	"later/pkg/model"
	"later/pkg/repository"
	"later/pkg/repository/util"
	"later/pkg/util/wrappers"
)

var userRepository repository.User

var user, err = model.NewUserFromSignUp(
	wrappers.NewNullStringFromString("test_username"),
	wrappers.NewNullStringFromString("test_email"),
	"1111111111")

func TestInsertUserAndById(t *testing.T) {
	beforeEach()

	userRepository.Insert(user)

	actual, _ := userRepository.ByID(user.ID)

	util.AssertEquals(t, actual, user)
}

func TestUsersByIDs(t *testing.T) {
	beforeEach()

	user2, _ := model.NewUserFromSignUp(
		wrappers.NewNullStringFromString("test_username_2"),
		wrappers.NewNullStringFromString("test_email_2"),
		"0000000000")

	userRepository.Insert(user)
	userRepository.Insert(user2)

	actual, err := userRepository.ByIDs([]uuid.UUID{user.ID, user2.ID})

	if err != nil {
		t.Error(err)
	}

	util.AssertEquals(t, actual, []model.User{*user, *user2})
}

func TestUserByPhoneNumber(t *testing.T) {
	beforeEach()

	userRepository.Insert(user)

	actual, err := userRepository.ByPhoneNumber(user.PhoneNumber)

	if err != nil {
		t.Error(err)
	}

	util.AssertEquals(t, actual, user)
}

func TestAllUsers(t *testing.T) {
	beforeEach()

	user2, _ := model.NewUserFromSignUp(
		wrappers.NewNullStringFromString("test_username_2"),
		wrappers.NewNullStringFromString("test_email_2"),
		"0000000000")

	userRepository.Insert(user)
	userRepository.Insert(user2)

	actual, err := userRepository.All(1)

	if err != nil {
		t.Error(err)
	}

	util.AssertEquals(t, actual, []model.User{*user2})
}
