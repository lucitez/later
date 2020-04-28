package server

import (
	"later/pkg/request"
	"later/pkg/service"
	"later/pkg/transfer"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Message ...
type Message struct {
	Service  service.Message
	Transfer transfer.Message
}

// NewMessage for wire generation
func NewMessage(
	service service.Message,
	transfer transfer.Message,
) Message {
	return Message{
		Service:  service,
		Transfer: transfer,
	}
}

func (server *Message) Prefix() string {
	return "/messages"
}

// Routes defines the routes for message API
func (server *Message) Routes(router *gin.RouterGroup) []gin.IRoutes {
	return []gin.IRoutes{
		router.GET("/by-chat-id", server.byChatID),
		router.POST("/send", server.sendMessage),
	}
}

func (server *Message) byChatID(c *gin.Context) {
	defaultLimit := "20"
	defaultOffset := "0"

	deser := NewDeser(
		c,
		QueryParameter{name: "chat_id", kind: UUID, required: true},
		QueryParameter{name: "limit", kind: Int, fallback: &defaultLimit},
		QueryParameter{name: "offset", kind: Int, fallback: &defaultOffset},
	)

	if parameters, ok := deser.DeserQueryParams(); ok {
		chatID := parameters["chat_id"].(*uuid.UUID)
		limit := parameters["limit"].(*int)
		offset := parameters["offset"].(*int)

		messages, err := server.Service.ByChatID(
			*chatID,
			*limit,
			*offset,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		wireMessages := server.Transfer.WireMessagesFrom(messages)

		c.JSON(http.StatusOK, wireMessages)
	}
}

func (server *Message) sendMessage(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)

	var body request.MessageSendRequestBody

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	message, err := server.Service.CreateFromMessage(
		body.ChatID,
		userID,
		body.Message,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, message)
}
