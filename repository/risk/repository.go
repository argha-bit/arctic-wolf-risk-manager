package risk

import (
	"arctic-wolf-risk-manager/adapter"
	"arctic-wolf-risk-manager/models"
	"arctic-wolf-risk-manager/repository"
	"log"
)

type RiskRepository struct {
	DB *adapter.RiskStorage
}

func (r *RiskRepository) GetRisk(riskModel *models.Risk) []models.Risk {
	return r.DB.Get(riskModel.Id)
}
func (r *RiskRepository) CreateRisk(riskModel *models.Risk) models.Risk {
	log.Println("creating risk record", riskModel)
	r.DB.Set(riskModel)
	return *riskModel
}

func NewRiskRepositoryHandler(RiskDB *adapter.RiskStorage) repository.RiskRepoSitoryHandler {
	return &RiskRepository{
		DB: RiskDB,
	}
}
