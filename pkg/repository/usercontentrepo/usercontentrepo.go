package usercontentrepo

import (
	"database/sql"
	"strconv"

	// Postgres driver
	"github.com/google/uuid"
	"later.co/pkg/later/entity"
	"later.co/pkg/repository/util"
	"later.co/pkg/response"
)

// DB is this repository's database connection
var DB *sql.DB

// Insert inserts a new userContent
func Insert(userContent *entity.UserContent) (*entity.UserContent, error) {

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

func scanRowIntoUserContent(userContent *entity.UserContent, row *sql.Row) error {
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

func scanRowsIntoUserContent(userContent *entity.UserContent, rows *sql.Rows) error {
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
func ByID(id uuid.UUID) (*entity.UserContent, error) {
	var userContent entity.UserContent

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
func All(limit int) ([]entity.UserContent, error) {
	userContents := []entity.UserContent{}

	rows, err := DB.Query(`SELECT * FROM user_content LIMIT $1`, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var userContent entity.UserContent
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
	archived *bool) ([]response.WireUserContent, error) {

	userIDString := userID.String()

	args := []string{userIDString}

	testStatement := `
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
		testStatement += `AND user_content.sender_type = $` + strconv.Itoa(counter) + ` `
		args = append(args, *senderType)
		counter++
	}

	if contentType != nil {
		testStatement += `AND user_content.content_type = $` + strconv.Itoa(counter) + ` `
		args = append(args, *contentType)
		counter++
	}

	if archived != nil && *archived == true {
		testStatement += `
		AND user_content.archived_at IS NOT NULL
		`
	}

	switch {
	case archived != nil && *archived == true:
		testStatement += `ORDER BY user_content.archived_at DESC`
	default:
		testStatement += `ORDER BY user_content.created_at DESC`
	}

	userContents := []response.WireUserContent{}

	rows, err := DB.Query(testStatement, util.GenerateArguments(args)...)

	if err != nil {
		return nil, err
	}

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

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return userContents, nil
}
