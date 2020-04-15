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
	"test_username",
	"first_name",
	wrappers.NewNullStringFromString("last_name"),
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

func TestAddFriendFilter(t *testing.T) {
	beforeEach(t)

	user1 := model.NewUserFromSignUp(
		"test_username",
		"first_name",
		wrappers.NewNullStringFromString("last_name"),
		wrappers.NewNullStringFromString("test_email"),
		"1111111111",
	)

	user2 := model.NewUserFromSignUp(
		"foo",
		"first_name",
		wrappers.NewNullStringFromString("last_name"),
		wrappers.NewNullStringFromString("foo"),
		"2222222222",
	)

	user3 := model.NewUserFromSignUp(
		"bar",
		"first_name",
		wrappers.NewNullStringFromString("last_name"),
		wrappers.NewNullStringFromString("bar"),
		"3333333333",
	)

	friend := model.NewFriend(
		user1.ID,
		user2.ID,
	)

	err := userRepo.Insert(user1)

	if err != nil {
		t.Error(err)
	}

	err = userRepo.Insert(user2)

	if err != nil {
		t.Error(err)
	}
	err = userRepo.Insert(user3)

	if err != nil {
		t.Error(err)
	}

	friendRepo.Insert(friend)

	actual := userRepo.AddFriendFilter(
		user1.ID,
		nil,
	)

	testUtil.Assert.Contains(actual, user3)
}

func TestAddFriendFilterWithSearch(t *testing.T) {
	beforeEach(t)

	user1 := model.NewUserFromSignUp(
		"test_username",
		"first_name",
		wrappers.NewNullStringFromString("last_name"),
		wrappers.NewNullStringFromString("test_email"),
		"1111111111",
	)

	user2 := model.NewUserFromSignUp(
		"foo",
		"first_name",
		wrappers.NewNullStringFromString("last_name"),
		wrappers.NewNullStringFromString("foo"),
		"2222222222",
	)

	user3 := model.NewUserFromSignUp(
		"bar",
		"first_name",
		wrappers.NewNullStringFromString("last_name"),
		wrappers.NewNullStringFromString("bar"),
		"3333333333",
	)

	user4 := model.NewUserFromSignUp(
		"baz",
		"first_name",
		wrappers.NewNullStringFromString("last_name"),
		wrappers.NewNullStringFromString("baz"),
		"4444444444",
	)

	friend := model.NewFriend(
		user1.ID,
		user2.ID,
	)

	userRepo.Insert(user1)
	userRepo.Insert(user2)
	userRepo.Insert(user3)
	userRepo.Insert(user4)

	friendRepo.Insert(friend)

	search := "bar"

	actual := userRepo.AddFriendFilter(
		user1.ID,
		&search,
	)

	testUtil.Assert.Contains(actual, user3)
}
