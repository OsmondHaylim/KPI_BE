package api

import (
	"goreact/model"
	"goreact/service"
	"net/http"
	"strconv"

	// "strings"

	"github.com/gin-gonic/gin")

type KpiAPI interface {
	AddAttendance(k *gin.Context)
	AddFactor(k *gin.Context)
	AddItem(k *gin.Context)
	AddMinipap(k *gin.Context)
	AddMonthly(k *gin.Context)
	AddResult(k *gin.Context)
	AddYearly(k *gin.Context)

	// AddEntireYearly(k *gin.Context)
	// AddEntireItem(k *gin.Context)
	// AddEntireResult(k *gin.Context)
	// AddEntireFactor(k *gin.Context)
	// AddEntireAttendance(k *gin.Context)

	UpdateAttendance(k *gin.Context)
	UpdateFactor(k *gin.Context)
	UpdateItem(k *gin.Context)
	UpdateMinipap(k *gin.Context)
	UpdateMonthly(k *gin.Context)
	UpdateResult(k *gin.Context)
	UpdateYearly(k *gin.Context)

	DeleteAttendance(k *gin.Context)
	DeleteFactor(k *gin.Context)
	DeleteItem(k *gin.Context)
	DeleteMinipap(k *gin.Context)
	DeleteMonthly(k *gin.Context)
	DeleteResult(k *gin.Context)
	DeleteYearly(k *gin.Context)

	// DeleteEntireYearly(k *gin.Context)
	// DeleteEntireItem(k *gin.Context)
	// DeleteEntireResult(k *gin.Context)
	// DeleteEntireFactor(k *gin.Context)
	// DeleteEntireAttendance(k *gin.Context)

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
	err := k.ShouldBindJSON(&newAttendance)
	if model.ErrorCheck(k, err) {return}
	err = ka.crudService.AddAttendance(&newAttendance)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Attendance success"})
}
func (ka *kpiAPI) AddFactor(k *gin.Context) {
	var newFactor model.Factor
	err := k.ShouldBindJSON(&newFactor)
	if model.ErrorCheck(k, err) {return}
	err = ka.crudService.AddFactor(&newFactor)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Factor success"})
}
func (ka *kpiAPI) AddItem(k *gin.Context) {
	var newItem model.Item
	err := k.ShouldBindJSON(&newItem)
	if model.ErrorCheck(k, err) {return}
	err = ka.crudService.AddItem(&newItem)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Item success"})
}
func (ka *kpiAPI) AddMinipap(k *gin.Context) {
	var newMinipap model.MiniPAP
	err := k.ShouldBindJSON(&newMinipap)
	if model.ErrorCheck(k, err) {return}
	err = ka.crudService.AddMinipap(&newMinipap)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Minipap success"})
}
func (ka *kpiAPI) AddMonthly(k *gin.Context) {
	var newMonthly model.Monthly
	err := k.ShouldBindJSON(&newMonthly)
	if model.ErrorCheck(k, err) {return}
	newMonthly = newMonthly.Reseted()
	err = ka.crudService.AddMonthly(&newMonthly)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Monthly success"})
}
func (ka *kpiAPI) AddResult(k *gin.Context) {
	var newResult model.Result
	err := k.ShouldBindJSON(&newResult)
	if model.ErrorCheck(k, err) {return}
	err = ka.crudService.AddResult(&newResult)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Result success"})
}
func (ka *kpiAPI) AddYearly(k *gin.Context) {
	var newYearly model.Yearly
	err := k.ShouldBindJSON(&newYearly)
	if model.ErrorCheck(k, err) {return}
	err = ka.crudService.AddYearly(&newYearly)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Yearly success"})
}

