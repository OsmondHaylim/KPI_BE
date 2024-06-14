package service

import (
	"errors"
	"fmt"
	// "fmt"
	"goreact/model"
)

// Create Functions
func (cs *crudService) AddAttendance(input *model.Attendance) error{
	newInput := input
	zero := 0
	if input.PlanID != &zero && (input.Plan == nil || input.Plan == &model.Monthly{}){newInput.PlanID = nil}
	if input.ActualID != &zero && (input.Actual == nil || input.Actual == &model.Monthly{}){newInput.ActualID = nil}
	if input.CutiID != &zero && (input.Cuti == nil || input.Cuti == &model.Monthly{}){newInput.CutiID = nil}
	if input.IzinID != &zero && (input.Izin == nil || input.Izin == &model.Monthly{}){newInput.IzinID = nil}
	if input.LainID != &zero && (input.Lain == nil || input.Lain == &model.Monthly{}){newInput.LainID = nil}
	return cs.attendanceRepo.Store(newInput)
}
func (cs *crudService) AddAnalisa(input *model.Analisa) error{return cs.analisaRepo.Store(input)}
func (cs *crudService) AddFactor(input *model.Factor) error{return cs.factorRepo.Store(input)}
func (cs *crudService) AddFile(input *model.UploadFile) error{return cs.fileRepo.Store(input)}
func (cs *crudService) AddItem(input *model.Item) error{return cs.itemRepo.Store(input)}
func (cs *crudService) AddMasalah(input *model.Masalah) error{return cs.masalahRepo.Store(input)}
func (cs *crudService) AddMinipap(input *model.MiniPAP) error{return cs.minipapRepo.Store(input)}
func (cs *crudService) AddMonthly(input *model.Monthly) error{
	newInput := input.Reseted()
	return cs.monthlyRepo.Store(&newInput)
}
func (cs *crudService) AddProject(input *model.Project) error{return cs.projectRepo.Store(input)}
func (cs *crudService) AddResult(input *model.Result) error{return cs.resultRepo.Store(input)}
func (cs *crudService) AddSummary(input *model.Summary) error{return cs.summaryRepo.Store(input)}
func (cs *crudService) AddUser(input *model.User) error {return cs.userRepo.Store(input)}
func (cs *crudService) AddYearly(input *model.Yearly) error{return cs.yearlyRepo.Store(input)}


