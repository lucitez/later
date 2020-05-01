package transfer

import (
	"fmt"

	"github.com/lucitez/later/pkg/model"
	"github.com/lucitez/later/pkg/response"
	"github.com/lucitez/later/pkg/service"
	"github.com/lucitez/later/pkg/util/wrappers"

	"github.com/google/uuid"
)

type Chat struct {
	UserService        service.User
	MessageService     service.Message
	UserMessageService service.UserMessage
}

func NewChat(
	userService service.User,
	messageService service.Message,
	userMessageService service.UserMessage,
) Chat {
	return Chat{
		userService,
		messageService,
		userMessageService,
	}
}

func (c *Chat) WireChatsFrom(chats []model.Chat, userID uuid.UUID) []response.WireChat {
	wireChats := make([]response.WireChat, len(chats))

	for i, chat := range chats {
		wireChats[i] = c.WireChatFromChat(chat, userID)
	}

	return wireChats
}

func (c *Chat) WireChatFromChat(chat model.Chat, userID uuid.UUID) (wireChat response.WireChat) {
	wireChat.ChatID = chat.ID
	var activityMessage = "sent a message"

	if conversationWith := getChatWithUserID(chat, userID); conversationWith != nil {
		if user := c.UserService.ByID(*conversationWith); user != nil {
			wireChat.OtherUserID = wrappers.NewNullUUIDFromUUID(user.ID)
			wireChat.OtherUserName = wrappers.NewNullStringFromString(user.Name)
			wireChat.OtherUserUsername = wrappers.NewNullStringFromString(user.Username)
		}
	} else {
		wireChat.GroupName = wrappers.NewNullStringFromString("PUT GROUP NAME HERE")
	}

	if messages, err := c.MessageService.ByChatID(chat.ID, 1, 0); err == nil && len(messages) > 0 {
		lastMessage := messages[0]
		if lastMessage.ContentID.Valid {
			activityMessage = "shared content"
		}

		if lastMessage.SentBy == userID {
			wireChat.Activity = fmt.Sprintf("You %s", activityMessage)
		} else if user := c.UserService.ByID(messages[0].SentBy); user != nil {
			wireChat.Activity = fmt.Sprintf("%s %s", user.Username, activityMessage)
		}
	}

	wireChat.LastMessageSentAt = chat.LastMessageSentAt
	wireChat.HasUnread = c.UserMessageService.UnreadByChatAndUser(chat.ID, userID)

	return
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
