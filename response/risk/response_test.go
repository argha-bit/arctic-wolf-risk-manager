package risk

import (
	"arctic-wolf-risk-manager/models"
	"arctic-wolf-risk-manager/requests"
	"arctic-wolf-risk-manager/utils/validator"
	"errors"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestProcessErrorResponse(t *testing.T) {
	responsehandler := NewRiskResponseHandler()
	testStruct := struct {
		State string `json:"state" validate:"required,checkValidRiskStatus"`
	}{
		State: "test",
	}
	err := validator.NewValidator().Validate(testStruct)
	testCases := []struct {
		Name             string
		Err              error
		ExpectedResponse ArcticWolfResponse
	}{
		{
			Name: "invalid state error",
			Err:  err,
			ExpectedResponse: ArcticWolfResponse{
				Code:    "ARW400",
				Message: "Invalid Input",
				Data:    []string{},
				Error: map[string]string{
					"state": "invalid risk status: only open/closed/accepted/investigating are accepted",
				},
			},
		},
		{
			Name: "generic error",
			Err:  errors.New("some generic error"),
			ExpectedResponse: ArcticWolfResponse{
				Code:    "ARW500",
				Message: "Internal Server Error",
				Data:    []string{},
				Error:   ErrrorResp{},
			},
		},
		{
			Name: "echo http Error",
			Err: &echo.HTTPError{
				Code: 415,
			},
			ExpectedResponse: ArcticWolfResponse{
				Code:    "ARW400",
				Message: "Invalid Request",
				Data:    []string{},
				Error: map[string]string{
					"error": "Unsupported Media Type. Please use application/json in request header Content-Type",
				},
			},
		},
	}
	assert := assert.New(t)
	for _, test := range testCases {
		actualResponse := responsehandler.ProcessErrorResponse(test.Err)
		assert.Equal(test.ExpectedResponse, actualResponse.(ArcticWolfResponse), test.Name)
	}

}
func TestProcessRiskResponse(t *testing.T) {
	responsehandler := NewRiskResponseHandler()
	testcases := []struct {
		Name             string
		Req              interface{}
		Data             interface{}
		ExpectedResponse interface{}
	}{
		{
			Name: "no data present",
			Req:  new(requests.GetRiskRequest),
			Data: []models.Risk{},
			ExpectedResponse: ArcticWolfResponse{
				Code:    "ARW200",
				Message: "No Risk Found!",
				Data:    []models.Risk{},
				Error:   ErrrorResp{},
			},
		},
		{
			Name: "single data present get request",
			Req:  new(requests.GetRiskRequest),
			Data: []models.Risk{{Id: "test", State: "test", RiskDescription: "test", RiskTitle: "test"}},
			ExpectedResponse: ArcticWolfResponse{
				Code:    "ARW200",
				Message: "Request Processed Successfully",
				Data: []models.Risk{
					{Id: "test", State: "test", RiskDescription: "test", RiskTitle: "test"},
				},
				Error: ErrrorResp{},
			},
		},
		{
			Name: "multiple data present get request",
			Req:  new(requests.GetRiskRequest),
			Data: []models.Risk{
				{Id: "test", State: "test", RiskDescription: "test", RiskTitle: "test"},
				{Id: "test", State: "test", RiskDescription: "test", RiskTitle: "test"},
			},
			ExpectedResponse: ArcticWolfResponse{
				Code:    "ARW200",
				Message: "Request Processed Successfully",
				Data: []models.Risk{
					{Id: "test", State: "test", RiskDescription: "test", RiskTitle: "test"},
					{Id: "test", State: "test", RiskDescription: "test", RiskTitle: "test"},
				},
				Error: ErrrorResp{},
			},
		},
		{
			Name: "create request",
			Req:  new(requests.CreateRiskRequest),
			Data: models.Risk{Id: "test", State: "test", RiskDescription: "test", RiskTitle: "test"},
			ExpectedResponse: ArcticWolfResponse{
				Code:    "ARW200",
				Message: "Risk Request Created Successfully",
				Data:    models.Risk{Id: "test", State: "test", RiskDescription: "test", RiskTitle: "test"},
				Error:   ErrrorResp{},
			},
		},
	}
	assert := assert.New(t)

	for _, test := range testcases {
		resp := responsehandler.ProcessRiskResponse(test.Req, test.Data)
		assert.Equal(test.ExpectedResponse.(ArcticWolfResponse), resp.(ArcticWolfResponse), test.Name)
	}
}
