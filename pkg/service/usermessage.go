package service

import (
	"log"

	"github.com/google/uuid"

	"github.com/lucitez/later/pkg/model"
	"github.com/lucitez/later/pkg/repository"
)

type UserMessage struct {
	Repo repository.UserMessage
	Chat Chat
}

func NewUserMessage(
	repo repository.UserMessage,
	chat Chat,
) UserMessage {
	return UserMessage{
		repo,
		chat,
	}
}

// CreateByMessage emits user messages for users after a message is sent to a chat.
// This is called via goroutine
// 1. Get consumers of chat [other user_id or all members of group of group chat]
// 2. Create UserMessage for each of these people
// 3. Send notifications
func (um *UserMessage) CreateByMessage(message model.Message) {
	chat, err := um.Chat.ByID(message.ChatID)

	if err != nil && chat == nil {
		log.Printf("[WARN] Error creating user message. Error: %v Chat:%v\n", err, chat)
	}

	targetUserIds := []uuid.UUID{}

	if chat.GroupID.Valid {
		// lookup users of group and add to target users
	} else {
		if chat.User1ID.ID == message.SentBy {
			targetUserIds = append(targetUserIds, chat.User2ID.ID)
		} else {
			targetUserIds = append(targetUserIds, chat.User1ID.ID)
		}
	}

	for _, userID := range targetUserIds {
		userMessage := model.NewUserMessage(
			chat.ID,
			userID,
			message.ID,
		)

		err := um.Repo.Insert(userMessage)

		if err != nil {
			log.Printf("[WARN] Error creating user message. Error: %v\n", err)
		}
	}
}

func (um *UserMessage) UnreadByChatAndUser(chatID uuid.UUID, userID uuid.UUID) bool {
	return um.Repo.UnreadByChatAndUser(chatID, userID)
}

func (um *UserMessage) MarkReadByChatAndUser(chatID uuid.UUID, userID uuid.UUID) {
	um.Repo.MarkReadByChatAndUser(chatID, userID)
}
