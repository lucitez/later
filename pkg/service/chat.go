package service

import (
	"github.com/lucitez/later/pkg/model"
	"github.com/lucitez/later/pkg/repository"

	"github.com/google/uuid"
)

type Chat struct {
	Repo repository.Chat
}

func NewChat(
	repo repository.Chat,
) Chat {
	return Chat{repo}
}

func (c *Chat) ByID(id uuid.UUID) (*model.Chat, error) {
	return c.Repo.ByID(id)
}

func (c *Chat) ForUser(userID uuid.UUID) ([]model.Chat, error) {
	return c.Repo.ByUserID(userID)
}

func (c *Chat) ByUserIDs(user1ID uuid.UUID, user2ID uuid.UUID) (*model.Chat, error) {
	return c.Repo.ByUserIDs(user1ID, user2ID)
}

func (c *Chat) FindOrCreateByUserIDs(user1ID uuid.UUID, user2ID uuid.UUID) (*model.Chat, error) {
	if existingChat, err := c.Repo.ByUserIDs(user1ID, user2ID); err != nil {
		return nil, err
	} else if existingChat != nil {
		return existingChat, nil
	}

	chat := model.NewUserChat(user1ID, user2ID)

	if err := c.Repo.Insert(chat); err != nil {
		return nil, err
	}

	return &chat, nil
}
