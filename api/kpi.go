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

type KpiAPI interface {
	AddAttendance(k *gin.Context)
	AddFactor(k *gin.Context)
	AddItem(k *gin.Context)
	AddMinipap(k *gin.Context)
	AddMonthly(k *gin.Context)
	// AddPap(k *gin.Context)
	AddResult(k *gin.Context)
	AddYearly(k *gin.Context)

	UpdateAttendance(k *gin.Context)
	UpdateFactor(k *gin.Context)
	UpdateItem(k *gin.Context)
	UpdateMinipap(k *gin.Context)
	UpdateMonthly(k *gin.Context)
	// UpdatePap(k *gin.Context)
	UpdateResult(k *gin.Context)
	UpdateYearly(k *gin.Context)

	DeleteAttendance(k *gin.Context)
	DeleteFactor(k *gin.Context)
	DeleteItem(k *gin.Context)
	DeleteMinipap(k *gin.Context)
	DeleteMonthly(k *gin.Context)
	// DeletePap(k *gin.Context)
	DeleteResult(k *gin.Context)
	DeleteYearly(k *gin.Context)

	GetAttendanceByID(k *gin.Context)
	GetFactorByID(k *gin.Context)
	GetItemByID(k *gin.Context)
	GetMinipapByID(k *gin.Context)
	GetMonthlyByID(k *gin.Context)
	// GetPapByID(k *gin.Context)
	GetResultByID(k *gin.Context)
	GetYearlyByID(k *gin.Context)

	GetAttendanceList(k *gin.Context)
	GetFactorList(k *gin.Context)
	GetItemList(k *gin.Context)
	GetMinipapList(k *gin.Context)
	GetMonthlyList(k *gin.Context)
	// GetPapList(k *gin.Context)
	GetResultList(k *gin.Context)
	GetYearlyList(k *gin.Context)
}

type kpiAPI struct {
	attendanceService 	service.AttendanceService
	factorService 		service.FactorService
	itemService			service.ItemService
	minipapService		service.MiniPAPService
	monthlyService 		service.MonthlyService
	// papService			service.PAPService
	resultService		service.ResultService
	yearlyService		service.YearlyService
}

func NewKpiAPI(
	attendanceService service.AttendanceService, 
	factorService service.FactorService, 
	itemService service.ItemService,
	minipapService service.MiniPAPService,
	monthlyService service.MonthlyService, 
	// papService service.PAPService,
	resultService service.ResultService,
	yearlyService service.YearlyService) *kpiAPI{
	return &kpiAPI{
		attendanceService, 
		factorService, 
		itemService,
		minipapService, 
		monthlyService, 
		// papService,
		resultService,
		yearlyService}
}

// Add
func (ka *kpiAPI) AddAttendance(k *gin.Context) {
	var newAttendance model.Attendance
	if err := k.ShouldBindJSON(&newAttendance); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	err := ka.attendanceService.Store(&newAttendance)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Attendance success"})
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
func (ka *kpiAPI) AddItem(k *gin.Context) {
	var newItem model.Item
	if err := k.ShouldBindJSON(&newItem); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	err := ka.itemService.Store(&newItem)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Item success"})
}
func (ka *kpiAPI) AddMinipap(k *gin.Context) {
	var newMinipap model.MiniPAP
	if err := k.ShouldBindJSON(&newMinipap); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	err := ka.minipapService.Store(&newMinipap)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Minipap success"})
}
func (ka *kpiAPI) AddMonthly(k *gin.Context) {
	var newMonthly model.Monthly
	if err := k.ShouldBindJSON(&newMonthly); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	newMonthly.Ytd = 
		newMonthly.Jan + 
		newMonthly.Feb +
		newMonthly.Mar +
		newMonthly.Apr +
		newMonthly.May +
		newMonthly.Jun +
		newMonthly.Jul +
		newMonthly.Aug +
		newMonthly.Sep +
		newMonthly.Oct +
		newMonthly.Nov + 
		newMonthly.Dec
	err := ka.monthlyService.Store(&newMonthly)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Monthly success"})
}
// func (ka *kpiAPI) AddPap(k *gin.Context) {
// 	var newPap model.PAP
// 	if err := k.ShouldBindJSON(&newPap); err != nil {
// 		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
// 		return
// 	}
// 	// model.ToPercentage(newPap)
// 	err := ka.papService.Store(&newPap)
// 	if err != nil {
// 		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
// 		return
// 	}
// 	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add PAP success"})
// }
func (ka *kpiAPI) AddResult(k *gin.Context) {
	var newResult model.Result
	if err := k.ShouldBindJSON(&newResult); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	err := ka.resultService.Store(&newResult)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Result success"})
}
func (ka *kpiAPI) AddYearly(k *gin.Context) {
	var newYearly model.Yearly
	if err := k.ShouldBindJSON(&newYearly); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	err := ka.yearlyService.Store(&newYearly)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Yearly success"})
}

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
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Attendance update success"})
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
	newFactor.Factor_ID = KpiID
	err = ka.factorService.Saves(newFactor)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Factor update success"})
}
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
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Item update success"})
}
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
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "MiniPAP update success"})
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
	// err = ka.monthlyService.Update(KpiID, newMonthly)
	newMonthly.Monthly_ID = KpiID
	err = ka.monthlyService.Saves(newMonthly)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Monthly update success"})
}
// func (ka *kpiAPI) UpdatePap(k *gin.Context) {
// 	var newPap model.PAP
// 	if err := k.ShouldBindJSON(&newPap); err != nil {
// 		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
// 		return
// 	}
// 	KpiID, err := strconv.Atoi(k.Param("id"))
// 	if err != nil {
// 		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Pap ID"})
// 		return
// 	}
// 	newPap.Pap_ID = KpiID
// 	err = ka.papService.Saves(newPap)	
// 	if err != nil {
// 		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
// 		return
// 	}
// 	k.JSON(http.StatusOK, model.SuccessResponse{Message: "PAP update success"})
// }
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
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Result update success"})
}
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
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Yearly update success"})
}

