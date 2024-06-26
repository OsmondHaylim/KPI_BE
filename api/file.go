package api

import (
	"goreact/model"
	"goreact/service"
	"net/http"
	// "path/filepath"
	// "io"
	// "os"
	// "strconv"

	// "time"
	// "strings"
	"github.com/gin-gonic/gin"
)

type FileAPI interface {
	SaveFile(f *gin.Context)
	KpiFileUpload(f *gin.Context)
	AnalisaFileUpload(f *gin.Context)
	SummaryFileUpload(f *gin.Context)
}

type fileAPI struct {
	crudService service.CrudService
	parseService service.ParseService
}

func NewFileAPI(crudService service.CrudService, parseService service.ParseService) *fileAPI {
	return &fileAPI{
		crudService,
		parseService,
	}
}

func (fa *fileAPI) SaveFile(f *gin.Context) {
	file, header, err := f.Request.FormFile("file")
	if model.ErrorCheck(f, err){return}
	defer file.Close()
	_, err = fa.parseService.SaveFile(file, header)
	if model.ErrorCheck(f, err) {return}
	f.JSON(http.StatusAccepted, model.SuccessResponse{Message: "File Saved"})
}

func (fa *fileAPI) KpiFileUpload(f *gin.Context) {
	file, header, err := f.Request.FormFile("file")
	if model.ErrorCheck(f, err){return}
	defer file.Close()
	if header.Header.Get("Content-Type") != "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
		f.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Not xlsx"})
		return
	}
	input, err := fa.parseService.ParseKpi(file)
	if model.ErrorCheck(f, err) {return}
	if model.ErrorCheck(f, fa.crudService.AddEntireYearly(input)) {return}
	// f.JSON(http.StatusAccepted, model.SuccessResponse{Message: "Inputted"})
	f.JSON(http.StatusAccepted, input)
}

func (fa *fileAPI) AnalisaFileUpload(f *gin.Context) {
	file, header, err := f.Request.FormFile("file")
	if model.ErrorCheck(f, err){return}
	defer file.Close()
	if header.Header.Get("Content-Type") != "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
		f.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Not xlsx"})
		return
	}
	input, err := fa.parseService.ParseAnalisis(file)
	if model.ErrorCheck(f, err) {return}
	if model.ErrorCheck(f, fa.crudService.AddEntireAnalisa(input)) {return}
	// f.JSON(http.StatusAccepted, model.SuccessResponse{Message: "Inputted"})
	f.JSON(http.StatusAccepted, input)
}

func (fa *fileAPI) SummaryFileUpload(f *gin.Context) {
	file, header, err := f.Request.FormFile("file")
	if model.ErrorCheck(f, err){return}
	defer file.Close()
	if header.Header.Get("Content-Type") != "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
		f.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Not xlsx"})
		return
	}
	input, err := fa.parseService.ParseSummary(file)
	if model.ErrorCheck(f, err) {return}
	if model.ErrorCheck(f, fa.crudService.AddEntireSummary(input)) {return}
	// f.JSON(http.StatusAccepted, model.SuccessResponse{Message: "Inputted"})
	f.JSON(http.StatusAccepted, input)
}

