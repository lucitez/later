package repository_test

import (
	"testing"

	"github.com/lucitez/later/pkg/model"

	"github.com/lucitez/later/pkg/repository"
)

var userMessageRepo repository.UserMessage

var userMessage = model.NewUserMessage(
	chat.ID,
	userID,
	message.ID,
)

func TestUnreadByUser(t *testing.T) {
	beforeEach(t)

	userMessageRepo.Insert(userMessage)

	actual := userMessageRepo.UnreadByChatAndUser(chat.ID, userID)

	testUtil.Assert.True(actual)
}

func TestMarkRead(t *testing.T) {
	beforeEach(t)

	userMessageRepo.Insert(userMessage)
	userMessageRepo.MarkReadByChatAndUser(chat.ID, userID)

	actual := userMessageRepo.UnreadByChatAndUser(chat.ID, userID)

	testUtil.Assert.False(actual)
}
