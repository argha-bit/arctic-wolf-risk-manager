package repository

import "arctic-wolf-risk-manager/models"

type RiskRepoSitoryHandler interface {
	GetRisk(riskModel *models.Risk) []models.Risk
	CreateRisk(riskModel *models.Risk) models.Risk
}
