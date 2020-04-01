package repository

import (
	"database/sql"
	"strconv"

	// Postgres driver
	"later/pkg/model"
	"later/pkg/repository/util"
	"later/pkg/response"

	"github.com/google/uuid"
)

// UserContent ...
type UserContent struct {
	DB *sql.DB
}

// NewUserContent ...
func NewUserContent(db *sql.DB) UserContent {
	return UserContent{db}
}

var selectUserContent = util.GenerateSelectStatement(model.UserContent{}, "user_content")

// Insert inserts a new userContent
func (repository *UserContent) Insert(userContent *model.UserContent) (*model.UserContent, error) {

	statement := util.GenerateInsertStatement(*userContent, "user_content")

	_, err := repository.DB.Exec(
		statement,
		userContent.ID,
		userContent.ShareID,
		userContent.ContentID,
		userContent.ContentType,
		userContent.UserID,
		userContent.SentByUserID,
		userContent.CreatedAt,
		userContent.UpdatedAt,
		userContent.ArchivedAt,
		userContent.DeletedAt)

	if err != nil {
		return nil, err
	}

	return userContent, nil
}

// ByID gets a userContent by id
func (repository *UserContent) ByID(id uuid.UUID) (*model.UserContent, error) {
	var userContent model.UserContent

	statement := selectUserContent + `
	WHERE id = $1;
	`

	row := repository.DB.QueryRow(statement, id)

	err := userContent.ScanRow(row)

	return &userContent, err
}

// All returns all userContents
func (repository *UserContent) All(limit int) ([]model.UserContent, error) {
	statement := selectUserContent + `
	LIMIT $1
	`
	rows, err := repository.DB.Query(statement, limit)

	if err != nil {
		return nil, err
	}

	return repository.scanRows(rows)
}

// Feed gets usercontent
// TODO: refactor me to remove join
func (repository *UserContent) Feed(
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

func (repository *UserContent) scanRows(rows *sql.Rows) ([]model.UserContent, error) {
	userContents := []model.UserContent{}

	defer rows.Close()

	for rows.Next() {
		var userContent model.UserContent
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

func (repository *UserContent) scanRowsIntoWireUserContent(rows *sql.Rows) ([]response.WireUserContent, error) {
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
