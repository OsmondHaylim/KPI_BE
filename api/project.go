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
type projectAPI struct {
	projectService		service.ProjectService
	summaryService 		service.SummaryService
}
func NewProjectAPI(
	projectService service.ProjectService,
	summaryService service.SummaryService) *projectAPI{
	return &projectAPI{
		projectService,
		summaryService}
}

func (aa *projectAPI) AddProject(k *gin.Context) {
	var newProject model.Project
	if err := k.ShouldBindJSON(&newProject); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	err := aa.projectService.Store(&newProject)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Project success"})
}
func (aa *projectAPI) AddSummary(k *gin.Context) {
	var newSummary model.Summary
	if err := k.ShouldBindJSON(&newSummary); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	err := aa.summaryService.Store(&newSummary)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Summary success"})
}

func (aa *projectAPI) UpdateProject(k *gin.Context) {
	var newProject model.Project
	if err := k.ShouldBindJSON(&newProject); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Project ID"})
		return
	}
	newProject.Project_ID = KpiID
	err = aa.projectService.Saves(newProject)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Project update success"})
}
func (aa *projectAPI) UpdateSummary(k *gin.Context) {
	var newSummary model.Summary
	if err := k.ShouldBindJSON(&newSummary); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Summary ID"})
		return
	}
	newSummary.Summary_ID = KpiID
	err = aa.summaryService.Saves(newSummary)	
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Summary update success"})
}

func (aa *projectAPI) DeleteProject(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Project ID"})
		return
	}
	err = aa.projectService.Delete(KpiID)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Project delete success"})
}
func (aa *projectAPI) DeleteSummary(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Summary ID"})
		return
	}
	err = aa.summaryService.Delete(KpiID)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Summary delete success"})
}

func (aa *projectAPI) GetProjectByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid Project Year"})
		return
	}
	Project, err := aa.projectService.GetByID(KpiID)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, Project)
}
func (aa *projectAPI) GetSummaryByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid Summary ID"})
		return
	}
	Summary, err := aa.summaryService.GetByID(KpiID)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	// Result := Summary.ToResponse()
	k.JSON(http.StatusOK, Summary)
}

func (aa *projectAPI) GetProjectList(k *gin.Context) {
	Project, err := aa.projectService.GetList()
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	var result model.ProjectArrayResponse
	result.Project = []model.ProjectResponse{}
	for _, data := range Project{
		result.Project = append(result.Project, data.ToResponse())
	}
	result.Message = "Getting All Projects Success"
	k.JSON(http.StatusOK, result)
}
func (aa *projectAPI) GetSummaryList(k *gin.Context) {
	Summary, err := aa.summaryService.GetList()
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	var result model.SummaryArrayResponse
	result.Summary = []model.SummaryResponse{}
	for _, data := range Summary{
		result.Summary = append(result.Summary, data.ToResponse())
	} 
	result.Message = "Getting All Summarys Success"
	k.JSON(http.StatusOK, result)
}

