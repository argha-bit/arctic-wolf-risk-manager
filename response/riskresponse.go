package response

type RiskResponse interface {
	ProcessRiskResponse(req interface{}, data interface{}) interface{}
	ProcessErrorResponse(err error) interface{}
}
