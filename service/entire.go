package service

import (
	"goreact/model"
)

// Add Entire
func (cs *crudService) AddEntireYearly(input *model.YearlyResponse) error {
	//Storing Yearlys
	var newYearly model.Yearly
	newYearly.Year = input.Year
	newYearly.AttendanceID = &input.Year
	//Store Attendance
	err := cs.AddEntireAttendance(input.Attendance, &newYearly.Year)
	if err != nil {return err}
	//Creating Yearly
	err = cs.AddYearly(&newYearly)
	if err != nil {return err}
	//Storing Items
	for _, item := range input.Items{
		err = cs.AddEntireItem(&item, &newYearly.Year)
		if err != nil {return err}
	}
	return nil
}
func (cs *crudService) AddEntireItem(input *model.ItemResponse, id *int) error {
	//Storing Items
	var newItem model.Item
	if id != nil {newItem.YearID = id}
	newItem.Name = input.Name
	//Creating Items to get id
	err := cs.AddItem(&newItem)
	if err != nil {return err}
	//Storing Results
	for _, result := range input.Results{
		err = cs.AddEntireResult(&result, &newItem.Item_ID)
		if err != nil {return err}
	}
	return nil
}
func (cs *crudService) AddEntireResult(input *model.ResultResponse, id *int) error {
	//Storing Results
	var newResult model.Result
	newResult.Name = input.Name
	if id != nil {newResult.ItemID = id}
	//Creating Results to get id
	err := cs.AddResult(&newResult)
	if err != nil {return err}
	//Storing Factors
	for _, factor := range input.Factors{
		err = cs.AddEntireFactor(&factor, &newResult.Result_ID)
		if err != nil {return err}
	}
	return nil
}
func (cs *crudService) AddEntireFactor(input *model.FactorResponse, id *int) error {
	var newFactor model.Factor
	if id != nil {newFactor.ResultID = id}
	newFactor.Title = input.Title
	newFactor.Unit = input.Unit
	newFactor.Target = input.Target
	if input.Plan != nil{
		//Storing MiniPAP Plan
		var newMinipap model.MiniPAP
		//Create MiniPAP to get id
		err := cs.AddMinipap(&newMinipap)
		if err != nil {return err}
		//Connect MiniPAP to Factor
		newFactor.PlanID = &newMinipap.MiniPAP_ID
		//Storing Plan Monthly
		for _, monthly := range input.Plan.Monthly{
			newMonthly := monthly.Reseted()
			newMonthly.MinipapID = &newMinipap.MiniPAP_ID
			err := cs.AddMonthly(&newMonthly)
			if err != nil {return err}
		}
	}
	if input.Actual != nil{
		//Storing MiniPAP Actual
		var newMinipap model.MiniPAP
		//Create MiniPAP to get id
		err := cs.AddMinipap(&newMinipap)
		if err != nil {return err}
		//Connect MiniPAP to Factor
		newFactor.ActualID = &newMinipap.MiniPAP_ID
		//Storing Actual Monthly
		for _, monthly := range input.Actual.Monthly{
			newMonthly := monthly.Reseted()
			newMonthly.MinipapID = &newMinipap.MiniPAP_ID
			err := cs.AddMonthly(&newMonthly)
			if err != nil {return err}
		}
	}
	err := cs.AddFactor(&newFactor)
	if err != nil {return err}
	return nil
}
func (cs *crudService) AddEntireAttendance(input *model.AttendanceResponse, id *int) error{
	// Storing Attendance
	var newAttendance model.Attendance
	if id != nil {newAttendance.Year = *id} else {newAttendance.Year = input.Year}
	
	// // Creating monthly from attendance(apparently the addAttendance already add the monthly)
	// if input.Plan != nil{
	// 	newMonthly := input.Plan.Reseted()
	// 	err := cs.AddMonthly(&newMonthly)
	// 	if err != nil {return err}
	// 	newAttendance.PlanID = &newMonthly.Monthly_ID
	// 	newAttendance.Plan = &newMonthly
	// }
	// if input.Actual != nil{
	// 	newMonthly := input.Actual.Reseted()
	// 	err := cs.AddMonthly(&newMonthly)
	// 	if err != nil {return err}
	// 	newAttendance.ActualID = &newMonthly.Monthly_ID
	// 	newAttendance.Actual = &newMonthly
	// }
	// if input.Cuti != nil{
	// 	newMonthly := input.Cuti.Reseted()
	// 	err := cs.AddMonthly(&newMonthly)
	// 	if err != nil {return err}
	// 	newAttendance.CutiID = &newMonthly.Monthly_ID
	// 	newAttendance.Cuti = &newMonthly
	// }
	// if input.Izin != nil{
	// 	newMonthly := input.Izin.Reseted()
	// 	err := cs.AddMonthly(&newMonthly)
	// 	if err != nil {return err}
	// 	newAttendance.IzinID = &newMonthly.Monthly_ID
	// 	newAttendance.Izin = &newMonthly
	// }
	// if input.Lain != nil{
	// 	newMonthly := input.Lain.Reseted()
	// 	err := cs.AddMonthly(&newMonthly)
	// 	if err != nil {return err}
	// 	newAttendance.LainID = &newMonthly.Monthly_ID
	// 	newAttendance.Lain = &newMonthly
	// }
	// Creating attendance
	err := cs.AddAttendance(&newAttendance)
	if err != nil {return err}
	return nil
}
// func (cs *crudService) AddEntireAnalisa(input *model.AnalisaResponse) error{
// 	var newAnalisa model.Analisa
// 	newAnalisa.Year = input.Year
// 	err := cs.AddAnalisa(&newAnalisa)
// 	if err != nil {return err}
// 	for _, data := range input.Masalah{
// 		var newMasalah = model.Masalah{
// 			Masalah: data.Masalah,
// 			Why: data.Why,
// 			Tindakan: data.Tindakan,
// 			Pic: data.Pic,
// 			Target: data.Target,
// 			Year: &newAnalisa.Year,
// 			//Default status here
// 		}
// 		err = cs.AddMasalah(&newMasalah)
// 		if err != nil {return err}
// 	}
// 	return nil
// }
func (cs *crudService) AddEntireSummary(input *model.SummaryResponse) error{
	var newSummary = model.Summary{
		// IssuedDate: input.IssuedDate,
	}
	err := cs.AddSummary(&newSummary)
	if err != nil {return err}
	for _, data := range input.Projects{
		var newProject = model.Project{
			Name: data.Name,
			Summary_ID: &newSummary.Summary_ID,
			INYS: data.Item["Not Yet Start Issued FR"],
			QNYS: data.Quantity["Not Yet Start Issued FR"],
			IDR: data.Item["DR"],
			QDR: data.Quantity["DR"],
			IPR: data.Item["PR to PO"],
			QPR: data.Quantity["PR to PO"],
			II: data.Item["Install"],
			QI: data.Quantity["Install"],
			IF: data.Item["Finish"],
			QF: data.Quantity["Finish"],
			IC: data.Item["Cancelled"],
			QC: data.Quantity["Cancelled"],
		}
		err = cs.AddProject(&newProject)
		if err != nil {return err}
	}
	return nil
}

