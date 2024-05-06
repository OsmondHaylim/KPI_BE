package service

import (
	"fmt"
	"goreact/model"
	"mime/multipart"
	"regexp"
	"strings"
	"time"

	// "path/filepath"
	"io"
	"os"
	"strconv"

	"errors"
	// "time"
	// "strings"
	"github.com/tealeg/xlsx"
)

var directory = "./uploads/"
func (ps *parseService) ParseKpi (input multipart.File) (*model.YearlyResponse, error){
	//Create temp file
	out, err := os.CreateTemp("", "upload-*.xlsx")
	if err != nil {return nil, err}	
	defer out.Close()
	//Copy File
	_, err = io.Copy(out, input)
	if err != nil {return nil, err}
	excel, err := xlsx.OpenFile(out.Name())
	if err != nil {return nil, err}
	test := ""
	for _, sheet := range excel.Sheets {
		//Find Sheet
		if sheet.Name == "KPI staff Design" { 
			//Get Year
			re := regexp.MustCompile(`(\d{4})$`)
			match := re.FindStringSubmatch(sheet.Rows[3].Cells[2].String())
			year := 0
			if len(match) == 2 {
				year, err = strconv.Atoi(match[1])
				if err != nil {return nil, err}
			}
			test += strconv.Itoa(year) + " "
			//Declare Response
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

			content := false	//Reading table status
			onHold := false		//Switch for 1 delay read before content
			prereq := false		//Length prerequisite status 
			attend := 0			//Countdown for Attendance Reading
			test += "/" + strconv.Itoa(len(sheet.Rows)) + "/"
			for _, row := range sheet.Rows {
				prereq = len(row.Cells) >= 22
				if (!content && !onHold && row.Cells[0].String() == "Item"){
					onHold = true
					continue
				}
				if (!content && onHold){
					content = true
					onHold = false
					continue
				}
				if (prereq && row.Cells[0].String() == "ABSENSI"){
					attend = 5
					tempResult.Factors = append(tempResult.Factors, tempFactor)
					tempFactor = model.FactorResponse{}
					tempItem.Results = append(tempItem.Results, tempResult)
					tempResult = model.ResultResponse{}
					KPI.Items = append(KPI.Items, tempItem)
					tempItem = model.ItemResponse{}
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
						if err != nil {return nil, err}
						monthly.Feb, err = row.Cells[10].Float()
						if err != nil {return nil, err}
						monthly.Mar, err = row.Cells[11].Float()
						if err != nil {return nil, err}
						monthly.Apr, err = row.Cells[12].Float()
						if err != nil {return nil, err}
						monthly.May, err = row.Cells[13].Float()
						if err != nil {return nil, err}
						monthly.Jun, err = row.Cells[14].Float()
						if err != nil {return nil, err}
						monthly.Jul, err = row.Cells[15].Float()
						if err != nil {return nil, err}
						monthly.Aug, err = row.Cells[16].Float()
						if err != nil {return nil, err}
						monthly.Sep, err = row.Cells[17].Float()
						if err != nil {return nil, err}
						monthly.Oct, err = row.Cells[18].Float()
						if err != nil {return nil, err}
						monthly.Nov, err = row.Cells[19].Float()
						if err != nil {return nil, err}
						monthly.Dec, err = row.Cells[20].Float()
						if err != nil {return nil, err}
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
					if err != nil {return nil, err}

					if row.Cells[10].Type() == xlsx.CellTypeNumeric {
						monthly.Feb, err = strconv.ParseFloat(row.Cells[10].Value, 64)
					}else if row.Cells[10].Value != ""{
						monthly.Feb, err = strconv.ParseFloat(row.Cells[10].Value[:len(row.Cells[10].Value)-1], 64)
					}
					if err != nil {return nil, err}

					if row.Cells[11].Type() == xlsx.CellTypeNumeric {
						monthly.Mar, err = strconv.ParseFloat(row.Cells[11].Value, 64)
					}else if row.Cells[11].Value != ""{
						monthly.Mar, err = strconv.ParseFloat(row.Cells[11].Value[:len(row.Cells[11].Value)-1], 64)
					}
					if err != nil {return nil, err}

					if row.Cells[12].Type() == xlsx.CellTypeNumeric {
						monthly.Apr, err = strconv.ParseFloat(row.Cells[12].Value, 64)
					}else if row.Cells[12].Value != ""{
						monthly.Apr, err = strconv.ParseFloat(row.Cells[12].Value[:len(row.Cells[12].Value)-1], 64)
					}
					if err != nil {return nil, err}

					if row.Cells[13].Type() == xlsx.CellTypeNumeric {
						monthly.May, err = strconv.ParseFloat(row.Cells[13].Value, 64)
					}else if row.Cells[13].Value != ""{
						monthly.May, err = strconv.ParseFloat(row.Cells[13].Value[:len(row.Cells[13].Value)-1], 64)
					}
					if err != nil {return nil, err}

					if row.Cells[14].Type() == xlsx.CellTypeNumeric {
						monthly.Jun, err = strconv.ParseFloat(row.Cells[14].Value, 64)
					}else if row.Cells[14].Value != ""{
						monthly.Jun, err = strconv.ParseFloat(row.Cells[14].Value[:len(row.Cells[14].Value)-1], 64)
					}
					if err != nil {return nil, err}

					if row.Cells[15].Type() == xlsx.CellTypeNumeric {
						monthly.Jul, err = strconv.ParseFloat(row.Cells[15].Value, 64)
					}else if row.Cells[15].Value != ""{
						monthly.Jul, err = strconv.ParseFloat(row.Cells[15].Value[:len(row.Cells[15].Value)-1], 64)
					}
					if err != nil {return nil, err}

					if row.Cells[16].Type() == xlsx.CellTypeNumeric {
						monthly.Aug, err = strconv.ParseFloat(row.Cells[16].Value, 64)
					}else if row.Cells[16].Value != ""{
						monthly.Aug, err = strconv.ParseFloat(row.Cells[16].Value[:len(row.Cells[16].Value)-1], 64)
					}
					if err != nil {return nil, err}

					if row.Cells[17].Type() == xlsx.CellTypeNumeric {
						monthly.Sep, err = strconv.ParseFloat(row.Cells[17].Value, 64)
					}else if row.Cells[17].Value != ""{
						monthly.Sep, err = strconv.ParseFloat(row.Cells[17].Value[:len(row.Cells[17].Value)-1], 64)
					}
					if err != nil {return nil, err}

					if row.Cells[18].Type() == xlsx.CellTypeNumeric {
						monthly.Oct, err = strconv.ParseFloat(row.Cells[18].Value, 64)
					}else if row.Cells[18].Value != ""{
						monthly.Oct, err = strconv.ParseFloat(row.Cells[18].Value[:len(row.Cells[18].Value)-1], 64)
					}
					if err != nil {return nil, err}

					if row.Cells[19].Type() == xlsx.CellTypeNumeric {
						monthly.Nov, err = strconv.ParseFloat(row.Cells[19].Value, 64)
					}else if row.Cells[19].Value != ""{
						monthly.Nov, err = strconv.ParseFloat(row.Cells[19].Value[:len(row.Cells[19].Value)-1], 64)
					}
					if err != nil {return nil, err}

					if row.Cells[20].Type() == xlsx.CellTypeNumeric {
						monthly.Dec, err = strconv.ParseFloat(row.Cells[20].Value, 64)
					}else if row.Cells[20].Value != ""{
						monthly.Dec, err = strconv.ParseFloat(row.Cells[20].Value[:len(row.Cells[20].Value)-1], 64)
					}
					if err != nil {return nil, err}
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
		return &KPI, nil
		}
	}
	return nil, nil
} 

func (ps *parseService) SaveFile (input multipart.File, header *multipart.FileHeader) (*model.UploadFile, error) {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		os.Mkdir(directory, os.ModePerm)
	}

	// Create a new file in the specified directory
	newFiles, err := os.Create(directory + header.Filename)
	if err != nil {return nil, err}
	defer newFiles.Close()

	// Copy the uploaded file to the newly created file
	written, err := io.Copy(newFiles, input)
	if err != nil {return nil, err}
	if written <= 0 {return nil, errors.New("nothing pasted")}
	
	content, err := io.ReadAll(newFiles)
	if err != nil {return nil, err}
	file := model.UploadFile{
		FileName: header.Filename,
		File: content,
	}
	return &file, nil
}

func (ps *parseService) ParseAnalisis (input multipart.File) (*model.Analisa, error){
	//Create temp file
	out, err := os.CreateTemp("", "upload-*.xlsx")
	if err != nil {return nil, err}	
	defer out.Close()
	//Copy File
	_, err = io.Copy(out, input)
	if err != nil {return nil, err}
	excel, err := xlsx.OpenFile(out.Name())
	if err != nil {return nil, err}
	test := ""
	for _, sheet := range excel.Sheets {
		if sheet.Name == "Analisa" {
			fmt.Println("found")
			re := regexp.MustCompile(`(\d{4})$`)
			match := re.FindStringSubmatch(sheet.Rows[1].Cells[0].String())
			year := 0
			if len(match) == 2 {
				year, err = strconv.Atoi(match[1])
				if err != nil {return nil, err}
			}
			test += strconv.Itoa(year) + " "
			Analisis := model.Analisa{
				Year: year,
			}
			tempMasalah := model.Masalah{}
			prereq := false
			content := false
			onHold := false
			fmt.Println("loop start")
			for _, row := range sheet.Rows{
				// fmt.Println(len(row.Cells))
				prereq = len(row.Cells) >= 12
				if (prereq && !content && row.Cells[0].String() == "No."){
					// fmt.Println("preparing to input")
					onHold = true
					continue
				} 
				if (!content && onHold){
					// fmt.Println("begin to input")
					content = true
					onHold = false
					continue
				}
				if (prereq && content && row.Cells[1].String() != "") {
					// fmt.Print("inputting")
					tempMasalah.Masalah = row.Cells[1].String()
					tempMasalah.Tindakan = row.Cells[7].String()
					tempMasalah.Pic = row.Cells[8].String()
					tempMasalah.Target = row.Cells[9].String()
					if row.Cells[10].String() != "" && row.Cells[10].String() != "-"{
						tempDate, err := time.Parse(row.Cells[10].String(), row.Cells[10].String())
						if err != nil{return nil, err}
						tempMasalah.FolDate = &tempDate
					}
					if row.Cells[11].String() != "" && row.Cells[11].String() != "-"{
						tempMasalah.Status = row.Cells[11].String()
					}
					for i := 2; i <= 6; i++{
						if row.Cells[i].String() != "" && row.Cells[i].String() != "-"{
							tempMasalah.Why = append(tempMasalah.Why, row.Cells[i].String()) 
						}
					}
					// fmt.Print(tempMasalah)
					Analisis.Masalah = append(Analisis.Masalah, tempMasalah)
					tempMasalah = model.Masalah{}
				}
				if (prereq && content && row.Cells[1].String() == "") {
					break
				}
			
			test += "/" + strconv.Itoa(len(sheet.Rows)) + "/"
			}
			return &Analisis, nil
		}
		
	}
	return nil, nil
}

func (ps *parseService) ParseSummary (input multipart.File) (*model.SummaryResponse, error){
	//Create temp file
	out, err := os.CreateTemp("", "upload-*.xlsx")
	if err != nil {return nil, err}	
	defer out.Close()
	//Copy File
	_, err = io.Copy(out, input)
	if err != nil {return nil, err}
	excel, err := xlsx.OpenFile(out.Name())
	if err != nil {return nil, err}
	for _, sheet := range excel.Sheets {
		if sheet.Name == "Summary Project" {
			fmt.Println("found")
			Summary := model.SummaryResponse{}
			
			
			for i := 3; i <= len(sheet.Rows[2].Cells); i+=2{
				if sheet.Rows[2].Cells[i].String() != "REMARKS" && !strings.Contains(sheet.Rows[2].Cells[i].String(), "vs"){
					tempProject := model.ProjectResponse{}
					item := make(map[string]int)
					qty := make(map[string]int)
					tempProject.Name = sheet.Rows[2].Cells[i].String()
					fmt.Println(tempProject.Name)
					item["Not Yet Start Issued FR"], err = sheet.Rows[5].Cells[i].Int()
					fmt.Println(item["Not Yet Start Issued FR"])
					if err != nil {return nil, err}
					qty["Not Yet Start Issued FR"], err = sheet.Rows[5].Cells[i+1].Int()
					if err != nil {return nil, err}
					item["DR"], err = sheet.Rows[6].Cells[i].Int()
					fmt.Println(item["DR"])
					if err != nil {return nil, err}
					qty["DR"], err = sheet.Rows[6].Cells[i+1].Int()
					if err != nil {return nil, err}
					item["PR to PO"], err = sheet.Rows[7].Cells[i].Int()
					fmt.Println(item["PR to PO"])
					if err != nil {return nil, err}
					qty["PR to PO"], err = sheet.Rows[7].Cells[i+1].Int()
					if err != nil {return nil, err}
					item["Install"], err = sheet.Rows[8].Cells[i].Int()
					fmt.Println(item["Install"])
					if err != nil {return nil, err}
					qty["Install"], err = sheet.Rows[8].Cells[i+1].Int()
					if err != nil {return nil, err}
					item["Finish"], err = sheet.Rows[9].Cells[i].Int()
					if err != nil {return nil, err}
					qty["Finish"], err = sheet.Rows[9].Cells[i+1].Int()
					if err != nil {return nil, err}
					item["Cancelled"], err = sheet.Rows[10].Cells[i].Int()
					if err != nil {return nil, err}
					qty["Cancelled"], err = sheet.Rows[10].Cells[i+1].Int()
					if err != nil {return nil, err}
					tempProject.Item = item
					tempProject.Quantity = qty
					Summary.Projects = append(Summary.Projects, tempProject)
				} else {break}
			}
			return &Summary, nil
		}
	}
	return nil, nil
}