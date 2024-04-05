package api

import (
	"goreact/model"
	"goreact/service"
	"net/http"
	"regexp"

	// "path/filepath"
	"io"
	"os"
	"strconv"

	// "time"
	// "strings"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

type FileAPI interface {
	FileUpload(f *gin.Context)
}

type fileAPI struct {
	fileService service.FileService
}

func NewFileAPI(fileService service.FileService) *fileAPI {
	return &fileAPI{
		fileService,
	}
}

func (fa *fileAPI) FileUpload(f *gin.Context) {
	//get File
	file, header, err := f.Request.FormFile("file")
	if model.ErrorCheck(f, err){return}
	defer file.Close()
	//Check format
	if header.Header.Get("Content-Type") != "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
		f.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "not xlsx"})
		return
	}
	//Create temp file
	out, err := os.CreateTemp("", "upload-*.xlsx")
	if model.ErrorCheck(f, err){return}
	defer out.Close()
	//Copy File
	_, err = io.Copy(out, file)
	if model.ErrorCheck(f, err){return}

	//Create Directory
	directory := "./uploads/"
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		os.Mkdir(directory, os.ModePerm)
	}

	// // Create a new file in the specified directory
	// out, err = os.Create(directory + header.Filename)
	// if err != nil {
	// 	f.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
	// 	return
	// }
	// defer out.Close()

	// // Copy the uploaded file to the newly created file
	// _, err = io.Copy(out, file)
	// if err != nil {
	// 	f.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
	// 	return
	// }

	//Open Excel
	excel, err := xlsx.OpenFile(out.Name())
	if model.ErrorCheck(f, err){return}
	test := ""
	for _, sheet := range excel.Sheets {
		if sheet.Name == "KPI staff Design" {
			re := regexp.MustCompile(`(\d{4})$`)
			match := re.FindStringSubmatch(sheet.Rows[3].Cells[2].String())
			year := 0
			if len(match) == 2 {
				year, err = strconv.Atoi(match[1])
				if model.ErrorCheck(f, err){return}
			}
			test += strconv.Itoa(year) + " "
			KPI := model.YearlyResponse{
				Year: year,
				Attendance: &model.AttendanceResponse{},
			}
			tempItem := model.ItemResponse{}
			tempResult := model.ResultResponse{}
			tempFactor := model.FactorResponse{}
			tempAttend := model.AttendanceResponse{
				Year: year,
			}
			// monthly := model.Monthly{}

			content := false
			onHold := false
			prereq := false
			attend := 0
			test += "/" + strconv.Itoa(len(sheet.Rows)) + "/"
			for _, row := range sheet.Rows {
				prereq = len(row.Cells) >= 22
				if (prereq && row.Cells[0].String() == "ABSENSI"){
					attend = 5
					tempResult.Factors = append(tempResult.Factors, tempFactor)
					tempFactor = model.FactorResponse{}
					tempItem.Results = append(tempItem.Results, tempResult)
					tempResult = model.ResultResponse{}
					KPI.Items = append(KPI.Items, tempItem)
					tempItem = model.ItemResponse{}
				}
				if (!content && !onHold && row.Cells[0].String() == "Item"){
					onHold = true
					continue
				}
				if (!content && onHold){
					content = true
					onHold = false
					continue
				}
				if (prereq && content && row.Cells[0].String() != "ABSENSI") {
					if row.Cells[3].String() != ""{
						// Input Factor to yearly, reset tempFactor
						if tempFactor.Title != ""{
							tempResult.Factors = append(tempResult.Factors, tempFactor)
							tempFactor = model.FactorResponse{}
						}
						//Inputting Factor
						tempFactor.Title = row.Cells[3].String()
						tempFactor.Unit = row.Cells[7].String()
						tempFactor.Target = row.Cells[8].String()
					}
					if row.Cells[1].String() != ""{
						// Input Result to item, reset tempResult
						if tempResult.Name != ""{
							tempItem.Results = append(tempItem.Results, tempResult)
							tempResult = model.ResultResponse{}
						}
						//Inputting Result
						tempResult.Name = row.Cells[1].String() + " " + row.Cells[2].String()
					}
					if row.Cells[0].String() != ""{
						// Input item to yearly, reset tempItem
						if tempItem.Name != ""{
							KPI.Items = append(KPI.Items, tempItem)
							tempItem = model.ItemResponse{}
						}
						//Inputting item
						tempItem.Name = row.Cells[0].String()
					}
					remarks := row.Cells[22].String() + row.Cells[23].String()
					monthly := model.Monthly{}
					//Extracting Monthly
					if len(row.Cells[6].String()) != 0 && row.Cells[6].String()[0] == '%'{
						continue
					}else if len(row.Cells[6].String()) != 0{
						monthly.Jan, err = row.Cells[9].Float()
						if model.ErrorCheck(f, err){return}
						monthly.Feb, err = row.Cells[10].Float()
						if model.ErrorCheck(f, err){return}
						monthly.Mar, err = row.Cells[11].Float()
						if model.ErrorCheck(f, err){return}
						monthly.Apr, err = row.Cells[12].Float()
						if model.ErrorCheck(f, err){return}
						monthly.May, err = row.Cells[13].Float()
						if model.ErrorCheck(f, err){return}
						monthly.Jun, err = row.Cells[14].Float()
						if model.ErrorCheck(f, err){return}
						monthly.Jul, err = row.Cells[15].Float()
						if model.ErrorCheck(f, err){return}
						monthly.Aug, err = row.Cells[16].Float()
						if model.ErrorCheck(f, err){return}
						monthly.Sep, err = row.Cells[17].Float()
						if model.ErrorCheck(f, err){return}
						monthly.Oct, err = row.Cells[18].Float()
						if model.ErrorCheck(f, err){return}
						monthly.Nov, err = row.Cells[19].Float()
						if model.ErrorCheck(f, err){return}
						monthly.Dec, err = row.Cells[20].Float()
						if model.ErrorCheck(f, err){return}
						monthly.Remarks = &remarks
						//Inputting MiniPAP & Monthly
						switch row.Cells[6].String()[0]{
						case 'P':
							if tempFactor.Plan == nil{
								tempFactor.Plan = &model.MiniPAP{}
							}
							tempFactor.Plan.Monthly = append(tempFactor.Plan.Monthly, monthly)
						case 'A':
							if tempFactor.Actual == nil{
								tempFactor.Actual = &model.MiniPAP{}
							}
							tempFactor.Actual.Monthly = append(tempFactor.Actual.Monthly, monthly)
						default:
							continue	
						}
					}
					monthly = model.Monthly{}
				}	
				if (prereq && attend > 0){
					monthly := model.Monthly{}
					remarks := row.Cells[22].String() + row.Cells[23].String()
					if row.Cells[9].Type() == xlsx.CellTypeNumeric {
						monthly.Jan, err = strconv.ParseFloat(row.Cells[9].Value, 64)
					}else if row.Cells[9].Value != ""{
						monthly.Jan, err = strconv.ParseFloat(row.Cells[9].Value[:len(row.Cells[9].Value)-1], 64)
					}
					if model.ErrorCheck(f, err){return}

					if row.Cells[10].Type() == xlsx.CellTypeNumeric {
						monthly.Feb, err = strconv.ParseFloat(row.Cells[10].Value, 64)
					}else if row.Cells[10].Value != ""{
						monthly.Feb, err = strconv.ParseFloat(row.Cells[10].Value[:len(row.Cells[10].Value)-1], 64)
					}
					if model.ErrorCheck(f, err){return}

					if row.Cells[11].Type() == xlsx.CellTypeNumeric {
						monthly.Mar, err = strconv.ParseFloat(row.Cells[11].Value, 64)
					}else if row.Cells[11].Value != ""{
						monthly.Mar, err = strconv.ParseFloat(row.Cells[11].Value[:len(row.Cells[11].Value)-1], 64)
					}
					if model.ErrorCheck(f, err){return}

					if row.Cells[12].Type() == xlsx.CellTypeNumeric {
						monthly.Apr, err = strconv.ParseFloat(row.Cells[12].Value, 64)
					}else if row.Cells[12].Value != ""{
						monthly.Apr, err = strconv.ParseFloat(row.Cells[12].Value[:len(row.Cells[12].Value)-1], 64)
					}
					if model.ErrorCheck(f, err){return}

					if row.Cells[13].Type() == xlsx.CellTypeNumeric {
						monthly.May, err = strconv.ParseFloat(row.Cells[13].Value, 64)
					}else if row.Cells[13].Value != ""{
						monthly.May, err = strconv.ParseFloat(row.Cells[13].Value[:len(row.Cells[13].Value)-1], 64)
					}
					if model.ErrorCheck(f, err){return}

					if row.Cells[14].Type() == xlsx.CellTypeNumeric {
						monthly.Jan, err = strconv.ParseFloat(row.Cells[14].Value, 64)
					}else if row.Cells[14].Value != ""{
						monthly.Jun, err = strconv.ParseFloat(row.Cells[14].Value[:len(row.Cells[14].Value)-1], 64)
					}
					if model.ErrorCheck(f, err){return}

					if row.Cells[15].Type() == xlsx.CellTypeNumeric {
						monthly.Jul, err = strconv.ParseFloat(row.Cells[15].Value, 64)
					}else if row.Cells[15].Value != ""{
						monthly.Jul, err = strconv.ParseFloat(row.Cells[15].Value[:len(row.Cells[15].Value)-1], 64)
					}
					if model.ErrorCheck(f, err){return}

					if row.Cells[16].Type() == xlsx.CellTypeNumeric {
						monthly.Aug, err = strconv.ParseFloat(row.Cells[16].Value, 64)
					}else if row.Cells[16].Value != ""{
						monthly.Aug, err = strconv.ParseFloat(row.Cells[16].Value[:len(row.Cells[16].Value)-1], 64)
					}
					if model.ErrorCheck(f, err){return}

					if row.Cells[17].Type() == xlsx.CellTypeNumeric {
						monthly.Jan, err = strconv.ParseFloat(row.Cells[17].Value, 64)
					}else if row.Cells[17].Value != ""{
						monthly.Sep, err = strconv.ParseFloat(row.Cells[17].Value[:len(row.Cells[17].Value)-1], 64)
					}
					if model.ErrorCheck(f, err){return}

					if row.Cells[18].Type() == xlsx.CellTypeNumeric {
						monthly.Oct, err = strconv.ParseFloat(row.Cells[18].Value, 64)
					}else if row.Cells[18].Value != ""{
						monthly.Oct, err = strconv.ParseFloat(row.Cells[18].Value[:len(row.Cells[18].Value)-1], 64)
					}
					if model.ErrorCheck(f, err){return}

					if row.Cells[19].Type() == xlsx.CellTypeNumeric {
						monthly.Nov, err = strconv.ParseFloat(row.Cells[19].Value, 64)
					}else if row.Cells[19].Value != ""{
						monthly.Nov, err = strconv.ParseFloat(row.Cells[19].Value[:len(row.Cells[19].Value)-1], 64)
					}
					if model.ErrorCheck(f, err){return}

					if row.Cells[20].Type() == xlsx.CellTypeNumeric {
						monthly.Dec, err = strconv.ParseFloat(row.Cells[20].Value, 64)
					}else if row.Cells[20].Value != ""{
						monthly.Dec, err = strconv.ParseFloat(row.Cells[20].Value[:len(row.Cells[20].Value)-1], 64)
					}
					if model.ErrorCheck(f, err){return}
					monthly.Remarks = &remarks
					// fmt.Print(monthly)
					switch attend{
					case 5:
						tempAttend.Plan = &monthly
					case 4:
						tempAttend.Actual = &monthly
					case 3:
						tempAttend.Cuti = &monthly
					case 2: 
						tempAttend.Izin = &monthly
					case 1:
						tempAttend.Lain = &monthly
						KPI.Attendance = &tempAttend
					}
					attend--
				}
			}
		
		f.JSON(http.StatusOK, KPI)	
		return
		}
	}
}
