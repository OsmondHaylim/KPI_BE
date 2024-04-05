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
	return &kpiAPI{crudService}}

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

// Add Entire
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
		err := ka.crudService.Store(&newMonthly)
		if model.ErrorCheck(k, err){return}
		newAttendance.PlanID = &newMonthly.Monthly_ID
	}
	if newYearly.Attendance.Actual != nil{
		newMonthly := newYearly.Attendance.Actual.Reseted()
		err := ka.crudService.Store(&newMonthly)
		if model.ErrorCheck(k, err){return}
		newAttendance.ActualID = &newMonthly.Monthly_ID
	}
	if newYearly.Attendance.Cuti != nil{
		newMonthly := newYearly.Attendance.Cuti.Reseted()
		err := ka.crudService.Store(&newMonthly)
		if model.ErrorCheck(k, err){return}
		newAttendance.CutiID = &newMonthly.Monthly_ID
	}
	if newYearly.Attendance.Izin != nil{
		newMonthly := newYearly.Attendance.Izin.Reseted()
		err := ka.crudService.Store(&newMonthly)
		if model.ErrorCheck(k, err){return}
		newAttendance.CutiID = &newMonthly.Monthly_ID
	}
	if newYearly.Attendance.Lain != nil{
		newMonthly := newYearly.Attendance.Lain.Reseted()
		err := ka.crudService.Store(&newMonthly)
		if model.ErrorCheck(k, err){return}
		newAttendance.LainID = &newMonthly.Monthly_ID
	}
	//Creating attendance
	err := ka.crudService.Store(&newAttendance)
	if model.ErrorCheck(k, err) {return}
	//Link Attendance to Yearly
	newYear.AttendanceID = &newAttendance.Year
	//Create Yearly
	err = ka.crudService.Store(&newYear)
	if model.ErrorCheck(k, err) {return}
	//Storing Items
	for _, item := range newYearly.Items{
		var newItem model.Item
		newItem.Name = item.Name
		//Connect Item to Yearly
		newItem.YearID = &newYear.Year
		//Creating Items to get id
		err := ka.crudService.Store(&newItem)
		if model.ErrorCheck(k, err){return}
		//Storing Results
		for _, result := range item.Results{
			var newResult model.Result
			newResult.Name = result.Name
			//Connect Result to Item
			newResult.ItemID = &newItem.Item_ID
			//Creating Results to get id
			err := ka.crudService.Store(&newResult)
			if model.ErrorCheck(k, err){return}
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
					err := ka.crudService.Store(&newMinipap)
					if model.ErrorCheck(k, err) {return}
					//Connect MiniPAP to Factor
					newFactor.PlanID = &newMinipap.MiniPAP_ID
					//Storing Plan Monthly
					for _, monthly := range factor.Plan.Monthly{
						newMonthly := monthly.Reseted()
						newMonthly.MinipapID = &newMinipap.MiniPAP_ID
						err := ka.crudService.Store(&newMonthly)
						if model.ErrorCheck(k, err) {return}
					}
				}
				if factor.Actual != nil{
					//Storing MiniPAP Actual
					var newMinipap model.MiniPAP
					//Create MiniPAP to get id
					err := ka.crudService.Store(&newMinipap)
					if model.ErrorCheck(k, err) {return}
					//Connect MiniPAP to Factor
					newFactor.ActualID = &newMinipap.MiniPAP_ID
					//Storing Actual Monthly
					for _, monthly := range factor.Actual.Monthly{
						newMonthly := monthly.Reseted()
						newMonthly.MinipapID = &newMinipap.MiniPAP_ID
						err := ka.crudService.Store(&newMonthly)
						if model.ErrorCheck(k, err) {return}
					}
				}
				err := ka.crudService.Store(&newFactor)
				if model.ErrorCheck(k, err){return}
			}
		}
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Entire Yearly success"})
}
func (ka *kpiAPI) AddEntireItem(k *gin.Context) {
	var response model.ItemResponse
	if err := k.ShouldBindJSON(&response); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	//Storing Items
	var newItem model.Item
	newItem.Name = response.Name
	//Creating Items to get id
	err := ka.crudService.Store(&newItem)
	if model.ErrorCheck(k, err){return}
	//Storing Results
	for _, result := range response.Results{
		var newResult model.Result
		newResult.Name = result.Name
		//Connect Result to Item
		newResult.ItemID = &newItem.Item_ID
		//Creating Results to get id
		err := ka.crudService.Store(&newResult)
		if model.ErrorCheck(k, err){return}
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
				err := ka.crudService.Store(&newMinipap)
				if model.ErrorCheck(k, err) {return}
				//Connect MiniPAP to Factor
				newFactor.PlanID = &newMinipap.MiniPAP_ID
				//Storing Plan Monthly
				for _, monthly := range factor.Plan.Monthly{
					newMonthly := monthly.Reseted()
					newMonthly.MinipapID = &newMinipap.MiniPAP_ID
					err := ka.crudService.Store(&newMonthly)
					if model.ErrorCheck(k, err) {return}
				}
			}
			if factor.Actual != nil{
				//Storing MiniPAP Actual
				var newMinipap model.MiniPAP
				//Create MiniPAP to get id
				err := ka.crudService.Store(&newMinipap)
				if model.ErrorCheck(k, err) {return}
				//Connect MiniPAP to Factor
				newFactor.ActualID = &newMinipap.MiniPAP_ID
				//Storing Actual Monthly
				for _, monthly := range factor.Actual.Monthly{
					newMonthly := monthly.Reseted()
					newMonthly.MinipapID = &newMinipap.MiniPAP_ID
					err := ka.crudService.Store(&newMonthly)
					if model.ErrorCheck(k, err) {return}
				}
			}
			err := ka.crudService.Store(&newFactor)
			if model.ErrorCheck(k, err){return}
		}
	}

	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Entire Item success"})
}
func (ka *kpiAPI) AddEntireResult(k *gin.Context) {
	var response model.ResultResponse
	if err := k.ShouldBindJSON(&response); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	//Storing Results
	var newResult model.Result
	newResult.Name = response.Name
	//Creating Results to get id
	err := ka.crudService.Store(&newResult)
	if model.ErrorCheck(k, err){return}
	//Storing Factors
	for _, factor := range response.Factors{
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
			err := ka.crudService.Store(&newMinipap)
			if model.ErrorCheck(k, err) {return}
			//Connect MiniPAP to Factor
			newFactor.PlanID = &newMinipap.MiniPAP_ID
			//Storing Plan Monthly
			for _, monthly := range factor.Plan.Monthly{
				newMonthly := monthly.Reseted()
				newMonthly.MinipapID = &newMinipap.MiniPAP_ID
				err := ka.crudService.Store(&newMonthly)
				if model.ErrorCheck(k, err) {return}
			}
		}
		if factor.Actual != nil{
			//Storing MiniPAP Actual
			var newMinipap model.MiniPAP
			//Create MiniPAP to get id
			err := ka.crudService.Store(&newMinipap)
			if model.ErrorCheck(k, err) {return}
			//Connect MiniPAP to Factor
			newFactor.ActualID = &newMinipap.MiniPAP_ID
			//Storing Actual Monthly
			for _, monthly := range factor.Actual.Monthly{
				newMonthly := monthly.Reseted()
				newMonthly.MinipapID = &newMinipap.MiniPAP_ID
				err := ka.crudService.Store(&newMonthly)
				if model.ErrorCheck(k, err) {return}
			}
		}
		err := ka.crudService.Store(&newFactor)
		if model.ErrorCheck(k, err){return}
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Entire Result success"})
}
func (ka *kpiAPI) AddEntireFactor(k *gin.Context) {
	var response model.FactorResponse
	if err := k.ShouldBindJSON(&response); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	var newFactor model.Factor
	newFactor.Title = response.Title
	newFactor.Unit = response.Unit
	newFactor.Target = response.Target
	if response.Plan != nil{
		//Storing MiniPAP Plan
		var newMinipap model.MiniPAP
		//Create MiniPAP to get id
		err := ka.crudService.Store(&newMinipap)
		if model.ErrorCheck(k, err) {return}
		//Connect MiniPAP to Factor
		newFactor.PlanID = &newMinipap.MiniPAP_ID
		//Storing Plan Monthly
		for _, monthly := range response.Plan.Monthly{
			newMonthly := monthly.Reseted()
			newMonthly.MinipapID = &newMinipap.MiniPAP_ID
			err := ka.crudService.Store(&newMonthly)
			if model.ErrorCheck(k, err) {return}
		}
	}
	if response.Actual != nil{
		//Storing MiniPAP Actual
		var newMinipap model.MiniPAP
		//Create MiniPAP to get id
		err := ka.crudService.Store(&newMinipap)
		if model.ErrorCheck(k, err) {return}
		//Connect MiniPAP to Factor
		newFactor.ActualID = &newMinipap.MiniPAP_ID
		//Storing Actual Monthly
		for _, monthly := range response.Actual.Monthly{
			newMonthly := monthly.Reseted()
			newMonthly.MinipapID = &newMinipap.MiniPAP_ID
			err := ka.crudService.Store(&newMonthly)
			if model.ErrorCheck(k, err) {return}
		}
	}
	err := ka.crudService.Store(&newFactor)
	if model.ErrorCheck(k, err){return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "add Entire Factor success"})
}
func (ka *kpiAPI) AddEntireAttendance(k *gin.Context) {
	var response model.AttendanceResponse
	if err := k.ShouldBindJSON(&response); err != nil {
		k.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	//Storing Attendance
	var newAttendance model.Attendance
	newAttendance.Year = response.Year
	//Creating monthly from attendance
	if response.Plan != nil{
		newMonthly := response.Plan.Reseted()
		err := ka.crudService.Store(&newMonthly)
		if model.ErrorCheck(k, err){return}
		newAttendance.PlanID = &newMonthly.Monthly_ID
	}
	if response.Actual != nil{
		newMonthly := response.Actual.Reseted()
		err := ka.crudService.Store(&newMonthly)
		if model.ErrorCheck(k, err){return}
		newAttendance.ActualID = &newMonthly.Monthly_ID
	}
	if response.Cuti != nil{
		newMonthly := response.Cuti.Reseted()
		err := ka.crudService.Store(&newMonthly)
		if model.ErrorCheck(k, err){return}
		newAttendance.CutiID = &newMonthly.Monthly_ID
	}
	if response.Izin != nil{
		newMonthly := response.Izin.Reseted()
		err := ka.crudService.Store(&newMonthly)
		if model.ErrorCheck(k, err){return}
		newAttendance.IzinID = &newMonthly.Monthly_ID
	}
	if response.Lain != nil{
		newMonthly := response.Lain.Reseted()
		err := ka.crudService.Store(&newMonthly)
		if model.ErrorCheck(k, err){return}
		newAttendance.LainID = &newMonthly.Monthly_ID
	}
	//Creating attendance
	err := ka.crudService.Store(&newAttendance)
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
	newAttendance.Year = KpiID
	err = ka.crudService.UpdateAttendance(newAttendance)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Attendance update success"})
}
func (ka *kpiAPI) UpdateFactor(k *gin.Context) {
	var newFactor model.Factor
	err := k.ShouldBindJSON(&newFactor)
	if model.ErrorCheck(k, err){return}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	newFactor.Factor_ID = KpiID
	err = ka.crudService.UpdateFactor(newFactor)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Factor update success"})
}
func (ka *kpiAPI) UpdateItem(k *gin.Context) {
	var newItem model.Item
	err := k.ShouldBindJSON(&newItem)
	if model.ErrorCheck(k, err){return}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	newItem.Item_ID = KpiID
	err = ka.crudService.UpdateItem(newItem)	
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Item update success"})
}
func (ka *kpiAPI) UpdateMinipap(k *gin.Context) {
	var newMinipap model.MiniPAP
	err := k.ShouldBindJSON(&newMinipap)
	if model.ErrorCheck(k, err){return}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	newMinipap.MiniPAP_ID = KpiID
	err = ka.crudService.UpdateMinipap(newMinipap)	
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "MiniPAP update success"})
}
func (ka *kpiAPI) UpdateMonthly(k *gin.Context) {
	var newMonthly model.Monthly
	err := k.ShouldBindJSON(&newMonthly)
	if model.ErrorCheck(k, err){return}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	// err = ka.crudService.Update(KpiID, newMonthly)
	newMonthly.Monthly_ID = KpiID
	err = ka.crudService.UpdateMonthly(newMonthly)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Monthly update success"})
}
func (ka *kpiAPI) UpdateResult(k *gin.Context) {
	var newResult model.Result
	err := k.ShouldBindJSON(&newResult)
	if model.ErrorCheck(k, err){return}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	newResult.Result_ID = KpiID
	err = ka.crudService.UpdateResult(newResult)	
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "Result update success"})
}
func (ka *kpiAPI) UpdateYearly(k *gin.Context) {
	var newYearly model.Yearly
	err := k.ShouldBindJSON(&newYearly)
	if model.ErrorCheck(k, err){return}
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	newYearly.Year = KpiID
	err = ka.crudService.UpdateYearly(newYearly)	
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

// Delete Entire
func (ka *kpiAPI) DeleteEntireYearly(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	Yearly, err := ka.crudService.GetByID(KpiID)
	if model.ErrorCheck(k, err) {return}
	
	
	//Delete Items
	for _, item := range Yearly.Items{
		for _, result := range item.Results{
			for _, factor := range result.Factors{
				for _, monthly := range factor.Plan.Monthly{
					err = ka.crudService.Delete(monthly.Monthly_ID)
					if model.ErrorCheck(k, err) {return}
				}
				err = ka.crudService.Delete(*factor.PlanID)
				if model.ErrorCheck(k, err) {return}
				for _, monthly := range factor.Actual.Monthly{
					err = ka.crudService.Delete(monthly.Monthly_ID)
					if model.ErrorCheck(k, err) {return}
				}
				err = ka.crudService.Delete(*factor.ActualID)
				if model.ErrorCheck(k, err) {return}
				err = ka.crudService.Delete(factor.Factor_ID)
				if model.ErrorCheck(k, err) {return}
			}
			err = ka.crudService.Delete(result.Result_ID)
			if model.ErrorCheck(k, err) {return}
		}
		err = ka.crudService.Delete(item.Item_ID)
		if model.ErrorCheck(k, err) {return}
	}
	//Delete Yearly
	err = ka.crudService.Delete(Yearly.Year)
	if model.ErrorCheck(k, err) {return}

	//Delete Attendance
	if Yearly.AttendanceID != nil {
		err = ka.crudService.Delete(Yearly.Attendance.Year)
		if model.ErrorCheck(k, err) {return}
	}
	//Delete Attendance Monthly
	if Yearly.Attendance.PlanID != nil {
		err = ka.crudService.Delete(*Yearly.Attendance.PlanID)
		if model.ErrorCheck(k, err) {return}
	}
	if Yearly.Attendance.ActualID != nil {
		err = ka.crudService.Delete(*Yearly.Attendance.ActualID)
		if model.ErrorCheck(k, err) {return}
	}
	if Yearly.Attendance.CutiID != nil {
		err = ka.crudService.Delete(*Yearly.Attendance.CutiID)
		if model.ErrorCheck(k, err) {return}
	}
	if Yearly.Attendance.IzinID != nil {
		err = ka.crudService.Delete(*Yearly.Attendance.IzinID)
		if model.ErrorCheck(k, err) {return}
	}
	if Yearly.Attendance.LainID != nil {
		err = ka.crudService.Delete(*Yearly.Attendance.LainID)
		if model.ErrorCheck(k, err) {return}
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "delete Entire Yearly success"})
}
func (ka *kpiAPI) DeleteEntireItem(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	item, err := ka.crudService.GetByID(KpiID)
	if model.ErrorCheck(k, err) {return}
	//Delete Items
	for _, result := range item.Results{
		for _, factor := range result.Factors{
			for _, monthly := range factor.Plan.Monthly{
				err = ka.crudService.Delete(monthly.Monthly_ID)
				if model.ErrorCheck(k, err) {return}
			}
			err = ka.crudService.Delete(*factor.PlanID)
			if model.ErrorCheck(k, err) {return}
			for _, monthly := range factor.Actual.Monthly{
				err = ka.crudService.Delete(monthly.Monthly_ID)
				if model.ErrorCheck(k, err) {return}
			}
			err = ka.crudService.Delete(*factor.ActualID)
			if model.ErrorCheck(k, err) {return}
			err = ka.crudService.Delete(factor.Factor_ID)
			if model.ErrorCheck(k, err) {return}
		}
		err = ka.crudService.Delete(result.Result_ID)
		if model.ErrorCheck(k, err) {return}
	}
	err = ka.crudService.Delete(item.Item_ID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "delete Entire Item success"})
}
func (ka *kpiAPI) DeleteEntireResult(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	result, err := ka.crudService.GetByID(KpiID)
	if model.ErrorCheck(k, err) {return}
	//Delete Results
	for _, factor := range result.Factors{
		for _, monthly := range factor.Plan.Monthly{
			err = ka.crudService.Delete(monthly.Monthly_ID)
			if model.ErrorCheck(k, err) {return}
		}
		err = ka.crudService.Delete(*factor.PlanID)
		if model.ErrorCheck(k, err) {return}
		for _, monthly := range factor.Actual.Monthly{
			err = ka.crudService.Delete(monthly.Monthly_ID)
			if model.ErrorCheck(k, err) {return}
		}
		err = ka.crudService.Delete(*factor.ActualID)
		if model.ErrorCheck(k, err) {return}
		err = ka.crudService.Delete(factor.Factor_ID)
		if model.ErrorCheck(k, err) {return}
	}
	err = ka.crudService.Delete(result.Result_ID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "delete Entire Result success"})
}
func (ka *kpiAPI) DeleteEntireFactor(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	factor, err := ka.crudService.GetByID(KpiID)
	if model.ErrorCheck(k, err) {return}
	//Delete Factors
	for _, monthly := range factor.Plan.Monthly{
		err = ka.crudService.Delete(monthly.Monthly_ID)
		if model.ErrorCheck(k, err) {return}
	}
	err = ka.crudService.Delete(*factor.PlanID)
	if model.ErrorCheck(k, err) {return}
	for _, monthly := range factor.Actual.Monthly{
		err = ka.crudService.Delete(monthly.Monthly_ID)
		if model.ErrorCheck(k, err) {return}
	}
	err = ka.crudService.Delete(*factor.ActualID)
	if model.ErrorCheck(k, err) {return}
	err = ka.crudService.Delete(factor.Factor_ID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "delete Entire Factor success"})
}
func (ka *kpiAPI) DeleteEntireAttendance(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	response, err := ka.crudService.GetByID(KpiID)
	if model.ErrorCheck(k, err) {return}

	//Delete Attendance
	err = ka.crudService.Delete(response.Year)
	if model.ErrorCheck(k, err) {return}
	//Delete Attendance Monthly
	if response.PlanID != nil {
		err = ka.crudService.Delete(*response.PlanID)
		if model.ErrorCheck(k, err) {return}
	}
	if response.ActualID != nil {
		err = ka.crudService.Delete(*response.ActualID)
		if model.ErrorCheck(k, err) {return}
	}
	if response.CutiID != nil {
		err = ka.crudService.Delete(*response.CutiID)
		if model.ErrorCheck(k, err) {return}
	}
	if response.IzinID != nil {
		err = ka.crudService.Delete(*response.IzinID)
		if model.ErrorCheck(k, err) {return}
	}
	if response.LainID != nil {
		err = ka.crudService.Delete(*response.LainID)
		if model.ErrorCheck(k, err) {return}
	}
	k.JSON(http.StatusOK, model.SuccessResponse{Message: "delete Entire Attendance success"})
}
// Get By ID
func (ka *kpiAPI) GetAttendanceByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	Attendance, err := ka.crudService.GetByID(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, Attendance.ToResponse())}
func (ka *kpiAPI) GetFactorByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	Factor, err := ka.crudService.GetByID(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, Factor.ToResponse())}
func (ka *kpiAPI) GetItemByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	Item, err := ka.crudService.GetByID(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, Item.ToResponse())}
func (ka *kpiAPI) GetMinipapByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	Minipap, err := ka.crudService.GetByID(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, Minipap)}
func (ka *kpiAPI) GetMonthlyByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	Monthly, err := ka.crudService.GetByID(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, Monthly)}
func (ka *kpiAPI) GetResultByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	Result, err := ka.crudService.GetByID(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, Result.ToResponse())}
func (ka *kpiAPI) GetYearlyByID(k *gin.Context) {
	KpiID, err := strconv.Atoi(k.Param("id"))
	if model.ErrorCheck(k, err){return}
	Yearly, err := ka.crudService.GetByID(KpiID)
	if model.ErrorCheck(k, err) {return}
	k.JSON(http.StatusOK, Yearly.ToResponse())}

