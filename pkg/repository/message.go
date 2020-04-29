package repository

import (
	"database/sql"
	"github.com/lucitez/later/pkg/model"
	"github.com/lucitez/later/pkg/repository/util"

	"github.com/google/uuid"

	// Postgres driver
	_ "github.com/lib/pq"
)

// Message ...
type Message struct {
	DB *sql.DB
}

// NewMessage for wire generation
func NewMessage(db *sql.DB) Message {
	return Message{db}
}

var messageSelectStatement = util.GenerateSelectStatement(model.Message{}, "messages")

// Insert inserts new message
func (repository *Message) Insert(message model.Message) error {
	_, err := repository.DB.Exec(
		util.GenerateInsertStatement(message, "messages"),
		util.GenerateInsertArguments(&message)...,
	)

	return err
}

// ByID gets a message by id
func (repository *Message) ByID(id uuid.UUID) (*model.Message, error) {
	var message model.Message

	statement := messageSelectStatement + " WHERE id = $1;"

	row := repository.DB.QueryRow(statement, id)

	err := util.ScanRowInto(row, &message)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &message, err
}

func (repository *Message) ByChatID(
	chatID uuid.UUID,
	limit int,
	offset int,
) ([]model.Message, error) {
	statement := messageSelectStatement + `
	WHERE chat_id = $1
	AND deleted_at IS NULL
	ORDER BY created_at desc
	LIMIT $2
	OFFSET $3;
	`

	rows, err := repository.DB.Query(
		statement,
		chatID,
		limit,
		offset,
	)

	if err != nil {
		return nil, err
	}

	return repository.scanRows(rows)
}

func (repository *Message) scanRows(rows *sql.Rows) ([]model.Message, error) {
	messages := []model.Message{}

	defer rows.Close()

	for rows.Next() {
		var message model.Message
		err := util.ScanRowsInto(rows, &message)

		if err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
