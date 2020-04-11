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

func TestFilterUsers(t *testing.T) {
	beforeEach(t)

	search := "test"
	userRepo.Insert(user)

	actual := userRepo.Filter(
		&search,
		1,
		0,
	)

	testUtil.Assert.Contains(actual, user)
}

func TestFilterUsersEmptySearch(t *testing.T) {
	beforeEach(t)

	search := ""
	userRepo.Insert(user)

	actual := userRepo.Filter(
		&search,
		1,
		0,
	)

	testUtil.Assert.Contains(actual, user)
}
