package repository

import (
	"database/sql"
	"later/pkg/model"

	"github.com/google/uuid"

	// Postgres driver
	_ "github.com/lib/pq"
)

// ContentRepository ...
type ContentRepository struct {
	DB *sql.DB
}

// NewContentRepository for wire generation
func NewContentRepository(db *sql.DB) ContentRepository {
	return ContentRepository{db}
}

// Insert inserts new content
func (repository *ContentRepository) Insert(content *model.Content) (*model.Content, error) {

	statement := `
	INSERT INTO content (
		id,
		title,
		description,
		image_url,
		content_type,
		url,
		domain,
		shares
	)
	VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8
	)
	`

	_, err := repository.DB.Exec(
		statement,
		content.ID,
		content.Title,
		content.Description,
		content.ImageURL,
		content.ContentType,
		content.URL,
		content.Domain,
		content.Shares)

	if err != nil {
		return nil, err
	}

	return content, nil
}

// ByID gets a content by id
func (repository *ContentRepository) ByID(id uuid.UUID) (*model.Content, error) {
	var content model.Content

	statement := `
	SELECT * FROM content 
	WHERE id = $1
	`

	row := repository.DB.QueryRow(statement, id)

	err := content.ScanRow(row)

	return &content, err
}

// All returns all content
func (repository *ContentRepository) All(limit int) ([]model.Content, error) {
	statement := `
	SELECT * FROM content
	WHERE deleted_at IS NULL
	LIMIT $1;
	`

	rows, err := repository.DB.Query(statement, limit)

	if err != nil {
		return nil, err
	}

	return repository.scanRows(rows)
}

func (repository *ContentRepository) scanRows(rows *sql.Rows) ([]model.Content, error) {
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
