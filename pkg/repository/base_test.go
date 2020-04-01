package repository_test

import (
	"fmt"
	"later/pkg/repository"
	"later/pkg/repository/util"
	"os"
	"testing"

	"github.com/google/uuid"
)

var tableNames = []string{
	"users",
	"user_content",
	"shares"}

var testUtil util.RepositoryTestUtil

/**
 * Declare some default vars to be reused across tests
 */
var shareID, _ = uuid.NewRandom()
var contentID, _ = uuid.NewRandom()
var userID, _ = uuid.NewRandom()

func TestMain(m *testing.M) {
	db, err := util.InitTestDB()

	if err != nil {
		fmt.Println(err)
		panic("Error creating test db connection")
	}

	defer testUtil.DB.Close()
	defer afterAll()

	userRepository = repository.NewUser(db)
	userContentRepository = repository.NewUserContent(db)
	shareRepo = repository.NewShare(db)
	testUtil = util.RepositoryTestUtil{DB: db}

	os.Exit(m.Run())
}

func beforeEach() {
	testUtil.TruncateTables(tableNames)
}

func afterAll() {
	testUtil.TruncateTables(tableNames)
}
