package model

type CompactAttendance struct{
	Year 		int			`gorm:"primaryKey" json:"Year"`
	PlanID 		*int		`json:"plan_id"`
	ActualID 	*int		`json:"actual_id"`
	CutiID 		*int		`json:"cuti_id"`
	IzinID 		*int		`json:"izin_id"`
	LainID 		*int		`json:"lain_id"`
}

type CompactPAP struct{
	Pap_ID		int				`gorm:"primaryKey;autoIncrement" json:"Pap_ID"`
	PlanID		*int			`json:"plan_id"`
	ActualID	*int			`json:"actual_id"`
}