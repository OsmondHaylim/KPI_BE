package api

import (

	// "fmt"
	"goreact/model"
	"goreact/service"
	"net/http"
	"regexp"
	// "path/filepath"
	"io"
	"os"
	"strconv"
	// "time"
	// "strings"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

type FileAPI interface {
	FileUpload(f *gin.Context)
}

type fileAPI struct {
	fileService service.FileService
}

func NewFileAPI(fileService service.FileService) *fileAPI {
	return &fileAPI{
		fileService,
	}
}

func (fa *fileAPI) FileUpload(f *gin.Context) {
	//get File
	file, header, err := f.Request.FormFile("file")
	if err != nil {
		f.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	defer file.Close()
	//Check format
	if header.Header.Get("Content-Type") != "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
		f.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "not xlsx"})
		return
	}
	//Create temp file
	out, err := os.CreateTemp("", "upload-*.xlsx")
	if err != nil {
		f.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	defer out.Close()
	//Copy File
	_, err = io.Copy(out, file)
	if err != nil {
		f.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	//Create Directory
	directory := "./uploads/"
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		os.Mkdir(directory, os.ModePerm)
	}

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

	//Open Excel
	excel, err := xlsx.OpenFile(out.Name())
	if err != nil {
		f.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	test := ""
	for _, sheet := range excel.Sheets {
		if sheet.Name == "KPI staff Design" {
			re := regexp.MustCompile(`(\d{4})$`)
			match := re.FindStringSubmatch(sheet.Rows[3].Cells[2].String())
			year := 0
			if len(match) == 2 {
				year, err = strconv.Atoi(match[1])
				if err != nil{
					f.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Failed extracting year"})
					return
				}
			}
			test += strconv.Itoa(year) + " "
			// KPI := model.Yearly{
			// 	Year: year,
			// }
			// tempItem := model.Item{}
			// tempResult := model.Result{}
			// tempFactor := model.Factor{}

			content := false
			test += "/" + strconv.Itoa(len(sheet.Rows)) + "/"
			for _, row := range sheet.Rows {
				if (!content && row.Cells[0].String() == "Item"){
					content = true
					continue
				}
				if content {
					if len(row.Cells) != 0 && row.Cells[0].String() != ""{
						// tempItem.Name = row.Cells[0].String()
						test += row.Cells[0].String()
					}
				}
				// test += "[" + strconv.Itoa(len(row.Cells)) + "]"
				
			}
		// f.JSON(http.StatusOK, KPI)	
		// return
		}
	}
	// f.JSON(http.StatusOK, model.SuccessResponse{Message: strconv.Itoa(test)})
	f.JSON(http.StatusOK, model.SuccessResponse{Message: test})
	
}
