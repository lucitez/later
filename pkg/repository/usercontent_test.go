package repository_test

import (
	"later/pkg/model"
	"later/pkg/repository"
	"later/pkg/util/wrappers"
	"testing"

	"github.com/google/uuid"
)

var userContentRepo repository.UserContent

var userContent = model.NewUserContent(
	shareID,
	contentID,
	wrappers.NewNullStringFromString("jpeg"),
	userID,
	userID,
)

func TestUserContentInsertAndByID(t *testing.T) {
	beforeEach(t)

	userContentRepo.Insert(userContent)

	actual := userContentRepo.ByID(userContent.ID)

	testUtil.Assert.Equal(*actual, userContent)
}

func TestUserContentByIDNull(t *testing.T) {
	beforeEach(t)

	id, _ := uuid.NewRandom()

	actual := userContentRepo.ByID(id)

	testUtil.Assert.Nil(actual)
}

func TestAllUserContent(t *testing.T) {
	beforeEach(t)

	userContentRepo.Insert(userContent)
	actual := userContentRepo.All(1)

	testUtil.Assert.Contains(actual, userContent)
}