// Add Entire (Not Done)
func (ka *kpiAPI) AddEntireYearly(k *gin.Context) {
	var newYearly model.YearlyResponse
	err := k.ShouldBindJSON(&newYearly)
	if model.ErrorCheck(k, err) {return}
	err = ka.crudService.AddEntireYearly(&newYearly)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Entire Yearly success"})
}
func (ka *kpiAPI) AddEntireItem(k *gin.Context) {
	var response model.ItemResponse
	err := k.ShouldBindJSON(&response)
	if model.ErrorCheck(k, err) {return}
	err = ka.crudService.AddEntireItem(&response, nil)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Entire Item success"})
}
func (ka *kpiAPI) AddEntireResult(k *gin.Context) {
	var response model.ResultResponse
	err := k.ShouldBindJSON(&response) 
	if model.ErrorCheck(k, err) {return}
	err = ka.crudService.AddEntireResult(&response, nil)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Entire Result success"})
}
func (ka *kpiAPI) AddEntireFactor(k *gin.Context) {
	var response model.FactorResponse
	err := k.ShouldBindJSON(&response) 
	if model.ErrorCheck(k, err) {return}
	err = ka.crudService.AddEntireFactor(&response, nil)
	if model.ErrorCheck(k, err){return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Entire Factor success"})
}
func (ka *kpiAPI) AddEntireAttendance(k *gin.Context) {
	var response model.AttendanceResponse
	err := k.ShouldBindJSON(&response)
	if model.ErrorCheck(k, err){return}
	err = ka.crudService.AddEntireAttendance(&response, nil)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Entire Attendance success"})
}

// Update
func (ka *kpiAPI) UpdateAttendance(k *gin.Context) {
	var newAttendance model.Attendance
	err := k.ShouldBindJSON(&newAttendance)
	if model.ErrorCheck(k, err){return}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	err = ka.crudService.UpdateAttendance(KpiID, newAttendance)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Attendance update success"})
}
func (ka *kpiAPI) UpdateFactor(k *gin.Context) {
	var newFactor model.Factor
	err := k.ShouldBindJSON(&newFactor)
	if model.ErrorCheck(k, err){return}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	err = ka.crudService.UpdateFactor(KpiID, newFactor)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Factor update success"})
}
func (ka *kpiAPI) UpdateItem(k *gin.Context) {
	var newItem model.Item
	err := k.ShouldBindJSON(&newItem)
	if model.ErrorCheck(k, err){return}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	err = ka.crudService.UpdateItem(KpiID, newItem)	
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Item update success"})
}
func (ka *kpiAPI) UpdateMinipap(k *gin.Context) {
	var newMinipap model.MiniPAP
	err := k.ShouldBindJSON(&newMinipap)
	if model.ErrorCheck(k, err){return}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	err = ka.crudService.UpdateMinipap(KpiID, newMinipap)	
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "MiniPAP update success"})
}
func (ka *kpiAPI) UpdateMonthly(k *gin.Context) {
	var newMonthly model.Monthly
	err := k.ShouldBindJSON(&newMonthly)
	if model.ErrorCheck(k, err){return}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	err = ka.crudService.UpdateMonthly(KpiID, newMonthly)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Monthly update success"})
}
func (ka *kpiAPI) UpdateResult(k *gin.Context) {
	var newResult model.Result
	err := k.ShouldBindJSON(&newResult)
	if model.ErrorCheck(k, err){return}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	err = ka.crudService.UpdateResult(KpiID, newResult)	
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Result update success"})
}
func (ka *kpiAPI) UpdateYearly(k *gin.Context) {
	var newYearly model.Yearly
	err := k.ShouldBindJSON(&newYearly)
	if model.ErrorCheck(k, err){return}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	err = ka.crudService.UpdateYearly(KpiID, newYearly)	
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Yearly update success"})
}

