package api_tests

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
	"github.com/younny/slobbo-backend/src/api/operations"
	"github.com/younny/slobbo-backend/src/types"
)

var (
	randomDuration = types.Duration{
		StartDate:    "2006-01-02T15:04:05Z",
		EndDate:      "2006-01-02T15:04:05Z",
		TotalInHours: 1,
	}
	testOrganiser = types.Organiser{
		Name: "orgniser1",
		Link: "abc.organiser.com",
	}
	testPrice = types.Price{
		Amount: 0,
	}

	testWorkshop1 = types.Workshop{
		ID:          0,
		Name:        "Workshop 1",
		Description: "ABC",
		Organiser:   &testOrganiser,
		Location:    "f",
		Duration:    &randomDuration,
		Capacity:    2,
		Price:       &testPrice,
		CreatedAt:   randomTime,
		UpdatedAt:   randomTime,
	}
	testWorkshop2 = types.Workshop{
		ID:          1,
		Name:        "Workshop 2",
		Description: "ABC",
		Organiser:   &testOrganiser,
		Location:    "f",
		Duration:    &randomDuration,
		Capacity:    2,
		Price:       &testPrice,
		CreatedAt:   randomTime,
		UpdatedAt:   randomTime,
	}
	newWorkshopRequest = types.Workshop{
		Name:        "New workshop",
		Description: "ABC",
		Organiser:   &testOrganiser,
		Location:    "f",
		Duration:    &randomDuration,
		Capacity:    2,
		Price:       &testPrice,
	}
	newWorkshopResponse = types.Workshop{
		ID:          3,
		Name:        "New workshop",
		Description: "ABC",
		Organiser:   &testOrganiser,
		Location:    "f",
		Duration:    &randomDuration,
		Capacity:    2,
		Price:       &testPrice,
		CreatedAt:   createdTime,
		UpdatedAt:   createdTime,
	}
	updatedWorkshopResponse = types.Workshop{
		ID:          0,
		Name:        "Hello World",
		Description: "ABC",
		Organiser:   &testOrganiser,
		Location:    "f",
		Duration:    &randomDuration,
		Capacity:    2,
		Price:       &testPrice,
		CreatedAt:   randomTime,
		UpdatedAt:   randomTime,
	}
	testWorkshop1Json, _           = json.Marshal(testWorkshop1)
	testWorkshop2Json, _           = json.Marshal(testWorkshop2)
	newWorkshopRequestJson, _      = json.Marshal(newWorkshopRequest)
	newWorkshopResponseJson, _     = json.Marshal(newWorkshopResponse)
	updatedWorkshopResponseJson, _ = json.Marshal(updatedWorkshopResponse)
)

func TestWorkshopEndpoints(t *testing.T) {
	s := operations.Server{}
	s.Set(getWorkshopDBClientMock(t))

	ts := httptest.NewServer(s.Router)
	defer ts.Close()

	tokenStr := Authenticate(s)

	testcasesInOrder := []string{
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
		"GET /workshops": {
			method:   http.MethodGet,
			path:     "/workshops",
			wantCode: http.StatusOK,
			wantBody: fmt.Sprintf(`{"items":[%s,%s]}`, testWorkshop1Json, testWorkshop2Json),
		},
		"GET /workshops?page_id=1": {
			method:   http.MethodGet,
			path:     "/workshops?page_id=1",
			wantCode: http.StatusOK,
			wantBody: fmt.Sprintf(`{"items":[%s]}`, testWorkshop2Json),
		},
		"GET /workshops/{id} 200": {
			method:   http.MethodGet,
			path:     "/workshops/0",
			wantCode: http.StatusOK,
			wantBody: fmt.Sprintf(`%s`, testWorkshop1Json),
		},
		"POST /workshops 200": {
			method: http.MethodPost,
			path:   "/workshops",
			body:   fmt.Sprintf(`%s`, newWorkshopRequestJson),
			header: http.Header{
				"Content-type":  {"application/json"},
				"Authorization": {tokenStr},
			},
			wantCode: http.StatusOK,
			wantBody: fmt.Sprintf(`%s`, newWorkshopResponseJson),
		},
		"PATCH /workshops": {
			method: http.MethodPatch,
			path:   "/workshops/0",
			body:   `{"name":"Hello World"}`,
			header: http.Header{
				"Content-type":  {"application/json"},
				"Authorization": {tokenStr},
			},
			wantCode: http.StatusOK,
			wantBody: fmt.Sprintf(`%s`, updatedWorkshopResponseJson),
		},
		"DELETE /workshops": {
			method: http.MethodDelete,
			path:   "/workshops/1",
			header: http.Header{
				"Content-type":  {"application/json"},
				"Authorization": {tokenStr},
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

func getWorkshopDBClientMock(t *testing.T) *mocks.MockClientInterface {
	ctrl := gomock.NewController(t)
	dbClient := mocks.NewMockClientInterface(ctrl)

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

	dbClient.EXPECT().GetUserByEmail(gomock.Any()).Return(&testUser).AnyTimes()

	return dbClient
}