// Get List
func (ka *kpiAPI) GetAttendanceList(k *gin.Context) {
	Attendance, err := ka.crudService.GetList()
	if model.ErrorCheck(k, err) {return}
	var result model.AttendanceArrayResponse
	result.Attendance = []model.AttendanceResponse{} 
	for _,att := range Attendance{
		result.Attendance = append(result.Attendance, att.ToResponse())	
	}
	result.Message = "Getting All Attendances Success"
	k.JSON(http.StatusOK, result)}
func (ka *kpiAPI) GetFactorList(k *gin.Context) {
	Factor, err := ka.crudService.GetList()
	if model.ErrorCheck(k, err) {return}
	var result model.FactorArrayResponse
	result.Factor = []model.FactorResponse{}
	for _, data := range Factor{
		result.Factor = append(result.Factor, data.ToResponse())
	} 
	result.Message = "Getting All Factors Success"
	k.JSON(http.StatusOK, result)}
func (ka *kpiAPI) GetItemList(k *gin.Context) {
	Item, err := ka.crudService.GetList()
	if model.ErrorCheck(k, err) {return}
	var result model.ItemArrayResponse
	result.Item = []model.ItemResponse{}
	for _, data := range Item{
		result.Item = append(result.Item, data.ToResponse())
	} 
	result.Message = "Getting All Items Success"
	k.JSON(http.StatusOK, result)}
