package model

type ErrorResponse struct {
	Error string `json:"message"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type AnalisaArrayResponse struct {
	Message  	string			`json:"message"`
	Analisa 	[]Analisa		`json:"data"`
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
type MasalahArrayResponse struct {
	Message  	string			`json:"message"`
	Masalah 	[]MasalahResponse 		`json:"data"`
}
type MinipapArrayResponse struct {
	Message  	string			`json:"message"`
	Minipap 	[]MiniPAP 		`json:"data"`
}
type MonthlyArrayResponse struct {
	Message  	string			`json:"message"`
	Monthly 	[]Monthly 		`json:"data"`
}
// type PapArrayResponse struct {
// 	Message  	string			`json:"message"`
// 	Pap 		[]PAPResponse 			`json:"data"`
// }
type ResultArrayResponse struct {
	Message  	string			`json:"message"`
	Result 		[]ResultResponse 		`json:"data"`
}
type YearlyArrayResponse struct {
	Message  	string			`json:"message"`
	Yearly 		[]YearlyResponse 		`json:"data"`
}

// type PAPResponse struct{
// 	Pap_ID		int				`json:"Pap_ID"`
// 	Plan 		*MiniPAP		`json:"Planned"`
// 	Actual 		*MiniPAP		`json:"Actual"`
// 	Percentage	[][]Monthly		`json:"Percentage"`
// }

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
	// Statistic 	*PAP		`json:"Statistic"`
	Plan 		*MiniPAP		`json:"Planned"`
	Actual 		*MiniPAP		`json:"Actual"`
	Percentage	[][]Monthly		`json:"Percentage"`
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

type MasalahResponse struct{
	Masalah_ID		int				`gorm:"primaryKey" json:"Masalah_ID"`
	Masalah 		string			`gorm:"notNull"`
	Why				[]string		`gorm:"type:text[]" json:"Why"`
	Tindakan		string
	Pic				string
	Target			string
}

type YearlyResponse struct{
	Year			int				`json:"Year"`
	Items			[]ItemResponse			`json:"Items"`
	Attendance 		*AttendanceResponse		`json:"Attendance"`
}
