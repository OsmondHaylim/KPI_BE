package api

import (
	"goreact/model"
	"goreact/service"
	"net/http"
	// "path/filepath"
	// "io"
	"os"
	// "strconv"

	// "time"
	// "strings"
	"github.com/gin-gonic/gin"
)

type FileAPI interface {
	KpiFileUpload(f *gin.Context)
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

func (fa *fileAPI) KpiFileUpload(f *gin.Context) {
	//get File
	file, header, err := f.Request.FormFile("file")
	if model.ErrorCheck(f, err){return}
	defer file.Close()
	//Check format
	if header.Header.Get("Content-Type") != "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
		f.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Not xlsx"})
		return
	}
	//Create temp file
	out, err := os.CreateTemp("", "upload-*.xlsx")
	if model.ErrorCheck(f, err){return}
	defer out.Close()
	// //Copy File
	// _, err = io.Copy(out, file)
	// if model.ErrorCheck(f, err){return}

	// //Create Directory
	// directory := "./uploads/"
	// if _, err := os.Stat(directory); os.IsNotExist(err) {
	// 	os.Mkdir(directory, os.ModePerm)
	// }

	// // Create a new file in the specified directory
	// out, err = os.Create(directory + header.Filename)
	// if err != nil {
	// 	f.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
	// 	return
	// }
	// defer out.Close()

	// // Copy the uploaded file to the newly created file
	// _, err = io.Copy(out, file)
	// if err != nil {
	// 	f.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
	// 	return
	// }

	input, err := fa.parseService.ParseKpi(out)
	if model.ErrorCheck(f, err) {return}

	err = fa.crudService.AddEntireYearly(input)
	if model.ErrorCheck(f, err) {return}

	f.JSON(http.StatusAccepted, model.SuccessResponse{Message: "Data Received"})
}
