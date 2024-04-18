package service

import (
	"goreact/model"
	repo "goreact/repository"
)

type CrudService interface {
	//Create independent
	AddAttendance(input *model.Attendance) error
	AddAnalisa(input *model.Analisa) error
	AddFactor(input *model.Factor) error
	AddFile(input *model.UploadFile) error
	AddItem(input *model.Item) error
	AddMasalah(input *model.Masalah) error
	AddMinipap(input *model.MiniPAP) error
	AddMonthly(input *model.Monthly) error
	AddProject(input *model.Project) error
	AddResult(input *model.Result) error
	AddSummary(input *model.Summary) error
	AddYearly(input *model.Yearly) error

	//Create cascade
	AddEntireYearly()
	AddEntireItem()
	AddEntireResult()
	AddEntireFactor()
	AddEntireAttendance()

	//Update
	UpdateAttendance(id int, input model.Attendance) error
	UpdateAnalisa(id int, input model.Analisa) error
	UpdateFactor(id int, input model.Factor) error
	UpdateFile(id int, input model.UploadFile) error
	UpdateItem(id int, input model.Item) error
	UpdateMasalah(id int, input model.Masalah) error
	UpdateMinipap(id int, input model.MiniPAP) error
	UpdateMonthly(id int, input model.Monthly) error
	UpdateProject(id int, input model.Project) error
	UpdateResult(id int, input model.Result) error
	UpdateSummary(id int, input model.Summary) error
	UpdateYearly(id int, input model.Yearly) error


	//Delete independent
	DeleteAttendance(input int) error
	DeleteAnalisa(input int) error
	DeleteFile(input int) error
	DeleteFactor(input int) error
	DeleteItem(input int) error
	DeleteMasalah(input int) error
	DeleteMinipap(input int) error
	DeleteMonthly(input int) error
	DeleteProject(input int) error
	DeleteResult(input int) error
	DeleteSummary(input int) error
	DeleteYearly(input int) error


	//Delete cascade
	DeleteEntireYearly()
	DeleteEntireItem()
	DeleteEntireResult()
	DeleteEntireFactor()
	DeleteEntireAttendance()

	//Get specified
	GetAttendanceByID(input int) (*model.AttendanceResponse, error)
	GetFactorByID(input int) (*model.FactorResponse, error)
	GetItemByID(input int) (*model.ItemResponse, error)
	GetMinipapByID(input int) (*model.MiniPAP, error)
	GetMonthlyByID(input int) (*model.Monthly, error)
	GetResultByID(input int) (*model.ResultResponse, error)
	GetYearlyByID(input int) (*model.YearlyResponse, error)
	GetAnalisaByID(input int) (*model.Analisa, error)
	GetMasalahByID(input int) (*model.MasalahResponse, error)
	GetProjectByID(input int) (*model.ProjectResponse, error)
	GetSummaryByID(input int) (*model.SummaryResponse, error)
	GetFileByID(input int) (*model.UploadFile, error)

	//Get batch
	GetAttendanceList()(model.AttendanceArrayResponse, error)
	GetFactorList()(model.FactorArrayResponse, error)
	GetItemList()(model.ItemArrayResponse, error)
	GetMinipapList()(model.MinipapArrayResponse, error)
	GetMonthlyList()(model.MonthlyArrayResponse, error)
	GetResultList()(model.ResultArrayResponse, error)
	GetYearlyList()(model.YearlyArrayResponse, error)
	GetAnalisaList()(model.AnalisaArrayResponse, error)
	GetMasalahList()(model.MasalahArrayResponse, error)
	GetProjectList()(model.ProjectArrayResponse, error)
	GetSummaryList()(model.SummaryArrayResponse, error)
	GetFileList()([]model.UploadFile, error)
}

type ParseService interface {
	ParseYearly()
	ParseItem()
	ParseResult()
	ParseFactor()
	ParseAttendance()

	ParseAnalisa()
	ParseMasalah()

	ParseSummary()
	ParseProject()
}

type crudService struct {
	attendanceRepo 	repo.AttendanceRepo
	analisaRepo		repo.AnalisaRepo
	factorRepo     	repo.FactorRepo
	fileRepo 		repo.FileRepo
	itemRepo       	repo.ItemRepo
	masalahRepo 	repo.MasalahRepo
	minipapRepo    	repo.MiniPAPRepo
	monthlyRepo    	repo.MonthlyRepo
	projectRepo		repo.ProjectRepo
	resultRepo     	repo.ResultRepo
	summaryRepo 	repo.SummaryRepo
	yearlyRepo     	repo.YearlyRepo
}

func NewCrudService(
	attendanceRepo 	repo.AttendanceRepo,
	analisaRepo 	repo.AnalisaRepo,
	factorRepo 		repo.FactorRepo,
	fileRepo		repo.FileRepo,
	itemRepo 		repo.ItemRepo,
	masalahRepo 	repo.MasalahRepo,
	minipapRepo 	repo.MiniPAPRepo,
	monthlyRepo 	repo.MonthlyRepo,
	projectRepo		repo.ProjectRepo,
	resultRepo 		repo.ResultRepo,
	summaryRepo 	repo.SummaryRepo,
	yearlyRepo 		repo.YearlyRepo,) *crudService {
	return &crudService{
		attendanceRepo,
		analisaRepo,
		factorRepo,
		fileRepo,
		itemRepo,
		masalahRepo,
		minipapRepo,
		monthlyRepo,
		projectRepo,
		resultRepo,
		summaryRepo,
		yearlyRepo}
}