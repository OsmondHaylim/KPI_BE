package service

import (
	"goreact/model"
	repo "goreact/repository"
	"mime/multipart"
)

type CrudService interface {
	//Create Independent
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
	AddEntireYearly(input *model.YearlyResponse) error
	AddEntireItem(input *model.ItemResponse, id *int) error
	AddEntireResult(input *model.ResultResponse, id *int) error
	AddEntireFactor(input *model.FactorResponse, id *int) error
	AddEntireAttendance(input *model.AttendanceResponse, year *int) error
	AddEntireAnalisa(input *model.AnalisaResponse) error
	AddEntireSummary(input *model.SummaryResponse) error

	//Update Independent
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

	// // Update Cascade
	// UpdateEntireAttendance(id int, input model.AttendanceResponse) error
	// UpdateEntireAnalisa(id int, input model.AnalisaResponse) error
	// UpdateEntireFactor(id int, input model.FactorResponse) error
	// UpdateEntireItem(id int, input model.ItemResponse) error
	// UpdateEntireResult(id int, input model.ResultResponse) error
	// UpdateEntireSummary(id int, input model.SummaryResponse) error
	// UpdateEntireYearly(id int, input model.YearlyResponse) error

	//Delete Independent
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


	//Delete Cascade
	DeleteEntireYearly(input int) error
	DeleteEntireItem(input int) error
	DeleteEntireResult(input int) error
	DeleteEntireFactor(input int) error
	DeleteEntireAttendance(input int) error
	DeleteEntireAnalisa(input int) error
	DeleteEntireSummary(input int) error

	//Get Specified
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

	//Get Batch
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

type ParseService interface {
	ParseKpi(input multipart.File) (*model.YearlyResponse, error)
	SaveFile(input multipart.File, header *multipart.FileHeader) (*model.UploadFile, error)
}
type parseService struct {
	fileRepo repo.FileRepo
}
func NewParseService(fileRepo repo.FileRepo) *parseService{
	return &parseService{
		fileRepo,
	}
}