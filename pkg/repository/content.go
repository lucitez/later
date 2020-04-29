package repository

import (
	"database/sql"
	"github.com/lucitez/later/pkg/model"
	"github.com/lucitez/later/pkg/repository/util"
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
		content.CreatedBy,
		content.CreatedAt,
		content.UpdatedAt,
	)

	return err
}

// ByID gets a content by id
func (repository *Content) ByID(id uuid.UUID) (*model.Content, error) {
	var content model.Content

	statement := contentSelectStatement + " WHERE id = $1;"

	row := repository.DB.QueryRow(statement, id)

	return content.ScanRow(row)
}

// All returns all content
func (repository *Content) All(limit int) []model.Content {
	statement := contentSelectStatement + `
	LIMIT $1;
	`

	rows, err := repository.DB.Query(statement, limit)

	if err != nil {
		log.Fatal(err)
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

	if rows, err := repository.DB.Query(statement, userID); err != nil {
		return 0, err
	} else {
		var taste int

		for rows.Next() {
			rows.Scan(&taste)
		}

		return taste, rows.Err()
	}
}

// IncrementShareCount does what it sounds like
func (repository *Content) IncrementShareCount(id uuid.UUID, amount int) error {
	statement := "UPDATE content SET shares = shares + $2 WHERE id = $1;"

	_, err := repository.DB.Exec(statement, id, amount)

	return err
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
