package repository_test

import (
	"github.com/lucitez/later/api/src/pkg/model"
	"github.com/lucitez/later/api/src/pkg/repository"
	"testing"
)

var friendRequestRepo repository.FriendRequest

var fr = model.NewFriendRequest(
	userID,
	userID2,
)

func TestInsertAndByID(t *testing.T) {
	beforeEach(t)
	friendRequestRepo.Insert(fr)

	actual := friendRequestRepo.ByID(fr.ID)

	testUtil.Assert.Equal(*actual, fr)
}

func TestPendingByUserID(t *testing.T) {
	beforeEach(t)
	friendRequestRepo.Insert(fr)

	actual := friendRequestRepo.PendingByUserID(userID2)

	testUtil.Assert.Contains(actual, fr)
}

func TestAccept(t *testing.T) {
	beforeEach(t)

	friendRequestRepo.Insert(fr)
	friendRequestRepo.Accept(fr.ID)
	pending := friendRequestRepo.PendingByUserID(userID2)
	accepted := friendRequestRepo.ByID(fr.ID)

	if !accepted.AcceptedAt.Valid {
		t.Error("Expected acceptedAt to not be null")
	}

	testUtil.Assert.Empty(pending)
}

func TestDecline(t *testing.T) {
	beforeEach(t)

	friendRequestRepo.Insert(fr)
	friendRequestRepo.Decline(fr.ID)
	pending := friendRequestRepo.PendingByUserID(userID2)
	declined := friendRequestRepo.ByID(fr.ID)

	if !declined.DeclinedAt.Valid {
		t.Error("Expected declinedAt to not be null")
	}

	testUtil.Assert.Empty(pending)
}
