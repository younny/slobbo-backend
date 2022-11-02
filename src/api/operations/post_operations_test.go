package operations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/younny/slobbo-backend/src/api/mocks"
	"github.com/younny/slobbo-backend/src/types"
)

var (
	testPost1 = types.Post{
		ID:        0,
		Title:     "H",
		SubTitle:  "S",
		Body:      "B",
		Author:    "Koo",
		Category:  1,
		Thumbnail: "w",
		CreatedAt: randomTime,
		UpdatedAt: randomTime,
	}
	testPost2 = types.Post{
		ID:        1,
		Title:     "H",
		SubTitle:  "S",
		Body:      "B",
		Author:    "Koo",
		Category:  0,
		Thumbnail: "w",
		CreatedAt: randomTime,
		UpdatedAt: randomTime,
	}

	newPost = types.Post{
		Title:     "New Post",
		SubTitle:  "New Post Sub",
		Body:      "Boo",
		Author:    "Koo",
		Category:  0,
		Thumbnail: "www",
	}

	newPostResponse = types.Post{
		ID:        3,
		Title:     "New Post",
		SubTitle:  "New Post Sub",
		Body:      "Boo",
		Author:    "Koo",
		Category:  0,
		Thumbnail: "www",
		CreatedAt: createdTime,
		UpdatedAt: createdTime,
	}

	updatePost = types.Post{
		ID:        0,
		Title:     "This title is updated",
		SubTitle:  "S",
		Body:      "B",
		Author:    "Koo",
		Category:  1,
		Thumbnail: "w",
	}

	updatePostResponse = types.Post{
		ID:        0,
		Title:     "This title is updated",
		SubTitle:  "S",
		Body:      "B",
		Author:    "Koo",
		Category:  1,
		Thumbnail: "w",
		CreatedAt: createdTime,
		UpdatedAt: createdTime,
	}

	emptyTitlePost = types.Post{
		Title:     "",
		SubTitle:  "S",
		Body:      "B",
		Author:    "Koo",
		Category:  1,
		Thumbnail: "w",
	}

	testPost1Json, _          = json.Marshal(testPost1)
	testPost2Json, _          = json.Marshal(testPost2)
	newPostInJson, _          = json.Marshal(newPost)
	newPostResponseJson, _    = json.Marshal(newPostResponse)
	updatePostJson, _         = json.Marshal(updatePost)
	updatePostResponseJson, _ = json.Marshal(updatePostResponse)
	emptyTitlePostJson, _     = json.Marshal(emptyTitlePost)
)

func TestPostEndpoints(t *testing.T) {
	s := Server{}
	s.Set(getPostDBClientMock(t))

	ts := httptest.NewServer(s.Router)
	defer ts.Close()

	tokenStr := Authenticate(s)

	testcasesInOrder := []string{
		"GET /posts",
		"GET /posts?page_id=1",
		"GET /posts/{id} 200",
		"POST /posts 200",
		"PATCH /posts/{id}",
		"DELETE /posts",
		"DELETE /post 404",
		"GET /posts/{id} 404",
		"POST /posts 400",
	}
	testcases := map[string]TestCase{
		"GET /posts": {
			method:   http.MethodGet,
			path:     "/posts",
			wantCode: http.StatusOK,
			wantBody: fmt.Sprintf(`{"items":[%s,%s]}`, testPost1Json, testPost2Json),
		},
		"GET /posts?page_id=1": {
			method:   http.MethodGet,
			path:     "/posts?page_id=1",
			wantCode: http.StatusOK,
			wantBody: fmt.Sprintf(`{"items":[%s]}`, testPost2Json),
		},
		"GET /posts/{id} 200": {
			method:   http.MethodGet,
			path:     "/posts/0",
			wantCode: http.StatusOK,
			wantBody: fmt.Sprintf(`%s`, testPost1Json),
		},
		"POST /posts 200": {
			method: http.MethodPost,
			path:   "/posts",
			body:   fmt.Sprintf(`%s`, newPostInJson),
			header: http.Header{
				"Content-type":  {"application/json"},
				"Authorization": {tokenStr},
			},
			wantCode: http.StatusOK,
			wantBody: fmt.Sprintf(`%s`, newPostResponseJson),
		},
		"PATCH /posts/{id}": {
			method: http.MethodPatch,
			path:   "/posts/0",
			body:   fmt.Sprintf(`%s`, updatePostJson),
			header: http.Header{
				"Content-type":  {"application/json"},
				"Authorization": {tokenStr},
			},
			wantCode: http.StatusOK,
			wantBody: fmt.Sprintf(`%s`, updatePostResponseJson),
		},
		"DELETE /posts": {
			method: http.MethodDelete,
			path:   "/posts/1",
			header: http.Header{
				"Content-type":  {"application/json"},
				"Authorization": {tokenStr},
			},
			wantCode: http.StatusOK,
		},
		"DELETE /post 404": {
			method: http.MethodDelete,
			path:   "/posts/3",
			header: http.Header{
				"Content-type":  {"application/json"},
				"Authorization": {tokenStr},
			},
			wantCode: http.StatusNotFound,
		},
		"GET /posts/{id} 404": {
			method:   http.MethodGet,
			path:     "/posts/3",
			wantCode: http.StatusNotFound,
		},
		"POST /posts 400": {
			method: http.MethodPost,
			path:   "/posts",
			body:   fmt.Sprintf(`%s`, emptyTitlePostJson),
			header: http.Header{
				"Content-type":  {"application/json"},
				"Authorization": {tokenStr},
			},
			wantCode: http.StatusBadRequest,
		},
	}

	for _, name := range testcasesInOrder {
		test := testcases[name]
		t.Run(name, func(t *testing.T) {
			body := bytes.NewReader([]byte(test.body))
			gotResponse, gotBody := RequestHandler(t, ts, test.method, test.path, body, test.header)
			assert.Equal(t, test.wantCode, gotResponse.StatusCode)
			if test.wantBody != "" {
				assert.Equal(t, test.wantBody, gotBody, "Body did not match")
			}
		})
	}
}

func getPostDBClientMock(t *testing.T) *mocks.MockClientInterface {
	ctrl := gomock.NewController(t)
	dbClient := mocks.NewMockClientInterface(ctrl)

	dbClient.EXPECT().GetPosts(gomock.Eq(0)).Return(&types.PostList{
		Items: []*types.Post{
			&testPost1,
			&testPost2,
		},
	})
	dbClient.EXPECT().GetPosts(gomock.Eq(1)).Return(&types.PostList{
		Items: []*types.Post{
			&testPost2,
		},
	})

	dbClient.EXPECT().GetPostByID(gomock.Eq(uint(0))).Return(&testPost1).AnyTimes()

	dbClient.EXPECT().GetPostByID(gomock.Eq(uint(1))).Return(&testPost2).AnyTimes()

	dbClient.EXPECT().GetPostByID(gomock.Eq(uint(3))).Return(nil).AnyTimes()

	dbClient.EXPECT().CreatePost(gomock.Any()).DoAndReturn(func(post *types.Post) error {
		post.ID = 3
		return nil
	}).AnyTimes()

	dbClient.EXPECT().GetUserByEmail(gomock.Any()).Return(&testUser).Times(1)

	dbClient.EXPECT().UpdatePost(gomock.Any()).AnyTimes()

	dbClient.EXPECT().DeletePost(gomock.Any()).AnyTimes()

	return dbClient
}
