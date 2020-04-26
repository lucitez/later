package transfer

import (
	"fmt"
	"later/pkg/model"
	"later/pkg/response"
	"later/pkg/service"

	"github.com/google/uuid"
)

type Chat struct {
	UserService    service.User
	MessageService service.Message
}

func NewChat(
	userService service.User,
	messageService service.Message,
) Chat {
	return Chat{
		userService,
		messageService,
	}
}

func (c *Chat) WireChatsFrom(chats []model.Chat, userID uuid.UUID) []response.WireChat {
	wireChats := make([]response.WireChat, len(chats))

	for i, chat := range chats {
		wireChats[i] = c.WireChatFromChat(chat, userID)
	}

	return wireChats
}

func (c *Chat) WireChatFromChat(chat model.Chat, userID uuid.UUID) response.WireChat {
	var display string
	var activity string
	var activityMessage = "sent a message"

	if conversationWith := getChatWithUserID(chat, userID); conversationWith != nil {
		if user := c.UserService.ByID(*conversationWith); user != nil {
			display = user.Username
		}
	} else {
		display = "GROUP NAME"
	}

	if messages, err := c.MessageService.ByChatID(chat.ID, 1, 0); err == nil && len(messages) > 0 {
		lastMessage := messages[0]
		if lastMessage.ContentID.Valid {
			activityMessage = "shared content"
		}

		if lastMessage.SentBy == userID {
			activity = fmt.Sprintf("You %s", activityMessage)
		}
		if user := c.UserService.ByID(messages[0].SentBy); user != nil {
			activity = fmt.Sprintf("%s %s", user.Username, activityMessage)
		}
	}

	return response.WireChat{
		ChatID:   chat.ID,
		Display:  display,
		Activity: activity,
	}
}

func getChatWithUserID(chat model.Chat, userID uuid.UUID) *uuid.UUID {
	if chat.User1ID.Valid && chat.User1ID.ID != userID {
		return &chat.User1ID.ID
	} else if chat.User2ID.Valid && chat.User2ID.ID != userID {
		return &chat.User2ID.ID
	} else {
		return nil
	}
}
