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
	loginRequestUser = &types.User{
		Email:    "abc@abc.com",
		Password: "1234",
	}

	loginRequestInvalidUser = &types.User{
		Email: "abc@abc.com",
	}

	loginRequestNotExistingUser = &types.User{
		Email:    "efh@aaa.com",
		Password: "1234",
	}

	loginResponse = &types.TokenResponse{
		Token: "abcdefg....",
	}

	loginRequestJson, _                = json.Marshal(loginRequestUser)
	loginRequestInvalidUserJson, _     = json.Marshal(loginRequestInvalidUser)
	loginRequestNotExistingUserJson, _ = json.Marshal(loginRequestNotExistingUser)
)

func TestLoginEndpoints(t *testing.T) {
	s := Server{}
	s.Set(getLoginDBClientMock(t))

	ts := httptest.NewServer(s.Router)
	defer ts.Close()

	testcasesInOrder := []string{
		"POST /login",
		"POST /login 400",
		"POST /login 404",
	}

	testcases := map[string]TestCase{
		"POST /login": {
			method: http.MethodPost,
			path:   "/login",
			header: http.Header{
				"Content-Type": {"application/json"},
			},
			body:     fmt.Sprintf(`%s`, loginRequestJson),
			wantCode: http.StatusOK,
		},
		"POST /login 400": {
			method: http.MethodPost,
			path:   "/login",
			header: http.Header{
				"Content-Type": {"application/json"},
			},
			body:     fmt.Sprintf(`%s`, loginRequestInvalidUserJson),
			wantCode: http.StatusBadRequest,
		},
		"POST /login 404": {
			method: http.MethodPost,
			path:   "/login",
			header: http.Header{
				"Content-Type": {"application/json"},
			},
			body:     fmt.Sprintf(`%s`, loginRequestNotExistingUserJson),
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

func getLoginDBClientMock(t *testing.T) *mocks.MockClientInterface {
	ctrl := gomock.NewController(t)
	dbClient := mocks.NewMockClientInterface(ctrl)

	dbClient.EXPECT().GetUserByEmail(gomock.Eq("abc@abc.com")).Return(loginRequestUser).Times(1)

	dbClient.EXPECT().GetUserByEmail(gomock.Eq("efh@aaa.com")).Return(nil).Times(1)

	return dbClient
}
