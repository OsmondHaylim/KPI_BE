package api

import (
	// "fmt"
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

	AddEntireYearly(k *gin.Context)

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

	DeleteEntireYearly(k *gin.Context)

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
	GetYearlyList(k *gin.Context)}
type kpiAPI struct {
	attendanceService 	service.AttendanceService
	factorService 		service.FactorService
	itemService			service.ItemService
	minipapService		service.MiniPAPService
	monthlyService 		service.MonthlyService
	resultService		service.ResultService
	yearlyService		service.YearlyService}
func NewKpiAPI(
	attendanceService service.AttendanceService, 
	factorService service.FactorService, 
	itemService service.ItemService,
	minipapService service.MiniPAPService,
	monthlyService service.MonthlyService, 
	resultService service.ResultService,
	yearlyService service.YearlyService) *kpiAPI{
	return &kpiAPI{
		attendanceService, 
		factorService, 
		itemService,
		minipapService, 
		monthlyService, 
		resultService,
		yearlyService}}

// Add
func (ka *kpiAPI) AddAttendance(k *gin.Context) {
	var newAttendance model.Attendance
	if err := k.ShouldBindJSON(&newAttendance); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	err := ka.attendanceService.Store(&newAttendance)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Attendance success"})}
func (ka *kpiAPI) AddFactor(k *gin.Context) {
	var newFactor model.Factor
	if err := k.ShouldBindJSON(&newFactor); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	err := ka.factorService.Store(&newFactor)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Factor success"})}
func (ka *kpiAPI) AddItem(k *gin.Context) {
	var newItem model.Item
	if err := k.ShouldBindJSON(&newItem); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	err := ka.itemService.Store(&newItem)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Item success"})}
func (ka *kpiAPI) AddMinipap(k *gin.Context) {
	var newMinipap model.MiniPAP
	if err := k.ShouldBindJSON(&newMinipap); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	err := ka.minipapService.Store(&newMinipap)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Minipap success"})}
func (ka *kpiAPI) AddMonthly(k *gin.Context) {
	var newMonthly model.Monthly
	if err := k.ShouldBindJSON(&newMonthly); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	newMonthly = newMonthly.Reseted()
	err := ka.monthlyService.Store(&newMonthly)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Monthly success"})}
func (ka *kpiAPI) AddResult(k *gin.Context) {
	var newResult model.Result
	if err := k.ShouldBindJSON(&newResult); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	err := ka.resultService.Store(&newResult)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Result success"})}
func (ka *kpiAPI) AddYearly(k *gin.Context) {
	var newYearly model.Yearly
	if err := k.ShouldBindJSON(&newYearly); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	err := ka.yearlyService.Store(&newYearly)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Yearly success"})}

func (ka *kpiAPI) AddEntireYearly(k *gin.Context) {
	var newYearly model.YearlyResponse
	if err := k.ShouldBindJSON(&newYearly); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	//Reformat Response to Yearly
	var newYear model.Yearly
	newYear.Year = newYearly.Year

	//Storing Attendance
	var newAttendance model.Attendance
	newAttendance.Year = newYearly.Year
	//Creating monthly from attendance
	if newYearly.Attendance.Plan != nil{
		newMonthly := newYearly.Attendance.Plan.Reseted()
		err := ka.monthlyService.Store(&newMonthly)
		if err != nil {
			k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
			return
		}
		newAttendance.PlanID = &newMonthly.Monthly_ID
	}
	if newYearly.Attendance.Actual != nil{
		newMonthly := newYearly.Attendance.Actual.Reseted()
		err := ka.monthlyService.Store(&newMonthly)
		if err != nil {
			k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
			return
		}
		newAttendance.ActualID = &newMonthly.Monthly_ID
	}
	if newYearly.Attendance.Cuti != nil{
		newMonthly := newYearly.Attendance.Cuti.Reseted()
		err := ka.monthlyService.Store(&newMonthly)
		if err != nil {
			k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
			return
		}
		newAttendance.CutiID = &newMonthly.Monthly_ID
	}
	if newYearly.Attendance.Izin != nil{
		newMonthly := newYearly.Attendance.Izin.Reseted()
		err := ka.monthlyService.Store(&newMonthly)
		if err != nil {
			k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
			return
		}
		newAttendance.CutiID = &newMonthly.Monthly_ID
	}
	if newYearly.Attendance.Lain != nil{
		newMonthly := newYearly.Attendance.Lain.Reseted()
		err := ka.monthlyService.Store(&newMonthly)
		if err != nil {
			k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
			return
		}
		newAttendance.LainID = &newMonthly.Monthly_ID
	}
	//Creating attendance
	err := ka.attendanceService.Store(&newAttendance)
	if model.ErrorCheck(k, err) {return}
	newYear.AttendanceID = &newAttendance.Year
	err = ka.yearlyService.Store(&newYear)
	if model.ErrorCheck(k, err) {return}
	//Storing Items
	for _, item := range newYearly.Items{
		var newItem model.Item
		newItem.Name = item.Name
		//Connect Item to Yearly
		newItem.YearID = &newYear.Year
		//Creating Items to get id
		err := ka.itemService.Store(&newItem)
		if err != nil {
			k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
			return
		}
		//Storing Results
		for _, result := range item.Results{
			var newResult model.Result
			newResult.Name = result.Name
			//Connect Result to Item
			newResult.ItemID = &newItem.Item_ID
			//Creating Results to get id
			err := ka.resultService.Store(&newResult)
			if err != nil {
				k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
				return
			}
			//Storing Factors
			for _, factor := range result.Factors{
				var newFactor model.Factor
				newFactor.Title = factor.Title
				newFactor.Unit = factor.Unit
				newFactor.Target = factor.Target
				//Connect Factor to Result
				newFactor.ResultID = &newResult.Result_ID
				if factor.Plan != nil{
					//Storing MiniPAP Plan
					var newMinipap model.MiniPAP
					//Create MiniPAP to get id
					err := ka.minipapService.Store(&newMinipap)
					if model.ErrorCheck(k, err) {return}
					//Connect MiniPAP to Factor
					newFactor.PlanID = &newMinipap.MiniPAP_ID
					//Storing Plan Monthly
					for _, monthly := range factor.Plan.Monthly{
						newMonthly := monthly.Reseted()
						newMonthly.MinipapID = &newMinipap.MiniPAP_ID
						err := ka.monthlyService.Store(&newMonthly)
						if model.ErrorCheck(k, err) {return}
					}
				}
				if factor.Actual != nil{
					//Storing MiniPAP Actual
					var newMinipap model.MiniPAP
					//Create MiniPAP to get id
					err := ka.minipapService.Store(&newMinipap)
					if model.ErrorCheck(k, err) {return}
					//Connect MiniPAP to Factor
					newFactor.ActualID = &newMinipap.MiniPAP_ID
					//Storing Actual Monthly
					for _, monthly := range factor.Actual.Monthly{
						newMonthly := monthly.Reseted()
						newMonthly.MinipapID = &newMinipap.MiniPAP_ID
						err := ka.monthlyService.Store(&newMonthly)
						if model.ErrorCheck(k, err) {return}
					}
				}
				err := ka.factorService.Store(&newFactor)
				if model.ErrorCheck(k, err){return}
			}
		}
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Entire Yearly success"})}

// Update
func (ka *kpiAPI) UpdateAttendance(k *gin.Context) {
	var newAttendance model.Attendance
	if err := k.ShouldBindJSON(&newAttendance); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Attendance ID"})
		return
	}
	newAttendance.Year = KpiID
	err = ka.attendanceService.Saves(newAttendance)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Attendance update success"})}
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
	newFactor.Factor_ID = KpiID
	err = ka.factorService.Saves(newFactor)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Factor update success"})}
func (ka *kpiAPI) UpdateItem(k *gin.Context) {
	var newItem model.Item
	if err := k.ShouldBindJSON(&newItem); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Item ID"})
		return
	}
	newItem.Item_ID = KpiID
	err = ka.itemService.Saves(newItem)	
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Item update success"})}
func (ka *kpiAPI) UpdateMinipap(k *gin.Context) {
	var newMinipap model.MiniPAP
	if err := k.ShouldBindJSON(&newMinipap); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Minipap ID"})
		return
	}
	newMinipap.MiniPAP_ID = KpiID
	err = ka.minipapService.Saves(newMinipap)	
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "MiniPAP update success"})}
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
	// err = ka.monthlyService.Update(KpiID, newMonthly)
	newMonthly.Monthly_ID = KpiID
	err = ka.monthlyService.Saves(newMonthly)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Monthly update success"})}
func (ka *kpiAPI) UpdateResult(k *gin.Context) {
	var newResult model.Result
	if err := k.ShouldBindJSON(&newResult); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Result ID"})
		return
	}
	newResult.Result_ID = KpiID
	err = ka.resultService.Saves(newResult)	
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Result update success"})}
func (ka *kpiAPI) UpdateYearly(k *gin.Context) {
	var newYearly model.Yearly
	if err := k.ShouldBindJSON(&newYearly); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Yearly ID"})
		return
	}
	newYearly.Year = KpiID
	err = ka.yearlyService.Saves(newYearly)	
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Yearly update success"})}

// Delete
func (ka *kpiAPI) DeleteAttendance(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Attendance ID"})
		return
	}
	err = ka.attendanceService.Delete(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Attendance delete success"})}
func (ka *kpiAPI) DeleteFactor(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Factor ID"})
		return
	}
	err = ka.factorService.Delete(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Factor delete success"})}
func (ka *kpiAPI) DeleteItem(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Item ID"})
		return
	}
	err = ka.itemService.Delete(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Item delete success"})}
func (ka *kpiAPI) DeleteMinipap(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid MiniPAP ID"})
		return
	}
	err = ka.minipapService.Delete(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "MiniPAP delete success"})}
func (ka *kpiAPI) DeleteMonthly(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Monthly ID"})
		return
	}
	err = ka.monthlyService.Delete(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Monthly delete success"})}
func (ka *kpiAPI) DeleteResult(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Result ID"})
		return
	}
	err = ka.resultService.Delete(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Result delete success"})}
func (ka *kpiAPI) DeleteYearly(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Yearly ID"})
		return
	}
	err = ka.yearlyService.Delete(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Yearly delete success"})}

func (ka *kpiAPI) DeleteEntireYearly(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Yearly ID"})
		return
	}
	Yearly, err := ka.yearlyService.GetByID(KpiID)
	if model.ErrorCheck(k, err) {return}
	
	
	//Delete Items
	for _, item := range Yearly.Items{
		for _, result := range item.Results{
			for _, factor := range result.Factors{
				for _, monthly := range factor.Plan.Monthly{
					err = ka.monthlyService.Delete(monthly.Monthly_ID)
					if model.ErrorCheck(k, err) {return}
				}
				err = ka.minipapService.Delete(*factor.PlanID)
				if model.ErrorCheck(k, err) {return}
				for _, monthly := range factor.Actual.Monthly{
					err = ka.monthlyService.Delete(monthly.Monthly_ID)
					if model.ErrorCheck(k, err) {return}
				}
				err = ka.minipapService.Delete(*factor.ActualID)
				if model.ErrorCheck(k, err) {return}
				err = ka.factorService.Delete(factor.Factor_ID)
				if model.ErrorCheck(k, err) {return}
			}
			err = ka.resultService.Delete(result.Result_ID)
			if model.ErrorCheck(k, err) {return}
		}
		err = ka.itemService.Delete(item.Item_ID)
		if model.ErrorCheck(k, err) {return}
	}
	//Delete Yearly
	err = ka.yearlyService.Delete(Yearly.Year)
	if model.ErrorCheck(k, err) {return}

	//Delete Attendance
	if Yearly.AttendanceID != nil {
		err = ka.attendanceService.Delete(Yearly.Attendance.Year)
		if model.ErrorCheck(k, err) {return}
	}
	//Delete Attendance Monthly
	if Yearly.Attendance.PlanID != nil {
		err = ka.monthlyService.Delete(*Yearly.Attendance.PlanID)
		if model.ErrorCheck(k, err) {return}
	}
	if Yearly.Attendance.ActualID != nil {
		err = ka.monthlyService.Delete(*Yearly.Attendance.ActualID)
		if model.ErrorCheck(k, err) {return}
	}
	if Yearly.Attendance.CutiID != nil {
		err = ka.monthlyService.Delete(*Yearly.Attendance.CutiID)
		if model.ErrorCheck(k, err) {return}
	}
	if Yearly.Attendance.IzinID != nil {
		err = ka.monthlyService.Delete(*Yearly.Attendance.IzinID)
		if model.ErrorCheck(k, err) {return}
	}
	if Yearly.Attendance.PlanID != nil {
		err = ka.monthlyService.Delete(*Yearly.Attendance.LainID)
		if model.ErrorCheck(k, err) {return}
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "delete Entire Yearly success"})
}

// Get By ID
func (ka *kpiAPI) GetAttendanceByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid Attendance Year"})
		return
	}
	Attendance, err := ka.attendanceService.GetByID(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, Attendance.ToResponse())}
func (ka *kpiAPI) GetFactorByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid Factor ID"})
		return
	}
	Factor, err := ka.factorService.GetByID(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, Factor.ToResponse())}
func (ka *kpiAPI) GetItemByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid Item ID"})
		return
	}
	Item, err := ka.itemService.GetByID(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, Item.ToResponse())}
func (ka *kpiAPI) GetMinipapByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid Minipap ID"})
		return
	}
	Minipap, err := ka.minipapService.GetByID(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, Minipap)}
