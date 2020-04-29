package repository

import (
	"database/sql"
	"github.com/lucitez/later/api/src/pkg/model"
	"github.com/lucitez/later/api/src/pkg/repository/util"

	"github.com/google/uuid"

	// Postgres driver
	_ "github.com/lib/pq"
)

// Chat ...
type Chat struct {
	DB *sql.DB
}

// NewChat for wire generation
func NewChat(db *sql.DB) Chat {
	return Chat{db}
}

var chatSelectStatement = util.GenerateSelectStatement(model.Chat{}, "chats")

// Insert inserts new chat
func (repository *Chat) Insert(chat model.Chat) error {
	_, err := repository.DB.Exec(
		util.GenerateInsertStatement(chat, "chats"),
		util.GenerateInsertArguments(&chat)...,
	)

	return err
}

// ByID gets a chat by id
func (repository *Chat) ByID(id uuid.UUID) (*model.Chat, error) {
	var chat model.Chat

	statement := chatSelectStatement + " WHERE id = $1;"

	row := repository.DB.QueryRow(statement, id)

	err := util.ScanRowInto(row, &chat)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &chat, err
}

func (repository *Chat) ByUserID(userID uuid.UUID) ([]model.Chat, error) {
	statement := chatSelectStatement + `
	WHERE (user1_id = $1 OR user2_id = $1)
	AND deleted_at IS NULL;
	`

	rows, err := repository.DB.Query(statement, userID)

	if err != nil {
		return nil, err
	}

	return repository.scanRows(rows)
}

func (repository *Chat) ByUserIDs(user1ID uuid.UUID, user2ID uuid.UUID) (*model.Chat, error) {
	var chat model.Chat

	statement := chatSelectStatement + `
	WHERE (
		(user1_id = $1 AND user2_id = $2)
		OR
		(user1_id = $2 AND user2_id = $1)
	)
	AND deleted_at IS NULL;
	`

	row := repository.DB.QueryRow(statement, user1ID, user2ID)

	err := util.ScanRowInto(row, &chat)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &chat, err
}

func (repository *Chat) scanRows(rows *sql.Rows) ([]model.Chat, error) {
	chats := []model.Chat{}

	defer rows.Close()

	for rows.Next() {
		var chat model.Chat
		err := util.ScanRowsInto(rows, &chat)

		if err != nil {
			return nil, err
		}

		chats = append(chats, chat)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return chats, nil
}
