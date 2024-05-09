package api

import (
	"goreact/model"
	"goreact/service"
	"net/http"
	"strconv"
	// "strings"

	"github.com/gin-gonic/gin"
)

type KpiAPI interface {
	AddAttendance(k *gin.Context)
	AddFactor(k *gin.Context)
	AddItem(k *gin.Context)
	AddMinipap(k *gin.Context)
	AddMonthly(k *gin.Context)
	AddResult(k *gin.Context)
	AddYearly(k *gin.Context)

	AddEntireYearly(k *gin.Context)
	AddEntireItem(k *gin.Context)
	AddEntireResult(k *gin.Context)
	AddEntireFactor(k *gin.Context)
	AddEntireAttendance(k *gin.Context)

	UpdateAttendance(k *gin.Context)
	UpdateFactor(k *gin.Context)
	UpdateItem(k *gin.Context)
	UpdateMinipap(k *gin.Context)
	UpdateMonthly(k *gin.Context)
	UpdateResult(k *gin.Context)
	UpdateYearly(k *gin.Context)

	UpdateEntireAttendance(k *gin.Context)
	UpdateEntireFactor(k *gin.Context)
	UpdateEntireItem(k *gin.Context)
	UpdateEntireResult(k *gin.Context)
	UpdateEntireYearly(k *gin.Context)

	DeleteAttendance(k *gin.Context)
	DeleteFactor(k *gin.Context)
	DeleteItem(k *gin.Context)
	DeleteMinipap(k *gin.Context)
	DeleteMonthly(k *gin.Context)
	DeleteResult(k *gin.Context)
	DeleteYearly(k *gin.Context)

	DeleteEntireYearly(k *gin.Context)
	DeleteEntireItem(k *gin.Context)
	DeleteEntireResult(k *gin.Context)
	DeleteEntireFactor(k *gin.Context)
	DeleteEntireAttendance(k *gin.Context)

	GetAttendanceByID(k *gin.Context)
	GetFactorByID(k *gin.Context)
	GetItemByID(k *gin.Context)
	GetMinipapByID(k *gin.Context)
	GetMonthlyByID(k *gin.Context)
	GetResultByID(k *gin.Context)
	GetYearlyByID(k *gin.Context)

	GetAttendanceList(k *gin.Context)
	GetFactorList(k *gin.Context)
	GetItemList(k *gin.Context)
	GetMinipapList(k *gin.Context)
	GetMonthlyList(k *gin.Context)
	GetResultList(k *gin.Context)
	GetYearlyList(k *gin.Context)
}
type kpiAPI struct {crudService service.CrudService}
func NewKpiAPI(crudService service.CrudService) *kpiAPI{
	return &kpiAPI{crudService}
}

// Add
func (ka *kpiAPI) AddAttendance(k *gin.Context) {
	var newAttendance model.Attendance
	if model.ErrorCheck(k, k.ShouldBindJSON(&newAttendance)) {return}
	if model.ErrorCheck(k, ka.crudService.AddAttendance(&newAttendance)) {return}
	k.JSON(http.StatusCreated, model.SuccessResponse{Message: "add Attendance success"})
}
func (ka *kpiAPI) AddFactor(k *gin.Context) {
	var newFactor model.Factor
	if model.ErrorCheck(k, k.ShouldBindJSON(&newFactor)) {return}
	if model.ErrorCheck(k, ka.crudService.AddFactor(&newFactor)) {return}
	k.JSON(http.StatusCreated, model.SuccessResponse{Message: "add Factor success"})
}
func (ka *kpiAPI) AddItem(k *gin.Context) {
	var newItem model.Item
	if model.ErrorCheck(k, k.ShouldBindJSON(&newItem)) {return}
	if model.ErrorCheck(k, ka.crudService.AddItem(&newItem)) {return}
	k.JSON(http.StatusCreated, model.SuccessResponse{Message: "add Item success"})
}
func (ka *kpiAPI) AddMinipap(k *gin.Context) {
	var newMinipap model.MiniPAP
	if model.ErrorCheck(k, k.ShouldBindJSON(&newMinipap)) {return}
	if model.ErrorCheck(k, ka.crudService.AddMinipap(&newMinipap)) {return}
	k.JSON(http.StatusCreated, model.SuccessResponse{Message: "add Minipap success"})
}
func (ka *kpiAPI) AddMonthly(k *gin.Context) {
	var newMonthly model.Monthly
	if model.ErrorCheck(k, k.ShouldBindJSON(&newMonthly)) {return}
	newMonthly = newMonthly.Reseted()
	if model.ErrorCheck(k, ka.crudService.AddMonthly(&newMonthly)) {return}
	k.JSON(http.StatusCreated, model.SuccessResponse{Message: "add Monthly success"})
}
func (ka *kpiAPI) AddResult(k *gin.Context) {
	var newResult model.Result
	if model.ErrorCheck(k, k.ShouldBindJSON(&newResult)) {return}
	if model.ErrorCheck(k, ka.crudService.AddResult(&newResult)) {return}
	k.JSON(http.StatusCreated, model.SuccessResponse{Message: "add Result success"})
}
func (ka *kpiAPI) AddYearly(k *gin.Context) {
	var newYearly model.Yearly
	if model.ErrorCheck(k, k.ShouldBindJSON(&newYearly)) {return}
	if model.ErrorCheck(k, ka.crudService.AddYearly(&newYearly)) {return}
	k.JSON(http.StatusCreated, model.SuccessResponse{Message: "add Yearly success"})
}