// Delete Entire
func (cs *crudService) DeleteEntireYearly(input int) error{
	temp, err := cs.GetYearlyByID(input)
	if err != nil {return err}
	if temp.Items != nil {
		for _, item := range temp.Items{
			err = cs.DeleteEntireItem(item.Item_ID)
			if err != nil {return err}
		}
	}
	if temp.Attendance != nil {
		cs.DeleteEntireAttendance(temp.Attendance.Year)
		if err != nil {return err}
	}
	
	err = cs.DeleteYearly(input)
	if err != nil {return err}

	
	return nil
}
func (cs *crudService) DeleteEntireItem(input int) error{
	temp, err := cs.GetItemByID(input)
	if err != nil {return err}
	if temp.Results != nil {
		for _, result := range temp.Results{
			err = cs.DeleteEntireResult(result.Result_ID)
			if err != nil {return err}
		}
	}
	err = cs.DeleteItem(input)
	if err != nil {return err}
	return nil
}
func (cs *crudService) DeleteEntireResult(input int) error{
	temp, err := cs.GetResultByID(input)
	if err != nil {return err}
	if temp.Factors != nil {
		for _, factor := range temp.Factors{
			err = cs.DeleteEntireFactor(factor.Factor_ID)
			if err != nil {return err}
		}
	}
	err = cs.DeleteResult(input)
	if err != nil {return err}
	return nil
}
func (cs *crudService) DeleteEntireFactor(input int) error{
	temp, err := cs.GetFactorByID(input)
	if err != nil {return err}
	
	if temp.Plan != nil {
		for _, monthly := range temp.Plan.Monthly{
			err = cs.DeleteMonthly(monthly.Monthly_ID)
			if err != nil {return err}
		}
		err = cs.DeleteMinipap(temp.Plan.MiniPAP_ID)
		if err != nil {return err}
	}
	if temp.Actual != nil {
		for _, monthly := range temp.Actual.Monthly{
			err = cs.DeleteMonthly(monthly.Monthly_ID)
			if err != nil {return err}
		}
		err = cs.DeleteMinipap(temp.Actual.MiniPAP_ID)
		if err != nil {return err}
	}
	err = cs.DeleteFactor(input)
	if err != nil {return err}
	return nil
}
func (cs *crudService) DeleteEntireAttendance(input int) error{
	temp, err := cs.GetAttendanceByID(input)
	if err != nil {return err}
	if temp.Plan != nil{
		err = cs.DeleteMonthly(temp.Plan.Monthly_ID)
		if err != nil {return err}
	}
	if temp.Actual != nil {
		err = cs.DeleteMonthly(temp.Actual.Monthly_ID)
		if err != nil {return err}
	}
	if temp.Cuti != nil {
		err = cs.DeleteMonthly(temp.Cuti.Monthly_ID)
		if err != nil {return err}
	}
	if temp.Izin != nil {
		err = cs.DeleteMonthly(temp.Izin.Monthly_ID)
		if err != nil {return err}
	}
	if temp.Lain != nil {
		err = cs.DeleteMonthly(temp.Lain.Monthly_ID)
		if err != nil {return err}
	}
	err = cs.DeleteAttendance(temp.Year)
	if err != nil {return err}
	return nil
}
// func (cs *crudService) DeleteEntireAnalisa(input int) error{
// 	temp, err := cs.GetAnalisaByID(input)
// 	if err != nil {return err}
// 	err = cs.DeleteAnalisa(temp.Year)
// 	if err != nil {return err}
// 	for _, data := range temp.Masalah{
// 		err = cs.DeleteMasalah(data.Masalah_ID)
// 		if err != nil {return err}
// 	}
// 	return nil
// }
func (cs *crudService) DeleteEntireSummary(input int) error{
	temp, err := cs.GetSummaryByID(input)
	if err != nil {return err}

	for _, data := range temp.Projects{
		err = cs.DeleteProject(data.Project_ID)
		if err != nil {return err}
	}
	
	err = cs.DeleteSummary(temp.Summary_ID)
	if err != nil {return err}

	
	return nil
}

// Update Entire