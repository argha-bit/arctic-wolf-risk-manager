package requests

import (
	"arctic-wolf-risk-manager/models"
	"log"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type RiskRequestHandler interface {
	Bind(c echo.Context, request interface{}, model *models.Risk) error
}

type RiskRequest struct{}
type GetRiskRequest struct {
	Id string `param:"id"`
}
type CreateRiskRequest struct {
	State           string `json:"state" validate:"required,checkValidRiskStatus"`
	RiskTitle       string `json:"title"`
	RiskDescription string `json:"description"`
}

func (r *RiskRequest) Bind(c echo.Context, req interface{}, model *models.Risk) error {
	var err error

	if err = c.Bind(req); err != nil {
		log.Println("Error in reading request", err.Error())
		return err
	}
	if err = c.Validate(req); err != nil {
		log.Println("error in validating request", err.Error())
		return err
	}
	switch v := req.(type) {
	case *GetRiskRequest:
		model.Id = req.(*GetRiskRequest).Id
		log.Println("request is ", req.(*GetRiskRequest))
	case *CreateRiskRequest:
		model.Id = uuid.NewString()
		model.State = req.(*CreateRiskRequest).State
		model.RiskTitle = req.(*CreateRiskRequest).RiskTitle
		model.RiskDescription = req.(*CreateRiskRequest).RiskDescription
	default:
		log.Println("request type Unknown for transformation", v)
	}
	return nil
}

func NewRiskRequestHandler() RiskRequestHandler {
	return &RiskRequest{}
}
