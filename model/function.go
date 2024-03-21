package model

import (
	"strconv"
)

func (p Factor) ToPercentage() [][]Monthly{
	percentMonthly := [][]Monthly{}
	temp := []Monthly{}
	if p.Plan != nil || p.Actual != nil{
		for _, act := range p.Actual.Monthly{
			for _, data := range p.Plan.Monthly{
				b := "Percentage for Plan MonthlyID " + strconv.Itoa(data.Monthly_ID) + " actual MonthlyID " + strconv.Itoa(act.Monthly_ID) + " in Factor ID " + strconv.Itoa(p.Factor_ID) 
				var a Monthly
				if p.Target == "Zero"{
					a = Monthly{
						Jan: ((data.Jan + 1) / (act.Jan + 1)) * 100,
						Feb: ((data.Feb + 1) / (act.Feb + 1)) * 100,
						Mar: ((data.Mar + 1) / (act.Mar + 1)) * 100,
						Apr: ((data.Apr + 1) / (act.Apr + 1)) * 100,
						May: ((data.May + 1) / (act.May + 1)) * 100,
						Jun: ((data.Jun + 1) / (act.Jun + 1)) * 100,
						Jul: ((data.Jul + 1) / (act.Jul + 1)) * 100,
						Aug: ((data.Aug + 1) / (act.Aug + 1)) * 100,
						Sep: ((data.Sep + 1) / (act.Sep + 1)) * 100,
						Oct: ((data.Oct + 1) / (act.Oct + 1)) * 100,
						Nov: ((data.Nov + 1) / (act.Nov + 1)) * 100,
						Dec: ((data.Dec + 1) / (act.Dec + 1)) * 100,
						Remarks: &b,	
					}
				} else{
					if data.Jan == 0 {a.Jan = (act.Jan / (data.Jan + 1)) * 100} else {a.Jan = (act.Jan / data.Jan) * 100}
					if data.Feb == 0 {a.Feb = (act.Feb / (data.Feb + 1)) * 100} else {a.Feb = (act.Feb / data.Feb) * 100}
					if data.Mar == 0 {a.Mar = (act.Mar / (data.Mar + 1)) * 100} else {a.Mar = (act.Mar / data.Mar) * 100}
					if data.Apr == 0 {a.Apr = (act.Apr / (data.Apr + 1)) * 100} else {a.Apr = (act.Apr / data.Apr) * 100}
					if data.May == 0 {a.May = (act.May / (data.May + 1)) * 100} else {a.May = (act.May / data.May) * 100}
					if data.Jun == 0 {a.Jun = (act.Jun / (data.Jun + 1)) * 100} else {a.Jun = (act.Jun / data.Jun) * 100}
					if data.Jul == 0 {a.Jul = (act.Jul / (data.Jul + 1)) * 100} else {a.Jul = (act.Jul / data.Jul) * 100}
					if data.Aug == 0 {a.Aug = (act.Aug / (data.Aug + 1)) * 100} else {a.Aug = (act.Aug / data.Aug) * 100}
					if data.Sep == 0 {a.Sep = (act.Sep / (data.Sep + 1)) * 100} else {a.Sep = (act.Sep / data.Sep) * 100}
					if data.Oct == 0 {a.Oct = (act.Oct / (data.Oct + 1)) * 100} else {a.Oct = (act.Oct / data.Oct) * 100}
					if data.Nov == 0 {a.Nov = (act.Nov / (data.Nov + 1)) * 100} else {a.Nov = (act.Nov / data.Nov) * 100}
					if data.Dec == 0 {a.Dec = (act.Dec / (data.Dec + 1)) * 100} else {a.Dec = (act.Dec / data.Dec) * 100}
					a.Remarks = &b
				}
				temp = append(temp, a)
			}
			percentMonthly = append(percentMonthly, temp)
			temp = []Monthly{}
		}
	}
	return percentMonthly
}

func (a Attendance) ToResponse() AttendanceResponse{
	return AttendanceResponse{
		Year: a.Year,
		Plan: a.Plan,
		Actual: a.Actual,
		Cuti: a.Cuti,
		Izin: a.Izin,
		Lain: a.Lain,
	}
}
// func (p PAP) ToResponse() PAPResponse{
// 	return PAPResponse{
// 		Pap_ID: p.Pap_ID,
// 		Plan: p.Plan,
// 		Actual: p.Actual,
// 		Percentage: p.ToPercentage(),
// 	}
// }
func (f Factor) ToResponse() FactorResponse{
	return FactorResponse{
		Factor_ID: f.Factor_ID,
		Title: f.Title,
		Unit: f.Unit,
		Target: f.Target,
		Plan: f.Plan,
		Actual: f.Actual,
		Percentage: f.ToPercentage(),	}
}
func (r Result) ToResponse() ResultResponse{
	newRes := ResultResponse{
		Result_ID: r.Result_ID,
		Name: r.Name,
	}
	for _, data := range r.Factors{
		// newFac := FactorResponse{
		// 	Factor_ID: data.Factor_ID,
		// 	Title: data.Title,
		// 	Unit: data.Unit,
		// 	Target: data.Target,
		// 	Plan: f.Plan,
		// 	Actual: f.Actual,
		// 	Percentage: f.ToPercentage(),
		// }
		newRes.Factors = append(newRes.Factors, data.ToResponse())
	}
	return newRes
}
func (i Item) ToResponse() ItemResponse{
	newItem := ItemResponse{
		Item_ID: i.Item_ID,
		Name: i.Name,
	}
	for _, Result := range i.Results{		
		// newRes := ResultResponse{
		// 	Result_ID: Result.Result_ID,
		// 	Name: Result.Name,
		// }
		// for _, data := range Result.Factors{
		// 	// newFac := FactorResponse{
		// 	// 	Factor_ID: data.Factor_ID,
		// 	// 	Title: data.Title,
		// 	// 	Unit: data.Unit,
		// 	// 	Target: data.Target,
		// 	// 	Statistic: data.Statistic,
		// 	// }
		// 	newRes.Factors = append(newRes.Factors, data.ToResponse())
		// }
		newItem.Results = append(newItem.Results, Result.ToResponse())
	}
	return newItem
}
func (y Yearly) ToResponse() YearlyResponse{
	newYear := YearlyResponse{
		Year: y.Year,
	}
	for _, Item := range y.Items{
		// newItem := ItemResponse{
		// 	Item_ID: Item.Item_ID,
		// 	Name: Item.Name,
		// }
		// for _, Result := range Item.Results{		
		// 	newRes := ResultResponse{
		// 		Result_ID: Result.Result_ID,
		// 		Name: Result.Name,
		// 	}
		// 	for _, data := range Result.Factors{
		// 		// newFac := FactorResponse{
		// 		// 	Factor_ID: data.Factor_ID,
		// 		// 	Title: data.Title,
		// 		// 	Unit: data.Unit,
		// 		// 	Target: data.Target,
		// 		// 	Statistic: data.Statistic,
		// 		// }
		// 		newRes.Factors = append(newRes.Factors, data.ToResponse())
		// 	}
		// 	newItem.Results = append(newItem.Results, newRes)
		// }
		newYear.Items = append(newYear.Items, Item.ToResponse())
	}
	newAtt := AttendanceResponse{
		Year: y.Attendance.Year,
		Plan: y.Attendance.Plan,
		Actual: y.Attendance.Actual,
		Cuti: y.Attendance.Cuti,
		Izin: y.Attendance.Izin,
		Lain: y.Attendance.Lain,
	}
	newYear.Attendance = &newAtt
	return newYear
}
