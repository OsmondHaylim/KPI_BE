package model

type ErrorResponse struct {
	Error string `json:"message"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type AttendanceArrayResponse struct {
	Message  	string			`json:"message"`
	Attendance 	[]Attendance	`json:"data"`
}
type FactorArrayResponse struct {
	Message  	string			`json:"message"`
	Factor 		[]Factor 		`json:"data"`
}
type ItemArrayResponse struct {
	Message  	string			`json:"message"`
	Item 		[]Item 			`json:"data"`
}
type MinipapArrayResponse struct {
	Message  	string			`json:"message"`
	Minipap 	[]MiniPAP 		`json:"data"`
}
type MonthlyArrayResponse struct {
	Message  	string			`json:"message"`
	Monthly 	[]Monthly 		`json:"data"`
}
type PapArrayResponse struct {
	Message  	string			`json:"message"`
	Pap 		[]PAP 			`json:"data"`
}
type ResultArrayResponse struct {
	Message  	string			`json:"message"`
	Result 		[]Result 		`json:"data"`
}
type YearlyArrayResponse struct {
	Message  	string			`json:"message"`
	Yearly 		[]Yearly 		`json:"data"`
}

