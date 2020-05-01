package repository

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/lucitez/later/pkg/model"
	"github.com/lucitez/later/pkg/repository/util"
)

type UserMessage struct {
	DB *sql.DB
}

func NewUserMessage(DB *sql.DB) UserMessage {
	return UserMessage{DB}
}

func (repo *UserMessage) Insert(userMessage model.UserMessage) error {
	_, err := repo.DB.Exec(
		util.GenerateInsertStatement(userMessage, "user_messages"),
		util.GenerateInsertArguments(&userMessage)...,
	)

	return err
}

func (repo *UserMessage) MarkReadByChatAndUser(chatID uuid.UUID, userID uuid.UUID) {
	statement := `
	UPDATE user_messages
	SET read_at = now()
	WHERE chat_id = $1
	AND user_id = $2
	AND deleted_at IS NULL;
	`

	_, err := repo.DB.Exec(statement, chatID, userID)

	if err != nil {
		log.Printf("[WARN] unhandled error updating user_messages: %v\n", err)
	}
}

// UnreadByChatAndUser returns true if there are any unread messages in a chat for the given user
func (repo *UserMessage) UnreadByChatAndUser(chatID uuid.UUID, userID uuid.UUID) bool {
	statement := `
	SELECT EXISTS 
		(
			SELECT * FROM user_messages
			WHERE chat_id = $1
			AND user_id = $2
			AND read_at IS NULL
			AND deleted_at IS NULL
		);
	`

	var hasUnread bool

	row := repo.DB.QueryRow(statement, chatID, userID)
	err := row.Scan(&hasUnread)

	if err != nil && err != sql.ErrNoRows {
		log.Printf("[WARN] unhandled error querying unread user_messages: %v\n", err)
	}

	return hasUnread // zero value is false
}
