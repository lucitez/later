package usercontentrepo

import (
	"database/sql"
	"fmt"

	// Postgres driver
	"github.com/google/uuid"
	"later.co/pkg/later/usercontent"
	"later.co/pkg/repository"
)

// DB is this repository's database connection
var DB *sql.DB

// Insert inserts a new userContent
func Insert(userContent *usercontent.UserContent) (*usercontent.UserContent, error) {

	statement := `
	INSERT INTO user_content (id, share_id, content_id, user_id, sent_by)
	VALUES (
		$1,
		$2,
		$3,
		$4,
		$5
	)
	`

	_, err := DB.Exec(
		statement,
		userContent.ID,
		userContent.ShareID,
		userContent.ContentID,
		userContent.UserID,
		userContent.SentBy)

	if err != nil {
		return nil, err
	}

	return userContent, nil
}

func scanRowIntoUserContent(userContent *usercontent.UserContent, row *sql.Row) error {
	err := row.Scan(
		&userContent.ID,
		&userContent.ShareID,
		&userContent.ContentID,
		&userContent.ContentType,
		&userContent.UserID,
		&userContent.SentBy,
		&userContent.CreatedAt,
		&userContent.UpdatedAt,
		&userContent.ArchivedAt,
		&userContent.DeletedAt)

	return err
}

func scanRowsIntoUserContent(userContent *usercontent.UserContent, rows *sql.Rows) error {
	err := rows.Scan(
		&userContent.ID,
		&userContent.ShareID,
		&userContent.ContentID,
		&userContent.ContentType,
		&userContent.UserID,
		&userContent.SentBy,
		&userContent.CreatedAt,
		&userContent.UpdatedAt,
		&userContent.ArchivedAt,
		&userContent.DeletedAt)

	return err
}

// ByID gets a userContent by id
func ByID(id uuid.UUID) (*usercontent.UserContent, error) {
	var userContent usercontent.UserContent

	statement := `
	SELECT * FROM user_content 
	WHERE id = $1
	`

	row := DB.QueryRow(statement, id)

	err := scanRowIntoUserContent(&userContent, row)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &userContent, nil
}

// All returns all userContents
func All(limit int) ([]usercontent.UserContent, error) {
	userContents := []usercontent.UserContent{}

	rows, err := DB.Query(`SELECT * FROM user_content LIMIT $1`, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var userContent usercontent.UserContent
		err := scanRowsIntoUserContent(&userContent, rows)

		if err != nil {
			return nil, err
		}
		userContents = append(userContents, userContent)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return userContents, nil
}

// Feed gets usercontent
func Feed(
	userID uuid.UUID,
	senderType *string,
	contentType *string,
	archived *bool) ([]usercontent.UserContent, error) {

	userContents := []usercontent.UserContent{}

	statement := `
	SELECT * FROM user_content
	WHERE user_id = $1
	AND deleted_at IS NULL
	`
	whens := []string{}
	arguments := []string{userID.String()}
	orders := []string{}

	if senderType != nil {
		whens = append(whens, `sender_type`)
		arguments = append(arguments, *senderType)
	}

	if contentType != nil {
		whens = append(whens, `content_type`)
		arguments = append(arguments, *contentType)
	}

	if archived != nil && *archived == true {
		statement = statement + `AND archived_at IS NOT NULL`
		orders = append(orders, `archived_at`)
	}

	statement = statement + repository.ConstructWhens(whens, 2)

	rows, err := repository.QueryRowsWithArguments(DB, statement, arguments)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var userContent usercontent.UserContent
		err := scanRowsIntoUserContent(&userContent, rows)

		if err != nil {
			return nil, err
		}
		userContents = append(userContents, userContent)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return userContents, nil
}