// Update Functions
func (cs *crudService) UpdateAttendance(id int, input model.Attendance) error{
	list, err := cs.attendanceRepo.GetList()
	if err != nil {return err}
	for _, data := range list{
		if data.Year == id{
			input.Year = id
			input.PlanID = data.PlanID
			input.Plan.Monthly_ID = *data.PlanID
			input.ActualID = data.ActualID
			input.Actual.Monthly_ID = *data.ActualID
			input.CutiID = data.CutiID
			input.Cuti.Monthly_ID = *data.CutiID
			input.IzinID = data.IzinID
			input.Izin.Monthly_ID = *data.IzinID
			input.LainID = data.LainID
			input.Lain.Monthly_ID = *data.LainID
			return cs.attendanceRepo.Saves(input)
		}
	}
	return errors.New("not found")
}
func (cs *crudService) UpdateAnalisa(id int, input model.Analisa) error{
	list, err := cs.analisaRepo.GetList()
	if err != nil {return err}
	for _, data := range list{
		if data.Year == id{
			input.Year = id
			if len(data.Masalah) < len(input.Masalah){
				diff := len(input.Masalah) - len(data.Masalah)
				for i := 0; i < diff; i++{
					if err := cs.AddMasalah(&input.Masalah[len(data.Masalah) + i]); err != nil {return err}
				}
			}
			for i, data2 := range data.Masalah{
				if i <= len(input.Masalah) - 1{
					input.Masalah[i].Masalah_ID = data2.Masalah_ID
				}else{
					if err := cs.DeleteMasalah(data2.Masalah_ID); err != nil {return err}
				}
			}
			return cs.analisaRepo.Saves(input)
		}
	}
	return errors.New("not found")
}
func (cs *crudService) UpdateFactor(id int, input model.Factor) error{
	list, err := cs.factorRepo.GetList()
	if err != nil {return err}
	for _, data := range list{
		if data.Factor_ID == id{
			input.Factor_ID = id
			input.PlanID = data.PlanID
			input.Plan.MiniPAP_ID = data.Plan.MiniPAP_ID
			input.ActualID = data.ActualID
			input.Actual.MiniPAP_ID = data.Actual.MiniPAP_ID
			return cs.factorRepo.Saves(input)
		}
	}
	return errors.New("not found")
}
func (cs *crudService) UpdateFile(id int, input model.UploadFile) error{
	newInput := input
	newInput.ID = uint(id)
	return cs.fileRepo.Saves(newInput)
}
func (cs *crudService) UpdateItem(id int, input model.Item) error{
	list, err := cs.itemRepo.GetList()
	if err != nil {return err}
	for _, data := range list{
		if data.Item_ID == id{
			input.Item_ID = id
			return cs.itemRepo.Saves(input)
		}
	}
	return errors.New("not found")
}
func (cs *crudService) UpdateMasalah(id int, input model.Masalah) error{
	list, err := cs.masalahRepo.GetList()
	if err != nil {return err}
	for _, data := range list{
		if data.Masalah_ID == id{
			input.Masalah_ID = id
			return cs.masalahRepo.Saves(input)
		}
	}
	return errors.New("not found")
}
func (cs *crudService) UpdateMinipap(id int, input model.MiniPAP) error{
	list, err := cs.minipapRepo.GetList()
	if err != nil {return err}
	for _, data := range list{
		if data.MiniPAP_ID == id{
			input.MiniPAP_ID = id
			return cs.minipapRepo.Saves(input)
		}
	}
	return errors.New("not found")
}
func (cs *crudService) UpdateMonthly(id int, input model.Monthly) error{
	list, err := cs.monthlyRepo.GetList()
	if err != nil {return err}
	for _, data := range list{
		if data.Monthly_ID == id{
			input.Monthly_ID = id
			return cs.monthlyRepo.Saves(input)
		}
	}
	return errors.New("not found")
}
func (cs *crudService) UpdateProject(id int, input model.Project) error{
	list, err := cs.projectRepo.GetList()
	if err != nil {return err}
	for _, data := range list{
		if data.Project_ID == id{
			input.Project_ID = id
			return cs.projectRepo.Saves(input)
		}
	}
	return errors.New("not found")
}
func (cs *crudService) UpdateResult(id int, input model.Result) error{
	list, err := cs.resultRepo.GetList()
	if err != nil {return err}
	for _, data := range list{
		if data.Result_ID == id{
			input.Result_ID = id
			return cs.resultRepo.Saves(input)
		}
	}
	return errors.New("not found")
}
func (cs *crudService) UpdateSummary(id int, input model.Summary) error{
	list, err := cs.summaryRepo.GetList()
	if err != nil {return err}
	for _, data := range list{
		if data.Summary_ID == id{
			input.Summary_ID = id
			if len(data.Projects) < len(input.Projects){
				diff := len(input.Projects) - len(data.Projects)
				for i := 0; i < diff; i++{
					if err := cs.AddProject(&input.Projects[len(data.Projects) + i]); err != nil {return err}
				}
			}
			for i, data2 := range data.Projects{
				if i <= len(input.Projects) - 1{
					input.Projects[i].Project_ID = data2.Project_ID
				}else{
					if err := cs.DeleteProject(data2.Project_ID); err != nil {return err}
				}
			}
			return cs.summaryRepo.Saves(input)
		}
	}
	return errors.New("not found")
}
func (cs *crudService) UpdateUser(id int, input model.User) error{
	list, err := cs.userRepo.GetList()
	if err != nil {return err}
	for _, data := range list{
		if data.ID == id{
			input.ID = id
			return cs.userRepo.Update(id, input)
		}
	}
	return errors.New("not found")
}
func (cs *crudService) UpdateYearly(id int, input model.Yearly) error{
	list, err := cs.yearlyRepo.GetList()
	if err != nil {return err}
	for _, data := range list{
		if data.Year == id{
			input.Year = id
			return cs.yearlyRepo.Saves(input)
		}
	}
	return errors.New("not found")
}

