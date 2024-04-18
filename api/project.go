package api

import (
	// "fmt"
	"goreact/model"
	"goreact/service"
	"net/http"
	"strconv"
	// "strings"

	"github.com/gin-gonic/gin"
)

type ProjectAPI interface {
	AddProject(k *gin.Context)
	AddSummary(k *gin.Context)

	UpdateProject(k *gin.Context)
	UpdateSummary(k *gin.Context)

	DeleteProject(k *gin.Context)
	DeleteSummary(k *gin.Context)

	GetProjectByID(k *gin.Context)
	GetSummaryByID(k *gin.Context)

	GetProjectList(k *gin.Context)
	GetSummaryList(k *gin.Context)
}
type projectAPI struct {crudService		service.CrudService}
func NewProjectAPI(crudService service.CrudService) *projectAPI{
	return &projectAPI{crudService,}}

func (aa *projectAPI) AddProject(k *gin.Context) {
	var newProject model.Project
	err := k.ShouldBindJSON(&newProject)
	if model.ErrorCheck(k, err){return}
	err = aa.crudService.AddProject(&newProject)
	if model.ErrorCheck(k, err){return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Project success"})
}
func (aa *projectAPI) AddSummary(k *gin.Context) {
	var newSummary model.Summary
	err := k.ShouldBindJSON(&newSummary)
	if model.ErrorCheck(k, err){return}
	err = aa.crudService.AddSummary(&newSummary)
	if model.ErrorCheck(k, err){return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Summary success"})
}

func (aa *projectAPI) UpdateProject(k *gin.Context) {
	var newProject model.Project
	err := k.ShouldBindJSON(&newProject)
	if model.ErrorCheck(k, err){return}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	err = aa.crudService.UpdateProject(KpiID, newProject)
	if model.ErrorCheck(k, err){return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Project update success"})
}
func (aa *projectAPI) UpdateSummary(k *gin.Context) {
	var newSummary model.Summary
	err := k.ShouldBindJSON(&newSummary)
	if model.ErrorCheck(k, err){return}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	err = aa.crudService.UpdateSummary(KpiID, newSummary)	
	if model.ErrorCheck(k, err){return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Summary update success"})
}

func (aa *projectAPI) DeleteProject(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	err = aa.crudService.DeleteProject(KpiID)
	if model.ErrorCheck(k, err){return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Project delete success"})
}
func (aa *projectAPI) DeleteSummary(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	err = aa.crudService.DeleteSummary(KpiID)
	if model.ErrorCheck(k, err){return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Summary delete success"})
}

func (aa *projectAPI) GetProjectByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	Project, err := aa.crudService.GetProjectByID(KpiID)
	if model.ErrorCheck(k, err){return}
	k.JSON(http.StatusOK, Project)
}
func (aa *projectAPI) GetSummaryByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	Summary, err := aa.crudService.GetSummaryByID(KpiID)
	if model.ErrorCheck(k, err){return}
	// Result := Summary.ToResponse()
	k.JSON(http.StatusOK, Summary)
}

func (aa *projectAPI) GetProjectList(k *gin.Context) {
	Project, err := aa.crudService.GetProjectList()
	if model.ErrorCheck(k, err){return}
	Project.Message = "Getting All Projects Success"
	k.JSON(http.StatusOK, Project)
}
func (aa *projectAPI) GetSummaryList(k *gin.Context) {
	Summary, err := aa.crudService.GetSummaryList()
	if model.ErrorCheck(k, err){return}
	Summary.Message = "Getting All Summarys Success"
	k.JSON(http.StatusOK, Summary)
}

