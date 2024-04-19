package service

import (
	"goreact/model"
)

// Create Functions
func (cs *crudService) AddAttendance(input *model.Attendance) error{return cs.attendanceRepo.Store(input)}
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
func (cs *crudService) AddYearly(input *model.Yearly) error{return cs.yearlyRepo.Store(input)}


// Update Functions
func (cs *crudService) UpdateAttendance(id int, input model.Attendance) error{
	newInput := input
	newInput.Year = id
	return cs.attendanceRepo.Saves(newInput)
}
func (cs *crudService) UpdateAnalisa(id int, input model.Analisa) error{
	newInput := input
	newInput.Year = id
	return cs.analisaRepo.Saves(newInput)
}
func (cs *crudService) UpdateFactor(id int, input model.Factor) error{
	newInput := input
	newInput.Factor_ID = id
	return cs.factorRepo.Saves(newInput)
}
func (cs *crudService) UpdateFile(id int, input model.UploadFile) error{
	newInput := input
	newInput.ID = uint(id)
	return cs.fileRepo.Saves(newInput)
}
func (cs *crudService) UpdateItem(id int, input model.Item) error{
	newInput := input
	newInput.Item_ID = id
	return cs.itemRepo.Saves(newInput)
}
func (cs *crudService) UpdateMasalah(id int, input model.Masalah) error{
	newInput := input
	newInput.Masalah_ID = id
	return cs.masalahRepo.Saves(newInput)
}
func (cs *crudService) UpdateMinipap(id int, input model.MiniPAP) error{
	newInput := input
	newInput.MiniPAP_ID = id
	return cs.minipapRepo.Saves(newInput)
}
func (cs *crudService) UpdateMonthly(id int, input model.Monthly) error{
	newInput := input
	newInput.Monthly_ID = id
	return cs.monthlyRepo.Saves(newInput)
}
func (cs *crudService) UpdateProject(id int, input model.Project) error{
	newInput := input
	newInput.Project_ID = id
	return cs.projectRepo.Saves(newInput)
}
func (cs *crudService) UpdateResult(id int, input model.Result) error{
	newInput := input
	newInput.Result_ID = id
	return cs.resultRepo.Saves(newInput)
}
func (cs *crudService) UpdateSummary(id int, input model.Summary) error{
	newInput := input
	newInput.Summary_ID = id
	return cs.summaryRepo.Saves(newInput)
}
func (cs *crudService) UpdateYearly(id int, input model.Yearly) error{
	newInput := input
	newInput.Year = id
	return cs.yearlyRepo.Saves(newInput)
}

// Delete Functions
func (cs *crudService) DeleteAttendance(input int) error{return cs.attendanceRepo.Delete(input)}
func (cs *crudService) DeleteAnalisa(input int) error{return cs.analisaRepo.Delete(input)}
func (cs *crudService) DeleteFactor(input int) error{return cs.factorRepo.Delete(input)}
func (cs *crudService) DeleteFile(input int) error{return cs.fileRepo.Delete(input)}
func (cs *crudService) DeleteItem(input int) error{return cs.itemRepo.Delete(input)}
func (cs *crudService) DeleteMasalah(input int) error{return cs.masalahRepo.Delete(input)}
func (cs *crudService) DeleteMinipap(input int) error{return cs.minipapRepo.Delete(input)}
func (cs *crudService) DeleteMonthly(input int) error{return cs.monthlyRepo.Delete(input)}
func (cs *crudService) DeleteProject(input int) error{return cs.projectRepo.Delete(input)}
func (cs *crudService) DeleteResult(input int) error{return cs.resultRepo.Delete(input)}
func (cs *crudService) DeleteSummary(input int) error{return cs.summaryRepo.Delete(input)}
func (cs *crudService) DeleteYearly(input int) error{return cs.yearlyRepo.Delete(input)}

// Read Specified Functions
func (cs *crudService) GetAttendanceByID(input int) (*model.AttendanceResponse, error){
	tempInput, err := cs.attendanceRepo.GetByID(input)
	newInput := tempInput.ToResponse()
	return &newInput, err
}
func (cs *crudService) GetAnalisaByID(input int) (*model.Analisa, error){return cs.analisaRepo.GetByID(input)}
func (cs *crudService) GetFactorByID(input int) (*model.FactorResponse, error){
	tempInput, err := cs.factorRepo.GetByID(input)
	newInput := tempInput.ToResponse()
	return &newInput, err
}
func (cs *crudService) GetFileByID(input int) (*model.UploadFile, error){return cs.fileRepo.GetByID(input)}
func (cs *crudService) GetItemByID(input int) (*model.ItemResponse, error){
	tempInput, err := cs.itemRepo.GetByID(input)
	newInput := tempInput.ToResponse()
	return &newInput, err
}
func (cs *crudService) GetMasalahByID(input int) (*model.MasalahResponse, error){
	tempInput, err := cs.masalahRepo.GetByID(input)
	newInput := tempInput.ToResponse()
	return &newInput, err
}
func (cs *crudService) GetMinipapByID(input int) (*model.MiniPAP, error){return cs.minipapRepo.GetByID(input)}
func (cs *crudService) GetMonthlyByID(input int) (*model.Monthly, error){return cs.monthlyRepo.GetByID(input)}
func (cs *crudService) GetProjectByID(input int) (*model.ProjectResponse, error){
	tempInput, err := cs.projectRepo.GetByID(input)
	newInput := tempInput.ToResponse()
	return &newInput, err
}
func (cs *crudService) GetResultByID(input int) (*model.ResultResponse, error){
	tempInput, err := cs.resultRepo.GetByID(input)
	newInput := tempInput.ToResponse()
	return &newInput, err
}
func (cs *crudService) GetSummaryByID(input int) (*model.SummaryResponse, error){
	tempInput, err := cs.summaryRepo.GetByID(input)
	newInput := tempInput.ToResponse()
	return &newInput, err
}
func (cs *crudService) GetYearlyByID(input int) (*model.YearlyResponse, error){
	tempInput, err := cs.yearlyRepo.GetByID(input)
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
	newInput.Masalah = []model.MasalahResponse{}
	for _, temp := range tempInput{
		newInput.Masalah = append(newInput.Masalah, temp.ToResponse())	
	}
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
	newInput.Project = []model.ProjectResponse{}
	for _, temp := range tempInput{
		newInput.Project = append(newInput.Project, temp.ToResponse())	
	}
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
	newInput.Summary = []model.SummaryResponse{}
	for _, temp := range tempInput{
		newInput.Summary = append(newInput.Summary, temp.ToResponse())	
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