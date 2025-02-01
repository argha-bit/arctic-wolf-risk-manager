package risk

import (
	"arctic-wolf-risk-manager/models"
	"arctic-wolf-risk-manager/repository"
	"arctic-wolf-risk-manager/usecase"
	"log"
)

type RiskUsecase struct {
	RiskRepo repository.RiskRepoSitoryHandler
}

func (r RiskUsecase) GetRisk(riskModel *models.Risk) []models.Risk {
	log.Println("Getting Risk Data")
	resp := r.RiskRepo.GetRisk(riskModel)
	log.Println("data fetched ", resp, riskModel)
	return resp
}
func (r RiskUsecase) CreateRisk(riskModel *models.Risk) models.Risk {
	log.Println("initiating risk creation", riskModel)
	resp := r.RiskRepo.CreateRisk(riskModel)
	return resp

}
func NewRiskUsecaseHandler(riskRepo repository.RiskRepoSitoryHandler) usecase.RiskUsecasehandler {
	return RiskUsecase{
		RiskRepo: riskRepo,
	}
}
