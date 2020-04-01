package repository_test

import (
	"later/pkg/model"
	"later/pkg/repository"
	"later/pkg/repository/util"
	"later/pkg/util/wrappers"
	"testing"
)

var userContentRepo repository.UserContent

var userContent, _ = model.NewUserContent(
	shareID,
	contentID,
	wrappers.NewNullStringFromString("jpeg"),
	userID,
	userID)

func TestUserContentInsertAndByID(t *testing.T) {
	beforeEach()

	userContentRepo.Insert(userContent)

	actual, _ := userContentRepo.ByID(userContent.ID)

	util.AssertEquals(t, actual, userContent)
}

func TestAllUserContent(t *testing.T) {
	beforeEach()

	userContentRepo.Insert(userContent)
	actual, _ := userContentRepo.All(1)

	util.AssertContainsOne(t, actual, *userContent)
}
