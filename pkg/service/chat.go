package service

import (
	"later/pkg/model"
	"later/pkg/repository"

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

func (c *Chat) ForUser(userID uuid.UUID) ([]model.Chat, error) {
	return c.Repo.ByUserID(userID)
}
