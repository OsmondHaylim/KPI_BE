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

type AnalisaAPI interface {
	AddAnalisa(k *gin.Context)
	AddMasalah(k *gin.Context)

	UpdateAnalisa(k *gin.Context)
	UpdateMasalah(k *gin.Context)

	DeleteAnalisa(k *gin.Context)
	DeleteMasalah(k *gin.Context)

	GetAnalisaByID(k *gin.Context)
	GetMasalahByID(k *gin.Context)

	GetAnalisaList(k *gin.Context)
	GetMasalahList(k *gin.Context)
}
type analisaAPI struct {
	analisaService		service.AnalisaService
	masalahService 		service.MasalahService
}
func NewAnalisaAPI(
	analisaService service.AnalisaService,
	masalahService service.MasalahService) *analisaAPI{
	return &analisaAPI{
		analisaService,
		masalahService}
}

func (aa *analisaAPI) AddAnalisa(k *gin.Context) {
	var newAnalisa model.Analisa
	if err := k.ShouldBindJSON(&newAnalisa); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	err := aa.analisaService.Store(&newAnalisa)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Analisa success"})
}
func (aa *analisaAPI) AddMasalah(k *gin.Context) {
	var newMasalah model.Masalah
	if err := k.ShouldBindJSON(&newMasalah); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	err := aa.masalahService.Store(&newMasalah)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Masalah success"})
}

func (aa *analisaAPI) UpdateAnalisa(k *gin.Context) {
	var newAnalisa model.Analisa
	if err := k.ShouldBindJSON(&newAnalisa); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Analisa ID"})
		return
	}
	newAnalisa.Year = KpiID
	err = aa.analisaService.Saves(newAnalisa)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Analisa update success"})
}
func (aa *analisaAPI) UpdateMasalah(k *gin.Context) {
	var newMasalah model.Masalah
	if err := k.ShouldBindJSON(&newMasalah); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Masalah ID"})
		return
	}
	newMasalah.Masalah_ID = KpiID
	err = aa.masalahService.Saves(newMasalah)	
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Masalah update success"})
}

func (aa *analisaAPI) DeleteAnalisa(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Analisa ID"})
		return
	}
	err = aa.analisaService.Delete(KpiID)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Analisa delete success"})
}
func (aa *analisaAPI) DeleteMasalah(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Masalah ID"})
		return
	}
	err = aa.masalahService.Delete(KpiID)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Masalah delete success"})
}

func (aa *analisaAPI) GetAnalisaByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid Analisa Year"})
		return
	}
	Analisa, err := aa.analisaService.GetByID(KpiID)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, Analisa)
}
func (aa *analisaAPI) GetMasalahByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid Masalah ID"})
		return
	}
	Masalah, err := aa.masalahService.GetByID(KpiID)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	// Result := Masalah.ToResponse()
	k.JSON(http.StatusOK, Masalah)
}

func (aa *analisaAPI) GetAnalisaList(k *gin.Context) {
	Analisa, err := aa.analisaService.GetList()
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	var result model.AnalisaArrayResponse
	result.Analisa = Analisa 
	result.Message = "Getting All Analisas Success"
	k.JSON(http.StatusOK, result)
}
func (aa *analisaAPI) GetMasalahList(k *gin.Context) {
	Masalah, err := aa.masalahService.GetList()
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	var result model.MasalahArrayResponse
	result.Masalah = []model.MasalahResponse{}
	for _, data := range Masalah{
		result.Masalah = append(result.Masalah, data.ToResponse())
	} 
	result.Message = "Getting All Masalahs Success"
	k.JSON(http.StatusOK, result)
}

