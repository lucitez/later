package service

import (
	"later/pkg/model"
	"later/pkg/repository"

	"github.com/google/uuid"
)

type Message struct {
	Repo repository.Message
}

func NewMessage(
	repo repository.Message,
) Message {
	return Message{repo}
}

func (c *Message) ByChatID(
	userID uuid.UUID,
	limit int,
	offset int,
) ([]model.Message, error) {
	return c.Repo.ByChatID(userID, limit, offset)
}
