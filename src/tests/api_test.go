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

	"github.com/younny/slobbo-backend/src/api/mocks"
	"github.com/younny/slobbo-backend/src/api/operations"
	"github.com/younny/slobbo-backend/src/types"
)

var randomTime = time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC)
var randomDuration = time.Duration(20)

var (
	testContacts = types.Contacts{
		Email:  "abc@example.com",
		Github: "github.com/younny",
	}
	testAbout = types.About{
		Title:    "About me",
		SubTitle: "About me sub",
		Body1:    "This is body 1",
		Body2:    "This is body 2",
		Contacts: &testContacts,
	}
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

	testWorkshop1 = types.Workshop{
		ID:        0,
		Name:      "Workshop 1",
		Details:   "ABC",
		Organiser: "f",
		Location:  "f",
		Datetime:  randomTime,
		Duration:  randomDuration,
		Capacity:  2,
		Price:     "free",
		CreatedAt: randomTime,
		UpdatedAt: randomTime,
	}
	testWorkshop2 = types.Workshop{
		ID:        1,
		Name:      "Workshop 2",
		Details:   "ABC",
		Organiser: "f",
		Location:  "f",
		Datetime:  randomTime,
		Duration:  randomDuration,
		Capacity:  2,
		Price:     "free",
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
	server := operations.Server{}
	server.Set(nil)
	r := server.Router
	testcases := map[string]struct {
		method string
		path   string
	}{
		"GET /about": {
			method: http.MethodGet,
			path:   "/about",
		},
		"PATCH /about": {
			method: http.MethodPatch,
			path:   "/about",
		},
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
		"GET /workshops": {
			method: http.MethodGet,
			path:   "/workshops",
		},
		"GET /workshops/{id}": {
			method: http.MethodGet,
			path:   "/workshops/id",
		},
		"POST /workshops": {
			method: http.MethodPost,
			path:   "/workshops",
		},
		"PATCH /workshops": {
			method: http.MethodPatch,
			path:   "/workshops",
		},
		"DELETE /workshops": {
			method: http.MethodDelete,
			path:   "/workshops",
		},
	}

	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			got := r.Match(chi.NewRouteContext(), test.method, test.path)
			assert.Equal(t, true, got, fmt.Sprintf("Not found: %s '%s'", test.method, test.path))
		})
	}
}

