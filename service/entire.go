package service

import (
	"goreact/model"
	"sync"
)

// Add Entire
func (cs *crudService) AddEntireYearly(input *model.YearlyResponse) error {
	wg, errs := model.GoRoutineInit()
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
		wg.Add(1)
		go cs.AddEntireItem(&wg, &item, &newYearly.Year, errs)
	}
	return model.SimpleErrorChanCheck(&wg, errs)
}
func (cs *crudService) AddEntireItem(wg *sync.WaitGroup, input *model.ItemResponse, id *int, errChan chan error) {
	defer wg.Done()
	//Storing Items
	var newItem model.Item
	if id != nil {newItem.YearID = id}
	newItem.Name = input.Name
	//Creating Items to get id
	err := cs.AddItem(&newItem)
	if err != nil {errChan <- err; return}
	//Storing Results
	wgs := sync.WaitGroup{}
	for _, result := range input.Results{
		wgs.Add(1)
		go cs.AddEntireResult(&wgs, &result, &newItem.Item_ID, errChan)
	}
	wgs.Wait()
}
func (cs *crudService) AddEntireResult(wg *sync.WaitGroup, input *model.ResultResponse, id *int, errChan chan error) {
	defer wg.Done()
	//Storing Results
	var newResult model.Result
	newResult.Name = input.Name
	if id != nil {newResult.ItemID = id}
	//Creating Results to get id
	err := cs.AddResult(&newResult)
	if err != nil {errChan <- err; return}
	//Storing Factors
	wgs := sync.WaitGroup{}
	for _, factor := range input.Factors{
		wgs.Add(1)
		go cs.AddEntireFactor(&wgs, &factor, &newResult.Result_ID, errChan)
	}
	wgs.Wait()
}
func (cs *crudService) AddEntireFactor(wg *sync.WaitGroup, input *model.FactorResponse, id *int, errChan chan error) {
	defer wg.Done()
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
		if err != nil {errChan <- err; return}
		//Connect MiniPAP to Factor
		newFactor.PlanID = &newMinipap.MiniPAP_ID
		//Storing Plan Monthly
		for _, monthly := range input.Plan.Monthly{
			newMonthly := monthly.Reseted()
			newMonthly.MinipapID = &newMinipap.MiniPAP_ID
			err := cs.AddMonthly(&newMonthly)
			if err != nil {errChan <- err; return}
		}
	}
	if input.Actual != nil{
		//Storing MiniPAP Actual
		var newMinipap model.MiniPAP
		//Create MiniPAP to get id
		err := cs.AddMinipap(&newMinipap)
		if err != nil {errChan <- err; return}
		//Connect MiniPAP to Factor
		newFactor.ActualID = &newMinipap.MiniPAP_ID
		//Storing Actual Monthly
		for _, monthly := range input.Actual.Monthly{
			newMonthly := monthly.Reseted()
			newMonthly.MinipapID = &newMinipap.MiniPAP_ID
			err := cs.AddMonthly(&newMonthly)
			if err != nil {errChan <- err; return}
		}
	}
	err := cs.AddFactor(&newFactor)
	if err != nil {errChan <- err; return}
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
func (cs *crudService) AddEntireAnalisa(input *model.AnalisaResponse) error{
	var newAnalisa model.Analisa
	newAnalisa.Year = input.Year

	err := cs.AddAnalisa(&newAnalisa)
	if err != nil {return err}

	wg, errs := model.GoRoutineInit()

	for _, data := range input.Masalah{
		wg.Add(1)
		go func(data model.MasalahResponse){
			var newMasalah = model.Masalah{
				Masalah: data.Masalah,
				Why: data.Why,
				Tindakan: data.Tindakan,
				Pic: data.Pic,
				Target: data.Target,
				Year: &newAnalisa.Year,
				//Default status here
			}
			err = cs.AddMasalah(&newMasalah)
			if err != nil {errs <- err; return}
		}(data)
	}
	return model.SimpleErrorChanCheck(&wg, errs)
}
func (cs *crudService) AddEntireSummary(input *model.SummaryResponse) error{
	var newSummary = model.Summary{
		IssuedDate: input.IssuedDate,
	}
	err := cs.AddSummary(&newSummary)
	if err != nil {return err}
	wg, errs := model.GoRoutineInit()
	for _, data := range input.Projects{
		wg.Add(1)
		go func(data model.ProjectResponse){
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
			err := cs.AddProject(&newProject)
			if err != nil {errs <- err; return}
		}(data)	
	}
	return model.SimpleErrorChanCheck(&wg, errs)
}

