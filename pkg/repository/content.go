package repository

import (
	"database/sql"
	"later/pkg/model"
	"later/pkg/repository/util"

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
func (repository *Content) Insert(content *model.Content) (*model.Content, error) {

	statement := util.GenerateInsertStatement(*content, "content")

	_, err := repository.DB.Exec(
		statement,
		content.ID,
		content.Title,
		content.Description,
		content.ImageURL,
		content.ContentType,
		content.URL,
		content.Domain,
		content.Shares,
		content.CreatedAt,
		content.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return content, nil
}

// ByID gets a content by id
func (repository *Content) ByID(id uuid.UUID) (*model.Content, error) {
	var content model.Content

	statement := contentSelectStatement + " WHERE id = $1;"

	row := repository.DB.QueryRow(statement, id)

	err := content.ScanRow(row)

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

func (repository *Content) scanRows(rows *sql.Rows) ([]model.Content, error) {
	contents := []model.Content{}

	defer rows.Close()

	for rows.Next() {
		var content model.Content
		err := content.ScanRows(rows)

		if err != nil {
			return nil, err
		}
		contents = append(contents, content)
	}

	err := rows.Err()
	if err != nil {
		return nil, err
	}

	return contents, nil
}
