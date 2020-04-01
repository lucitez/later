package repository_test

import (
	"later/pkg/model"
	"later/pkg/repository"
	"later/pkg/repository/util"
	"testing"
)

var shareRepo repository.Share

var share, _ = model.NewShare(
	contentID,
	userID,
	userID)

func TestShareInsertAndByID(t *testing.T) {
	beforeEach()
	shareRepo.Insert(share)

	actual, _ := shareRepo.ByID(share.ID)

	util.AssertEquals(t, actual, share)
}

func TestShareAll(t *testing.T) {
	beforeEach()
	shareRepo.Insert(share)

	actual, _ := shareRepo.All(1)

	util.AssertContainsOne(t, actual, *share)
}
