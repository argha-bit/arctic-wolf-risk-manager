package requests

import (
	"arctic-wolf-risk-manager/models"
	"arctic-wolf-risk-manager/utils/validator"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type MockValidator struct{}

func TestBind(t *testing.T) {
	testCases := []struct {
		Name          string
		Model         *models.Risk
		Req           interface{}
		Param         string
		ReqMethod     string
		ReqUrl        string
		ExpectedModel *models.Risk
		ExpectedError error
		ReqBody       string
		ErrorExpected bool
	}{
		{
			Name:          "Get Valid Request",
			Model:         &models.Risk{},
			Req:           &GetRiskRequest{},
			ReqMethod:     http.MethodGet,
			ReqUrl:        "/v1/risks",
			ExpectedModel: &models.Risk{},
			ExpectedError: nil,
			ErrorExpected: false,
		},
		{
			Name:      "Get Valid Request with ID",
			Model:     &models.Risk{},
			Req:       &GetRiskRequest{},
			ReqMethod: http.MethodGet,
			ReqUrl:    "/v1/risks",
			ExpectedModel: &models.Risk{
				Id: "test1",
			},
			ErrorExpected: false,
			Param:         "test1",
		},
		{
			Name:          "valid post",
			Model:         &models.Risk{},
			Req:           &CreateRiskRequest{},
			ReqMethod:     http.MethodPost,
			ReqUrl:        "/v1/risks",
			ErrorExpected: false,
			ReqBody:       `{"state": "open","title": "Test Risk","description": "Test description"}`,
		},
		{
			Name:          "invalid json",
			Model:         &models.Risk{},
			Req:           &CreateRiskRequest{},
			ReqMethod:     http.MethodPost,
			ReqUrl:        "/v1/risks",
			ErrorExpected: true,
			ReqBody:       `{"state": "open","title": "Test Risk","description": "Test description"`,
		},
		{
			Name:          "invalid state",
			Model:         &models.Risk{},
			Req:           &CreateRiskRequest{},
			ReqMethod:     http.MethodPost,
			ReqUrl:        "/v1/risks",
			ErrorExpected: true,
			ReqBody:       `{"state": "test","title": "Test Risk","description": "Test description"}`,
		},
		{
			Name:          "missing state",
			Model:         &models.Risk{},
			Req:           &CreateRiskRequest{},
			ReqMethod:     http.MethodPost,
			ReqUrl:        "/v1/risks",
			ErrorExpected: true,
			ReqBody:       `{"title": "Test Risk","description": "Test description"}`,
		},
		{
			Name:          "unknown Request type",
			Model:         &models.Risk{},
			Req:           &RiskRequest{},
			ReqMethod:     http.MethodPost,
			ReqUrl:        "/v1/risks",
			ErrorExpected: false,
			ReqBody:       `{"title": "Test Risk","description": "Test description"}`,
		},
	}
	assert := assert.New(t)

	for _, test := range testCases {
		var req *http.Request
		switch test.ReqMethod {
		case http.MethodGet:
			req = httptest.NewRequest(test.ReqMethod, test.ReqUrl, nil)
		case http.MethodPost:
			req = httptest.NewRequest(test.ReqMethod, test.ReqUrl, bytes.NewBufferString(test.ReqBody))
			req.Header.Add("Content-Type", "application/json")
		}
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = validator.NewValidator()
		c := e.NewContext(req, rec)
		if test.Param != "" {
			c.SetParamNames("id")
			c.SetParamValues(test.Param)
		}
		r := NewRiskRequestHandler()
		err := r.Bind(c, test.Req, test.Model)
		switch test.ReqMethod {
		case http.MethodGet:
			assert.Equal(test.ExpectedModel, test.Model, test.Name)
		case http.MethodPost:
			if !test.ErrorExpected {
				if _, ok := test.Req.(*CreateRiskRequest); ok {
					assert.NotEmpty(test.Model.Id)
				}
			}
		}
		switch test.ErrorExpected {
		case true:
			assert.Error(err, test.Name)
		case false:
			assert.NoError(err, test.ExpectedError, test.Name)
		}
	}
}
