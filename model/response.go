package model

import "time"

// Success/Error Response
type ErrorResponse struct {
	Error string `json:"message"`}
type SuccessResponse struct {
	Message string `json:"message"`}
// Array Response
type AnalisaArrayResponse struct {
	Message  	string			`json:"message"`
	Analisa 	[]Analisa		`json:"data"`}
type AttendanceArrayResponse struct {
	Message  	string					`json:"message"`
	Attendance 	[]AttendanceResponse	`json:"data"`}
type FactorArrayResponse struct {
	Message  	string					`json:"message"`
	Factor 		[]FactorResponse 		`json:"data"`}
type ItemArrayResponse struct {
	Message  	string			`json:"message"`
	Item 		[]ItemResponse 			`json:"data"`}
type MasalahArrayResponse struct {
	Message  	string					`json:"message"`
	Masalah 	[]MasalahResponse 		`json:"data"`}
type MinipapArrayResponse struct {
	Message  	string			`json:"message"`
	Minipap 	[]MiniPAP 		`json:"data"`}
type MonthlyArrayResponse struct {
	Message  	string			`json:"message"`
	Monthly 	[]Monthly 		`json:"data"`}
type ProjectArrayResponse struct {
	Message  	string				`json:"message"`
	Project		[]ProjectResponse 	`json:"data"`}
type ResultArrayResponse struct {
	Message  	string					`json:"message"`
	Result 		[]ResultResponse 		`json:"data"`}
type SummaryArrayResponse struct{
	Message  	string					`json:"message"`
	Summary 	[]SummaryResponse		`json:"data"`}
type YearlyArrayResponse struct {
	Message  	string					`json:"message"`
	Yearly 		[]YearlyResponse 		`json:"data"`}
type UserArrayResponse struct {
	Message string `json:"message"`
	Users   []User_compact `json:"data"`}
type SessionArrayResponse struct {
	Message  string    `json:"message"`
	Sessions []Session `json:"data"`}
// Simplified Response
type AttendanceResponse struct{
	Year 		int			`json:"Year"`
	Plan 		*Monthly	`json:"Planned"`
	Actual 		*Monthly	`json:"Actual"`
	Cuti 		*Monthly	`json:"Cuti"`
	Izin 		*Monthly	`json:"Izin"`
	Lain 		*Monthly	`json:"Lain"`}
type FactorResponse struct{
	Factor_ID	int			`json:"Factor_ID"`
	Title 		string		`json:"Title"`
	Unit 		string		`json:"Unit"`
	Target 		string		`json:"Target"`
	Plan 		*MiniPAP		`json:"Planned"`
	Actual 		*MiniPAP		`json:"Actual"`
	Percentage	[][]Monthly		`json:"Percentage"`}
type ResultResponse struct{
	Result_ID	int					`json:"Result_ID"`
	Name		string				`json:"Name"`
	Factors 	[]FactorResponse	`json:"Factors"`}
type ItemResponse struct{
	Item_ID		int					`json:"Item_ID"`
	Name		string				`json:"Name"`
	Results		[]ResultResponse	`json:"Results"`}
type MasalahResponse struct{
	Masalah_ID		int				`gorm:"primaryKey" json:"Masalah_ID"`
	Masalah 		string			`gorm:"notNull"`
	Why				[]string		`gorm:"type:text[]" json:"Why"`
	Tindakan		string
	Pic				string
	Target			string}
type AnalisaResponse struct{
	Year 			int
	Masalah			[]MasalahResponse}
type ProjectResponse struct{
	Project_ID		int				
	Name 			string		
	Item			map[string]int	
	Quantity		map[string]int}
type SummaryResponse struct{
	Summary_ID		int				
	Projects		[]ProjectResponse
	IssuedDate		*time.Time}
type YearlyResponse struct{
	Year			int						`json:"Year"`
	Items			[]ItemResponse			`json:"Items"`
	Attendance 		*AttendanceResponse		`json:"Attendance"`}

type SessionResponse struct {
	Message string  `json:"message"`
	Session Session `json:"data"`
}
type LoginResponse struct {
	Message 	string					`json:"message"`
	Data		struct{
		ApiKey 		string    			`json:"apiKey"`
		User   		User_compact	    `json:"user"`
	}									`json:"data"`
}
type RegisterInput struct {
	Username         string
	Email            string
	Password         string
	Confirm_password string
}
type UserResponse struct {
	Message string `json:"message"`
	User    User_compact   `json:"data"`
}

