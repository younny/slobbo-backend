package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/younny/slobbo-backend/src/types"
)

var (
	testUser = types.User{
		Username: "Foo",
		Email:    "abc@abc.com",
		Password: "1234",
	}

	secondUser = types.User{
		Username: "Boo",
		Email:    "def@def.com",
		Password: "1234",
	}
)

func TestClient_User(t *testing.T) {
	testClient.Client.DropTable(&types.User{})
	testClient.AutoMigrate(&types.User{})
	first := testUser
	err := testClient.CreateUser(&first)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), first.ID)

	second := secondUser
	err = testClient.CreateUser(&second)
	assert.NoError(t, err)
	assert.Equal(t, uint(2), second.ID)

	users := testClient.GetUsers()
	assert.Equal(t, len(users.Items), 2, "")

	update := first
	update.Username = "Tester"
	err = testClient.UpdateUser(&update)
	assert.NoError(t, err)

	got := testClient.GetUserByID(1)
	assert.Equal(t, testUser.Email, got.Email, "")
	assert.Equal(t, update.Username, got.Username, "")

	got = testClient.GetUserByEmail("abc@abc.com")
	assert.Equal(t, uint(1), got.ID, "")

	got = testClient.GetUserByEmail("abcdefg@abc.com")
	assert.Empty(t, got)

	err = testClient.DeleteUser(1)
	assert.NoError(t, err)

	got = testClient.GetUserByID(1)
	assert.Empty(t, got)

	err = testClient.DeleteUser(3)
	assert.Error(t, err)
}
