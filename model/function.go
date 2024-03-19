package model

import "strconv"

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