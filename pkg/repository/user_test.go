package repository_test

import (
	"fmt"
	"os"
	"testing"

	// Postgres driver
	"github.com/google/uuid"
	_ "github.com/lib/pq"

	"later/pkg/model"
	"later/pkg/repository"
	"later/pkg/repository/util"
	"later/pkg/util/wrappers"
)

var testUtil util.RepositoryTestUtil
var repo repository.UserRepository
var tableNames = []string{"users"}

var user, err = model.NewUserFromSignUp(
	wrappers.NewNullStringFromString("test_username"),
	wrappers.NewNullStringFromString("test_email"),
	"1111111111")

func TestMain(m *testing.M) {
	db, err := util.InitTestDB()

	if err != nil {
		fmt.Println(err)
		panic("Error creating test db connection")
	}

	defer repo.DB.Close()
	defer afterAll()

	repo = repository.UserRepository{db}
	testUtil = util.RepositoryTestUtil{DB: db}

	os.Exit(m.Run())
}

func beforeEach() {
	testUtil.TruncateTables(tableNames)
}

func afterAll() {
	testUtil.TruncateTables(tableNames)
}

func TestInsertAndByID(t *testing.T) {
	beforeEach()

	repo.Insert(user)

	actual, err := repo.ByID(user.ID)

	if err != nil {
		t.Error(err)
	}

	util.AssertEquals(t, actual, user)
}

func TestByIDs(t *testing.T) {
	beforeEach()

	user2, _ := model.NewUserFromSignUp(
		wrappers.NewNullStringFromString("test_username_2"),
		wrappers.NewNullStringFromString("test_email_2"),
		"0000000000")

	repo.Insert(user)
	repo.Insert(user2)

	actual, err := repo.ByIDs([]uuid.UUID{user.ID, user2.ID})

	if err != nil {
		t.Error(err)
	}

	util.AssertEquals(t, actual, []model.User{*user, *user2})
}

func TestByPhoneNumber(t *testing.T) {
	beforeEach()

	repo.Insert(user)

	actual, err := repo.ByPhoneNumber(user.PhoneNumber)

	if err != nil {
		t.Error(err)
	}

	util.AssertEquals(t, actual, user)
}

func TestAll(t *testing.T) {
	beforeEach()

	user2, _ := model.NewUserFromSignUp(
		wrappers.NewNullStringFromString("test_username_2"),
		wrappers.NewNullStringFromString("test_email_2"),
		"0000000000")

	repo.Insert(user)
	repo.Insert(user2)

	actual, err := repo.All(1)

	if err != nil {
		t.Error(err)
	}

	util.AssertEquals(t, actual, []model.User{*user2})
}
