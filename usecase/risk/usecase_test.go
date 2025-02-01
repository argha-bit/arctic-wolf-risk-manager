package risk

import (
	"arctic-wolf-risk-manager/adapter"
	"arctic-wolf-risk-manager/models"
	"arctic-wolf-risk-manager/repository/risk"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRisk(t *testing.T) {
	testCase := []struct {
		Name             string
		RiskModel        *models.Risk
		ExpectedResponse []models.Risk
	}{
		{
			Name: "No List Present",
			RiskModel: &models.Risk{
				Id: "test1",
			},
			ExpectedResponse: []models.Risk{},
		},
		{
			Name: "Get Risk by Id List",
			RiskModel: &models.Risk{
				Id: "test1",
			},
			ExpectedResponse: []models.Risk{
				{
					Id:              "test1",
					State:           "test",
					RiskDescription: "test",
					RiskTitle:       "test",
				},
			},
		},
		{
			Name:      "Get Entire List",
			RiskModel: &models.Risk{},
			ExpectedResponse: []models.Risk{
				{
					Id:              "test1",
					State:           "test",
					RiskDescription: "test",
					RiskTitle:       "test",
				}, {
					Id:              "test2",
					State:           "test",
					RiskDescription: "test",
					RiskTitle:       "test",
				},
			},
		},
	}
	storage := adapter.GetInstance()
	repository := risk.NewRiskRepositoryHandler(storage)
	getUsecase := NewRiskUsecaseHandler(repository)

	assert := assert.New(t)
	for i, test := range testCase {
		actualResponse := getUsecase.GetRisk(test.RiskModel)
		assert.Equal(actualResponse, test.ExpectedResponse, test.Name)
		if i == 0 {
			repository.CreateRisk(&models.Risk{
				Id:              "test1",
				State:           "test",
				RiskDescription: "test",
				RiskTitle:       "test",
			})
			repository.CreateRisk(&models.Risk{
				Id:              "test2",
				State:           "test",
				RiskDescription: "test",
				RiskTitle:       "test",
			})
		}
	}

}

func TestCreateRisk(t *testing.T) {
	testcases := []struct {
		Name             string
		RiskModel        *models.Risk
		ExpectedResponse models.Risk
	}{
		{
			Name: "create risk",
			RiskModel: &models.Risk{
				Id:              "test1",
				State:           "test",
				RiskDescription: "test",
				RiskTitle:       "test",
			},
			ExpectedResponse: models.Risk{
				Id:              "test1",
				State:           "test",
				RiskDescription: "test",
				RiskTitle:       "test",
			},
		},
	}
	assert := assert.New(t)

	for _, test := range testcases {
		storage := adapter.GetInstance()
		repository := risk.NewRiskRepositoryHandler(storage)
		createUsecase := NewRiskUsecaseHandler(repository)
		actualResponse := createUsecase.CreateRisk(&models.Risk{
			Id:              "test1",
			State:           "test",
			RiskDescription: "test",
			RiskTitle:       "test",
		})
		assert.Equal(test.ExpectedResponse, actualResponse, test.Name)
	}
}
