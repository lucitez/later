package repository

import (
	"database/sql"

	"github.com/google/uuid"

	// Postgres driver
	_ "github.com/lib/pq"

	"later.co/pkg/later/entity"
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
func (repository *ContentRepository) Insert(content *entity.Content) (*entity.Content, error) {

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
func (repository *ContentRepository) ByID(id uuid.UUID) (*entity.Content, error) {
	var content entity.Content

	statement := `
	SELECT * FROM content 
	WHERE id = $1
	`

	row := repository.DB.QueryRow(statement, id)

	err := content.ScanRow(row)

	return &content, err
}

// All returns all content
func (repository *ContentRepository) All(limit int) ([]entity.Content, error) {
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

func (repository *ContentRepository) scanRows(rows *sql.Rows) ([]entity.Content, error) {
	contents := []entity.Content{}

	defer rows.Close()

	for rows.Next() {
		var content entity.Content
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