// Add Entire
func (ka *kpiAPI) AddEntireYearly(k *gin.Context) {
	var newYearly model.YearlyResponse
	if model.ErrorCheck(k, k.ShouldBindJSON(&newYearly)) {return}
	if model.ErrorCheck(k, ka.crudService.AddEntireYearly(&newYearly)) {return}
	k.JSON(http.StatusCreated, model.SuccessResponse{Message: "add Entire Yearly success"})
}
func (ka *kpiAPI) AddEntireItem(k *gin.Context) {
	var response model.ItemResponse
	if model.ErrorCheck(k, k.ShouldBindJSON(&response)) {return}
	wg, errs := model.GoRoutineInit()
	wg.Add(1)
	go ka.crudService.AddEntireItem(&wg, &response, nil, errs)
	wg.Wait()
	if model.ErrorChanCheck(k, errs) {return}
	k.JSON(http.StatusCreated, model.SuccessResponse{Message: "add Entire Item success"})
}
func (ka *kpiAPI) AddEntireResult(k *gin.Context) {
	var response model.ResultResponse
	if model.ErrorCheck(k, k.ShouldBindJSON(&response) ) {return}
	if model.ErrorCheck(k, ka.crudService.AddEntireResult(&response, nil)) {return}
	k.JSON(http.StatusCreated, model.SuccessResponse{Message: "add Entire Result success"})
}
func (ka *kpiAPI) AddEntireFactor(k *gin.Context) {
	var response model.FactorResponse
	if model.ErrorCheck(k, k.ShouldBindJSON(&response) ) {return}
	if model.ErrorCheck(k, ka.crudService.AddEntireFactor(&response, nil)) {return}
	k.JSON(http.StatusCreated, model.SuccessResponse{Message: "add Entire Factor success"})
}
func (ka *kpiAPI) AddEntireAttendance(k *gin.Context) {
	var response model.AttendanceResponse
	if model.ErrorCheck(k, k.ShouldBindJSON(&response)){return}
	if model.ErrorCheck(k, ka.crudService.AddEntireAttendance(&response, nil)) {return}
	k.JSON(http.StatusCreated, model.SuccessResponse{Message: "add Entire Attendance success"})
}

