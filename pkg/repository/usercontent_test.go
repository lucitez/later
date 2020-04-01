package repository_test

import (
	"later/pkg/model"
	"later/pkg/repository"
	"later/pkg/repository/util"
	"later/pkg/util/wrappers"
	"testing"
)

var userContentRepository repository.UserContent

var userContent, _ = model.NewUserContent(
	shareID,
	contentID,
	wrappers.NewNullStringFromString("jpeg"),
	userID,
	userID)

func TestUserContentInsertAndByID(t *testing.T) {
	beforeEach()

	userContentRepository.Insert(userContent)

	actual, _ := userContentRepository.ByID(userContent.ID)

	util.AssertEquals(t, actual, userContent)
}

func TestAllUserContent(t *testing.T) {
	beforeEach()

	userContentRepository.Insert(userContent)
	actual, _ := userContentRepository.All(1)

	util.AssertContainsOne(t, actual, *userContent)
}