// Delete
func (ka *kpiAPI) DeleteAttendance(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	err = ka.crudService.DeleteAttendance(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Attendance delete success"})
}
func (ka *kpiAPI) DeleteFactor(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	err = ka.crudService.DeleteFactor(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Factor delete success"})
}
func (ka *kpiAPI) DeleteItem(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	err = ka.crudService.DeleteItem(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Item delete success"})
}
func (ka *kpiAPI) DeleteMinipap(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	err = ka.crudService.DeleteMinipap(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "MiniPAP delete success"})
}
func (ka *kpiAPI) DeleteMonthly(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	err = ka.crudService.DeleteMonthly(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Monthly delete success"})
}
func (ka *kpiAPI) DeleteResult(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	err = ka.crudService.DeleteResult(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Result delete success"})
}
func (ka *kpiAPI) DeleteYearly(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	err = ka.crudService.DeleteYearly(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Yearly delete success"})
}

// // Delete Entire (Not Done)
// func (ka *kpiAPI) DeleteEntireYearly(k *gin.Context) {
// 	KpiID, err := strconv.Atoi(k.Param("id"))
// 	if model.ErrorCheck(k, err){return}
// 	Yearly, err := ka.crudService.GetByID(KpiID)
// 	if model.ErrorCheck(k, err) {return}
	
	
// 	//Delete Items
// 	for _, item := range Yearly.Items{
// 		for _, result := range item.Results{
// 			for _, factor := range result.Factors{
// 				for _, monthly := range factor.Plan.Monthly{
// 					err = ka.crudService.Delete(monthly.Monthly_ID)
// 					if model.ErrorCheck(k, err) {return}
// 				}
// 				err = ka.crudService.Delete(*factor.PlanID)
// 				if model.ErrorCheck(k, err) {return}
// 				for _, monthly := range factor.Actual.Monthly{
// 					err = ka.crudService.Delete(monthly.Monthly_ID)
// 					if model.ErrorCheck(k, err) {return}
// 				}
// 				err = ka.crudService.Delete(*factor.ActualID)
// 				if model.ErrorCheck(k, err) {return}
// 				err = ka.crudService.Delete(factor.Factor_ID)
// 				if model.ErrorCheck(k, err) {return}
// 			}
// 			err = ka.crudService.Delete(result.Result_ID)
// 			if model.ErrorCheck(k, err) {return}
// 		}
// 		err = ka.crudService.Delete(item.Item_ID)
// 		if model.ErrorCheck(k, err) {return}
// 	}
// 	//Delete Yearly
// 	err = ka.crudService.Delete(Yearly.Year)
// 	if model.ErrorCheck(k, err) {return}

// 	//Delete Attendance
// 	if Yearly.AttendanceID != nil {
// 		err = ka.crudService.Delete(Yearly.Attendance.Year)
// 		if model.ErrorCheck(k, err) {return}
// 	}
// 	//Delete Attendance Monthly
// 	if Yearly.Attendance.PlanID != nil {
// 		err = ka.crudService.Delete(*Yearly.Attendance.PlanID)
// 		if model.ErrorCheck(k, err) {return}
// 	}
// 	if Yearly.Attendance.ActualID != nil {
// 		err = ka.crudService.Delete(*Yearly.Attendance.ActualID)
// 		if model.ErrorCheck(k, err) {return}
// 	}
// 	if Yearly.Attendance.CutiID != nil {
// 		err = ka.crudService.Delete(*Yearly.Attendance.CutiID)
// 		if model.ErrorCheck(k, err) {return}
// 	}
// 	if Yearly.Attendance.IzinID != nil {
// 		err = ka.crudService.Delete(*Yearly.Attendance.IzinID)
// 		if model.ErrorCheck(k, err) {return}
// 	}
// 	if Yearly.Attendance.LainID != nil {
// 		err = ka.crudService.Delete(*Yearly.Attendance.LainID)
// 		if model.ErrorCheck(k, err) {return}
// 	}
// 	k.JSON(http.StatusOK, model.SuccessResponse{Message: "delete Entire Yearly success"})
// }
// func (ka *kpiAPI) DeleteEntireItem(k *gin.Context) {
// 	KpiID, err := strconv.Atoi(k.Param("id"))
// 	if model.ErrorCheck(k, err){return}
// 	item, err := ka.crudService.GetByID(KpiID)
// 	if model.ErrorCheck(k, err) {return}
// 	//Delete Items
// 	for _, result := range item.Results{
// 		for _, factor := range result.Factors{
// 			for _, monthly := range factor.Plan.Monthly{
// 				err = ka.crudService.Delete(monthly.Monthly_ID)
// 				if model.ErrorCheck(k, err) {return}
// 			}
// 			err = ka.crudService.Delete(*factor.PlanID)
// 			if model.ErrorCheck(k, err) {return}
// 			for _, monthly := range factor.Actual.Monthly{
// 				err = ka.crudService.Delete(monthly.Monthly_ID)
// 				if model.ErrorCheck(k, err) {return}
// 			}
// 			err = ka.crudService.Delete(*factor.ActualID)
// 			if model.ErrorCheck(k, err) {return}
// 			err = ka.crudService.Delete(factor.Factor_ID)
// 			if model.ErrorCheck(k, err) {return}
// 		}
// 		err = ka.crudService.Delete(result.Result_ID)
// 		if model.ErrorCheck(k, err) {return}
// 	}
// 	err = ka.crudService.Delete(item.Item_ID)
// 	if model.ErrorCheck(k, err) {return}
// 	k.JSON(http.StatusOK, model.SuccessResponse{Message: "delete Entire Item success"})
// }
// func (ka *kpiAPI) DeleteEntireResult(k *gin.Context) {
// 	KpiID, err := strconv.Atoi(k.Param("id"))
// 	if model.ErrorCheck(k, err){return}
// 	result, err := ka.crudService.GetByID(KpiID)
// 	if model.ErrorCheck(k, err) {return}
// 	//Delete Results
// 	for _, factor := range result.Factors{
// 		for _, monthly := range factor.Plan.Monthly{
// 			err = ka.crudService.Delete(monthly.Monthly_ID)
// 			if model.ErrorCheck(k, err) {return}
// 		}
// 		err = ka.crudService.Delete(*factor.PlanID)
// 		if model.ErrorCheck(k, err) {return}
// 		for _, monthly := range factor.Actual.Monthly{
// 			err = ka.crudService.Delete(monthly.Monthly_ID)
// 			if model.ErrorCheck(k, err) {return}
// 		}
// 		err = ka.crudService.Delete(*factor.ActualID)
// 		if model.ErrorCheck(k, err) {return}
// 		err = ka.crudService.Delete(factor.Factor_ID)
// 		if model.ErrorCheck(k, err) {return}
// 	}
// 	err = ka.crudService.Delete(result.Result_ID)
// 	if model.ErrorCheck(k, err) {return}
// 	k.JSON(http.StatusOK, model.SuccessResponse{Message: "delete Entire Result success"})
// }
// func (ka *kpiAPI) DeleteEntireFactor(k *gin.Context) {
// 	KpiID, err := strconv.Atoi(k.Param("id"))
// 	if model.ErrorCheck(k, err){return}
// 	factor, err := ka.crudService.GetByID(KpiID)
// 	if model.ErrorCheck(k, err) {return}
// 	//Delete Factors
// 	for _, monthly := range factor.Plan.Monthly{
// 		err = ka.crudService.Delete(monthly.Monthly_ID)
// 		if model.ErrorCheck(k, err) {return}
// 	}
// 	err = ka.crudService.Delete(*factor.PlanID)
// 	if model.ErrorCheck(k, err) {return}
// 	for _, monthly := range factor.Actual.Monthly{
// 		err = ka.crudService.Delete(monthly.Monthly_ID)
// 		if model.ErrorCheck(k, err) {return}
// 	}
// 	err = ka.crudService.Delete(*factor.ActualID)
// 	if model.ErrorCheck(k, err) {return}
// 	err = ka.crudService.Delete(factor.Factor_ID)
// 	if model.ErrorCheck(k, err) {return}
// 	k.JSON(http.StatusOK, model.SuccessResponse{Message: "delete Entire Factor success"})
// }
// func (ka *kpiAPI) DeleteEntireAttendance(k *gin.Context) {
// 	KpiID, err := strconv.Atoi(k.Param("id"))
// 	if model.ErrorCheck(k, err){return}
// 	response, err := ka.crudService.GetByID(KpiID)
// 	if model.ErrorCheck(k, err) {return}

// 	//Delete Attendance
// 	err = ka.crudService.Delete(response.Year)
// 	if model.ErrorCheck(k, err) {return}
// 	//Delete Attendance Monthly
// 	if response.PlanID != nil {
// 		err = ka.crudService.Delete(*response.PlanID)
// 		if model.ErrorCheck(k, err) {return}
// 	}
// 	if response.ActualID != nil {
// 		err = ka.crudService.Delete(*response.ActualID)
// 		if model.ErrorCheck(k, err) {return}
// 	}
// 	if response.CutiID != nil {
// 		err = ka.crudService.Delete(*response.CutiID)
// 		if model.ErrorCheck(k, err) {return}
// 	}
// 	if response.IzinID != nil {
// 		err = ka.crudService.Delete(*response.IzinID)
// 		if model.ErrorCheck(k, err) {return}
// 	}
// 	if response.LainID != nil {
// 		err = ka.crudService.Delete(*response.LainID)
// 		if model.ErrorCheck(k, err) {return}
// 	}
// 	k.JSON(http.StatusOK, model.SuccessResponse{Message: "delete Entire Attendance success"})
// }


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

// Get List (Not Done)
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