package service

import (
	"goreact/model"
)

func (cs *crudService) AddEntireYearly(input *model.YearlyResponse) error {
	//Storing Yearlys
	var newYearly model.Yearly
	newYearly.Year = input.Year
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
	
	// Creating monthly from attendance
	if input.Plan != nil{
		newMonthly := input.Plan.Reseted()
		err := cs.AddMonthly(&newMonthly)
		if err != nil {return err}
		newAttendance.PlanID = &newMonthly.Monthly_ID
	}
	if input.Actual != nil{
		newMonthly := input.Actual.Reseted()
		err := cs.AddMonthly(&newMonthly)
		if err != nil {return err}
		newAttendance.ActualID = &newMonthly.Monthly_ID
	}
	if input.Cuti != nil{
		newMonthly := input.Cuti.Reseted()
		err := cs.AddMonthly(&newMonthly)
		if err != nil {return err}
		newAttendance.CutiID = &newMonthly.Monthly_ID
	}
	if input.Izin != nil{
		newMonthly := input.Izin.Reseted()
		err := cs.AddMonthly(&newMonthly)
		if err != nil {return err}
		newAttendance.IzinID = &newMonthly.Monthly_ID
	}
	if input.Lain != nil{
		newMonthly := input.Lain.Reseted()
		err := cs.AddMonthly(&newMonthly)
		if err != nil {return err}
		newAttendance.LainID = &newMonthly.Monthly_ID
	}
	// Creating attendance
	err := cs.AddAttendance(&newAttendance)
	if err != nil {return err}
	return nil
}