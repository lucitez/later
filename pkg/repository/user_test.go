package repository_test

import (
	"testing"

	// Postgres driver
	"github.com/google/uuid"
	_ "github.com/lib/pq"

	"later/pkg/model"
	"later/pkg/repository"
	"later/pkg/util/wrappers"
)

var userRepo repository.User

var user = model.NewUserFromSignUp(
	wrappers.NewNullStringFromString("test_username"),
	wrappers.NewNullStringFromString("test_email"),
	"1111111111",
)

func TestInsertUserAndById(t *testing.T) {
	beforeEach(t)

	userRepo.Insert(user)

	actual := userRepo.ByID(user.ID)

	testUtil.Assert.Equal(*actual, user)
}

func TestUsersByIDs(t *testing.T) {
	beforeEach(t)

	userRepo.Insert(user)

	actual := userRepo.ByIDs([]uuid.UUID{user.ID})

	testUtil.Assert.Contains(actual, user)
}

func TestUserByPhoneNumber(t *testing.T) {
	beforeEach(t)

	userRepo.Insert(user)

	actual := userRepo.ByPhoneNumber(user.PhoneNumber)

	testUtil.Assert.Equal(*actual, user)
}

func TestAllUsers(t *testing.T) {
	beforeEach(t)

	user2 := model.NewUserFromSignUp(
		wrappers.NewNullStringFromString("test_username_2"),
		wrappers.NewNullStringFromString("test_email_2"),
		"0000000000",
	)

	userRepo.Insert(user)
	userRepo.Insert(user2)

	actual := userRepo.All(1)

	testUtil.Assert.Contains(actual, user2)
}
