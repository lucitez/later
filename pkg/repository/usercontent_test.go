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
	wrappers.NewNullStringFromString("watch"),
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

func TestArchiveUserContent(t *testing.T) {
	beforeEach(t)

	userContentRepo.Insert(userContent)
	userContentRepo.Archive(
		userContent.ID,
		wrappers.NewNullStringFromString("memes"),
	)

	actual := userContentRepo.ByID(userContent.ID)

	testUtil.Assert.Equal(actual.Tag.String, "memes")
	testUtil.Assert.True(actual.ArchivedAt.Valid)
}

func TestDeleteUserContent(t *testing.T) {
	beforeEach(t)

	userContentRepo.Insert(userContent)
	userContentRepo.Delete(userContent.ID)

	actual := userContentRepo.ByID(userContent.ID)

	testUtil.Assert.True(actual.DeletedAt.Valid)
}

func TestFilter(t *testing.T) {
	contentType := "watch"
	beforeEach(t)

	userContentRepo.Insert(userContent)

	actual := userContentRepo.Filter(
		userContent.UserID,
		nil,
		&contentType,
		false,
		1,
	)

	testUtil.Assert.Contains(actual, userContent)
}

func TestFilterArchived(t *testing.T) {
	contentType := "watch"
	beforeEach(t)

	userContentRepo.Insert(userContent)

	actual := userContentRepo.Filter(
		userContent.UserID,
		nil,
		&contentType,
		true, // archived
		1,
	)

	testUtil.Assert.Empty(actual)

	userContentRepo.Archive(userContent.ID, wrappers.NewNullStringFromString("memes"))

	actual = userContentRepo.Filter(
		userContent.UserID,
		nil,
		&contentType,
		true, // archived
		1,
	)

	testUtil.Assert.NotEmpty(actual)
}
