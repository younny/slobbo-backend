package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/younny/slobbo-backend/src/types"
)

var (
	testAbout = types.About{
		Title:    "This is about",
		SubTitle: "this is about",
		Body1:    "Blah !",
		Body2:    "Blah blah ",
		Contacts: &types.Contacts{
			Email:  "abc@abc.com",
			Github: "githubbio.com/blah",
		},
	}
)

func TestClient_About(t *testing.T) {
	testClient.Client.DropTable(&types.About{})
	testClient.AutoMigrate(&types.About{})

	got := testClient.GetAbout()
	assert.Empty(t, got)

	new := testAbout
	err := testClient.CreateAbout(&new)
	assert.NoError(t, err)

	got = testClient.GetAbout()
	assert.Equal(t, testAbout.Body1, got.Body1, "")
	assert.Equal(t, uint(1), got.ID, "")

	update := got
	update.Body1 = "This is updated"
	testClient.UpdateAbout(update)

	got = testClient.GetAbout()
	assert.Equal(t, update.Body1, got.Body1)

}
