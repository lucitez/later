package repository

import (
	"database/sql"
	"later/pkg/service/body"
	"log"
	"strconv"

	// Postgres driver
	"later/pkg/model"
	"later/pkg/repository/util"
	"later/pkg/util/wrappers"

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
func (repository *UserContent) Insert(userContent model.UserContent) error {
	statement := util.GenerateInsertStatement(userContent, "user_content")

	_, err := repository.DB.Exec(
		statement,
		userContent.ID,
		userContent.ShareID,
		userContent.ContentID,
		userContent.UserID,
		userContent.SentByUserID,
		userContent.Tag,
		userContent.CreatedAt,
		userContent.UpdatedAt,
		userContent.SavedAt,
		userContent.DeletedAt,
	)

	return err
}

// ByID gets a userContent by id
func (repository *UserContent) ByID(id uuid.UUID) *model.UserContent {
	var userContent model.UserContent

	statement := selectUserContent + `
	WHERE id = $1;
	`

	row := repository.DB.QueryRow(statement, id)

	return userContent.ScanRow(row)
}

// All returns all userContents
func (repository *UserContent) All(limit int) []model.UserContent {
	statement := selectUserContent + `
	WHERE deleted_at IS NULL
	LIMIT $1;
	`
	rows, err := repository.DB.Query(statement, limit)

	if err != nil {
		log.Fatal(err)
	}

	return repository.scanRows(rows)
}

// Save a post, optionally update the tag
func (repository *UserContent) Save(
	id uuid.UUID,
	tag wrappers.NullString,
) error {
	statement := `
	UPDATE user_content
	SET saved_at = now(),
		tag = $1
	WHERE id = $2;
	`

	_, err := repository.DB.Exec(
		statement,
		tag,
		id,
	)

	return err
}

// Delete a post, optionally update the tag
func (repository *UserContent) Delete(id uuid.UUID) error {
	statement := `
	UPDATE user_content
	SET deleted_at = now()
	WHERE id = $1;
	`

	_, err := repository.DB.Exec(
		statement,
		id,
	)

	return err
}

// Update user_content
func (repository *UserContent) Update(body body.UserContentUpdateBody) {
	statement := `
	UPDATE user_content
	SET tag = $1
	WHERE id = $2;
	`
	_, err := repository.DB.Exec(
		statement,
		body.Tag,
		body.ID,
	)

	if err != nil {
		panic(err)
	}
}

// Filter gets usercontent
func (repository *UserContent) Filter(
	userID uuid.UUID,
	tag *string,
	contentType *string,
	saved bool,
	search *string,
	limit int,
) []model.UserContent {
	statement := selectUserContent

	statement += `
		JOIN content c ON user_content.content_id = c.id
		WHERE user_content.user_id = $1
		AND user_content.deleted_at IS NULL`

	counter := 2
	var fuzzySearch *string = nil

	if search != nil {
		statement += `
		AND (
			c.title ILIKE $2
			OR c.domain ILIKE $2
			OR user_content.tag ILIKE $2
		)
		`

		counter++
		fuzzySearch = search
		*fuzzySearch = "%" + *fuzzySearch + "%"
	}

	if tag != nil {
		statement += `
		AND tag = $` + strconv.Itoa(counter)
		counter++
	}

	if contentType != nil {
		statement += `
		AND c.content_type = $` + strconv.Itoa(counter)
		counter++
	}

	if saved {
		statement += `
		AND saved_at IS NOT NULL`
	} else {
		statement += `
		AND saved_at IS NULL`
	}

	statement += `
	ORDER BY user_content.created_at DESC
	LIMIT $` + strconv.Itoa(counter) + `;`

	args := util.GenerateArguments(
		userID,
		fuzzySearch,
		tag,
		contentType,
		limit,
	)

	rows, err := repository.DB.Query(statement, args...)

	if err != nil {
		panic(err)
	}

	return repository.scanRows(rows)
}

func (repository *UserContent) scanRows(rows *sql.Rows) []model.UserContent {
	userContents := []model.UserContent{}

	defer rows.Close()

	for rows.Next() {
		var userContent model.UserContent
		userContent.ScanRows(rows)

		userContents = append(userContents, userContent)
	}

	err := rows.Err()
	if err != nil {
		panic(err)
	}

	return userContents
}
