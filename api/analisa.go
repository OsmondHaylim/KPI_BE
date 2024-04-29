package api

// import (
// 	// "fmt"
// 	"goreact/model"
// 	"goreact/service"
// 	"net/http"
// 	"strconv"
// 	// "strings"

// 	"github.com/gin-gonic/gin"
// )

// type AnalisaAPI interface {
// 	AddAnalisa(k *gin.Context)
// 	AddMasalah(k *gin.Context)

// 	UpdateAnalisa(k *gin.Context)
// 	UpdateMasalah(k *gin.Context)

// 	DeleteAnalisa(k *gin.Context)
// 	DeleteMasalah(k *gin.Context)

// 	GetAnalisaByID(k *gin.Context)
// 	GetMasalahByID(k *gin.Context)

// 	GetAnalisaList(k *gin.Context)
// 	GetMasalahList(k *gin.Context)

// 	AddEntireAnalisa(k *gin.Context)
// 	DeleteEntireAnalisa(k *gin.Context)
// }
// type analisaAPI struct {crudService service.CrudService}
// func NewAnalisaAPI(crudService service.CrudService) *analisaAPI{
// 	return &analisaAPI{crudService}}

// func (aa *analisaAPI) AddAnalisa(k *gin.Context) {
// 	var newAnalisa model.Analisa
// 	err := k.ShouldBindJSON(&newAnalisa)
// 	if model.ErrorCheck(k, err){return}
// 	err = aa.crudService.AddAnalisa(&newAnalisa)
// 	if model.ErrorCheck(k, err){return}
// 	k.JSON(http.StatusCreated, model.SuccessResponse{Message: "add Analisa success"})
// }
// func (aa *analisaAPI) AddMasalah(k *gin.Context) {
// 	var newMasalah model.Masalah
// 	err := k.ShouldBindJSON(&newMasalah)
// 	if model.ErrorCheck(k, err){return}
// 	err = aa.crudService.AddMasalah(&newMasalah)
// 	if model.ErrorCheck(k, err){return}
// 	k.JSON(http.StatusCreated, model.SuccessResponse{Message: "add Masalah success"})
// }

// func (aa *analisaAPI) UpdateAnalisa(k *gin.Context) {
// 	var newAnalisa model.Analisa
// 	err := k.ShouldBindJSON(&newAnalisa)
// 	if model.ErrorCheck(k, err){return}
// 	KpiID, err := strconv.Atoi(k.Param("id"))
// 	if model.ErrorCheck(k, err){return}
// 	err = aa.crudService.UpdateAnalisa(KpiID, newAnalisa)
// 	if model.ErrorCheck(k, err){return}
// 	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Analisa update success"})
// }
// func (aa *analisaAPI) UpdateMasalah(k *gin.Context) {
// 	var newMasalah model.Masalah
// 	err := k.ShouldBindJSON(&newMasalah)
// 	if model.ErrorCheck(k, err){return}
// 	KpiID, err := strconv.Atoi(k.Param("id"))
// 	if model.ErrorCheck(k, err){return}
// 	err = aa.crudService.UpdateMasalah(KpiID, newMasalah)	
// 	if model.ErrorCheck(k, err){return}
// 	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Masalah update success"})
// }

// func (aa *analisaAPI) DeleteAnalisa(k *gin.Context) {
// 	KpiID, err := strconv.Atoi(k.Param("id"))
// 	if model.ErrorCheck(k, err){return}
// 	err = aa.crudService.DeleteAnalisa(KpiID)
// 	if model.ErrorCheck(k, err){return}
// 	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Analisa delete success"})
// }
// func (aa *analisaAPI) DeleteMasalah(k *gin.Context) {
// 	KpiID, err := strconv.Atoi(k.Param("id"))
// 	if model.ErrorCheck(k, err){return}
// 	err = aa.crudService.DeleteMasalah(KpiID)
// 	if model.ErrorCheck(k, err){return}
// 	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Masalah delete success"})
// }

// func (aa *analisaAPI) GetAnalisaByID(k *gin.Context) {
// 	KpiID, err := strconv.Atoi(k.Param("id"))
// 	if model.ErrorCheck(k, err){return}
// 	Analisa, err := aa.crudService.GetAnalisaByID(KpiID)
// 	if model.ErrorCheck(k, err){return}
// 	k.JSON(http.StatusOK, Analisa)
// }
// func (aa *analisaAPI) GetMasalahByID(k *gin.Context) {
// 	KpiID, err := strconv.Atoi(k.Param("id"))
// 	if model.ErrorCheck(k, err){return}
// 	Masalah, err := aa.crudService.GetMasalahByID(KpiID)
// 	if model.ErrorCheck(k, err){return}
// 	// Result := Masalah.ToResponse()
// 	k.JSON(http.StatusOK, Masalah)
// }

// func (aa *analisaAPI) GetAnalisaList(k *gin.Context) {
// 	Analisa, err := aa.crudService.GetAnalisaList()
// 	if model.ErrorCheck(k, err){return}
// 	Analisa.Message = "Getting All Analisas Success"
// 	k.JSON(http.StatusOK, Analisa)
// }
// func (aa *analisaAPI) GetMasalahList(k *gin.Context) {
// 	Masalah, err := aa.crudService.GetMasalahList()
// 	if model.ErrorCheck(k, err){return}
// 	Masalah.Message = "Getting All Masalahs Success"
// 	k.JSON(http.StatusOK, Masalah)
// }

// func (aa *analisaAPI) AddEntireAnalisa(k *gin.Context){
// 	var newAnalisa model.AnalisaResponse
// 	err := k.ShouldBindJSON(&newAnalisa)
// 	if model.ErrorCheck(k, err){return}
// 	err = aa.crudService.AddEntireAnalisa(&newAnalisa)
// 	if model.ErrorCheck(k, err){return}
// 	k.JSON(http.StatusCreated, model.SuccessResponse{Message: "add Analisa success"})
// }
// func (aa *analisaAPI) DeleteEntireAnalisa(k *gin.Context) {
// 	KpiID, err := strconv.Atoi(k.Param("id"))
// 	if model.ErrorCheck(k, err){return}
// 	err = aa.crudService.DeleteEntireAnalisa(KpiID)
// 	if model.ErrorCheck(k, err){return}
// 	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Analisa delete success"})
// }