func TestEndpoints(t *testing.T) {
	s := operations.Server{}
	s.Set(getDBClientMock(t))

	ts := httptest.NewServer(s.Router)
	defer ts.Close()

	testcasesInOrder := []string{
		"GET /about",
		"PATCH /about",
		"GET /posts",
		"GET /posts?page_id=1",
		"GET /posts/{id} 200",
		"POST /posts 200",
		"PATCH /posts",
		"DELETE /posts",
		"GET /posts/{id} 404",
		"POST /posts 400",
		"GET /workshops",
		"GET /workshops?page_id=1",
		"GET /workshops/{id} 200",
		"POST /workshops 200",
		"PATCH /workshops",
		"DELETE /workshops",
		"GET /workshops/{id} 404",
		"POST /workshops 400",
	}
	testcases := map[string]TestCase{
		"GET /about": {
			method:   http.MethodGet,
			path:     "/about",
			wantCode: http.StatusOK,
			wantBody: `{"id":0,"title":"About me","sub_title":"About me sub","body_1":"This is body 1","body_2":"This is body 2","contacts":{"email":"abc@example.com","github":"github.com/younny"}}`,
		},
		"PATCH /about": {
			method: http.MethodPatch,
			path:   "/about",
			body:   `{"title":"About me !!!"}`,
			header: map[string][]string{
				"Content-type": {"application/json"},
			},
			wantCode: http.StatusOK,
			wantBody: `{"id":0,"title":"About me !!!","sub_title":"About me sub","body_1":"This is body 1","body_2":"This is body 2","contacts":{"email":"abc@example.com","github":"github.com/younny"}}`,
		},
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
		"POST /posts 200": {
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
		"POST /posts 400": {
			method: http.MethodPost,
			path:   "/posts",
			body:   `{"title":"","sub_title":"S","body":"B","author":"Koo","category":0,"thumbnail":"w"}`,
			header: map[string][]string{
				"Content-type": {"application/json"},
			},
			wantCode: http.StatusBadRequest,
		},
		"GET /workshops": {
			method:   http.MethodGet,
			path:     "/workshops",
			wantCode: http.StatusOK,
			wantBody: `{"items":[{"id":0,"name":"Workshop 1","details":"ABC","organiser":"f","location":"f","datetime":"1969-12-31T00:00:00Z","duration":20,"capacity":2,"price":"free","createdAt":"1969-12-31T00:00:00Z","updatedAt":"1969-12-31T00:00:00Z"},{"id":1,"name":"Workshop 2","details":"ABC","organiser":"f","location":"f","datetime":"1969-12-31T00:00:00Z","duration":20,"capacity":2,"price":"free","createdAt":"1969-12-31T00:00:00Z","updatedAt":"1969-12-31T00:00:00Z"}]}`,
		},
		"GET /workshops?page_id=1": {
			method:   http.MethodGet,
			path:     "/workshops?page_id=1",
			wantCode: http.StatusOK,
			wantBody: `{"items":[{"id":1,"name":"Workshop 2","details":"ABC","organiser":"f","location":"f","datetime":"1969-12-31T00:00:00Z","duration":20,"capacity":2,"price":"free","createdAt":"1969-12-31T00:00:00Z","updatedAt":"1969-12-31T00:00:00Z"}]}`,
		},
		"GET /workshops/{id} 200": {
			method:   http.MethodGet,
			path:     "/workshops/0",
			wantCode: http.StatusOK,
			wantBody: `{"id":0,"name":"Workshop 1","details":"ABC","organiser":"f","location":"f","datetime":"1969-12-31T00:00:00Z","duration":20,"capacity":2,"price":"free","createdAt":"1969-12-31T00:00:00Z","updatedAt":"1969-12-31T00:00:00Z"}`,
		},
		"POST /workshops 200": {
			method: http.MethodPost,
			path:   "/workshops",
			body:   `{"name":"Workshop 3","details":"ABC","organiser":"f","location":"f","datetime":"1969-12-31T00:00:00Z","duration":20,"capacity":2,"price":"free","createdAt":"1969-12-31T00:00:00Z","updatedAt":"1969-12-31T00:00:00Z"}`,
			header: map[string][]string{
				"Content-type": {"application/json"},
			},
			wantCode: http.StatusOK,
			wantBody: `{"id":3,"name":"Workshop 3","details":"ABC","organiser":"f","location":"f","datetime":"1969-12-31T00:00:00Z","duration":20,"capacity":2,"price":"free","createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"}`,
		},
		"PATCH /workshops": {
			method: http.MethodPatch,
			path:   "/workshops/0",
			body:   `{"name":"Hello World"}`,
			header: map[string][]string{
				"Content-type": {"application/json"},
			},
			wantCode: http.StatusOK,
			wantBody: `{"id":0,"name":"Hello World","details":"ABC","organiser":"f","location":"f","datetime":"1969-12-31T00:00:00Z","duration":20,"capacity":2,"price":"free","createdAt":"1969-12-31T00:00:00Z","updatedAt":"1969-12-31T00:00:00Z"}`,
		},
		"DELETE /workshops": {
			method: http.MethodDelete,
			path:   "/workshops/1",
			header: map[string][]string{
				"Content-type": {"application/json"},
			},
			wantCode: http.StatusOK,
		},
		"GET /workshops/{id} 404": {
			method:   http.MethodGet,
			path:     "/workshops/3",
			wantCode: http.StatusNotFound,
		},
		"POST /workshops 400": {
			method: http.MethodPost,
			path:   "/workshops",
			body:   `{"title":"","sub_title":"S","body":"B","author":"Koo","category":0,"thumbnail":"w"}`,
			header: map[string][]string{
				"Content-type": {"application/json"},
			},
			wantCode: http.StatusBadRequest,
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

func getDBClientMock(t *testing.T) *mocks.MockClientInterface {
	ctrl := gomock.NewController(t)
	dbClient := mocks.NewMockClientInterface(ctrl)

	dbClient.EXPECT().GetAbout().Return(&testAbout).Times(2)

	dbClient.EXPECT().UpdateAbout(gomock.Any()).AnyTimes()

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

	dbClient.EXPECT().GetWorkshops(gomock.Eq(0)).Return(&types.WorkshopList{
		Items: []*types.Workshop{
			&testWorkshop1,
			&testWorkshop2,
		},
	}).AnyTimes()
	dbClient.EXPECT().GetWorkshops(gomock.Eq(1)).Return(&types.WorkshopList{
		Items: []*types.Workshop{
			&testWorkshop2,
		},
	}).AnyTimes()

	dbClient.EXPECT().GetWorkshopByID(gomock.Eq(uint(0))).Return(&testWorkshop1).AnyTimes()

	dbClient.EXPECT().GetWorkshopByID(gomock.Eq(uint(1))).Return(&testWorkshop2).AnyTimes()

	dbClient.EXPECT().GetWorkshopByID(gomock.Eq(uint(3))).Return(nil).AnyTimes()

	dbClient.EXPECT().CreateWorkshop(gomock.Any()).DoAndReturn(func(workshop *types.Workshop) error {
		workshop.ID = 3
		return nil
	}).AnyTimes()

	dbClient.EXPECT().UpdateWorkshop(gomock.Any()).AnyTimes()

	dbClient.EXPECT().DeleteWorkshop(gomock.Eq(uint(1))).AnyTimes()

	return dbClient
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
