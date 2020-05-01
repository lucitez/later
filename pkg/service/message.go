package service

import (
	"errors"

	"github.com/lucitez/later/pkg/model"
	"github.com/lucitez/later/pkg/repository"

	"github.com/google/uuid"
)

type Message struct {
	Repo        repository.Message
	Chat        Chat
	UserMessage UserMessage
}

func NewMessage(
	repo repository.Message,
	chat Chat,
	userMessage UserMessage,
) Message {
	return Message{repo, chat, userMessage}
}

func (c *Message) ByChatID(
	chatID uuid.UUID,
	limit int,
	offset int,
) ([]model.Message, error) {
	return c.Repo.ByChatID(chatID, limit, offset)
}

// TODO update pubsub
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

	go c.UserMessage.CreateByMessage(newMessage)

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

		go c.UserMessage.CreateByMessage(message)

		return &message, nil
	}
}
