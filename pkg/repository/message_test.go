package repository_test

import (
	"github.com/lucitez/later/pkg/model"
	"github.com/lucitez/later/pkg/repository"
	"testing"
)

var messageRepo repository.Message

var message = model.NewMessage(
	chat.ID,
	userID,
	"test",
)

func TestMessageInsertAndByID(t *testing.T) {
	beforeEach(t)

	messageRepo.Insert(message)

	actual, _ := messageRepo.ByID(message.ID)

	testUtil.Assert.Equal(*actual, message)
}

func TestMessagesByChatID(t *testing.T) {
	beforeEach(t)

	messageRepo.Insert(message)
	actual, _ := messageRepo.ByChatID(chat.ID, 1, 0)

	testUtil.Assert.Contains(actual, message)
}
