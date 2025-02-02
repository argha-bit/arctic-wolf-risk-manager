package risk

import (
	"arctic-wolf-risk-manager/models"
	"arctic-wolf-risk-manager/requests"
	"arctic-wolf-risk-manager/response"
	"arctic-wolf-risk-manager/utils/validator"
	"strings"

	"github.com/labstack/echo/v4"
)

type RiskResponse struct {
}
type ArcticWolfResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}
type ErrrorResp struct {
}

func (r RiskResponse) ProcessRiskResponse(req interface{}, data interface{}) interface{} {
	resp := ArcticWolfResponse{}
	switch req.(type) {
	case *requests.GetRiskRequest:
		resp.Code = "ARW200"
		resp.Message = "Request Processed Successfully"
		if len(data.([]models.Risk)) == 0 {
			resp.Message = "No Risk Found!"
		}
		resp.Data = data
		resp.Error = ErrrorResp{}
	case *requests.CreateRiskRequest:
		resp.Code = "ARW200"
		resp.Message = "Risk Request Created Successfully"
		resp.Data = data
		resp.Error = ErrrorResp{}
	}
	return resp
}
func (r RiskResponse) ProcessErrorResponse(err error) interface{} {
	resp := ArcticWolfResponse{}
	switch err.(type) {
	case *validator.ValidationError:
		resp.Code = "ARW400"
		resp.Message = "Invalid Input"
		resp.Data = []string{}
		resp.Error = err.(*validator.ValidationError).Fields
	case *echo.HTTPError:
		if err.(*echo.HTTPError).Code == 415 {
			resp.Code = "ARW400"
			resp.Message = "Invalid Request"
			resp.Data = []string{}
			resp.Error = map[string]string{
				"error": "Unsupported Media Type. Please use application/json in request header Content-Type",
			}
		}

	default:
		resp.Code = "ARW500"
		resp.Message = "Internal Server Error"
		resp.Data = []string{}
		strings.Split(err.Error(), "")
		resp.Error = ErrrorResp{}
	}
	return resp
}

func NewRiskResponseHandler() response.RiskResponse {
	return RiskResponse{}
}
