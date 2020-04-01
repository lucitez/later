package repository_test

import (
	"later/pkg/model"
	"later/pkg/repository"
	"later/pkg/repository/util"
	"later/pkg/util/wrappers"
	"testing"
)

var contentRepo repository.Content

var content, _ = model.NewContent(
	"title",
	wrappers.NewNullStringFromString("description"),
	wrappers.NewNullStringFromString("image.com"),
	wrappers.NewNullStringFromString("jpeg"),
	"glump.com",
	"glump")

func TestContentInsertAndByID(t *testing.T) {
	beforeEach()
	contentRepo.Insert(content)
	actual, _ := contentRepo.ByID(content.ID)

	util.AssertEquals(t, actual, content)
}

func TestAll(t *testing.T) {
	beforeEach()
	contentRepo.Insert(content)
	actual, _ := contentRepo.All(1)

	util.AssertContainsOne(t, actual, *content)
}
