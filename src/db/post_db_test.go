package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/younny/slobbo-backend/src/types"
)

var (
	testPost = types.Post{
		Title:    "Hello World",
		SubTitle: "hello world",
		Body:     "Blah Blah",
		Author:   "Boo",
		Category: 0,
	}
)

func TestClient_Post(t *testing.T) {
	testClient.Client.DropTable(&types.Post{})
	testClient.AutoMigrate(&types.Post{})
	first := testPost
	err := testClient.CreatePost(&first)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), first.ID)

	second := testPost
	err = testClient.CreatePost(&second)
	assert.NoError(t, err)
	assert.Equal(t, uint(2), second.ID)

	update := first
	update.Title = "Hello Germany"
	err = testClient.UpdatePost(&update)
	assert.NoError(t, err)

	got := testClient.GetPostByID(1)
	assert.Equal(t, testPost.SubTitle, got.SubTitle, "")
	assert.Equal(t, update.Title, got.Title, "")

	testClient.DeletePost(1)
	got = testClient.GetPostByID(1)
	assert.Empty(t, got)

	err = testClient.DeletePost(3)
	assert.Error(t, err)
}

func TestClient_PaginatePosts(t *testing.T) {
	testClient.DropTable(&types.Post{})
	testClient.AutoMigrate(&types.Post{})

	for i := 0; i < pageSize+2; i++ {
		post := testPost
		_ = testClient.CreatePost(&post)
	}
	got := testClient.GetPosts(0)
	assert.Equal(t, 10, len(got.Items))
	assert.Equal(t, uint(11), got.NextPageID)

	got = testClient.GetPosts(11)
	assert.Equal(t, 2, len(got.Items))
	assert.Equal(t, uint(0), got.NextPageID)
}
