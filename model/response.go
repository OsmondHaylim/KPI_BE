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

// type Factor struct{
// 	Factor_ID	int			`gorm:"primaryKey;autoIncrement" json:"Factor_ID"`
// 	Title 		string		`gorm:"notNull" json:"Title"`
// 	Unit 		string		`gorm:"notNull" json:"Unit"`
// 	Target 		string		`gorm:"notNull" json:"Target"`
// 	StatID		*int     	`json:"stat_id"`
// 	Statistic 	*PAP		`gorm:"foreignKey:stat_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"Statistic"`
// 	ResultID 	*int		`json:"result_id"`
// 	Result		*Result		`gorm:"foreignKey:ResultID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
// }

// type Result struct{
// 	Result_ID	int				`gorm:"primaryKey;autoIncrement" json:"Result_ID"`
// 	Name		string			`gorm:"notNull" json:"Name"`
// 	Factors 	[]Factor		`gorm:"notNull;foreignKey:ResultID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"Factors"`
// 	ItemID		*int			`json:"item_id"`
// 	Item		*Item			`gorm:"foreignKey:ItemID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
// }

// type Item struct{
// 	Item_ID		int				`gorm:"primaryKey;autoIncrement" json:"Item_ID"`
// 	Name		string			`gorm:"notNull" json:"Name"`
// 	Results		[]Result		`gorm:"notNull;foreignKey:ItemID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"Results"`
// 	YearID 		*int			`json:"year_id"`
// 	Yearly 		*Yearly			`gorm:"foreignKey:YearID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"Yearly"`
// }

// type Yearly struct{
// 	Year			int				`gorm:"primaryKey" json:"Year"`
// 	Items			[]Item			`gorm:"notNull;foreignKey:YearID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"Items"`
// 	AttendanceID	*int			`json:"attendance_id"`
// 	Attendance 		*Attendance		`gorm:"foreignKey:AttendanceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"Attendance"`
// }
