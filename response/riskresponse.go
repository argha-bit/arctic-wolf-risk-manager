package response

type RiskResponse interface {
	ProcessRiskResponse(req interface{}, data interface{}) (interface{}, error)
	ProcessErrorResponse(err error) (interface{}, error)
}
