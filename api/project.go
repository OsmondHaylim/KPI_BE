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
	newProject.Project_ID = KpiID
	err = aa.crudService.UpdateProject(newProject)
	if model.ErrorCheck(k, err){return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Project update success"})
}
func (aa *projectAPI) UpdateSummary(k *gin.Context) {
	var newSummary model.Summary
	err := k.ShouldBindJSON(&newSummary)
	if model.ErrorCheck(k, err){return}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	newSummary.Summary_ID = KpiID
	err = aa.crudService.UpdateSummary(newSummary)	
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
	Project, err := aa.crudService.GetByID(KpiID)
	if model.ErrorCheck(k, err){return}
	k.JSON(http.StatusOK, Project)
}
func (aa *projectAPI) GetSummaryByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	Summary, err := aa.crudService.GetByID(KpiID)
	if model.ErrorCheck(k, err){return}
	// Result := Summary.ToResponse()
	k.JSON(http.StatusOK, Summary)
}

func (aa *projectAPI) GetProjectList(k *gin.Context) {
	Project, err := aa.crudService.GetList()
	if model.ErrorCheck(k, err){return}
	var result model.ProjectArrayResponse
	result.Project = []model.ProjectResponse{}
	for _, data := range Project{
		result.Project = append(result.Project, data.ToResponse())
	}
	result.Message = "Getting All Projects Success"
	k.JSON(http.StatusOK, result)
}
func (aa *projectAPI) GetSummaryList(k *gin.Context) {
	Summary, err := aa.crudService.GetList()
	if model.ErrorCheck(k, err){return}
	var result model.SummaryArrayResponse
	result.Summary = []model.SummaryResponse{}
	for _, data := range Summary{
		result.Summary = append(result.Summary, data.ToResponse())
	} 
	result.Message = "Getting All Summarys Success"
	k.JSON(http.StatusOK, result)
}

