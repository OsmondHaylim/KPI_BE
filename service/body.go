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
	UpdateAttendance(input model.Attendance) error
	UpdateAnalisa(input model.Analisa) error
	UpdateFactor(input model.Factor) error
	UpdateFile(input model.UploadFile) error
	UpdateItem(input model.Item) error
	UpdateMasalah(input model.Masalah) error
	UpdateMinipap(input model.MiniPAP) error
	UpdateMonthly(input model.Monthly) error
	UpdateProject(input model.Project) error
	UpdateResult(input model.Result) error
	UpdateSummary(input model.Summary) error
	UpdateYearly(input model.Yearly) error


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
	GetAttendanceByID()
	GetFactorByID()
	GetItemByID()
	GetMinipapByID()
	GetMonthlyByID()
	GetResultByID()
	GetYearlyByID()
	GetAnalisaByID()
	GetMasalahByID()
	GetProjectByID()
	GetSummaryByID()
	GetFileByID()

	//Get batch
	GetAttendanceList()
	GetFactorList()
	GetItemList()
	GetMinipapList()
	GetMonthlyList()
	GetResultList()
	GetYearlyList()
	GetAnalisaList()
	GetMasalahList()
	GetProjectList()
	GetSummaryList()
	GetFileList()
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
	factorRepo     	repo.FactorRepo
	itemRepo       	repo.ItemRepo
	minipapRepo    	repo.MiniPAPRepo
	monthlyRepo    	repo.MonthlyRepo
	resultRepo     	repo.ResultRepo
	yearlyRepo     	repo.YearlyRepo
	projectRepo		repo.ProjectRepo
	summaryRepo 	repo.SummaryRepo
	analisaRepo		repo.AnalisaRepo
	masalahRepo 	repo.MasalahRepo
	fileRepo 		repo.FileRepo
}

func NewCrudService(
	attendanceRepo 	repo.AttendanceRepo,
	factorRepo 		repo.FactorRepo,
	itemRepo 		repo.ItemRepo,
	minipapRepo 	repo.MiniPAPRepo,
	monthlyRepo 	repo.MonthlyRepo,
	resultRepo 		repo.ResultRepo,
	yearlyRepo 		repo.YearlyRepo,
	projectRepo		repo.ProjectRepo,
	summaryRepo 	repo.SummaryRepo,
	analisaRepo 	repo.AnalisaRepo,
	masalahRepo 	repo.MasalahRepo,
	fileRepo		repo.FileRepo) *crudService {
	return &crudService{
		attendanceRepo,
		factorRepo,
		itemRepo,
		minipapRepo,
		monthlyRepo,
		resultRepo,
		yearlyRepo,
		projectRepo,
		summaryRepo,
		analisaRepo,
		masalahRepo,
		fileRepo}
}