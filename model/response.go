package model

type ErrorResponse struct {
	Error string `json:"message"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type AttendanceArrayResponse struct {
	Message  	string			`json:"message"`
	Attendance 	[]AttendanceResponse	`json:"data"`
}
type FactorArrayResponse struct {
	Message  	string			`json:"message"`
	Factor 		[]FactorResponse 		`json:"data"`
}
type ItemArrayResponse struct {
	Message  	string			`json:"message"`
	Item 		[]ItemResponse 			`json:"data"`
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
	Pap 		[]PAPResponse 			`json:"data"`
}
type ResultArrayResponse struct {
	Message  	string			`json:"message"`
	Result 		[]ResultResponse 		`json:"data"`
}
type YearlyArrayResponse struct {
	Message  	string			`json:"message"`
	Yearly 		[]YearlyResponse 		`json:"data"`
}

type PAPResponse struct{
	Pap_ID		int				`json:"Pap_ID"`
	Plan 		*MiniPAP		`json:"Planned"`
	Actual 		*MiniPAP		`json:"Actual"`
	Percentage	[][]Monthly		`json:"Monthlies"`
}

type AttendanceResponse struct{
	Year 		int			`json:"Year"`
	Plan 		*Monthly	`json:"Planned"`
	Actual 		*Monthly	`json:"Actual"`
	Cuti 		*Monthly	`json:"Cuti"`
	Izin 		*Monthly	`json:"Izin"`
	Lain 		*Monthly	`json:"Lain"`
}

type FactorResponse struct{
	Factor_ID	int			`json:"Factor_ID"`
	Title 		string		`json:"Title"`
	Unit 		string		`json:"Unit"`
	Target 		string		`json:"Target"`
	Statistic 	*PAP		`json:"Statistic"`
}

type ResultResponse struct{
	Result_ID	int				`json:"Result_ID"`
	Name		string			`json:"Name"`
	Factors 	[]FactorResponse		`json:"Factors"`
}

type ItemResponse struct{
	Item_ID		int				`json:"Item_ID"`
	Name		string			`json:"Name"`
	Results		[]ResultResponse		`json:"Results"`
}

type YearlyResponse struct{
	Year			int				`json:"Year"`
	Items			[]ItemResponse			`json:"Items"`
	Attendance 		*AttendanceResponse		`json:"Attendance"`
}
