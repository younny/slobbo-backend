package operations

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"

	"github.com/younny/slobbo-backend/src/types"
)

var (
	randomTime       = time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC)
	createdTime, err = time.Parse(time.RFC3339, "0001-01-01T00:00:00Z")
	testUser         = types.User{
		ID:        1,
		Username:  "yuna",
		Email:     "abc@slobbo.com",
		Password:  "1234",
		Role:      "admin",
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
	server := Server{}
	server.Set(nil)
	r := server.Router
	testcases := map[string]struct {
		method string
		path   string
	}{
		"GET /users": {
			method: http.MethodGet,
			path:   "/users",
		},
		"GET /users/{id}": {
			method: http.MethodGet,
			path:   "/users/id",
		},
		"POST /users": {
			method: http.MethodPost,
			path:   "/users",
		},
		"PATCH /users": {
			method: http.MethodPatch,
			path:   "/users",
		},
		"DELETE /users": {
			method: http.MethodDelete,
			path:   "/users",
		},
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

func RequestHandler(t *testing.T, ts *httptest.Server, method, path string, body io.Reader, header http.Header) (*http.Response, string) {
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

func Authenticate(s Server) string {
	token, err := s.SignIn(testUser.Email, testUser.Password)
	if err != nil {
		log.Fatalf("cannot login: %v\n", err)
	}
	tokenStr := fmt.Sprintf("Bearer %v", token)

	return tokenStr
}
