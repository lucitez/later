package repository_test

import (
	"fmt"
	"later/pkg/repository"
	"later/pkg/repository/util"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var tableNames = []string{
	"users",
	"user_content",
	"shares",
	"friends",
	"friend_requests",
	"content",
	"domains",
}

var testUtil util.RepositoryTestUtil

/**
 * Declare some default vars to be reused across tests
 */
var userID, _ = uuid.NewRandom()
var userID2, _ = uuid.NewRandom()

func TestMain(m *testing.M) {
	db, err := util.InitTestDB()

	if err != nil {
		fmt.Println(err)
		panic("Error creating test db connection")
	}

	defer testUtil.DB.Close()
	defer afterAll()

	contentRepo = repository.NewContent(db)
	domainRepo = repository.NewDomain(db)
	friendRepo = repository.NewFriend(db)
	friendRequestRepo = repository.NewFriendRequest(db)
	shareRepo = repository.NewShare(db)
	testUtil = util.RepositoryTestUtil{DB: db}
	userRepo = repository.NewUser(db)
	userContentRepo = repository.NewUserContent(db)

	os.Exit(m.Run())
}

func beforeEach(t *testing.T) {
	testUtil.TruncateTables(tableNames)
	testUtil.Assert = assert.New(t)
}

func afterAll() {
	testUtil.TruncateTables(tableNames)
}
