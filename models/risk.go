package models

type Risk struct {
	Id              string `json:"id"`
	State           string `json:"state"`
	RiskTitle       string `json:"title"`
	RiskDescription string `json:"description"`
}