// Delete Functions
func (cs *crudService) DeleteAttendance(input int) error{
	list, err := cs.attendanceRepo.GetList()
	if err != nil {return err}
	for _, data := range list{
		if data.Year == input{
			return cs.attendanceRepo.Delete(input)
		}
	}
	return errors.New("not found")
}
func (cs *crudService) DeleteAnalisa(input int) error{
	list, err := cs.analisaRepo.GetList()
	if err != nil {return err}
	for _, data := range list{
		if data.Year == input{
			return cs.analisaRepo.Delete(input)
		}
	}
	return errors.New("not found")
}
func (cs *crudService) DeleteFactor(input int) error{
	list, err := cs.factorRepo.GetList()
	if err != nil {return err}
	for _, data := range list{
		if data.Factor_ID == input{
			return cs.factorRepo.Delete(input)
		}
	}
	return errors.New("not found")
}
func (cs *crudService) DeleteFile(input int) error{
	list, err := cs.fileRepo.GetList()
	if err != nil {return err}
	for _, data := range list{
		if int(data.ID) == input{
			return cs.fileRepo.Delete(input)
		}
	}
	return errors.New("not found")
}
func (cs *crudService) DeleteItem(input int) error{
	list, err := cs.itemRepo.GetList()
	if err != nil {return err}
	for _, data := range list{
		if data.Item_ID == input{
			return cs.itemRepo.Delete(input)
		}
	}
	return errors.New("not found")
}
func (cs *crudService) DeleteMasalah(input int) error{
	list, err := cs.masalahRepo.GetList()
	if err != nil {return err}
	for _, data := range list{
		if data.Masalah_ID == input{
			return cs.masalahRepo.Delete(input)
		}
	}
	return errors.New("not found")
}
func (cs *crudService) DeleteMinipap(input int) error{
	list, err := cs.minipapRepo.GetList()
	if err != nil {return err}
	for _, data := range list{
		if data.MiniPAP_ID == input{
			return cs.minipapRepo.Delete(input)
		}
	}
	return errors.New("not found")
}
func (cs *crudService) DeleteMonthly(input int) error{
	list, err := cs.monthlyRepo.GetList()
	if err != nil {return err}
	for _, data := range list{
		if data.Monthly_ID == input{
			return cs.monthlyRepo.Delete(input)
		}
	}
	fmt.Println("Not Found")
	return errors.New("not found")
}
func (cs *crudService) DeleteProject(input int) error{
	list, err := cs.projectRepo.GetList()
	if err != nil {return err}
	for _, data := range list{
		if data.Project_ID == input{
			return cs.projectRepo.Delete(input)
		}
	}
	return errors.New("not found")
}
func (cs *crudService) DeleteResult(input int) error{
	list, err := cs.resultRepo.GetList()
	if err != nil {return err}
	for _, data := range list{
		if data.Result_ID == input{
			return cs.resultRepo.Delete(input)
		}
	}
	return errors.New("not found")
}
func (cs *crudService) DeleteSummary(input int) error{
	list, err := cs.summaryRepo.GetList()
	if err != nil {return err}
	for _, data := range list{
		if data.Summary_ID == input{
			return cs.summaryRepo.Delete(input)
		}
	}
	return errors.New("not found")
}
func (cs *crudService) DeleteUser(input int) error{
	list, err := cs.userRepo.GetList()
	if err != nil {return err}
	for _, data := range list{
		if data.ID == input{
			return cs.userRepo.Delete(input)
		}
	}
	return errors.New("not found")
}
func (cs *crudService) DeleteYearly(input int) error{
	list, err := cs.yearlyRepo.GetList()
	if err != nil {return err}
	for _, data := range list{
		if data.Year == input{
			return cs.yearlyRepo.Delete(input)
		}
	}
	return errors.New("not found")
}

// Read Specified Functions
func (cs *crudService) GetAttendanceByID(input int) (*model.AttendanceResponse, error){
	tempInput, err := cs.attendanceRepo.GetByID(input)
	if err != nil {return nil, err}
	newInput := tempInput.ToResponse()
	return &newInput, err
}
func (cs *crudService) GetAnalisaByID(input int) (*model.Analisa, error){return cs.analisaRepo.GetByID(input)}
func (cs *crudService) GetFactorByID(input int) (*model.FactorResponse, error){
	tempInput, err := cs.factorRepo.GetByID(input)
	if err != nil {return nil, err}
	newInput := tempInput.ToResponse()
	return &newInput, err
}
func (cs *crudService) GetFileByID(input int) (*model.UploadFile, error){return cs.fileRepo.GetByID(input)}
func (cs *crudService) GetItemByID(input int) (*model.ItemResponse, error){
	tempInput, err := cs.itemRepo.GetByID(input)
	if err != nil {return nil, err}
	newInput := tempInput.ToResponse()
	return &newInput, err
}
func (cs *crudService) GetMasalahByID(input int) (*model.Masalah, error){return cs.masalahRepo.GetByID(input)}
func (cs *crudService) GetMinipapByID(input int) (*model.MiniPAP, error){return cs.minipapRepo.GetByID(input)}
func (cs *crudService) GetMonthlyByID(input int) (*model.Monthly, error){return cs.monthlyRepo.GetByID(input)}
func (cs *crudService) GetProjectByID(input int) (*model.Project, error){return cs.projectRepo.GetByID(input)}
func (cs *crudService) GetResultByID(input int) (*model.ResultResponse, error){
	tempInput, err := cs.resultRepo.GetByID(input)
	if err != nil {return nil, err}
	newInput := tempInput.ToResponse()
	return &newInput, err
}
func (cs *crudService) GetSummaryByID(input int) (*model.Summary, error){return cs.summaryRepo.GetByID(input)}
func (cs *crudService) GetUserByID(input int) (*model.UserResponse, error){
	tempInput, err := cs.userRepo.GetByID(input)
	if err != nil {return nil, err}
	newInput := tempInput.ToResponse()
	return &newInput, err
}
func (cs *crudService) GetYearlyByID(input int) (*model.YearlyResponse, error){
	tempInput, err := cs.yearlyRepo.GetByID(input)
	if err != nil {return nil, err}
	newInput := tempInput.ToResponse()
	return &newInput, err
}

