package model

type CompactAttendance1 struct{
	Year 		int			`gorm:"primaryKey" json:"Year"`
	Plan 		*Monthly	`gorm:"foreignKey:plan_id" json:"Planned"`
	Actual 		*Monthly	`gorm:"foreignKey:actual_id" json:"Actual"`
	Cuti 		*Monthly	`gorm:"foreignKey:cuti_id" json:"Cuti"`
	Izin 		*Monthly	`gorm:"foreignKey:izin_id" json:"Izin"`
	Lain 		*Monthly	`gorm:"foreignKey:lain_id" json:"Lain"`
}

type CompactAttendance2 struct{
	Year 		int			`gorm:"primaryKey" json:"Year"`
	PlanID 		*int		`json:"plan_id"`
	ActualID 	*int		`json:"actual_id"`
	CutiID 		*int		`json:"cuti_id"`
	IzinID 		*int		`json:"izin_id"`
	LainID 		*int		`json:"lain_id"`
}

type CompactPAP1 struct{
	Pap_ID		int				`gorm:"primaryKey;autoIncrement" json:"Pap_ID"`
	Plan 		*MiniPAP		`gorm:"foreignKey:plan_id" json:"Planned"`
	Actual 		*MiniPAP		`gorm:"foreignKey:actual_id" json:"Actual"`
}

type CompactPAP2 struct{
	Pap_ID		int				`gorm:"primaryKey;autoIncrement" json:"Pap_ID"`
	PlanID		*int			`json:"plan_id"`
	ActualID	*int			`json:"actual_id"`
}