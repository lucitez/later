package repository_test

import (
	"later/pkg/model"
	"later/pkg/repository"
	"later/pkg/repository/util"
	"later/pkg/util/wrappers"
	"testing"
)

var friendRepo repository.Friend

var friend, _ = model.NewFriend(
	userID,
	userID2)

func TestInsertAndByUserID(t *testing.T) {
	beforeEach()
	_, err := friendRepo.Insert(friend)

	if err != nil {
		t.Error(err)
	}

	actual, err := friendRepo.ByUserID(friend.UserID)

	if err != nil {
		t.Error(err)
	}

	util.AssertContainsOne(t, actual, *friend)
}

func TestSearchByUserID(t *testing.T) {
	beforeEach()
	user, _ := model.NewUserFromSignUp(
		wrappers.NewNullStringFromString("test"),
		wrappers.NewNullStringFromString("test"),
		"2222222222")

	friend, _ := model.NewFriend(
		userID,
		user.ID)

	friendRepo.Insert(friend)
	userRepo.Insert(user)

	actual, _ := friendRepo.SearchByUserID(friend.UserID, "TES")

	util.AssertContainsOne(t, actual, *friend)
}