// Read All Functions
func (cs *crudService) GetAttendanceList() (model.AttendanceArrayResponse, error){
	tempInput, err := cs.attendanceRepo.GetList()
	var newInput model.AttendanceArrayResponse
	newInput.Attendance = []model.AttendanceResponse{}
	for _, temp := range tempInput{
		newInput.Attendance = append(newInput.Attendance, temp.ToResponse())	
	}
	return newInput, err
}
func (cs *crudService) GetAnalisaList() (model.AnalisaArrayResponse, error){
	tempInput, err := cs.analisaRepo.GetList()
	var newInput model.AnalisaArrayResponse
	newInput.Analisa = append(newInput.Analisa, tempInput...)
	return newInput, err
}
func (cs *crudService) GetFactorList() (model.FactorArrayResponse, error){
	tempInput, err :=  cs.factorRepo.GetList()
	var newInput model.FactorArrayResponse
	newInput.Factor = []model.FactorResponse{}
	for _, temp := range tempInput{
		newInput.Factor = append(newInput.Factor, temp.ToResponse())	
	}
	return newInput, err
}
func (cs *crudService) GetFileList() ([]model.UploadFile, error){return cs.fileRepo.GetList()}
func (cs *crudService) GetItemList() (model.ItemArrayResponse, error){
	tempInput, err := cs.itemRepo.GetList()
	var newInput model.ItemArrayResponse
	newInput.Item = []model.ItemResponse{}
	for _, temp := range tempInput{
		newInput.Item = append(newInput.Item, temp.ToResponse())	
	}
	return newInput, err
}
func (cs *crudService) GetMasalahList() (model.MasalahArrayResponse, error){
	tempInput, err := cs.masalahRepo.GetList()
	var newInput model.MasalahArrayResponse
	newInput.Masalah = []model.Masalah{}
	newInput.Masalah = append(newInput.Masalah, tempInput...)	
	return newInput, err
}
func (cs *crudService) GetMinipapList() (model.MinipapArrayResponse, error){
	tempInput, err := cs.minipapRepo.GetList()
	var newInput model.MinipapArrayResponse
	newInput.Minipap = append(newInput.Minipap, tempInput...)
	return newInput, err
}
func (cs *crudService) GetMonthlyList() (model.MonthlyArrayResponse, error){
	tempInput, err := cs.monthlyRepo.GetList()
	var newInput model.MonthlyArrayResponse
	newInput.Monthly = append(newInput.Monthly, tempInput...)
	return newInput, err
}
func (cs *crudService) GetProjectList() (model.ProjectArrayResponse, error){
	tempInput, err := cs.projectRepo.GetList()
	var newInput model.ProjectArrayResponse
	newInput.Project = []model.Project{}
	newInput.Project = append(newInput.Project, tempInput...)
	return newInput, err
}
func (cs *crudService) GetResultList() (model.ResultArrayResponse, error){
	tempInput, err := cs.resultRepo.GetList()
	var newInput model.ResultArrayResponse
	newInput.Result = []model.ResultResponse{}
	for _, temp := range tempInput{
		newInput.Result = append(newInput.Result, temp.ToResponse())	
	}
	return newInput, err
}
func (cs *crudService) GetSummaryList() (model.SummaryArrayResponse, error){
	tempInput, err := cs.summaryRepo.GetList()
	var newInput model.SummaryArrayResponse
	newInput.Summary = []model.Summary{}
	newInput.Summary = append(newInput.Summary, tempInput...)	
	return newInput, err
}
func (cs *crudService) GetUserList() (model.UserArrayResponse, error){
	tempInput, err := cs.userRepo.GetList()
	var newInput model.UserArrayResponse
	newInput.Users = []model.UserResponse{}
	for _, temp := range tempInput{
		newInput.Users = append(newInput.Users, temp.ToResponse())	
	}
	return newInput, err
}
func (cs *crudService) GetYearlyList() (model.YearlyArrayResponse, error){
	tempInput, err := cs.yearlyRepo.GetList()
	var newInput model.YearlyArrayResponse
	newInput.Yearly = []model.YearlyResponse{}
	for _, temp := range tempInput{
		newInput.Yearly = append(newInput.Yearly, temp.ToResponse())	
	}
	return newInput, err
}