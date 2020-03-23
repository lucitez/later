package repository

import (
	"database/sql"

	"github.com/google/uuid"

	"later.co/pkg/later/entity"
)

// FriendRequestRepository ...
type FriendRequestRepository interface {
	Insert(friendRequest *entity.FriendRequest) (*entity.FriendRequest, error)
	PendingByUserID(userID uuid.UUID) ([]entity.FriendRequest, error)
	Accept(ID uuid.UUID) error
	Decline(ID uuid.UUID) error
}

type FriendRequestRepositoryImpl struct {
	DB *sql.DB
}

// Insert inserts a new friend
func (repository *FriendRequestRepositoryImpl) Insert(friendRequest *entity.FriendRequest) (*entity.FriendRequest, error) {

	statement := `
	INSERT INTO friend_requests (id, sent_by_user_id, recipient_user_id)
	VALUES (
		$1,
		$2,
		$3
	);
	`

	_, err := repository.DB.Exec(
		statement,
		friendRequest.ID,
		friendRequest.SentByUserID,
		friendRequest.RecipientUserID)

	if err != nil {
		return nil, err
	}

	return friendRequest, nil
}

// PendingByUserID gets all pending friend requests for a user
func (repository *FriendRequestRepositoryImpl) PendingByUserID(userID uuid.UUID) ([]entity.FriendRequest, error) {
	statement := `
	SELECT * FROM friend_requests 
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
func (repository *FriendRequestRepositoryImpl) Accept(ID uuid.UUID) error {
	statement := `
	UPDATE friend_requests
	SET accepted_at = now()
	WHERE id = $1;
	`

	_, err := repository.DB.Exec(statement, ID)

	return err
}

// Decline updates accepted_at
func (repository *FriendRequestRepositoryImpl) Decline(ID uuid.UUID) error {
	statement := `
	UPDATE friend_requests
	SET declined_at = now()
	WHERE id = $1;
	`

	_, err := repository.DB.Exec(statement, ID)

	return err
}

func (repository *FriendRequestRepositoryImpl) scanRows(rows *sql.Rows) ([]entity.FriendRequest, error) {
	friendRequests := []entity.FriendRequest{}

	defer rows.Close()

	for rows.Next() {
		var friendRequest entity.FriendRequest
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
