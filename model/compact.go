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