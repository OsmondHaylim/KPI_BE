package api

import (

	"fmt"
	"goreact/model"
	"goreact/service"
	"net/http"
	"path/filepath"
	"os"
	"io"
// "strconv"
	"time"
	"strings"
	"github.com/gin-gonic/gin"
// "github.com/tealeg/xlsx"
)

type FileAPI interface{
	FileUpload(f *gin.Context)
}

type fileAPI struct{
	fileService		service.FileService
}

func NewFileAPI(fileService service.FileService) *fileAPI{
	return &fileAPI{
		fileService,
	}
}

func (fa *fileAPI) FileUpload(f *gin.Context) {
	file, header, err := f.Request.FormFile("file")
	if err != nil{
		f.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	defer file.Close()
	fileExt := filepath.Ext(header.Filename)
    originalFileName := strings.TrimSuffix(filepath.Base(header.Filename), filepath.Ext(header.Filename))
    now := time.Now()
    filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
    filePath := "http://localhost:8080/files/" + filename

	out, err := os.Create("public/excel/" + filename)
	if err != nil {
		f.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		f.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	f.JSON(http.StatusOK, gin.H{"filename": filename, "filepath": filePath})}