package model

type Monthly struct{
	Monthly_ID 		int			`gorm:"primaryKey;autoIncrement" json:"Monthly_ID"`	
	MiniPAP			*MiniPAP	`gorm:"foreignKey:MinipapID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Jan 			float64		`gorm:"notNull" json:"January"`
	Feb 			float64		`gorm:"notNull" json:"February"`
	Mar 			float64		`gorm:"notNull" json:"March"`
	Apr 			float64		`gorm:"notNull" json:"April"`
	May 			float64		`gorm:"notNull" json:"May"`
	Jun 			float64		`gorm:"notNull" json:"June"`
	Jul 			float64		`gorm:"notNull" json:"July"`
	Aug 			float64		`gorm:"notNull" json:"August"`
	Sep 			float64		`gorm:"notNull" json:"September"`
	Oct 			float64		`gorm:"notNull" json:"October"`
	Nov 			float64		`gorm:"notNull" json:"November"`
	Dec 			float64		`gorm:"notNull" json:"December"`
	Ytd 			float64		`gorm:"notNull" json:"YTD"`
	Remarks 		*string		`json:"Remarks"`
	MinipapID		*int			/* `json:"minipap_id"` */
}

type Attendance struct{
	Year 		int			`gorm:"primaryKey" json:"Year"`
	PlanID 		*int		`json:"plan_id"`
	Plan 		*Monthly	`gorm:"foreignKey:plan_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"Planned"`
	ActualID 	*int		`json:"actual_id"`
	Actual 		*Monthly	`gorm:"foreignKey:actual_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"Actual"`
	CutiID 		*int		`json:"cuti_id"`
	Cuti 		*Monthly	`gorm:"foreignKey:cuti_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"Cuti"`
	IzinID 		*int		`json:"izin_id"`
	Izin 		*Monthly	`gorm:"foreignKey:izin_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"Izin"`
	LainID 		*int		`json:"lain_id"`
	Lain 		*Monthly	`gorm:"foreignKey:lain_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"Lain"`
}

type MiniPAP struct{
	MiniPAP_ID	int			`gorm:"primaryKey;autoIncrement" json:"Minipap_ID"`
	Monthly		[]Monthly	`gorm:"foreignKey:MinipapID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"Monthly"`
}

// type PAP struct{
// 	Pap_ID		int				`gorm:"primaryKey;autoIncrement" json:"Pap_ID"`
// 	PlanID		*int			`json:"plan_id"`
// 	Plan 		*MiniPAP		`gorm:"foreignKey:plan_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"Planned"`
// 	ActualID	*int			`json:"actual_id"`
// 	Actual 		*MiniPAP		`gorm:"foreignKey:actual_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"Actual"`
// }

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

type Factor struct{
	Factor_ID	int			`gorm:"primaryKey;autoIncrement" json:"Factor_ID"`
	Title 		string		`gorm:"notNull" json:"Title"`
	Unit 		string		`gorm:"notNull" json:"Unit"`
	Target 		string		`gorm:"notNull" json:"Target"`
	PlanID		*int		`json:"plan_id"`
	Plan 		*MiniPAP	`gorm:"foreignKey:plan_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"Planned"`
	ActualID	*int		`json:"actual_id"`
	Actual 		*MiniPAP	`gorm:"foreignKey:actual_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"Actual"`
	ResultID 	*int		`json:"result_id"`
	Result		*Result		`gorm:"foreignKey:ResultID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Result struct{
	Result_ID	int				`gorm:"primaryKey;autoIncrement" json:"Result_ID"`
	Name		string			`gorm:"notNull" json:"Name"`
	Factors 	[]Factor		`gorm:"notNull;foreignKey:ResultID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"Factors"`
	ItemID		*int			`json:"item_id"`
	Item		*Item			`gorm:"foreignKey:ItemID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Item struct{
	Item_ID		int				`gorm:"primaryKey;autoIncrement" json:"Item_ID"`
	Name		string			`gorm:"notNull" json:"Name"`
	Results		[]Result		`gorm:"notNull;foreignKey:ItemID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"Results"`
	YearID 		*int			`json:"year_id"`
	Yearly 		*Yearly			`gorm:"foreignKey:YearID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"Yearly"`
}

type Yearly struct{
	Year			int				`gorm:"primaryKey" json:"Year"`
	Items			[]Item			`gorm:"notNull;foreignKey:YearID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"Items"`
	AttendanceID	*int			`json:"attendance_id"`
	Attendance 		*Attendance		`gorm:"foreignKey:AttendanceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"Attendance"`
}


