package repository_test

import (
	"later/pkg/model"
	"later/pkg/repository"
	"later/pkg/repository/util"
	"testing"
)

var friendRequestRepo repository.FriendRequest

var fr, err = model.NewFriendRequest(
	userID,
	userID2)

func TestInsertAndByID(t *testing.T) {
	beforeEach()
	friendRequestRepo.Insert(fr)

	actual, _ := friendRequestRepo.ByID(fr.ID)

	util.AssertEquals(t, actual, fr)
}

func TestPendingByUserID(t *testing.T) {
	beforeEach()
	friendRequestRepo.Insert(fr)

	actual, _ := friendRequestRepo.PendingByUserID(userID2)

	util.AssertContainsOne(t, actual, *fr)
}

func TestAccept(t *testing.T) {
	beforeEach()

	friendRequestRepo.Insert(fr)
	friendRequestRepo.Accept(fr.ID)
	pending, _ := friendRequestRepo.PendingByUserID(userID2)
	accepted, _ := friendRequestRepo.ByID(fr.ID)

	if !accepted.AcceptedAt.Valid {
		t.Error("Expected acceptedAt to not be null")
	}

	if len(pending) != 0 {
		t.Error("Expected no pending friend requests")
	}
}

func TestDecline(t *testing.T) {
	beforeEach()

	friendRequestRepo.Insert(fr)
	friendRequestRepo.Decline(fr.ID)
	pending, _ := friendRequestRepo.PendingByUserID(userID2)
	declined, _ := friendRequestRepo.ByID(fr.ID)

	if !declined.DeclinedAt.Valid {
		t.Error("Expected declinedAt to not be null")
	}

	if len(pending) != 0 {
		t.Error("Expected no pending friend requests")
	}
}
