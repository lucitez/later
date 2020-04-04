package repository_test

import (
	"later/pkg/model"
	"later/pkg/repository"
	"testing"
)

var shareRepo repository.Share

var share = model.NewShare(
	contentID,
	userID,
	userID)

func TestShareInsertAndByID(t *testing.T) {
	beforeEach(t)
	shareRepo.Insert(share)

	actual := shareRepo.ByID(share.ID)

	testUtil.Assert.Equal(*actual, share)
}

func TestShareAll(t *testing.T) {
	beforeEach(t)
	shareRepo.Insert(share)

	actual := shareRepo.All(1)

	testUtil.Assert.Contains(actual, share)
}
