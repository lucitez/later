package repository_test

import (
	"later/pkg/service/body"
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
	"name",
	wrappers.NewNullStringFromString("test_email"),
	"1111111111",
	"pass",
)

func TestInsertUserAndById(t *testing.T) {
	beforeEach(t)

	userRepo.Insert(user)

	actual := userRepo.ByID(user.ID)

	testUtil.Assert.Equal(actual.ID, user.ID)
}

func TestUsersByIDs(t *testing.T) {
	beforeEach(t)

	userRepo.Insert(user)

	actual := userRepo.ByIDs([]uuid.UUID{user.ID})

	testUtil.Assert.Len(actual, 1)
}

func TestUserByPhoneNumber(t *testing.T) {
	beforeEach(t)

	userRepo.Insert(user)

	actual := userRepo.ByPhoneNumber(user.PhoneNumber)

	testUtil.Assert.Equal(actual.ID, user.ID)
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

	testUtil.Assert.Len(actual, 1)
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

	testUtil.Assert.Len(actual, 1)
}

func TestAddFriendFilter(t *testing.T) {
	beforeEach(t)

	user1 := model.NewUserFromSignUp(
		"test_username",
		"name",
		wrappers.NewNullStringFromString("test_email"),
		"1111111111",
		"pass",
	)

	user2 := model.NewUserFromSignUp(
		"foo",
		"name",
		wrappers.NewNullStringFromString("foo"),
		"2222222222",
		"pass",
	)

	user3 := model.NewUserFromSignUp(
		"bar",
		"name",
		wrappers.NewNullStringFromString("bar"),
		"3333333333",
		"pass",
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

	testUtil.Assert.Equal(actual[0].ID, user3.ID)
}

func TestAddFriendFilterWithSearch(t *testing.T) {
	beforeEach(t)

	user1 := model.NewUserFromSignUp(
		"test_username",
		"name",
		wrappers.NewNullStringFromString("test_email"),
		"1111111111",
		"pass",
	)

	user2 := model.NewUserFromSignUp(
		"foo",
		"name",
		wrappers.NewNullStringFromString("foo"),
		"2222222222",
		"pass",
	)

	user3 := model.NewUserFromSignUp(
		"bar",
		"name",
		wrappers.NewNullStringFromString("bar"),
		"3333333333",
		"pass",
	)

	user4 := model.NewUserFromSignUp(
		"baz",
		"name",
		wrappers.NewNullStringFromString("baz"),
		"4444444444",
		"pass",
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

	testUtil.Assert.Equal(actual[0].ID, user3.ID)
}

func TestUpdateUser(t *testing.T) {
	beforeEach(t)

	userRepo.Insert(user)

	updateBody := body.UserUpdate{
		ID:   user.ID,
		Name: wrappers.NewNullStringFromString("glump"),
	}

	userRepo.Update(updateBody)

	actual := userRepo.ByID(user.ID)

	testUtil.Assert.Equal(actual.Name.String, "glump")
}
