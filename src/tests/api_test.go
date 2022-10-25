package tests

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/younny/slobbo-backend/src/api"
	"github.com/younny/slobbo-backend/src/api/mocks"
	"github.com/younny/slobbo-backend/src/types"
)

var randomTime = time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC)

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
)

type TestCase struct {
	method   string
	path     string
	body     string
	header   http.Header
	wantCode int
	wantBody string
}

func TestGetRouter(t *testing.T) {
	log, _ := zap.NewProduction(zap.WithCaller(false))
	r := api.GetRouter(log, nil)

	testcases := map[string]struct {
		method string
		path   string
	}{
		"GET /posts": {
			method: http.MethodGet,
			path:   "/posts",
		},
		"GET /posts/{id}": {
			method: http.MethodGet,
			path:   "/posts/id",
		},
		"POST /posts": {
			method: http.MethodPost,
			path:   "/posts",
		},
		"PATCH /posts": {
			method: http.MethodPatch,
			path:   "/posts",
		},
		"DELETE /posts": {
			method: http.MethodDelete,
			path:   "/posts",
		},
	}

	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			got := r.Match(chi.NewRouteContext(), test.method, test.path)
			assert.Equal(t, true, got, fmt.Sprintf("Not found: %s '%s'", test.method, test.path))
		})
	}
}

func getDBClientMock(t *testing.T) *mocks.MockClientInterface {
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

	dbClient.EXPECT().UpdatePost(gomock.Any()).AnyTimes()

	dbClient.EXPECT().DeletePost(gomock.Eq(uint(1))).AnyTimes()

	return dbClient
}

func TestEndpoints(t *testing.T) {
	r := api.GetRouter(nil, getDBClientMock(t))
	ts := httptest.NewServer(r)
	defer ts.Close()

	testcasesInOrder := []string{
		"GET /posts",
		"GET /posts?page_id=1",
		"GET /posts/{id} 200",
		"POST /posts",
		"PATCH /posts",
		"DELETE /posts",
		"GET /posts/{id} 404",
	}
	testcases := map[string]TestCase{
		"GET /posts": {
			method:   http.MethodGet,
			path:     "/posts",
			wantCode: http.StatusOK,
			wantBody: `{"items":[{"id":0,"title":"H","sub_title":"S","body":"B","author":"Koo","category":1,"thumbnail":"w","createdAt":"1969-12-31T00:00:00Z","updatedAt":"1969-12-31T00:00:00Z"},{"id":1,"title":"H","sub_title":"S","body":"B","author":"Koo","category":0,"thumbnail":"w","createdAt":"1969-12-31T00:00:00Z","updatedAt":"1969-12-31T00:00:00Z"}]}`,
		},
		"GET /posts?page_id=1": {
			method:   http.MethodGet,
			path:     "/posts?page_id=1",
			wantCode: http.StatusOK,
			wantBody: `{"items":[{"id":1,"title":"H","sub_title":"S","body":"B","author":"Koo","category":0,"thumbnail":"w","createdAt":"1969-12-31T00:00:00Z","updatedAt":"1969-12-31T00:00:00Z"}]}`,
		},
		"GET /posts/{id} 200": {
			method:   http.MethodGet,
			path:     "/posts/0",
			wantCode: http.StatusOK,
			wantBody: `{"id":0,"title":"H","sub_title":"S","body":"B","author":"Koo","category":1,"thumbnail":"w","createdAt":"1969-12-31T00:00:00Z","updatedAt":"1969-12-31T00:00:00Z"}`,
		},
		"POST /posts": {
			method: http.MethodPost,
			path:   "/posts",
			body:   `{"title":"Hello","sub_title":"S","body":"B","author":"Koo","category":0,"thumbnail":"w"}`,
			header: map[string][]string{
				"Content-type": {"application/json"},
			},
			wantCode: http.StatusOK,
			wantBody: `{"id":3,"title":"Hello","sub_title":"S","body":"B","author":"Koo","category":0,"thumbnail":"w","createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"}`,
		},
		"PATCH /posts": {
			method: http.MethodPatch,
			path:   "/posts/0",
			body:   `{"sub_title":"Hello World"}`,
			header: map[string][]string{
				"Content-type": {"application/json"},
			},
			wantCode: http.StatusOK,
			wantBody: `{"id":0,"title":"H","sub_title":"Hello World","body":"B","author":"Koo","category":1,"thumbnail":"w","createdAt":"1969-12-31T00:00:00Z","updatedAt":"1969-12-31T00:00:00Z"}`,
		},
		"DELETE /posts": {
			method: http.MethodDelete,
			path:   "/posts/1",
			header: map[string][]string{
				"Content-type": {"application/json"},
			},
			wantCode: http.StatusOK,
		},
		"GET /posts/{id} 404": {
			method:   http.MethodGet,
			path:     "/posts/3",
			wantCode: http.StatusNotFound,
		},
	}

	for _, name := range testcasesInOrder {
		test := testcases[name]
		t.Run(name, func(t *testing.T) {
			body := bytes.NewReader([]byte(test.body))
			gotResponse, gotBody := testRequest(t, ts, test.method, test.path, body, test.header)
			assert.Equal(t, test.wantCode, gotResponse.StatusCode)
			if test.wantBody != "" {
				assert.Equal(t, test.wantBody, gotBody, "Body did not match")
			}
		})
	}
}

func testRequest(t *testing.T, ts *httptest.Server, method, path string, body io.Reader, header http.Header) (*http.Response, string) {
	t.Helper()
	req, err := http.NewRequest(method, ts.URL+path, body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}
	req.Header = header

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	respBody = bytes.TrimSpace(respBody)
	return resp, string(respBody)

}
