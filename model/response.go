package model

// import "time"

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
	Message  	string			`json:"message"`
	Masalah 	[]Masalah 		`json:"data"`}
type MinipapArrayResponse struct {
	Message  	string			`json:"message"`
	Minipap 	[]MiniPAP 		`json:"data"`}
type MonthlyArrayResponse struct {
	Message  	string			`json:"message"`
	Monthly 	[]Monthly 		`json:"data"`}
type ProjectArrayResponse struct {
	Message  	string				`json:"message"`
	Project		[]Project 	`json:"data"`}
type ResultArrayResponse struct {
	Message  	string					`json:"message"`
	Result 		[]ResultResponse 		`json:"data"`}
type SummaryArrayResponse struct{
	Message  	string					`json:"message"`
	Summary 	[]Summary		`json:"data"`}
type YearlyArrayResponse struct {
	Message  	string					`json:"message"`
	Yearly 		[]YearlyResponse 		`json:"data"`}
type UserArrayResponse struct {
	Message string `json:"message"`
	Users   []UserResponse `json:"data"`}
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
	Lain 		*Monthly	`json:"Lain"`
}
type FactorResponse struct{
	Factor_ID	int			`json:"Factor_ID"`
	Title 		string		`json:"Title"`
	Unit 		string		`json:"Unit"`
	Target 		string		`json:"Target"`
	Plan 		*MiniPAP	`json:"Planned"`
	Actual 		*MiniPAP	`json:"Actual"`
	Percentage	[][]Monthly	`json:"Percentage"`
	Result_ID	*int		`json:"Result_ID"`}
type ResultResponse struct{
	Result_ID	int					`json:"Result_ID"`
	Name		string				`json:"Name"`
	Factors 	[]FactorResponse	`json:"Factors"`
	Item_ID		*int 				`json:"Item_ID"`
}
type ItemResponse struct{
	Item_ID		int					`json:"Item_ID"`
	Name		string				`json:"Name"`
	Results		[]ResultResponse	`json:"Results"`
	Year 		*int				`json:"Year"`
}
// type ProjectResponse struct{
// 	Project_ID		int				
// 	Name 			string		
// 	Item			map[string]int	
// 	Quantity		map[string]int
// 	Summary_ID 		*int
// }
// type SummaryResponse struct{
// 	Summary_ID		int				
// 	Projects		[]ProjectResponse
// 	IssuedDate		*time.Time
// }
type YearlyResponse struct{
	Year			int						`json:"Year"`
	Items			[]ItemResponse			`json:"Items"`
	Attendance 		*AttendanceResponse		`json:"Attendance"`}

type LoginResponse struct {
	Message 	string					`json:"message"`
	Token		string    				`json:"token"`
}

type UserResponse struct {
	ID 			int				`gorm:"primaryKey" json:"id"`
	Username	string			`gorm:"notNull" json:"username"`	
	Email		string			`gorm:"notNull" json:"email"`
	Role		string			`gorm:"notNull" json:"role"`
}

