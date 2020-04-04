package repository

import (
	"database/sql"
	"later/pkg/model"
	"later/pkg/repository/util"
	"log"

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
	statement := util.GenerateInsertStatement(content, "content")

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
		content.UpdatedAt,
	)

	return err
}

// ByID gets a content by id
func (repository *Content) ByID(id uuid.UUID) *model.Content {
	var content model.Content

	statement := contentSelectStatement + " WHERE id = $1;"

	row := repository.DB.QueryRow(statement, id)

	content.ScanRow(row)

	return &content
}

// All returns all content
func (repository *Content) All(limit int) []model.Content {
	statement := contentSelectStatement + `
	WHERE deleted_at IS NULL
	LIMIT $1;
	`

	rows, err := repository.DB.Query(statement, limit)

	if err != nil {
		log.Fatal(err)
	}

	return repository.scanRows(rows)
}

func (repository *Content) scanRows(rows *sql.Rows) []model.Content {
	contents := []model.Content{}

	defer rows.Close()

	for rows.Next() {
		var content model.Content
		content.ScanRows(rows)

		contents = append(contents, content)
	}

	err := rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return contents
}