func (ka *kpiAPI) GetMonthlyByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid Monthly ID"})
		return
	}
	Monthly, err := ka.monthlyService.GetByID(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, Monthly)}
func (ka *kpiAPI) GetResultByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid Result ID"})
		return
	}
	Result, err := ka.resultService.GetByID(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, Result.ToResponse())}
func (ka *kpiAPI) GetYearlyByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid Yearly ID"})
		return
	}
	Yearly, err := ka.yearlyService.GetByID(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, Yearly.ToResponse())}

// Get List
func (ka *kpiAPI) GetAttendanceList(k *gin.Context) {
	Attendance, err := ka.attendanceService.GetList()
	if model.ErrorCheck(k, err) {return}
	var result model.AttendanceArrayResponse
	result.Attendance = []model.AttendanceResponse{} 
	for _,att := range Attendance{
		result.Attendance = append(result.Attendance, att.ToResponse())	
	}
	result.Message = "Getting All Attendances Success"
	k.JSON(http.StatusOK, result)}
func (ka *kpiAPI) GetFactorList(k *gin.Context) {
	Factor, err := ka.factorService.GetList()
	if model.ErrorCheck(k, err) {return}
	var result model.FactorArrayResponse
	result.Factor = []model.FactorResponse{}
	for _, data := range Factor{
		result.Factor = append(result.Factor, data.ToResponse())
	} 
	result.Message = "Getting All Factors Success"
	k.JSON(http.StatusOK, result)}