// Delete Entire
func (cs *crudService) DeleteEntireYearly(input int) error{
	temp, err := cs.GetYearlyByID(input)
	if err != nil {return err}
	if temp.Items != nil {
		wg, errs := model.GoRoutineInit()
		for _, item := range temp.Items{
			wg.Add(1)
			go cs.DeleteEntireItem(&wg, item.Item_ID, errs)
		}
		err = model.SimpleErrorChanCheck(&wg, errs)
		if err != nil {return err}
	}
	if temp.Attendance != nil {
		err = cs.DeleteEntireAttendance(temp.Attendance.Year)
		if err != nil {return err}
	}
	err = cs.DeleteYearly(input)
	if err != nil {return err}
	return nil
}
func (cs *crudService) DeleteEntireItem(wg *sync.WaitGroup, input int, errs chan error){
	defer wg.Done()
	temp, err := cs.GetItemByID(input)
	if err != nil {errs <- err; return}
	if temp.Results != nil {
		wgs := sync.WaitGroup{}
		for _, result := range temp.Results{
			wgs.Add(1)
			cs.DeleteEntireResult(&wgs, result.Result_ID, errs)
		}
		wgs.Wait()
	}
	err = cs.DeleteItem(input)
	if err != nil {errs <- err; return}
}
func (cs *crudService) DeleteEntireResult(wg *sync.WaitGroup, input int, errs chan error){
	defer wg.Done()
	temp, err := cs.GetResultByID(input)
	if err != nil {errs <- err; return}
	if temp.Factors != nil {
		wgs := sync.WaitGroup{}
		for _, factor := range temp.Factors{
			wgs.Add(1)
			go cs.DeleteEntireFactor(&wgs, factor.Factor_ID, errs)
		}
		wgs.Wait()
	}
	err = cs.DeleteResult(input)
	if err != nil {errs <- err; return}
}
func (cs *crudService) DeleteEntireFactor(wg *sync.WaitGroup, input int, errs chan error){
	defer wg.Done()
	temp, err := cs.GetFactorByID(input)
	if err != nil {errs <- err; return}
	
	if temp.Plan != nil {
		for _, monthly := range temp.Plan.Monthly{
			err = cs.DeleteMonthly(monthly.Monthly_ID)
			if err != nil {errs <- err; return}
		}
		err = cs.DeleteMinipap(temp.Plan.MiniPAP_ID)
		if err != nil {errs <- err; return}
	}
	if temp.Actual != nil {
		for _, monthly := range temp.Actual.Monthly{
			err = cs.DeleteMonthly(monthly.Monthly_ID)
			if err != nil {errs <- err; return}
		}
		err = cs.DeleteMinipap(temp.Actual.MiniPAP_ID)
		if err != nil {errs <- err; return}
	}
	err = cs.DeleteFactor(input)
	if err != nil {errs <- err; return}
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
func (cs *crudService) DeleteEntireAnalisa(input int) error{
	temp, err := cs.GetAnalisaByID(input)
	if err != nil {return err}
	wg, errs := model.GoRoutineInit()
	wg.Add(1)
	go func(){
		defer wg.Done()
		for _, data := range temp.Masalah{	
			err = cs.DeleteMasalah(data.Masalah_ID)
			if err != nil {errs <- err; return}
		}
	}()
	err = model.SimpleErrorChanCheck(&wg, errs)
	if err != nil {return err}
	return cs.DeleteAnalisa(temp.Year)
}
func (cs *crudService) DeleteEntireSummary(input int) error{
	temp, err := cs.GetSummaryByID(input)
	if err != nil {return err}
	wg, errs := model.GoRoutineInit()
	wg.Add(1)
	go func(){
		defer wg.Done()
		for _, data := range temp.Projects{
			err = cs.DeleteProject(data.Project_ID)
			if err != nil {errs <- err; return}
		}
	}()
	wg.Wait()
	err = model.SimpleErrorChanCheck(&wg, errs)
	if err != nil {return err}
	return cs.DeleteSummary(input)
}

// Update Entire
func (cs *crudService) UpdateEntireYearly(id int, input model.YearlyResponse) error {
	wg, errs := model.GoRoutineInit()
	before, err := cs.yearlyRepo.GetByID(id)
	if err != nil {return err}
	newYearly := input.Back()
	//Update Attendance
	err = cs.UpdateEntireAttendance(newYearly.Year, *input.Attendance)
	if err != nil {return err}
	//Updating Yearly
	err = cs.UpdateYearly(id, newYearly)
	if err != nil {return err}
	//Delete Item (no id in response)
	for _, item := range before.Items{
		wg.Add(1)
		go cs.DeleteEntireItem(&wg, item.Item_ID, errs)
	}
	err = model.SimpleErrorChanCheck(&wg, errs)
	if err != nil {return err}
	// Add Back items
	for _, item := range input.Items{
		wg.Add(1)
		go cs.AddEntireItem(&wg, &item, &id, errs)
	}
	return model.SimpleErrorChanCheck(&wg, errs)
}
func (cs *crudService) UpdateEntireItem(id int, input model.ItemResponse) error{
	wg, errs := model.GoRoutineInit()
	before, err := cs.itemRepo.GetByID(id)
	if err != nil {return err}
	newItem := input.Back()
	//Updating Yearly
	err = cs.UpdateItem(id, newItem)
	if err != nil {return err}
	//Delete Item (no id in response)
	for _, result := range before.Results{
		wg.Add(1)
		go cs.DeleteEntireResult(&wg, result.Result_ID, errs)
	}
	err = model.SimpleErrorChanCheck(&wg, errs)
	if err != nil {return err}
	// Add Back Results
	for _, result := range input.Results{
		wg.Add(1)
		go cs.AddEntireResult(&wg, &result, &id, errs)
	}
	return model.SimpleErrorChanCheck(&wg, errs)
}
func (cs *crudService) UpdateEntireResult(id int, input model.ResultResponse) error{
	wg, errs := model.GoRoutineInit()
	before, err := cs.resultRepo.GetByID(id)
	if err != nil {return err}
	newResult := input.Back()
	//Updating Yearly
	err = cs.UpdateResult(id, newResult)
	if err != nil {return err}
	//Delete Factor (no id in response)
	for _, factor := range before.Factors{
		wg.Add(1)
		go cs.DeleteEntireFactor(&wg, factor.Factor_ID, errs)
	}
	err = model.SimpleErrorChanCheck(&wg, errs)
	if err != nil {return err}
	// Add Back Factors
	for _, factor := range input.Factors{
		wg.Add(1)
		go cs.AddEntireFactor(&wg, &factor, &id, errs)
	}
	return model.SimpleErrorChanCheck(&wg, errs)
}
func (cs *crudService) UpdateEntireFactor(id int, input model.FactorResponse) error{
	wg, errs := model.GoRoutineInit()
	before, err := cs.factorRepo.GetByID(id)
	if err != nil {return err}
	newFactor := input.Back()
	if before.PlanID != nil{
		wg.Add(1)
		go func (){
			defer wg.Done()
			newFactor.PlanID = before.PlanID
			for _, monthly := range before.Plan.Monthly{
				err = cs.DeleteMonthly(monthly.Monthly_ID)
				if err != nil {errs <- err; return}
			}
		}()
	}
	if before.ActualID != nil{
		wg.Add(1)
		go func (){
			defer wg.Done()
			newFactor.ActualID = before.ActualID
			for _, monthly := range before.Actual.Monthly{
				err = cs.DeleteMonthly(monthly.Monthly_ID)
				if err != nil {errs <- err; return}
			}
		}()
	}
	err = model.SimpleErrorChanCheck(&wg, errs)
	if err != nil {return err}

	if newFactor.Plan != nil{
		wg.Add(1)
		go func (){
			defer wg.Done()
			for _, month := range newFactor.Plan.Monthly{
				if before.PlanID != nil{
					month.MinipapID = before.PlanID
				}else {
					newMini := model.MiniPAP{}
					err := cs.AddMinipap(&newMini)
					if err != nil {errs <- err; return}
					newFactor.PlanID = &newMini.MiniPAP_ID
					month.MinipapID = &newMini.MiniPAP_ID
				}
				err := cs.AddMonthly(&month)
				if err != nil {errs <- err; return}
			}
		}()
	}
	if newFactor.Actual != nil{
		wg.Add(1)
		go func (){
			defer wg.Done()
			for _, month := range newFactor.Actual.Monthly{
				if before.ActualID != nil{
					month.MinipapID = before.ActualID
				}else {
					newMini := model.MiniPAP{}
					err := cs.AddMinipap(&newMini)
					if err != nil {errs <- err}
					newFactor.ActualID = &newMini.MiniPAP_ID
					month.MinipapID = &newMini.MiniPAP_ID
				}
				err := cs.AddMonthly(&month)
				if err != nil {errs <- err}
			}
		}()
	}
	err = model.SimpleErrorChanCheck(&wg, errs)
	if err != nil {return err}
	return cs.UpdateFactor(id, newFactor)
}
func (cs *crudService) UpdateEntireAttendance(id int, input model.AttendanceResponse) error{
	// Storing Attendance
	before, err := cs.attendanceRepo.GetByID(id)
	if err != nil {return err}
	newAtt := input.Back()

	if newAtt.Plan != nil{
		if before.PlanID != nil{
			cs.UpdateMonthly(*before.PlanID, *newAtt.Plan)
			if err != nil {return err}
		}else{
			cs.AddMonthly(newAtt.Plan)
			if err != nil {return err}
		}
	}else if before.PlanID != nil{
		cs.DeleteMonthly(*before.PlanID)
		if err != nil {return err}
	}
	if newAtt.Actual != nil{
		if before.ActualID != nil{
			cs.UpdateMonthly(*before.ActualID, *newAtt.Actual)
			if err != nil {return err}
		}else{
			cs.AddMonthly(newAtt.Actual)
			if err != nil {return err}
		}
	}else if before.ActualID != nil{
		cs.DeleteMonthly(*before.ActualID)
		if err != nil {return err}
	}
	if newAtt.Cuti != nil{
		if before.CutiID != nil{
			cs.UpdateMonthly(*before.CutiID, *newAtt.Cuti)
			if err != nil {return err}
		}else{
			cs.AddMonthly(newAtt.Cuti)
			if err != nil {return err}
		}
	}else if before.CutiID != nil{
		cs.DeleteMonthly(*before.CutiID)
		if err != nil {return err}
	}
	if newAtt.Izin != nil{
		if before.IzinID != nil{
			cs.UpdateMonthly(*before.IzinID, *newAtt.Izin)
			if err != nil {return err}
		}else{
			cs.AddMonthly(newAtt.Izin)
			if err != nil {return err}
		}
	}else if before.IzinID != nil{
		cs.DeleteMonthly(*before.IzinID)
		if err != nil {return err}
	}
	if newAtt.Lain != nil{
		if before.LainID != nil{
			cs.UpdateMonthly(*before.LainID, *newAtt.Lain)
			if err != nil {return err}
		}else{
			cs.AddMonthly(newAtt.Lain)
			if err != nil {return err}
		}
	}else if before.LainID != nil{
		cs.DeleteMonthly(*before.LainID)
		if err != nil {return err}
	}
	err = cs.UpdateAttendance(id, newAtt)
	if err != nil {return err}
	return nil
}
func (cs *crudService) UpdateEntireAnalisa(id int, input model.AnalisaResponse) error{
	before, err := cs.analisaRepo.GetByID(id)
	if err != nil {return err}
	newAnalisa := input.Back()

	if before.Masalah != nil{
		for _, data := range before.Masalah{
			err = cs.DeleteMasalah(data.Masalah_ID)
			if err != nil {return err}
		}
	}
	if newAnalisa.Masalah != nil{
		for _, data := range newAnalisa.Masalah{
			data.Year = &id
			err = cs.AddMasalah(&data)
			if err != nil {return err}
		}
	}
	return cs.UpdateAnalisa(id, newAnalisa)
}
func (cs *crudService) UpdateEntireSummary(id int, input model.SummaryResponse) error{
	before, err := cs.summaryRepo.GetByID(id)
	if err != nil {return err}
	newSummary := input.Back()

	if before.Projects != nil{
		for _, data := range before.Projects{
			err = cs.DeleteProject(data.Project_ID)
			if err != nil {return err}
		}
	}
	if newSummary.Projects != nil{
		for _, data := range newSummary.Projects{
			data.Summary_ID = &id
			err = cs.AddProject(&data)
			if err != nil {return err}
		}
	}
	return cs.UpdateSummary(id, newSummary)
}