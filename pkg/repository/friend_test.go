package repository_test

import (
	"github.com/lucitez/later/pkg/model"
	"github.com/lucitez/later/pkg/repository"
	"github.com/lucitez/later/pkg/util/wrappers"
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
		"test",
		"test",
		wrappers.NewNullStringFromString("first"),
		"2222222222",
		"pass",
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
