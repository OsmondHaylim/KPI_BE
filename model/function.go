package model

import (
	"strconv"
	"net/http"

	"github.com/gin-gonic/gin"
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
						Ytd: (a.Jan + a.Feb + a.Mar + a.Apr + a.May + a.Jun + a.Jul + a.Aug + a.Sep + a.Oct + a.Nov + a.Dec) / 12,
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
					a.Ytd = ((data.Jan + data.Feb + data.Mar + data.Apr + data.May + data.Jun + data.Jul + data.Aug + data.Sep + data.Oct + data.Nov + data.Dec) / (act.Jan + act.Feb + act.Mar + act.Apr + act.May + act.Jun + act.Jul + act.Aug + act.Sep + act.Oct + act.Nov + act.Dec)) * 100
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



// func PercentParse(a float64, v string, err error) error {
// 	if len(v) > 5{
// 		a, err = strconv.ParseFloat(v[:5], 64)
// 	}else if len(v) >= 1{
// 		a, err = strconv.ParseFloat(v[:len(v)-1], 64)
// 	}
// 	return err
// }
// func Input (jan float64, feb float64, mar float64, apr float64, may float64, jun float64, jul float64, aug float64, sep float64, oct float64, nov float64, dec float64, remarks string) Monthly{
// 	result := Monthly{}
// 	result.Jan = jan
// 	result.Feb = feb
// 	result.Mar = mar
// 	result.Apr = apr
// 	result.May = may
// 	result.Jun = jun
// 	result.Jul = jul
// 	result.Aug = aug
// 	result.Sep = sep
// 	result.Oct = oct
// 	result.Nov = nov
// 	result.Dec = dec
// 	result.Remarks = &remarks
// 	return result
// }

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
func (a Analisa) ToResponse() AnalisaResponse{
	newAnalisa := AnalisaResponse{
		Year: a.Year,
	}
	for _, masalah := range a.Masalah{
		newAnalisa.Masalah = append(newAnalisa.Masalah, masalah.ToResponse())
	}
	return newAnalisa
}
func (m Masalah) ToResponse() MasalahResponse{
	return MasalahResponse{
		Masalah_ID: m.Masalah_ID,
		Masalah: m.Masalah,
		Why: m.Why,
		Tindakan: m.Tindakan,
		Pic: m.Pic,
		Target: m.Target,
	}
}
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
		newItem.Results = append(newItem.Results, Result.ToResponse())
	}
	return newItem
}
func (y Yearly) ToResponse() YearlyResponse{
	newYear := YearlyResponse{
		Year: y.Year,
	}
	for _, Item := range y.Items{
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
func (p Project) ToResponse() ProjectResponse{
	return ProjectResponse{
		Project_ID: p.Project_ID,				
		Name: p.Name,		
		Item: map[string]int{
			"Not Yet Start Issued FR":p.INYS,
			"DR":p.IDR,
			"PR to PO":p.IPR,
			"Install":p.II,
			"Finish":p.IF,
			"Cancelled":p.IC,
		},
		Quantity: map[string]int{
			"Not Yet Start Issued FR":p.QNYS,
			"DR":p.QDR,
			"PR to PO":p.QPR,
			"Install":p.QI,
			"Finish":p.QF,
			"Cancelled":p.QC,
		},
	}
}
func (s Summary) ToResponse() SummaryResponse{
	newSummary := SummaryResponse{
		Summary_ID: s.Summary_ID,
		IssuedDate: s.IssuedDate,
	}
	for _, Project := range s.Projects{
		newSummary.Projects = append(newSummary.Projects, Project.ToResponse())
	}
	return newSummary
}

func (m Monthly) Reseted() Monthly{
	return Monthly{
		Jan: m.Jan,
		Feb: m.Feb,
		Mar: m.Mar,
		Apr: m.Apr,
		May: m.May,
		Jun: m.Jun,
		Jul: m.Jul,
		Aug: m.Aug,
		Sep: m.Sep,
		Oct: m.Oct,
		Nov: m.Nov,
		Dec: m.Dec,
		Ytd: m.Jan + 
		m.Feb +
		m.Mar +
		m.Apr +
		m.May +
		m.Jun +
		m.Jul +
		m.Aug +
		m.Sep +
		m.Oct +
		m.Nov + 
		m.Dec,
		Remarks: m.Remarks,
		MinipapID: m.MinipapID,
	}
}

func ErrorCheck(k *gin.Context, err error) bool {
	if err != nil {
		k.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return true
	}
	return false
}