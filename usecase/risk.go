package usecase

import "arctic-wolf-risk-manager/models"

type RiskUsecasehandler interface {
	GetRisk(riskModel *models.Risk) []models.Risk
	CreateRisk(riskModel *models.Risk) models.Risk
}