func (ka *kpiAPI) GetItemList(k *gin.Context) {
	Item, err := ka.itemService.GetList()
	if model.ErrorCheck(k, err) {return}
	var result model.ItemArrayResponse
	result.Item = []model.ItemResponse{}
	for _, data := range Item{
		result.Item = append(result.Item, data.ToResponse())
	} 
	result.Message = "Getting All Items Success"
	k.JSON(http.StatusOK, result)}
func (ka *kpiAPI) GetMinipapList(k *gin.Context) {
	Minipap, err := ka.minipapService.GetList()
	if model.ErrorCheck(k, err) {return}
	var result model.MinipapArrayResponse
	result.Minipap = Minipap 
	result.Message = "Getting All MiniPAPs Success"
	k.JSON(http.StatusOK, result)}
func (ka *kpiAPI) GetMonthlyList(k *gin.Context) {
	Monthly, err := ka.monthlyService.GetList()
	if model.ErrorCheck(k, err) {return}
	var result model.MonthlyArrayResponse
	result.Monthly = Monthly 
	result.Message = "Getting All Monthlys Success"
	k.JSON(http.StatusOK, result)}
func (ka *kpiAPI) GetResultList(k *gin.Context) {
	Result, err := ka.resultService.GetList()
	if model.ErrorCheck(k, err) {return}
	var result model.ResultArrayResponse
	result.Result = []model.ResultResponse{}
	for _, data := range Result{
		result.Result = append(result.Result, data.ToResponse())
	} 
	result.Message = "Getting All Results Success"
	k.JSON(http.StatusOK, result)}
func (ka *kpiAPI) GetYearlyList(k *gin.Context) {
	Yearly, err := ka.yearlyService.GetList()
	if model.ErrorCheck(k, err) {return}
	var result model.YearlyArrayResponse
	result.Yearly = []model.YearlyResponse{}
	for _, data := range Yearly{
		result.Yearly = append(result.Yearly, data.ToResponse())
	} 
	result.Message = "Getting All Yearlys Success"
	k.JSON(http.StatusOK, result)}