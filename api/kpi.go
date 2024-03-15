package api

import (
	"goreact/model"
	"goreact/service" 
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type KpiAPI interface {
	AddFactor(k *gin.Context)
	UpdateFactor(k *gin.Context)
	DeleteFactor(k *gin.Context)
	GetFactorByID(k *gin.Context)
	GetFactorList(k *gin.Context)
	AddMonthly(k *gin.Context)
	UpdateMonthly(k *gin.Context)
	DeleteMonthly(k *gin.Context)
	GetMonthlyByID(k *gin.Context)
	GetMonthlyList(k *gin.Context)
}

type kpiAPI struct {
	attendanceService 	service.AttendanceService
	factorService 		service.FactorService
	monthlyService 		service.MonthlyService
	minipapService		service.MiniPAPService
	resultService		service.ResultService
}

func NewKpiAPI(
	attendanceService service.AttendanceService, 
	factorService service.FactorService, 
	monthlyService service.MonthlyService, 
	minipapService service.MiniPAPService,
	resultService service.ResultService) *kpiAPI{
	return &kpiAPI{
		attendanceService, 
		factorService, 
		monthlyService, 
		minipapService, 
		resultService}
}

func (ka *kpiAPI) AddFactor(k *gin.Context) {
	var newFactor model.Factor
	if err := k.ShouldBindJSON(&newFactor); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err := ka.factorService.Store(&newFactor)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Factor success"})
}

func (ka *kpiAPI) UpdateFactor(k *gin.Context) {
	var newFactor model.Factor
	if err := k.ShouldBindJSON(&newFactor); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Factor ID"})
		return
	}
	err = ka.factorService.Update(KpiID, newFactor)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Factor update success"})
}

func (ka *kpiAPI) DeleteFactor(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Factor ID"})
		return
	}
	err = ka.factorService.Delete(KpiID)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Factor delete success"})
}

func (ka *kpiAPI) GetFactorByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid Factor ID"})
		return
	}

	Factor, err := ka.factorService.GetByID(KpiID)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	k.JSON(http.StatusOK, Factor)
}

func (ka *kpiAPI) GetFactorList(k *gin.Context) {
	Factor, err := ka.factorService.GetList()
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	var result model.KpiArrayResponse
	result.Kpis = Factor 
	result.Message = "Getting All Kpis Success"

	k.JSON(http.StatusOK, result)
}

func (ka *kpiAPI) AddMonthly(k *gin.Context) {
	var newMonthly model.Monthly
	if err := k.ShouldBindJSON(&newMonthly); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err := ka.monthlyService.Store(&newMonthly)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Factor success"})
}

func (ka *kpiAPI) UpdateMonthly(k *gin.Context) {
	var newMonthly model.Monthly
	if err := k.ShouldBindJSON(&newMonthly); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Monthly ID"})
		return
	}
	err = ka.monthlyService.Update(KpiID, newMonthly)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Monthly update success"})
}

func (ka *kpiAPI) DeleteMonthly(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Monthly ID"})
		return
	}
	err = ka.monthlyService.Delete(KpiID)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Monthly delete success"})
}

func (ka *kpiAPI) GetMonthlyByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid Monthly ID"})
		return
	}

	Monthly, err := ka.monthlyService.GetByID(KpiID)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	k.JSON(http.StatusOK, Monthly)
}

func (ka *kpiAPI) GetMonthlyList(k *gin.Context) {
	Monthly, err := ka.monthlyService.GetList()
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	var result model.MonthlyArrayResponse
	result.Monthly = Monthly 
	result.Message = "Getting All Monthly Success"

	k.JSON(http.StatusOK, result)
}