package repository_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"

	"github.com/lucitez/later/pkg/inits"
	"github.com/lucitez/later/pkg/repository"
	"github.com/lucitez/later/pkg/repository/util"

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
	"hostnames",
}

var testUtil util.RepositoryTestUtil

/**
 * Declare some default vars to be reused across tests
 */
var userID, _ = uuid.NewRandom()
var userID2, _ = uuid.NewRandom()

func TestMain(m *testing.M) {
	godotenv.Load("../../.env.local")
	db := inits.TestDB()

	defer testUtil.DB.Close()
	defer afterAll()

	contentRepo = repository.NewContent(db)
	hostnameRepo = repository.NewHostname(db)
	friendRepo = repository.NewFriend(db)
	friendRequestRepo = repository.NewFriendRequest(db)
	shareRepo = repository.NewShare(db)
	testUtil = util.RepositoryTestUtil{DB: db}
	userRepo = repository.NewUser(db)
	userContentRepo = repository.NewUserContent(db)
	chatRepo = repository.NewChat(db)
	messageRepo = repository.NewMessage(db)
	userMessageRepo = repository.NewUserMessage(db)

	os.Exit(m.Run())
}

func beforeEach(t *testing.T) {
	testUtil.TruncateTables(tableNames)
	testUtil.Assert = assert.New(t)
}

func afterAll() {
	testUtil.TruncateTables(tableNames)
}
