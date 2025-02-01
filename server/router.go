package server

import (
	"arctic-wolf-risk-manager/adapter"
	"arctic-wolf-risk-manager/controllers/health"
	riskController "arctic-wolf-risk-manager/controllers/risk"
	riskrepo "arctic-wolf-risk-manager/repository/risk"
	"arctic-wolf-risk-manager/requests"
	riskResp "arctic-wolf-risk-manager/response/risk"
	riskUc "arctic-wolf-risk-manager/usecase/risk"
	"arctic-wolf-risk-manager/utils/validator"

	"github.com/labstack/echo/v4"
)

func newRouter() *echo.Echo {
	e := echo.New()
	e.Validator = validator.NewValidator()
	health.NewHealthController(e)

	storage := adapter.GetInstance()
	riskReq := requests.NewRiskRequestHandler()
	riskRepo := riskrepo.NewRiskRepositoryHandler(storage)
	riskUsecase := riskUc.NewRiskUsecaseHandler(riskRepo)
	riskResponse := riskResp.NewRiskResponseHandler()
	riskController.NewRiskController(e, riskReq, riskUsecase, riskResponse)
	return e
}
