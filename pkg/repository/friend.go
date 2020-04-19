package repository

import (
	"database/sql"
	"log"
	"strconv"

	"later/pkg/model"
	"later/pkg/repository/util"

	"github.com/google/uuid"

	// Postgres driver
	_ "github.com/lib/pq"
)

// Friend ...
type Friend struct {
	DB *sql.DB
}

// NewFriend for wire generation
func NewFriend(db *sql.DB) Friend {
	return Friend{db}
}

var friendSelectStatement = util.GenerateSelectStatement(model.Friend{}, "friends")

// Insert inserts a new friend
func (repository *Friend) Insert(friend model.Friend) error {

	statement := util.GenerateInsertStatement(friend, "friends")

	_, err := repository.DB.Exec(
		statement,
		friend.ID,
		friend.UserID,
		friend.FriendUserID,
		friend.CreatedAt,
		friend.UpdatedAt,
		friend.DeletedAt,
	)

	return err
}

// ForUser filter a user's friends
func (repository *Friend) ForUser(
	userID uuid.UUID,
	search *string,
	limit int,
	offset int,
) []model.Friend {
	statement := friendSelectStatement
	counter := 2
	var fuzzySearch *string = nil

	if search != nil {
		statement = statement + `
		JOIN users ON users.id = friends.friend_user_id
		WHERE friends.user_id = $1
		AND (
			users.username ILIKE $2
			OR users.email ILIKE $2
			OR users.first_name ILIKE $2
			OR users.last_name ILIKE $2
		)
		AND friends.deleted_at IS NULL
		`
		counter++
		fuzzySearch = search
		*fuzzySearch = "%" + *fuzzySearch + "%"
	} else {
		statement = statement + `
		WHERE user_id = $1
		AND deleted_at IS NULL
		`
	}

	statement += `
	LIMIT $` + strconv.Itoa(counter)

	counter++

	statement += `
	OFFSET $` + strconv.Itoa(counter)

	args := util.GenerateArguments(
		userID,
		fuzzySearch,
		limit,
		offset,
	)

	rows, err := repository.DB.Query(statement, args...)

	if err != nil {
		log.Fatal(err)
	}

	return repository.scanRows(rows)
}

func (repository *Friend) ByUserIDAndFriendUserID(
	userID uuid.UUID,
	friendUserID uuid.UUID,
) *model.Friend {
	var friend model.Friend

	statement := friendSelectStatement + `
	WHERE user_id = $1
	AND friend_user_id = $2
	AND deleted_at IS NULL;`

	row := repository.DB.QueryRow(statement, userID, friendUserID)

	return friend.ScanRow(row)
}

func (repository *Friend) DeleteByUserIDs(
	userID1 uuid.UUID,
	userID2 uuid.UUID,
) {
	statement := `
	UPDATE friends
	SET deleted_at = now()
	WHERE (
		user_id = $1
		AND friend_user_id = $2
	)
	OR (
		user_id = $2
		AND friend_user_id = $1
	);
	`

	_, err := repository.DB.Exec(statement, userID1, userID2)

	if err != nil {
		log.Fatal(err)
	}
}

func (repository *Friend) scanRows(rows *sql.Rows) []model.Friend {
	friends := []model.Friend{}

	defer rows.Close()

	for rows.Next() {
		var friend model.Friend
		friend.ScanRows(rows)

		friends = append(friends, friend)
	}

	err := rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return friends
}