// Delete
func (ka *kpiAPI) DeleteAttendance(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Attendance ID"})
		return
	}
	err = ka.attendanceService.Delete(KpiID)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Attendance delete success"})
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
func (ka *kpiAPI) DeleteItem(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Item ID"})
		return
	}
	err = ka.itemService.Delete(KpiID)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Item delete success"})
}
func (ka *kpiAPI) DeleteMinipap(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid MiniPAP ID"})
		return
	}
	err = ka.minipapService.Delete(KpiID)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "MiniPAP delete success"})
}
func (ka *kpiAPI) DeleteMonthly(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Monthly ID"})
		return
	}
	err = ka.monthlyService.Delete(KpiID)
	if err != nil {
		// if strings.Contains(err.Error(), "foreign key constraint") {
		// 	Attendance, where, err2 := ka.attendanceService.GetAttendanceFromMonthly(KpiID)
		// 	if err2 != nil {
		// 		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		// 		return
		// 	}
		// 	were := *where
		// 	switch were{
		// 	case "plan_id":
		// 		Attendance.PlanID = nil
		// 		Attendance.Plan = nil
		// 	case "actual_id":
		// 		Attendance.ActualID = nil
		// 		Attendance.Actual = nil
		// 	case "cuti_id":
		// 		Attendance.CutiID = nil
		// 		Attendance.Cuti = nil
		// 	case "izin_id":	
		// 		Attendance.IzinID = nil
		// 		Attendance.Izin = nil				
		// 	case "lain_id":	
		// 		Attendance.LainID = nil
		// 		Attendance.Lain = nil
		// 	default:
		// 	}
		// 	newAtt := *Attendance
		// 	err2 = ka.attendanceService.Saves(newAtt)
		// 	if err2 != nil {
		// 		year := strconv.Itoa(Attendance.Year)
		// 		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error() + ", Attendance " + year + "'s " + were + " Deleted"})
		// 		return
		// 	}
		// 	err = ka.monthlyService.Delete(KpiID)
		// 	if err != nil {
		// 		year := strconv.Itoa(Attendance.Year)
		// 		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error() + ", Attendance " + year + "'s " + were + " Updated"})
		// 		return
		// 	}
		// 	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Monthly delete success"})
		// 	return
		// }
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Monthly delete success"})
}
// func (ka *kpiAPI) DeletePap(k *gin.Context) {
// 	KpiID, err := strconv.Atoi(k.Param("id"))
// 	if err != nil {
// 		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid PAP ID"})
// 		return
// 	}
// 	err = ka.papService.Delete(KpiID)
// 	if err != nil {
// 		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
// 		return
// 	}
// 	k.JSON(http.StatusOK, model.SuccessResponse{Message: "PAP delete success"})
// }
func (ka *kpiAPI) DeleteResult(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Result ID"})
		return
	}
	err = ka.resultService.Delete(KpiID)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Result delete success"})
}
func (ka *kpiAPI) DeleteYearly(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Yearly ID"})
		return
	}
	err = ka.yearlyService.Delete(KpiID)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Yearly delete success"})
}