// Update
func (ka *kpiAPI) UpdateAttendance(k *gin.Context) {
	var newAttendance model.Attendance
	if model.ErrorCheck(k, k.ShouldBindJSON(&newAttendance)){return}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	if model.ErrorCheck(k, ka.crudService.UpdateAttendance(KpiID, newAttendance)) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Attendance update success"})
}
func (ka *kpiAPI) UpdateFactor(k *gin.Context) {
	var newFactor model.Factor
	if model.ErrorCheck(k, k.ShouldBindJSON(&newFactor)){return}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	if model.ErrorCheck(k, ka.crudService.UpdateFactor(KpiID, newFactor)) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Factor update success"})
}
func (ka *kpiAPI) UpdateItem(k *gin.Context) {
	var newItem model.Item
	if model.ErrorCheck(k, k.ShouldBindJSON(&newItem)){return}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	if model.ErrorCheck(k, ka.crudService.UpdateItem(KpiID, newItem)	) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Item update success"})
}
func (ka *kpiAPI) UpdateMinipap(k *gin.Context) {
	var newMinipap model.MiniPAP
	if model.ErrorCheck(k, k.ShouldBindJSON(&newMinipap)){return}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	if model.ErrorCheck(k, ka.crudService.UpdateMinipap(KpiID, newMinipap)	) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "MiniPAP update success"})
}
func (ka *kpiAPI) UpdateMonthly(k *gin.Context) {
	var newMonthly model.Monthly
	if model.ErrorCheck(k, k.ShouldBindJSON(&newMonthly)){return}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	if model.ErrorCheck(k, ka.crudService.UpdateMonthly(KpiID, newMonthly)) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Monthly update success"})
}
func (ka *kpiAPI) UpdateResult(k *gin.Context) {
	var newResult model.Result
	if model.ErrorCheck(k, k.ShouldBindJSON(&newResult)){return}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	if model.ErrorCheck(k, ka.crudService.UpdateResult(KpiID, newResult)) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Result update success"})
}
func (ka *kpiAPI) UpdateYearly(k *gin.Context) {
	var newYearly model.Yearly
	if model.ErrorCheck(k, k.ShouldBindJSON(&newYearly)){return}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	if model.ErrorCheck(k, ka.crudService.UpdateYearly(KpiID, newYearly)	) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Yearly update success"})
}

// Update Entire
func (ka *kpiAPI) UpdateEntireAttendance(k *gin.Context) {
	var newAttendance model.AttendanceResponse
	if model.ErrorCheck(k, k.ShouldBindJSON(&newAttendance)){return}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	if model.ErrorCheck(k, ka.crudService.UpdateEntireAttendance(KpiID, newAttendance)) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Attendance update success"})
}
func (ka *kpiAPI) UpdateEntireFactor(k *gin.Context) {
	var newFactor model.FactorResponse
	if model.ErrorCheck(k, k.ShouldBindJSON(&newFactor)){return}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	if model.ErrorCheck(k, ka.crudService.UpdateEntireFactor(KpiID, newFactor)) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Factor update success"})
}
func (ka *kpiAPI) UpdateEntireItem(k *gin.Context) {
	var newItem model.ItemResponse
	if model.ErrorCheck(k, k.ShouldBindJSON(&newItem)){return}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	if model.ErrorCheck(k, ka.crudService.UpdateEntireItem(KpiID, newItem)	) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Item update success"})
}
func (ka *kpiAPI) UpdateEntireResult(k *gin.Context) {
	var newResult model.ResultResponse
	if model.ErrorCheck(k, k.ShouldBindJSON(&newResult)){return}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	if model.ErrorCheck(k, ka.crudService.UpdateEntireResult(KpiID, newResult)	) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Result update success"})
}
func (ka *kpiAPI) UpdateEntireYearly(k *gin.Context) {
	var newYearly model.YearlyResponse
	if model.ErrorCheck(k, k.ShouldBindJSON(&newYearly)){return}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	if model.ErrorCheck(k, ka.crudService.UpdateEntireYearly(KpiID, newYearly)) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Yearly update success"})
}

// Delete
func (ka *kpiAPI) DeleteAttendance(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	if model.ErrorCheck(k, ka.crudService.DeleteAttendance(KpiID)) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Attendance delete success"})
}
func (ka *kpiAPI) DeleteFactor(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	if model.ErrorCheck(k, ka.crudService.DeleteFactor(KpiID)) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Factor delete success"})
}
func (ka *kpiAPI) DeleteItem(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	if model.ErrorCheck(k, ka.crudService.DeleteItem(KpiID)) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Item delete success"})
}
func (ka *kpiAPI) DeleteMinipap(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	if model.ErrorCheck(k, ka.crudService.DeleteMinipap(KpiID)) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "MiniPAP delete success"})
}
func (ka *kpiAPI) DeleteMonthly(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	if model.ErrorCheck(k, ka.crudService.DeleteMonthly(KpiID)) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Monthly delete success"})
}
func (ka *kpiAPI) DeleteResult(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	if model.ErrorCheck(k, ka.crudService.DeleteResult(KpiID)) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Result delete success"})
}
func (ka *kpiAPI) DeleteYearly(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	if model.ErrorCheck(k, ka.crudService.DeleteYearly(KpiID)) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Yearly delete success"})
}

