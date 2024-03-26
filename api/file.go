package api

import (

	// "fmt"
	"goreact/model"
	"goreact/service"
	"net/http"
	// "path/filepath"
	"io"
	"os"
	// "strconv"
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
			for i, row := range sheet.Rows {
				if i > 100 {
					break
				}
				if i < 14 {
					continue
				}
				test += row.Cells[0].String()
				// Assuming your data has some structure, you can process each row accordingly
				// For example, if you have columns 'name', 'age', 'email' in your XLSX
				// You can access them as row.Cells[index].Value
				// For example:
				// name := row.Cells[0].String()
				// age := row.Cells[1].Int()
				// email := row.Cells[2].String()

				// Now, you can use GORM to insert this data into your table
				// Example:
				// db.Create(&YourModel{...})
			}
		}
	}
	f.JSON(http.StatusOK, model.SuccessResponse{Message: test})
}
