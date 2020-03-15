package friendrequestrepo

import (
	"database/sql"

	"github.com/google/uuid"

	"later.co/pkg/later/entity"
)

// DB is this repository's database connection
var DB *sql.DB

// Insert inserts a new friend
func Insert(friendRequest *entity.FriendRequest) (*entity.FriendRequest, error) {

	statement := `
	INSERT INTO friend_requests (id, sent_by_user_id, recipient_user_id)
	VALUES (
		$1,
		$2,
		$3
	);
	`

	_, err := DB.Exec(
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
func PendingByUserID(userID uuid.UUID) ([]entity.FriendRequest, error) {
	statement := `
	SELECT * FROM friend_requests 
	WHERE recipient_user_id = $1
	AND accepted_at IS NULL
	AND declined_at IS NULL
	AND deleted_at IS NULL;
	`

	rows, err := DB.Query(statement, userID)

	if err != nil {
		return nil, err
	}

	friendRequests, err := scanRows(rows)

	if err != nil {
		return nil, err
	}

	return friendRequests, nil
}

// Accept updates accepted_at
func Accept(ID uuid.UUID) error {
	statement := `
	UPDATE friend_requests
	SET accepted_at = now()
	WHERE id = $1;
	`

	_, err := DB.Exec(statement, ID)

	return err
}

// Decline updates accepted_at
func Decline(ID uuid.UUID) error {
	statement := `
	UPDATE friend_requests
	SET declined_at = now()
	WHERE id = $1;
	`

	_, err := DB.Exec(statement, ID)

	return err
}

func scanRows(rows *sql.Rows) ([]entity.FriendRequest, error) {
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
