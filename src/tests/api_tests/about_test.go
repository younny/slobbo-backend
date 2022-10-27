package api_tests

import (
	"bytes"
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
	testAbout = types.About{
		Title:    "About me",
		SubTitle: "About me sub",
		Body1:    "This is body 1",
		Body2:    "This is body 2",
		Contacts: &types.Contacts{
			Email:  "abc@example.com",
			Github: "github.com/younny",
		},
	}
)

func TestAboutEndpoints(t *testing.T) {
	s := operations.Server{}
	s.Set(getAboutDBClientMock(t))

	ts := httptest.NewServer(s.Router)
	defer ts.Close()

	testcasesInOrder := []string{
		"GET /about",
		"PATCH /about",
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

func getAboutDBClientMock(t *testing.T) *mocks.MockClientInterface {
	ctrl := gomock.NewController(t)
	dbClient := mocks.NewMockClientInterface(ctrl)

	dbClient.EXPECT().GetAbout().Return(&testAbout).Times(2)

	dbClient.EXPECT().UpdateAbout(gomock.Any()).AnyTimes()
	return dbClient
}
