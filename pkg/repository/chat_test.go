package repository_test

import (
	"testing"

	"github.com/google/uuid"

	"github.com/lucitez/later/pkg/model"
	"github.com/lucitez/later/pkg/repository"
)

var chatRepo repository.Chat

var chat = model.NewUserChat(
	userID,
	userID2,
)

func TestChatInsertAndByID(t *testing.T) {
	beforeEach(t)

	chatRepo.Insert(chat)

	actual, _ := chatRepo.ByID(chat.ID)

	testUtil.Assert.Equal(*actual, chat)
}

func TestChatsByUserID(t *testing.T) {
	beforeEach(t)

	chatRepo.Insert(chat)
	actual, _ := chatRepo.ByUserID(userID)

	testUtil.Assert.Contains(actual, chat)
}

func TestChatsByUserIDOrdering(t *testing.T) {
	beforeEach(t)
	id, _ := uuid.NewRandom()

	chat2 := model.NewUserChat(
		userID,
		id,
	)

	chatRepo.Insert(chat)
	chatRepo.Insert(chat2)

	message2 := model.NewMessage(
		chat2.ID,
		userID,
		"hi",
	)

	messageRepo.Insert(message2)

	actual, _ := chatRepo.ByUserID(userID)

	testUtil.Assert.Equal(actual[0].ID, chat2.ID)
	testUtil.Assert.NotNil(actual[0].LastMessageSentAt)
}

func TestChatByUserIDs(t *testing.T) {
	beforeEach(t)

	chatRepo.Insert(chat)
	actual, err := chatRepo.ByUserIDs(userID, userID2)

	if err != nil {
		t.Error(err)
	}

	testUtil.Assert.Equal(*actual, chat)
}
