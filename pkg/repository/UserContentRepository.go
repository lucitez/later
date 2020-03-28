package repository

import (
	"database/sql"
	"strconv"

	// Postgres driver
	"github.com/google/uuid"
	"later.co/pkg/later/entity"
	"later.co/pkg/repository/util"
	"later.co/pkg/response"
)

type UserContentRepository interface {
	Insert(userContent *entity.UserContent) (*entity.UserContent, error)
	ByID(id uuid.UUID) (*entity.UserContent, error)
	All(limit int) ([]entity.UserContent, error)
	Feed(
		userID uuid.UUID,
		senderType *string,
		contentType *string,
		archived *bool) ([]response.WireUserContent, error)
}

type UserContentRepositoryImpl struct {
	DB *sql.DB
}

// Insert inserts a new userContent
func (repository *UserContentRepositoryImpl) Insert(userContent *entity.UserContent) (*entity.UserContent, error) {

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

	_, err := repository.DB.Exec(
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

// ByID gets a userContent by id
func (repository *UserContentRepositoryImpl) ByID(id uuid.UUID) (*entity.UserContent, error) {
	var userContent entity.UserContent

	statement := `
	SELECT * FROM user_content 
	WHERE id = $1
	`

	row := repository.DB.QueryRow(statement, id)

	err := userContent.ScanRow(row)

	return &userContent, err
}

// All returns all userContents
func (repository *UserContentRepositoryImpl) All(limit int) ([]entity.UserContent, error) {
	rows, err := repository.DB.Query(`SELECT * FROM user_content LIMIT $1`, limit)

	if err != nil {
		return nil, err
	}

	return repository.scanRows(rows)
}

// Feed gets usercontent
func (repository *UserContentRepositoryImpl) Feed(
	userID uuid.UUID,
	senderType *string,
	contentType *string,
	archived *bool) ([]response.WireUserContent, error) {

	userIDString := userID.String()

	args := []string{userIDString}

	statement := `
		SELECT 	user_content.id,
				content.id as content_id,
				content.title,
				content.description,
				content.image_url,
				content.content_type,
				content.domain,
				user_content.sent_by,
				user_content.created_at
		FROM user_content
		JOIN content ON content.id = user_content.content_id
		WHERE user_content.user_id = $1
		`

	counter := 2

	if senderType != nil {
		statement += `AND user_content.sender_type = $` + strconv.Itoa(counter) + ` `
		args = append(args, *senderType)
		counter++
	}

	if contentType != nil {
		statement += `AND user_content.content_type = $` + strconv.Itoa(counter) + ` `
		args = append(args, *contentType)
		counter++
	}

	if archived != nil && *archived == true {
		statement += `
		AND user_content.archived_at IS NOT NULL
		`
	}

	switch {
	case archived != nil && *archived == true:
		statement += `ORDER BY user_content.archived_at DESC`
	default:
		statement += `ORDER BY user_content.created_at DESC`
	}

	rows, err := repository.DB.Query(statement, util.GenerateArguments(args)...)

	if err != nil {
		return nil, err
	}

	return repository.scanRowsIntoWireUserContent(rows)
}

func (repository *UserContentRepositoryImpl) scanRows(rows *sql.Rows) ([]entity.UserContent, error) {
	userContents := []entity.UserContent{}

	defer rows.Close()

	for rows.Next() {
		var userContent entity.UserContent
		err := userContent.ScanRows(rows)

		if err != nil {
			return nil, err
		}
		userContents = append(userContents, userContent)
	}

	err := rows.Err()
	if err != nil {
		return nil, err
	}

	return userContents, nil
}

func (repository *UserContentRepositoryImpl) scanRowsIntoWireUserContent(rows *sql.Rows) ([]response.WireUserContent, error) {
	userContents := []response.WireUserContent{}

	defer rows.Close()

	for rows.Next() {
		var userContent response.WireUserContent
		err := rows.Scan(
			&userContent.ID,
			&userContent.ContentID,
			&userContent.Title,
			&userContent.Description,
			&userContent.ImageURL,
			&userContent.ContentType,
			&userContent.Domain,
			&userContent.SentBy,
			&userContent.CreatedAt)

		if err != nil {
			return nil, err
		}
		userContents = append(userContents, userContent)
	}

	err := rows.Err()
	if err != nil {
		return nil, err
	}

	return userContents, nil
}