// Get By ID
func (ka *kpiAPI) GetAttendanceByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid Attendance Year"})
		return
	}
	Attendance, err := ka.attendanceService.GetByID(KpiID)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, Attendance.ToResponse())
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
	k.JSON(http.StatusOK, Factor.ToResponse())
}
func (ka *kpiAPI) GetItemByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid Item ID"})
		return
	}
	Item, err := ka.itemService.GetByID(KpiID)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, Item.ToResponse())
}
func (ka *kpiAPI) GetMinipapByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid Minipap ID"})
		return
	}
	Minipap, err := ka.minipapService.GetByID(KpiID)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, Minipap)
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
// func (ka *kpiAPI) GetPapByID(k *gin.Context) {
// 	KpiID, err := strconv.Atoi(k.Param("id"))
// 	if err != nil {
// 		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid Pap ID"})
// 		return
// 	}
// 	Pap, err := ka.papService.GetByID(KpiID)
// 	if err != nil {
// 		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
// 		return
// 	}
// 	k.JSON(http.StatusOK, Pap.ToResponse())
// }
func (ka *kpiAPI) GetResultByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid Result ID"})
		return
	}
	Result, err := ka.resultService.GetByID(KpiID)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, Result.ToResponse())
}
func (ka *kpiAPI) GetYearlyByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid Yearly ID"})
		return
	}
	Yearly, err := ka.yearlyService.GetByID(KpiID)
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	k.JSON(http.StatusOK, Yearly.ToResponse())
}

// Get List
func (ka *kpiAPI) GetAttendanceList(k *gin.Context) {
	Attendance, err := ka.attendanceService.GetList()
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	var result model.AttendanceArrayResponse
	result.Attendance = []model.AttendanceResponse{} 
	for _,att := range Attendance{
		result.Attendance = append(result.Attendance, att.ToResponse())	
	}
	result.Message = "Getting All Attendances Success"
	k.JSON(http.StatusOK, result)
}
func (ka *kpiAPI) GetFactorList(k *gin.Context) {
	Factor, err := ka.factorService.GetList()
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	var result model.FactorArrayResponse
	result.Factor = []model.FactorResponse{}
	for _, data := range Factor{
		result.Factor = append(result.Factor, data.ToResponse())
	} 
	result.Message = "Getting All Factors Success"
	k.JSON(http.StatusOK, result)
}
func (ka *kpiAPI) GetItemList(k *gin.Context) {
	Item, err := ka.itemService.GetList()
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	var result model.ItemArrayResponse
	result.Item = []model.ItemResponse{}
	for _, data := range Item{
		result.Item = append(result.Item, data.ToResponse())
	} 
	result.Message = "Getting All Items Success"
	k.JSON(http.StatusOK, result)
}
func (ka *kpiAPI) GetMinipapList(k *gin.Context) {
	Minipap, err := ka.minipapService.GetList()
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	var result model.MinipapArrayResponse
	result.Minipap = Minipap 
	result.Message = "Getting All MiniPAPs Success"
	k.JSON(http.StatusOK, result)
}
func (ka *kpiAPI) GetMonthlyList(k *gin.Context) {
	Monthly, err := ka.monthlyService.GetList()
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	var result model.MonthlyArrayResponse
	result.Monthly = Monthly 
	result.Message = "Getting All Monthlys Success"
	k.JSON(http.StatusOK, result)
}
// func (ka *kpiAPI) GetPapList(k *gin.Context) {
// 	Pap, err := ka.papService.GetList()
// 	if err != nil {
// 		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
// 		return
// 	}
// 	var result model.PapArrayResponse
// 	result.Pap = []model.PAPResponse{}
// 	for _, p := range Pap{
// 		result.Pap = append(result.Pap, p.ToResponse())
// 	} 
// 	result.Message = "Getting All PAPs Success"
// 	k.JSON(http.StatusOK, result)
// }
func (ka *kpiAPI) GetResultList(k *gin.Context) {
	Result, err := ka.resultService.GetList()
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	var result model.ResultArrayResponse
	result.Result = []model.ResultResponse{}
	for _, data := range Result{
		result.Result = append(result.Result, data.ToResponse())
	} 
	result.Message = "Getting All Results Success"
	k.JSON(http.StatusOK, result)
}
func (ka *kpiAPI) GetYearlyList(k *gin.Context) {
	Yearly, err := ka.yearlyService.GetList()
	if err != nil {
		k.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	var result model.YearlyArrayResponse
	result.Yearly = []model.YearlyResponse{}
	for _, data := range Yearly{
		result.Yearly = append(result.Yearly, data.ToResponse())
	} 
	result.Message = "Getting All Yearlys Success"
	k.JSON(http.StatusOK, result)
}