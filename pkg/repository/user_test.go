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

var userRepo repository.User

var user, _ = model.NewUserFromSignUp(
	wrappers.NewNullStringFromString("test_username"),
	wrappers.NewNullStringFromString("test_email"),
	"1111111111")

func TestInsertUserAndById(t *testing.T) {
	beforeEach()

	userRepo.Insert(user)

	actual, _ := userRepo.ByID(user.ID)

	util.AssertEquals(t, actual, user)
}

func TestUsersByIDs(t *testing.T) {
	beforeEach()

	userRepo.Insert(user)

	actual, err := userRepo.ByIDs([]uuid.UUID{user.ID})

	if err != nil {
		t.Error(err)
	}

	util.AssertContainsOne(t, actual, *user)
}

func TestUserByPhoneNumber(t *testing.T) {
	beforeEach()

	userRepo.Insert(user)

	actual, err := userRepo.ByPhoneNumber(user.PhoneNumber)

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

	userRepo.Insert(user)
	userRepo.Insert(user2)

	actual, err := userRepo.All(1)

	if err != nil {
		t.Error(err)
	}

	util.AssertEquals(t, actual, []model.User{*user2})
}
