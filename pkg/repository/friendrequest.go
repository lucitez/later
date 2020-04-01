package repository

import (
	"database/sql"

	"github.com/google/uuid"

	"later/pkg/model"
	"later/pkg/repository/util"
)

// FriendRequest ...
type FriendRequest struct {
	DB *sql.DB
}

// NewFriendRequest for wire generation
func NewFriendRequest(db *sql.DB) FriendRequest {
	return FriendRequest{db}
}

var friendRequestSelectStatement = util.GenerateSelectStatement(model.FriendRequest{}, "friend_requests")

// Insert inserts a new friend
func (repository *FriendRequest) Insert(friendRequest *model.FriendRequest) (*model.FriendRequest, error) {

	statement := util.GenerateInsertStatement(*friendRequest, "friend_requests")

	_, err := repository.DB.Exec(
		statement,
		friendRequest.ID,
		friendRequest.SentByUserID,
		friendRequest.RecipientUserID,
		friendRequest.CreatedAt,
		friendRequest.UpdatedAt,
		friendRequest.AcceptedAt,
		friendRequest.DeclinedAt,
		friendRequest.DeletedAt)

	if err != nil {
		return nil, err
	}

	return friendRequest, nil
}

// ByID gets a friend request by id
func (repository *FriendRequest) ByID(id uuid.UUID) (*model.FriendRequest, error) {
	var friendRequest model.FriendRequest

	statement := friendRequestSelectStatement + ` WHERE id = $1;`

	row := repository.DB.QueryRow(statement, id)

	err := friendRequest.ScanRow(row)

	return &friendRequest, err
}

// PendingByUserID gets all pending friend requests for a user
func (repository *FriendRequest) PendingByUserID(userID uuid.UUID) ([]model.FriendRequest, error) {
	statement := friendRequestSelectStatement + `
	WHERE recipient_user_id = $1
	AND accepted_at IS NULL
	AND declined_at IS NULL
	AND deleted_at IS NULL;
	`

	rows, err := repository.DB.Query(statement, userID)

	if err != nil {
		return nil, err
	}

	friendRequests, err := repository.scanRows(rows)

	if err != nil {
		return nil, err
	}

	return friendRequests, nil
}

// Accept updates accepted_at
func (repository *FriendRequest) Accept(ID uuid.UUID) error {
	statement := `
	UPDATE friend_requests
	SET accepted_at = now()
	WHERE id = $1;
	`

	_, err := repository.DB.Exec(statement, ID)

	return err
}

// Decline updates accepted_at
func (repository *FriendRequest) Decline(ID uuid.UUID) error {
	statement := `
	UPDATE friend_requests
	SET declined_at = now()
	WHERE id = $1;
	`

	_, err := repository.DB.Exec(statement, ID)

	return err
}

func (repository *FriendRequest) scanRows(rows *sql.Rows) ([]model.FriendRequest, error) {
	friendRequests := []model.FriendRequest{}

	defer rows.Close()

	for rows.Next() {
		var friendRequest model.FriendRequest
		err := friendRequest.ScanRows(rows)

		if err != nil {
			return nil, err
		}
		friendRequests = append(friendRequests, friendRequest)
	}

	err := rows.Err()
	if err != nil {
		return nil, err
	}

	return friendRequests, nil
}
