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

func TestInsertAndByUserID(t *testing.T) {
	beforeEach(t)

	friendRepo.Insert(friend)

	actual := friendRepo.ByUserID(friend.UserID)

	testUtil.Assert.Contains(actual, friend)
}

func TestSearchByUserID(t *testing.T) {
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

	actual := friendRepo.SearchByUserID(friend.UserID, "TES")

	testUtil.Assert.Contains(actual, friend)
}
