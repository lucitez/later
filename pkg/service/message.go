package service

import (
	"errors"
	"later/pkg/model"
	"later/pkg/repository"

	"github.com/google/uuid"
)

type Message struct {
	Repo repository.Message
	Chat Chat
}

func NewMessage(
	repo repository.Message,
	chat Chat,
) Message {
	return Message{repo, chat}
}

func (c *Message) ByChatID(
	chatID uuid.UUID,
	limit int,
	offset int,
) ([]model.Message, error) {
	return c.Repo.ByChatID(chatID, limit, offset)
}

func (c *Message) CreateFromMessage(
	chatID uuid.UUID,
	sentBy uuid.UUID,
	message string,
) (*model.Message, error) {
	newMessage := model.NewMessage(
		chatID,
		sentBy,
		message,
	)

	if err := c.Repo.Insert(newMessage); err != nil {
		return nil, err
	}

	return &newMessage, nil
}

func (c *Message) CreateFromShare(
	share model.Share,
) (*model.Message, error) {
	// find or create
	if chat, err := c.Chat.FindOrCreateByUserIDs(share.SentByUserID, share.RecipientUserID); err != nil {
		return nil, err
	} else if chat == nil {
		return nil, errors.New("Could not create chat")
	} else {
		message := model.NewMessageFromContent(
			chat.ID,
			share.SentByUserID,
			share.ContentID,
		)

		if err := c.Repo.Insert(message); err != nil {
			return nil, err
		}

		return &message, nil
	}
}
