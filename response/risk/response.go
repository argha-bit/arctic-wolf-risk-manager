package risk

import (
	"arctic-wolf-risk-manager/models"
	"arctic-wolf-risk-manager/requests"
	"arctic-wolf-risk-manager/response"
	"arctic-wolf-risk-manager/utils/validator"
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

func (r RiskResponse) ProcessRiskResponse(req interface{}, data interface{}) (interface{}, error) {
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
	return resp, nil
}
func (r RiskResponse) ProcessErrorResponse(err error) (interface{}, error) {
	resp := ArcticWolfResponse{}
	switch err.(type) {
	case *validator.ValidationError:
		resp.Code = "ARW400"
		resp.Message = "Invalid Input"
		resp.Data = []string{}
		resp.Error = err.(*validator.ValidationError).Fields

	default:
		resp.Code = "ARW500"
		resp.Message = "Internal Server Error"
		resp.Data = []string{}
		resp.Error = ErrrorResp{}
	}
	return resp, nil
}

func NewRiskResponseHandler() response.RiskResponse {
	return RiskResponse{}
}
