package model

type CompactAttendance struct{
	Year 		int			`json:"Year"`
	PlanID 		*int		`json:"plan_id"`
	ActualID 	*int		`json:"actual_id"`
	CutiID 		*int		`json:"cuti_id"`
	IzinID 		*int		`json:"izin_id"`
	LainID 		*int		`json:"lain_id"`
}

type CompactPAP struct{
	Pap_ID		int				`json:"Pap_ID"`
	PlanID		*int			`json:"plan_id"`
	ActualID	*int			`json:"actual_id"`
}

type User_compact struct {
	ID 			int				`gorm:"primaryKey" json:"id"`
	Username	string			`gorm:"notNull" json:"username"`	
	Email		string			`gorm:"notNull" json:"email"`
	Role		string			`gorm:"notNull" json:"role"`
}