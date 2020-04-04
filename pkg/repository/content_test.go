package repository_test

import (
	"later/pkg/model"
	"later/pkg/repository"
	"later/pkg/util/wrappers"
	"testing"
)

var contentRepo repository.Content

var content = model.NewContent(
	wrappers.NewNullStringFromString("title"),
	wrappers.NewNullStringFromString("description"),
	wrappers.NewNullStringFromString("image.com"),
	wrappers.NewNullStringFromString("jpeg"),
	"glump.com",
	"glump",
)

func TestContentInsertAndByID(t *testing.T) {
	beforeEach(t)
	contentRepo.Insert(content)
	actual := contentRepo.ByID(content.ID)

	testUtil.Assert.Equal(*actual, content)
}

func TestAll(t *testing.T) {
	beforeEach(t)
	contentRepo.Insert(content)
	actual := contentRepo.All(1)

	testUtil.Assert.Contains(actual, content)
}
