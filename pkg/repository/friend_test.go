package repository_test

import (
	"later/pkg/model"
	"later/pkg/repository"
	"later/pkg/util/wrappers"
	"testing"
)

var friendRepo repository.Friend

var friend = model.NewFriend(
	userID,
	userID2,
)

func TestInsertAndForUser(t *testing.T) {
	beforeEach(t)

	friendRepo.Insert(friend)

	actual := friendRepo.ForUser(
		friend.UserID,
		nil,
		1,
		0,
	)

	testUtil.Assert.Contains(actual, friend)
}

func TestForUserWithSearch(t *testing.T) {
	beforeEach(t)
	user := model.NewUserFromSignUp(
		wrappers.NewNullStringFromString("test"),
		wrappers.NewNullStringFromString("test"),
		"2222222222",
	)

	friend := model.NewFriend(
		userID,
		user.ID,
	)

	friendRepo.Insert(friend)
	userRepo.Insert(user)

	search := "test"

	actual := friendRepo.ForUser(
		friend.UserID,
		&search,
		1,
		0,
	)

	testUtil.Assert.Contains(actual, friend)
}
