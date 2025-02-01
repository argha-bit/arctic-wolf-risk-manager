package adapter

import (
	"arctic-wolf-risk-manager/models"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	storage := GetInstance()
	testcases := []struct {
		Name             string
		Search           string
		ExpectedResponse []models.Risk
	}{
		{
			Name:             "empty list",
			ExpectedResponse: []models.Risk{},
		},
		{
			Name: "return all the available risks",
			ExpectedResponse: []models.Risk{
				{
					Id:              "test2",
					RiskTitle:       "test",
					RiskDescription: "test",
					State:           "test",
				},
				{
					Id:              "test1",
					RiskTitle:       "test",
					RiskDescription: "test",
					State:           "test",
				},
			},
		},
		{
			Name:   "get by Id",
			Search: "test1",
			ExpectedResponse: []models.Risk{
				{
					Id:              "test1",
					RiskTitle:       "test",
					RiskDescription: "test",
					State:           "test",
				},
			},
		},
		{
			Name:             "record unavailable",
			ExpectedResponse: []models.Risk{},
			Search:           uuid.NewString(),
		},
	}
	assert := assert.New(t)
	for i, test := range testcases {
		actualResponse := storage.Get(test.Search)
		assert.Equal(test.ExpectedResponse, actualResponse, test.Name)
		if i == 0 {
			storage.Set(&models.Risk{
				Id:              "test2",
				RiskTitle:       "test",
				RiskDescription: "test",
				State:           "test",
			})
			storage.Set(&models.Risk{
				Id:              "test1",
				RiskTitle:       "test",
				RiskDescription: "test",
				State:           "test",
			})
		}
	}
}
func TestSet(t *testing.T) {
	assert := assert.New(t)

	storage := GetInstance()
	risk := models.Risk{
		Id:              "test1",
		RiskDescription: "test",
		RiskTitle:       "test",
		State:           "risk",
	}
	storage.Set(&risk)
	checkVal := storage.Get(risk.Id)
	assert.Equal([]models.Risk{risk}, checkVal, "value getting Set")

}
func TestGetInstance(t *testing.T) {

	instance1 := GetInstance()
	instance2 := GetInstance()

	assert := assert.New(t)
	assert.Equal(instance1, instance2, "assert return same instance of storage always")

}
