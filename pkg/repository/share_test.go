package repository_test

import (
	"github.com/lucitez/later/pkg/model"
	"github.com/lucitez/later/pkg/repository"
	"testing"
)

var shareRepo repository.Share

var share = model.NewShare(
	content.ID,
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
