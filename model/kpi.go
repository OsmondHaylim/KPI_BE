package model

import (
	// "time"
	// "github.com/lib/pq"
)

type Monthly struct{
	Monthly_ID 		int			`gorm:"primaryKey;autoIncrement;uniqueIndex" json:"Monthly_ID"`	
	MiniPAP			*MiniPAP	`gorm:"foreignKey:MinipapID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
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
	Ytd 			float64		`json:"YTD"`
	Remarks 		*string		`json:"Remarks"`
	MinipapID		*int			/* `json:"minipap_id"` */
}

type Attendance struct{
	Year 		int			`gorm:"primaryKey;uniqueIndex" json:"Year"`
	PlanID 		*int		`json:"plan_id"`
	Plan 		*Monthly	`gorm:"foreignKey:plan_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"Planned"`
	ActualID 	*int		`json:"actual_id"`
	Actual 		*Monthly	`gorm:"foreignKey:actual_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"Actual"`
	CutiID 		*int		`json:"cuti_id"`
	Cuti 		*Monthly	`gorm:"foreignKey:cuti_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"Cuti"`
	IzinID 		*int		`json:"izin_id"`
	Izin 		*Monthly	`gorm:"foreignKey:izin_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"Izin"`
	LainID 		*int		`json:"lain_id"`
	Lain 		*Monthly	`gorm:"foreignKey:lain_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"Lain"`
}

type MiniPAP struct{
	MiniPAP_ID	int			`gorm:"primaryKey;autoIncrement;uniqueIndex" json:"Minipap_ID"`
	Monthly		[]Monthly	`gorm:"foreignKey:MinipapID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"Monthly"`
}

type Factor struct{
	Factor_ID	int			`gorm:"primaryKey;autoIncrement;uniqueIndex" json:"Factor_ID"`
	Title 		string		`gorm:"notNull" json:"Title"`
	Unit 		string		`gorm:"notNull" json:"Unit"`
	Target 		string		`gorm:"notNull" json:"Target"`
	PlanID		*int		`json:"plan_id"`
	Plan 		*MiniPAP	`gorm:"foreignKey:plan_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"Planned"`
	ActualID	*int		`json:"actual_id"`
	Actual 		*MiniPAP	`gorm:"foreignKey:actual_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"Actual"`
	ResultID 	*int		`json:"result_id"`
	Result		*Result		`gorm:"foreignKey:ResultID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type Result struct{
	Result_ID	int				`gorm:"primaryKey;autoIncrement;uniqueIndex" json:"Result_ID"`
	Name		string			`gorm:"notNull" json:"Name"`
	Factors 	[]Factor		`gorm:"notNull;foreignKey:ResultID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"Factors"`
	ItemID		*int			`json:"item_id"`
	Item		*Item			`gorm:"foreignKey:ItemID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type Item struct{
	Item_ID		int				`gorm:"primaryKey;autoIncrement;uniqueIndex" json:"Item_ID"`
	Name		string			`gorm:"notNull" json:"Name"`
	Results		[]Result		`gorm:"notNull;foreignKey:ItemID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"Results"`
	YearID 		*int			`json:"year_id"`
	Yearly 		*Yearly			`gorm:"foreignKey:YearID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"Yearly"`
}

type Yearly struct{
	Year			int				`gorm:"primaryKey;uniqueIndex" json:"Year"`
	Items			[]Item			`gorm:"notNull;foreignKey:YearID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"Items"`
	AttendanceID	*int			`json:"attendance_id"`
	Attendance 		*Attendance		`gorm:"foreignKey:AttendanceID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"Attendance"`
}

// type Analisa struct{
// 	Year 			int				`gorm:"primaryKey;uniqueIndex" json:"Year"`
// 	Masalah 		[]Masalah		`gorm:"foreignKey:Year;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"Masalah"`
// }

// type Masalah struct{
// 	Masalah_ID		int				`gorm:"primaryKey;autoIncrement;uniqueIndex" json:"Masalah_ID"`
// 	Masalah 		string			`gorm:"notNull"`
// 	Why				pq.StringArray	`gorm:"type:text[]" json:"Why"`
// 	Tindakan		string
// 	Pic				string
// 	Target			string
// 	FolDate			*time.Time
// 	Status			string		
// 	Year			*int			
// }

type Project struct{
	Project_ID		int				`gorm:"primaryKey;autoIncrement;uniqueIndex"`
	Name 			string			`gorm:"notNull" json:"Name"`
	Summary_ID		*int			`gorm:"foreignKey:Summary_ID"`
	INYS			int
	QNYS			int
	IDR				int
	QDR				int
	IPR				int
	QPR				int
	II				int
	QI				int
	IF				int
	QF				int
	IC				int
	QC				int
}

type Summary struct{
	Summary_ID		int			`gorm:"primaryKey;autoIncrement;uniqueIndex"`
	Projects		[]Project	`gorm:"foreignKey:Summary_ID"`
	// IssuedDate		*time.Time
}