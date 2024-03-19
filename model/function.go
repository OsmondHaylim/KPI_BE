package model

import (
	"strconv"
)

func (p PAP) ToPercentage() [][]Monthly{
	percentMonthly := [][]Monthly{}
	temp := []Monthly{}
	for _, act := range p.Actual.Monthly{
		for _, data := range p.Plan.Monthly{
			b := "Percentage for Plan MonthlyID " + strconv.Itoa(data.Monthly_ID) + " actual MonthlyID " + strconv.Itoa(act.Monthly_ID) + " in PAP ID " + strconv.Itoa(p.Pap_ID) 
			a := Monthly{
				Jan: (data.Jan / act.Jan) * 100,
				Feb: (data.Feb / act.Feb) * 100,
				Mar: (data.Mar / act.Mar) * 100,
				Apr: (data.Apr / act.Apr) * 100,
				May: (data.May / act.May) * 100,
				Jun: (data.Jun / act.Jun) * 100,
				Jul: (data.Jul / act.Jul) * 100,
				Aug: (data.Aug / act.Aug) * 100,
				Sep: (data.Sep / act.Sep) * 100,
				Oct: (data.Oct / act.Oct) * 100,
				Nov: (data.Nov / act.Nov) * 100,
				Dec: (data.Dec / act.Dec) * 100,
				Remarks: &b,
			}
			temp = append(temp, a)
		}
		percentMonthly = append(percentMonthly, temp)
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
func (p PAP) ToResponse() PAPResponse{
	return PAPResponse{
		Pap_ID: p.Pap_ID,
		Plan: p.Plan,
		Actual: p.Actual,
		Percentage: p.ToPercentage(),
	}
}
func (f Factor) ToResponse() FactorResponse{
	return FactorResponse{
		Factor_ID: f.Factor_ID,
		Title: f.Title,
		Unit: f.Unit,
		Target: f.Target,
		Statistic: f.Statistic,
	}
}
func (r Result) ToResponse() ResultResponse{
	newRes := ResultResponse{
		Result_ID: r.Result_ID,
		Name: r.Name,
	}
	for _, data := range r.Factors{
		newFac := FactorResponse{
			Factor_ID: data.Factor_ID,
			Title: data.Title,
			Unit: data.Unit,
			Target: data.Target,
			Statistic: data.Statistic,
		}
		newRes.Factors = append(newRes.Factors, newFac)
	}
	return newRes
}
func (i Item) ToResponse() ItemResponse{
	newItem := ItemResponse{
		Item_ID: i.Item_ID,
		Name: i.Name,
	}
	for _, Result := range i.Results{		
		newRes := ResultResponse{
			Result_ID: Result.Result_ID,
			Name: Result.Name,
		}
		for _, data := range Result.Factors{
			newFac := FactorResponse{
				Factor_ID: data.Factor_ID,
				Title: data.Title,
				Unit: data.Unit,
				Target: data.Target,
				Statistic: data.Statistic,
			}
			newRes.Factors = append(newRes.Factors, newFac)
		}
		newItem.Results = append(newItem.Results, newRes)
	}
	return newItem
}
func (y Yearly) ToResponse() YearlyResponse{
	newYear := YearlyResponse{
		Year: y.Year,
	}
	for _, Item := range y.Items{
		newItem := ItemResponse{
			Item_ID: Item.Item_ID,
			Name: Item.Name,
		}
		for _, Result := range Item.Results{		
			newRes := ResultResponse{
				Result_ID: Result.Result_ID,
				Name: Result.Name,
			}
			for _, data := range Result.Factors{
				newFac := FactorResponse{
					Factor_ID: data.Factor_ID,
					Title: data.Title,
					Unit: data.Unit,
					Target: data.Target,
					Statistic: data.Statistic,
				}
				newRes.Factors = append(newRes.Factors, newFac)
			}
			newItem.Results = append(newItem.Results, newRes)
		}
		newYear.Items = append(newYear.Items, newItem)
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
