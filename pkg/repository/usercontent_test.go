package repository_test

import (
	"later/pkg/model"
	"later/pkg/repository"
	"later/pkg/service/body"
	"later/pkg/util/wrappers"
	"testing"

	"github.com/google/uuid"
)

var userContentRepo repository.UserContent

var userContent = model.NewUserContent(
	share.ID,
	content.ID,
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

func TestSaveUserContent(t *testing.T) {
	beforeEach(t)

	userContentRepo.Insert(userContent)
	userContentRepo.Save(
		userContent.ID,
		wrappers.NewNullStringFromString("memes"),
	)

	actual := userContentRepo.ByID(userContent.ID)

	testUtil.Assert.Equal(actual.Tag.String, "memes")
	testUtil.Assert.True(actual.SavedAt.Valid)
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
	contentRepo.Insert(content)

	actual := userContentRepo.Filter(
		userContent.UserID,
		nil,
		&contentType,
		false,
		nil,
		1,
	)

	testUtil.Assert.Contains(actual, userContent)
}

func TestFilterSaved(t *testing.T) {
	contentType := "watch"
	beforeEach(t)

	userContentRepo.Insert(userContent)
	contentRepo.Insert(content)

	actual := userContentRepo.Filter(
		userContent.UserID,
		nil,
		&contentType,
		true, // saved
		nil,
		1,
	)

	testUtil.Assert.Empty(actual)

	userContentRepo.Save(userContent.ID, wrappers.NewNullStringFromString("memes"))

	actual = userContentRepo.Filter(
		userContent.UserID,
		nil,
		&contentType,
		true, // saved
		nil,
		1,
	)

	testUtil.Assert.NotEmpty(actual)
}

func TestFilterSearchTag(t *testing.T) {
	beforeEach(t)

	userContentRepo.Insert(userContent)
	contentRepo.Insert(content)

	updateBody := body.UserContentUpdateBody{
		ID:  userContent.ID,
		Tag: wrappers.NewNullStringFromString("memes"),
	}

	userContentRepo.Update(updateBody)

	search := "memes"

	actual := userContentRepo.Filter(
		userContent.UserID,
		nil,
		nil,
		false,
		&search,
		1,
	)

	testUtil.Assert.NotEmpty(actual)
}

func TestUpdate(t *testing.T) {
	beforeEach(t)

	userContentRepo.Insert(userContent)

	updateBody := body.UserContentUpdateBody{
		ID:  userContent.ID,
		Tag: wrappers.NewNullStringFromString("memes"),
	}

	userContentRepo.Update(updateBody)
	actual := userContentRepo.ByID(userContent.ID)

	testUtil.Assert.Equal(actual.Tag.String, "memes")
}

func TestFilterTags(t *testing.T) {
	beforeEach(t)

	userContentRepo.Insert(userContent)

	updateBody := body.UserContentUpdateBody{
		ID:  userContent.ID,
		Tag: wrappers.NewNullStringFromString("memes"),
	}

	userContentRepo.Update(updateBody)
	actual, _ := userContentRepo.FilterTags(
		userID,
		nil,
	)

	testUtil.Assert.Contains(actual, "memes")
}

func TestFilterTagsSearchPositive(t *testing.T) {
	beforeEach(t)

	userContentRepo.Insert(userContent)

	updateBody := body.UserContentUpdateBody{
		ID:  userContent.ID,
		Tag: wrappers.NewNullStringFromString("memes"),
	}

	userContentRepo.Update(updateBody)

	search := "mem"
	actual, _ := userContentRepo.FilterTags(
		userID,
		&search,
	)

	testUtil.Assert.Contains(actual, "memes")
}

func TestFilterTagsSearchNegative(t *testing.T) {
	beforeEach(t)

	userContentRepo.Insert(userContent)

	updateBody := body.UserContentUpdateBody{
		ID:  userContent.ID,
		Tag: wrappers.NewNullStringFromString("memes"),
	}

	userContentRepo.Update(updateBody)

	search := "something"
	actual, _ := userContentRepo.FilterTags(
		userID,
		&search,
	)

	t.Error(actual)

	testUtil.Assert.Empty(actual)
}

func TestFilterTagsSearchEmpty(t *testing.T) {
	beforeEach(t)

	userContentRepo.Insert(userContent)

	actual, _ := userContentRepo.FilterTags(
		userID,
		nil,
	)

	testUtil.Assert.Empty(actual)
}

func TestByTag(t *testing.T) {
	beforeEach(t)

	userContentRepo.Insert(userContent)

	updateBody := body.UserContentUpdateBody{
		ID:  userContent.ID,
		Tag: wrappers.NewNullStringFromString("memes"),
	}

	userContentRepo.Update(updateBody)

	actual, _ := userContentRepo.ByTag(
		userID,
		"memes",
	)

	testUtil.Assert.NotEmpty(actual)
	testUtil.Assert.Equal(actual[0].ID, userContent.ID)
}
