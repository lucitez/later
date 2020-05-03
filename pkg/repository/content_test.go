package repository_test

import (
	"testing"

	"github.com/lucitez/later/pkg/model"
	"github.com/lucitez/later/pkg/repository"
	"github.com/lucitez/later/pkg/util/wrappers"
)

var contentRepo repository.Content

var content = model.NewContent(
	wrappers.NewNullStringFromString("title"),
	wrappers.NewNullStringFromString("description"),
	wrappers.NewNullStringFromString("thumbnail.jpg"),
	wrappers.NewNullStringFromString("watch"),
	"youtube.com",
	"youtube",
	userID,
)

func TestContentInsertAndByID(t *testing.T) {
	beforeEach(t)
	contentRepo.Insert(content)

	actual, _ := contentRepo.ByID(content.ID)

	testUtil.Assert.Equal(*actual, content)
}

func TestAll(t *testing.T) {
	beforeEach(t)
	contentRepo.Insert(content)
	actual, _ := contentRepo.All(1)

	testUtil.Assert.Contains(actual, content)
}

func TestTasteByUserID(t *testing.T) {
	beforeEach(t)

	contentRepo.Insert(content)
	contentRepo.IncrementShareCount(content.ID, 1)
	actual, _ := contentRepo.TasteByUserID(content.CreatedBy)

	testUtil.Assert.Equal(actual, 1)
}

func TestTasteByUserIDNoShares(t *testing.T) {
	beforeEach(t)

	actual, _ := contentRepo.TasteByUserID(content.CreatedBy)

	testUtil.Assert.Equal(actual, 0)
}
