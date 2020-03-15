package contentrepo

import (
	"database/sql"

	"github.com/google/uuid"

	// Postgres driver
	_ "github.com/lib/pq"

	"later.co/pkg/later/entity"
)

// DB is this repository's database connection
var DB *sql.DB

// Insert inserts new content
func Insert(content *entity.Content) (*entity.Content, error) {

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

	_, err := DB.Exec(
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
func ByID(id uuid.UUID) (*entity.Content, error) {
	var content entity.Content

	statement := `
	SELECT * FROM content 
	WHERE id = $1
	`

	row := DB.QueryRow(statement, id)

	err := row.Scan(
		&content.ID,
		&content.Title,
		&content.Description,
		&content.ImageURL,
		&content.ContentType,
		&content.URL,
		&content.Domain,
		&content.Shares,
		&content.CreatedAt,
		&content.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &content, nil
}

// All returns all content
func All(limit int) ([]entity.Content, error) {
	allContent := []entity.Content{}

	rows, err := DB.Query(`SELECT * FROM content LIMIT $1`, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var content entity.Content
		err := rows.Scan(
			&content.ID,
			&content.Title,
			&content.Description,
			&content.ImageURL,
			&content.ContentType,
			&content.URL,
			&content.Domain,
			&content.Shares,
			&content.CreatedAt,
			&content.UpdatedAt)

		if err != nil {
			return nil, err
		}
		allContent = append(allContent, content)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return allContent, nil
}
