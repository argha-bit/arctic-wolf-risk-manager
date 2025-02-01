package risk

import (
	controller "arctic-wolf-risk-manager/controllers"
	"arctic-wolf-risk-manager/models"
	"arctic-wolf-risk-manager/requests"
	"arctic-wolf-risk-manager/response"
	"arctic-wolf-risk-manager/usecase"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	req     requests.RiskRequestHandler
	usecase usecase.RiskUsecasehandler
	resp    response.RiskResponse
}

func (risk Controller) GetRisk(c echo.Context) error {
	var err error
	req := new(requests.GetRiskRequest)
	dataModel := new(models.Risk)
	var data []models.Risk
	if err = risk.req.Bind(c, req, dataModel); err != nil {
		log.Println("error in reading request", err.Error())
		resp, _ := risk.resp.ProcessErrorResponse(err)
		return c.JSON(http.StatusBadRequest, resp)
	}
	data = risk.usecase.GetRisk(dataModel)
	resp, _ := risk.resp.ProcessRiskResponse(req, data)

	log.Println("response returned ", resp)
	return c.JSON(http.StatusOK, resp)
}
func (risk Controller) CreateRisk(c echo.Context) error {
	var err error
	req := new(requests.CreateRiskRequest)
	dataModel := new(models.Risk)
	var data models.Risk
	if err = risk.req.Bind(c, req, dataModel); err != nil {
		log.Println("error in reading request", err.Error())
		resp, _ := risk.resp.ProcessErrorResponse(err)
		return c.JSON(http.StatusBadRequest, resp)
	}
	data = risk.usecase.CreateRisk(dataModel)
	resp, _ := risk.resp.ProcessRiskResponse(req, data)

	log.Println("response returned ", resp)
	return c.JSON(http.StatusOK, resp)
}
func NewRiskController(e *echo.Echo, req requests.RiskRequestHandler, usecase usecase.RiskUsecasehandler, resp response.RiskResponse) controller.RiskController {
	riskController := Controller{
		req:     req,
		usecase: usecase,
		resp:    resp,
	}

	e.GET("/v1/risks/:id", riskController.GetRisk)
	e.GET("/v1/risks", riskController.GetRisk)
	e.POST("/v1/risks", riskController.CreateRisk)
	return e
}