func (ka *kpiAPI) GetMinipapList(k *gin.Context) {
	Minipap, err := ka.crudService.GetList()
	if model.ErrorCheck(k, err) {return}
	var result model.MinipapArrayResponse
	result.Minipap = Minipap 
	result.Message = "Getting All MiniPAPs Success"
	k.JSON(http.StatusOK, result)}
func (ka *kpiAPI) GetMonthlyList(k *gin.Context) {
	Monthly, err := ka.crudService.GetList()
	if model.ErrorCheck(k, err) {return}
	var result model.MonthlyArrayResponse
	result.Monthly = Monthly 
	result.Message = "Getting All Monthlys Success"
	k.JSON(http.StatusOK, result)}
func (ka *kpiAPI) GetResultList(k *gin.Context) {
	Result, err := ka.crudService.GetList()
	if model.ErrorCheck(k, err) {return}
	var result model.ResultArrayResponse
	result.Result = []model.ResultResponse{}
	for _, data := range Result{
		result.Result = append(result.Result, data.ToResponse())
	} 
	result.Message = "Getting All Results Success"
	k.JSON(http.StatusOK, result)}
func (ka *kpiAPI) GetYearlyList(k *gin.Context) {
	Yearly, err := ka.crudService.GetList()
	if model.ErrorCheck(k, err) {return}
	var result model.YearlyArrayResponse
	result.Yearly = []model.YearlyResponse{}
	for _, data := range Yearly{
		result.Yearly = append(result.Yearly, data.ToResponse())
	} 
	result.Message = "Getting All Yearlys Success"
	k.JSON(http.StatusOK, result)}