// Delete Entire
func (ka *kpiAPI) DeleteEntireYearly(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	if model.ErrorCheck(k, ka.crudService.DeleteEntireYearly(KpiID)){return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "delete Entire Yearly success"})
}
func (ka *kpiAPI) DeleteEntireItem(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	wg, errs := model.GoRoutineInit()
	wg.Add(1)
	go ka.crudService.DeleteEntireItem(&wg, KpiID, errs)
	wg.Wait()
	if model.ErrorChanCheck(k, errs) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "delete Entire Item success"})
}
func (ka *kpiAPI) DeleteEntireResult(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	if model.ErrorCheck(k, ka.crudService.DeleteEntireResult(KpiID)) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "delete Entire Result success"})
}
func (ka *kpiAPI) DeleteEntireFactor(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	if model.ErrorCheck(k, ka.crudService.DeleteEntireFactor(KpiID)) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "delete Entire Factor success"})
}
func (ka *kpiAPI) DeleteEntireAttendance(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err) {return}
	if model.ErrorCheck(k, ka.crudService.DeleteEntireAttendance(KpiID)) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "delete Entire Attendance success"})
}

// Get By ID
func (ka *kpiAPI) GetAttendanceByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	Attendance, err := ka.crudService.GetAttendanceByID(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, Attendance)
}
func (ka *kpiAPI) GetFactorByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	Factor, err := ka.crudService.GetFactorByID(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, Factor)
}
func (ka *kpiAPI) GetItemByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	Item, err := ka.crudService.GetItemByID(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, Item)
}
func (ka *kpiAPI) GetMinipapByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	Minipap, err := ka.crudService.GetMinipapByID(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, Minipap)
}
func (ka *kpiAPI) GetMonthlyByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	Monthly, err := ka.crudService.GetMonthlyByID(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, Monthly)
}
func (ka *kpiAPI) GetResultByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	Result, err := ka.crudService.GetResultByID(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, Result)
}
func (ka *kpiAPI) GetYearlyByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	Yearly, err := ka.crudService.GetYearlyByID(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, Yearly)
}

// Get List
func (ka *kpiAPI) GetAttendanceList(k *gin.Context) {
	Attendance, err := ka.crudService.GetAttendanceList()
	if model.ErrorCheck(k, err) {return}
	Attendance.Message = "Getting All Attendances Success"
	k.JSON(http.StatusOK, Attendance)
}
func (ka *kpiAPI) GetFactorList(k *gin.Context) {
	Factor, err := ka.crudService.GetFactorList()
	if model.ErrorCheck(k, err) {return}
	Factor.Message = "Getting All Factors Success"
	k.JSON(http.StatusOK, Factor)
}
func (ka *kpiAPI) GetItemList(k *gin.Context) {
	Item, err := ka.crudService.GetItemList()
	if model.ErrorCheck(k, err) {return}
	Item.Message = "Getting All Items Success"
	k.JSON(http.StatusOK, Item)
}
func (ka *kpiAPI) GetMinipapList(k *gin.Context) {
	Minipap, err := ka.crudService.GetMinipapList()
	if model.ErrorCheck(k, err) {return}
	Minipap.Message = "Getting All MiniPAPs Success"
	k.JSON(http.StatusOK, Minipap)
}
func (ka *kpiAPI) GetMonthlyList(k *gin.Context) {
	Monthly, err := ka.crudService.GetMonthlyList()
	if model.ErrorCheck(k, err) {return}
	Monthly.Message = "Getting All Monthlys Success"
	k.JSON(http.StatusOK, Monthly)
}
func (ka *kpiAPI) GetResultList(k *gin.Context) {
	Result, err := ka.crudService.GetResultList()
	if model.ErrorCheck(k, err) {return}
	Result.Message = "Getting All Results Success"
	k.JSON(http.StatusOK, Result)
}
func (ka *kpiAPI) GetYearlyList(k *gin.Context) {
	Yearly, err := ka.crudService.GetYearlyList()
	if model.ErrorCheck(k, err) {return}
	Yearly.Message = "Getting All Yearlys Success"
	k.JSON(http.StatusOK, Yearly)
}