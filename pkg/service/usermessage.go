package service

import (
	"log"

	"github.com/google/uuid"

	"github.com/lucitez/later/pkg/model"
	"github.com/lucitez/later/pkg/repository"
)

type UserMessage struct {
	repo                repository.UserMessage
	userService         User
	notificationService Notification
	chat                Chat
}

func NewUserMessage(
	repo repository.UserMessage,
	userService User,
	notificationService Notification,
	chat Chat,
) UserMessage {
	return UserMessage{
		repo,
		userService,
		notificationService,
		chat,
	}
}

// CreateByMessage emits user messages for users after a message is sent to a chat.
// This is called via goroutine
// 1. Get consumers of chat [other user_id or all members of group of group chat]
// 2. Create UserMessage for each of these people
// 3. Send notifications
func (um *UserMessage) CreateByMessage(message model.Message) {
	chat, err := um.chat.ByID(message.ChatID)

	if err != nil || chat == nil {
		log.Printf("[WARN] Error creating user message. Error: %v Chat:%v\n", err, chat)
		return
	}

	sender, err := um.userService.ByID(message.SentBy)

	if err != nil || sender == nil {
		log.Printf("[WARN] Error getting message sender: %s\n", err.Error())
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

		err := um.repo.Insert(userMessage)

		if err != nil {
			log.Printf("[WARN] Error creating user message. Error: %v\n", err)
		}

		if !message.ContentID.Valid {
			go um.sendMessageSentNotification(message, *sender, userID)
		}
	}
}

func (um *UserMessage) sendMessageSentNotification(message model.Message, from model.User, to uuid.UUID) {
	notificationMessage := PushMessage{
		To:    to,
		Title: from.Name + " sent you a message",
		Body:  message.Message.String,
	}

	um.notificationService.SendMessage(notificationMessage)
}

func (um *UserMessage) UnreadByChatAndUser(chatID uuid.UUID, userID uuid.UUID) bool {
	return um.repo.UnreadByChatAndUser(chatID, userID)
}

func (um *UserMessage) MarkReadByChatAndUser(chatID uuid.UUID, userID uuid.UUID) {
	um.repo.MarkReadByChatAndUser(chatID, userID)
}
