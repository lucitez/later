package repository

import (
	"database/sql"

	"github.com/lucitez/later/pkg/model"
	"github.com/lucitez/later/pkg/repository/util"

	"github.com/google/uuid"

	// Postgres driver
	_ "github.com/lib/pq"
)

// Content ...
type Content struct {
	DB *sql.DB
}

// NewContent for wire generation
func NewContent(db *sql.DB) Content {
	return Content{db}
}

var contentSelectStatement = util.GenerateSelectStatement(model.Content{}, "content")

// Insert inserts new content
func (repository *Content) Insert(content model.Content) error {
	_, err := repository.DB.Exec(
		util.GenerateInsertStatement(content, "content"),
		util.GenerateInsertArguments(&content)...,
	)

	return err
}

// ByID gets a content by id
func (repository *Content) ByID(id uuid.UUID) (*model.Content, error) {
	var content model.Content

	statement := contentSelectStatement + " WHERE id = $1;"

	row := repository.DB.QueryRow(statement, id)

	err := util.ScanRowInto(row, &content)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &content, err
}

// All returns all content
func (repository *Content) All(limit int) ([]model.Content, error) {
	statement := contentSelectStatement + `
	LIMIT $1;
	`

	rows, err := repository.DB.Query(statement, limit)

	if err != nil {
		return nil, err
	}

	return repository.scanRows(rows)
}

// TasteByUserID sum of unique shares for all content created [initial share] by a user.
func (repository *Content) TasteByUserID(userID uuid.UUID) (int, error) {
	statement := `
	SELECT SUM(shares)
	FROM content
	WHERE created_by = $1;
	`

	rows, err := repository.DB.Query(statement, userID)

	if err != nil {
		return 0, err
	}

	var taste int

	for rows.Next() {
		rows.Scan(&taste)
	}

	return taste, rows.Err()

}

// IncrementShareCount does what it sounds like
func (repository *Content) IncrementShareCount(id uuid.UUID, amount int) error {
	statement := "UPDATE content SET shares = shares + $2 WHERE id = $1;"

	_, err := repository.DB.Exec(statement, id, amount)

	return err
}

// func (repo *Message) Popular(
// 	limit int,
// 	offset int,
// ) ([]model.Message, error) {

// 	statement := `
// 	WITH most_shared AS (
// 		SELECT content_id, count(*) FROM (
// 			SELECT content_id, recipient_user_id
// 			FROM shares
// 			GROUP BY content_id, recipient_user_id
// 		)

// 		SELECT content_id, count(distinct recipient_user_id) FROM shares
// 		WHERE created_at > now() = interval '1 day'
// 		ORDER BY count(*)
// 	)
// 	`

// 	statement := messageSelectStatement + `
// 	WHERE
// 	ORDER BY created_at desc
// 	LIMIT $2
// 	OFFSET $3;
// 	`

// 	rows, err := repository.DB.Query(
// 		statement,
// 		chatID,
// 		limit,
// 		offset,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return repository.scanRows(rows)
// }

func (repository *Content) scanRows(rows *sql.Rows) ([]model.Content, error) {
	contents := []model.Content{}

	defer rows.Close()

	for rows.Next() {
		var content model.Content
		err := util.ScanRowsInto(rows, &content)

		if err != nil {
			return nil, err
		}

		contents = append(contents, content)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return contents, nil
}
