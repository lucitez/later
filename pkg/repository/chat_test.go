package repository_test

import (
	"later/pkg/model"
	"later/pkg/repository"
	"testing"
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
