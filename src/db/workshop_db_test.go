package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/younny/slobbo-backend/src/types"
)

var (
	testWorkshop = types.Workshop{
		Name:        "Hello",
		Description: "Blah Blah",
		Organiser: &types.Organiser{
			Name: "orgniser1",
			Link: "abc.organiser.com",
		},
		Location: "Koo",
		Duration: &types.Duration{
			StartDate:    "2006-01-02T15:04:05Z",
			EndDate:      "2006-01-02T15:04:05Z",
			TotalInHours: 1,
		},
		Capacity: 1,
		Price: &types.Price{
			Amount: 0,
		},
		Thumbnail: "www",
	}
)

func TestClient_Workshop(t *testing.T) {
	testClient.Client.DropTable(&types.Workshop{})
	testClient.AutoMigrate(&types.Workshop{})
	first := testWorkshop
	err := testClient.CreateWorkshop(&first)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), first.ID)

	second := testWorkshop
	err = testClient.CreateWorkshop(&second)
	assert.NoError(t, err)
	assert.Equal(t, uint(2), second.ID)

	update := first
	update.Name = "Hello Germany"
	err = testClient.UpdateWorkshop(&update)
	assert.NoError(t, err)

	got := testClient.GetWorkshopByID(1)
	assert.Equal(t, testWorkshop.Description, got.Description, "")
	assert.Equal(t, update.Name, got.Name, "")

	testClient.DeleteWorkshop(1)
	got = testClient.GetWorkshopByID(1)
	assert.Empty(t, got)

	err = testClient.DeleteWorkshop(3)
	assert.Error(t, err)
}

func TestClient_PaginateWorkshops(t *testing.T) {
	testClient.DropTable(&types.Workshop{})
	testClient.AutoMigrate(&types.Workshop{})

	for i := 0; i < pageSize+2; i++ {
		workshop := testWorkshop
		_ = testClient.CreateWorkshop(&workshop)
	}
	got := testClient.GetWorkshops(0)
	assert.Equal(t, 10, len(got.Items))
	assert.Equal(t, uint(11), got.NextPageID)

	got = testClient.GetWorkshops(11)
	assert.Equal(t, 2, len(got.Items))
	assert.Equal(t, uint(0), got.NextPageID)
}
