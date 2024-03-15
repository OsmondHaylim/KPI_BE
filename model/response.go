package model

type ErrorResponse struct {
	Error string `json:"message"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type KpiArrayResponse struct {
	Message  	string			`json:"message"`
	Kpis 		[]Factor 	`json:"data"`
}

type MonthlyArrayResponse struct {
	Message  	string			`json:"message"`
	Monthly 	[]Monthly 	`json:"data"`
}