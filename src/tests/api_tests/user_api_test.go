package api_tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/younny/slobbo-backend/src/api/mocks"
	"github.com/younny/slobbo-backend/src/api/operations"
	"github.com/younny/slobbo-backend/src/types"
)

var (
	testUser = types.User{
		ID:        1,
		Username:  "Oskar",
		Email:     "abc@example.com",
		Password:  "Slobbodibo",
		CreatedAt: randomTime,
		UpdatedAt: randomTime,
	}

	newUserRequest = types.User{
		Username: "Tester 2",
		Email:    "tester2@test.com",
		Password: "abcdefg",
	}

	newUserResponse = types.User{
		ID:        2,
		Username:  "Tester 2",
		Email:     "tester2@test.com",
		Password:  "abcdefg",
		CreatedAt: createdTime,
		UpdatedAt: createdTime,
	}

	newUserRequest2 = types.User{
		Email:    "tester2@test.com",
		Password: "abcdefg",
	}

	updateUserRequest = types.User{
		ID:       1,
		Username: "Foo",
		Password: "abcdefg",
		Email:    "tester@test.com",
	}

	updateUserResponse = types.User{
		ID:        1,
		Username:  "Foo",
		Email:     "tester@test.com",
		Password:  "abcdefg",
		CreatedAt: createdTime,
		UpdatedAt: createdTime,
	}

	testUserInJson, _           = json.Marshal(testUser)
	newUserRequestInJson, _     = json.Marshal(newUserRequest)
	newUserResponseInJson, _    = json.Marshal(newUserResponse)
	newUserRequest2InJson, _    = json.Marshal(newUserRequest2)
	updateUserRequestInJson, _  = json.Marshal(updateUserRequest)
	updateUserResponseInJson, _ = json.Marshal(updateUserResponse)
)

func TestUserEndpoints(t *testing.T) {
	s := operations.Server{}
	s.Set(getUserDBClientMock(t))

	ts := httptest.NewServer(s.Router)
	defer ts.Close()

	token, err := s.SignIn(testUser.Email, testUser.Password)
	if err != nil {
		log.Fatalf("cannot login: %v\n", err)
	}
	tokenStr := fmt.Sprintf("Bearer %v", token)

	testcasesInOrder := []string{
		"GET /users",
		"GET /users/{id} 200",
		"GET /users/{id} 404",
		"POST /users 200",
		"POST /users 400",
		"PATCH /users/{id} 200",
		"PATCH /users/{id} 400",
		"DELETE /users/{id} 200",
		"DELETE /users/{id} 400",
	}

	testcases := map[string]TestCase{
		"GET /users": {
			method:   http.MethodGet,
			path:     "/users",
			wantCode: http.StatusOK,
			wantBody: fmt.Sprintf(`{"items":[%s]}`, testUserInJson),
		},
		"GET /users/{id} 200": {
			method:   http.MethodGet,
			path:     "/users/1",
			wantCode: http.StatusOK,
			wantBody: fmt.Sprintf(`%s`, testUserInJson),
		},
		"GET /users/{id} 404": {
			method:   http.MethodGet,
			path:     "/users/2",
			wantCode: http.StatusNotFound,
		},
		"POST /users 200": {
			method: http.MethodPost,
			path:   "/users",
			body:   fmt.Sprintf(`%s`, newUserRequestInJson),
			header: http.Header{
				"Content-type": {"application/json"},
			},
			wantCode: http.StatusOK,
			wantBody: fmt.Sprintf(`%s`, newUserResponseInJson),
		},
		"POST /users 400": {
			method: http.MethodPost,
			path:   "/users",
			body:   fmt.Sprintf(`%s`, newUserRequest2InJson),
			header: http.Header{
				"Content-type": {"application/json"},
			},
			wantCode: http.StatusBadRequest,
		},
		"PATCH /users/{id} 200": {
			method: http.MethodPatch,
			path:   "/users/1",
			body:   fmt.Sprintf(`%s`, updateUserRequestInJson),
			header: http.Header{
				"Content-type":  {"application/json"},
				"Authorization": {tokenStr},
			},
			wantCode: http.StatusOK,
			wantBody: fmt.Sprintf(`%s`, updateUserResponseInJson),
		},
		"PATCH /users/{id} 400": {
			method: http.MethodPatch,
			path:   "/users/2",
			body:   fmt.Sprintf(`%s`, updateUserRequestInJson),
			header: http.Header{
				"Content-type":  {"application/json"},
				"Authorization": {tokenStr},
			},
			wantCode: http.StatusNotFound,
		},
		"DELETE /users/{id} 200": {
			method: http.MethodDelete,
			path:   "/users/1",
			header: http.Header{
				"Authorization": {tokenStr},
			},
			wantCode: http.StatusOK,
		},
		"DELETE /users/{id} 400": {
			method: http.MethodDelete,
			path:   "/users/2",
			header: http.Header{
				"Authorization": {tokenStr},
			},
			wantCode: http.StatusNotFound,
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

func getUserDBClientMock(t *testing.T) *mocks.MockClientInterface {
	ctrl := gomock.NewController(t)
	dbClient := mocks.NewMockClientInterface(ctrl)

	dbClient.EXPECT().GetUsers().Return(&types.UserList{
		Items: []*types.User{
			&testUser,
		},
	}).AnyTimes()

	dbClient.EXPECT().GetUserByID(gomock.Eq(uint(1))).Return(&testUser).AnyTimes()
	dbClient.EXPECT().GetUserByID(gomock.Eq(uint(2))).Return(nil).AnyTimes()
	dbClient.EXPECT().GetUserByEmail(gomock.Any()).Return(&testUser).AnyTimes()

	dbClient.EXPECT().CreateUser(gomock.Any()).DoAndReturn(func(user *types.User) error {
		user.ID = 2
		return nil
	}).AnyTimes()

	dbClient.EXPECT().UpdateUser(gomock.Any()).AnyTimes()

	dbClient.EXPECT().DeleteUser(gomock.Eq(uint(1))).Times(1)

	return dbClient
}
