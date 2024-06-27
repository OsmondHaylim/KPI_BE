package service

import (
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
						if err != nil {monthly.Jan = 0}
						monthly.Feb, err = row.Cells[10].Float()
						if err != nil {monthly.Feb = 0}
						monthly.Mar, err = row.Cells[11].Float()
						if err != nil {monthly.Mar = 0}
						monthly.Apr, err = row.Cells[12].Float()
						if err != nil {monthly.Apr = 0}
						monthly.May, err = row.Cells[13].Float()
						if err != nil {monthly.May = 0}
						monthly.Jun, err = row.Cells[14].Float()
						if err != nil {monthly.Jun = 0}
						monthly.Jul, err = row.Cells[15].Float()
						if err != nil {monthly.Jul = 0}
						monthly.Aug, err = row.Cells[16].Float()
						if err != nil {monthly.Aug = 0}
						monthly.Sep, err = row.Cells[17].Float()
						if err != nil {monthly.Sep = 0}
						monthly.Oct, err = row.Cells[18].Float()
						if err != nil {monthly.Oct = 0}
						monthly.Nov, err = row.Cells[19].Float()
						if err != nil {monthly.Nov = 0}
						monthly.Dec, err = row.Cells[20].Float()
						if err != nil {monthly.Dec = 0}
						switch tempFactor.Unit{
						case "MRp.", "mRp.", "mrp.", "MRP.","MRp", "mRp", "mrp", "MRP":
							monthly.Jan *= 1000000
							monthly.Feb *= 1000000
							monthly.Mar *= 1000000
							monthly.Apr *= 1000000
							monthly.May *= 1000000
							monthly.Jun *= 1000000
							monthly.Jul *= 1000000
							monthly.Aug *= 1000000
							monthly.Sep *= 1000000
							monthly.Oct *= 1000000
							monthly.Nov *= 1000000
							monthly.Dec *= 1000000	
							tempFactor.Unit = "Rp."
						default:
							continue
						}
						monthly.Remarks = &remarks
						//Inputting MiniPAP & Monthly
						if strings.Contains(row.Cells[6].String(), "vs") || strings.Contains(row.Cells[6].String(), "VS"){continue}
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
					}else if row.Cells[9].Value != "" && row.Cells[9].String()[len(row.Cells[9].String()) - 1] == '%' {
						monthly.Jan, err = strconv.ParseFloat(row.Cells[9].Value[:len(row.Cells[9].Value)-1], 64)
					}
					if err != nil {monthly.Jan = 0; err = nil}

					if row.Cells[10].Type() == xlsx.CellTypeNumeric {
						monthly.Feb, err = strconv.ParseFloat(row.Cells[10].Value, 64)
					}else if row.Cells[10].Value != "" && row.Cells[10].String()[len(row.Cells[10].String()) - 1] == '%' {
						monthly.Feb, err = strconv.ParseFloat(row.Cells[10].Value[:len(row.Cells[10].Value)-1], 64)
					}
					if err != nil {monthly.Feb = 0; err = nil}

					if row.Cells[11].Type() == xlsx.CellTypeNumeric {
						monthly.Mar, err = strconv.ParseFloat(row.Cells[11].Value, 64)
					}else if row.Cells[11].Value != "" && row.Cells[11].String()[len(row.Cells[11].String()) - 1] == '%' {
						monthly.Mar, err = strconv.ParseFloat(row.Cells[11].Value[:len(row.Cells[11].Value)-1], 64)
					}else{
						monthly.Mar = 0
					}
					if err != nil {monthly.Mar = 0; err = nil}

					if row.Cells[12].Type() == xlsx.CellTypeNumeric {
						monthly.Apr, err = strconv.ParseFloat(row.Cells[12].Value, 64)
					}else if row.Cells[12].Value != "" && row.Cells[11].String()[len(row.Cells[11].String()) - 1] == '%' {
						monthly.Apr, err = strconv.ParseFloat(row.Cells[12].Value[:len(row.Cells[12].Value)-1], 64)
					}
					if err != nil {monthly.Apr = 0; err = nil}

					if row.Cells[13].Type() == xlsx.CellTypeNumeric {
						monthly.May, err = strconv.ParseFloat(row.Cells[13].Value, 64)
					}else if row.Cells[13].Value != "" && row.Cells[11].String()[len(row.Cells[11].String()) - 1] == '%' {
						monthly.May, err = strconv.ParseFloat(row.Cells[13].Value[:len(row.Cells[13].Value)-1], 64)
					}
					if err != nil {monthly.May = 0; err = nil}

					if row.Cells[14].Type() == xlsx.CellTypeNumeric {
						monthly.Jun, err = strconv.ParseFloat(row.Cells[14].Value, 64)
					}else if row.Cells[14].Value != "" && row.Cells[11].String()[len(row.Cells[11].String()) - 1] == '%' {
						monthly.Jun, err = strconv.ParseFloat(row.Cells[14].Value[:len(row.Cells[14].Value)-1], 64)
					}
					if err != nil {monthly.Jun = 0; err = nil}

					if row.Cells[15].Type() == xlsx.CellTypeNumeric {
						monthly.Jul, err = strconv.ParseFloat(row.Cells[15].Value, 64)
					}else if row.Cells[15].Value != "" && row.Cells[11].String()[len(row.Cells[11].String()) - 1] == '%' {
						monthly.Jul, err = strconv.ParseFloat(row.Cells[15].Value[:len(row.Cells[15].Value)-1], 64)
					}
					if err != nil {monthly.Jul = 0; err = nil}

					if row.Cells[16].Type() == xlsx.CellTypeNumeric {
						monthly.Aug, err = strconv.ParseFloat(row.Cells[16].Value, 64)
					}else if row.Cells[16].Value != "" && row.Cells[11].String()[len(row.Cells[11].String()) - 1] == '%' {
						monthly.Aug, err = strconv.ParseFloat(row.Cells[16].Value[:len(row.Cells[16].Value)-1], 64)
					}
					if err != nil {monthly.Aug = 0; err = nil}

					if row.Cells[17].Type() == xlsx.CellTypeNumeric {
						monthly.Sep, err = strconv.ParseFloat(row.Cells[17].Value, 64)
					}else if row.Cells[17].Value != "" && row.Cells[11].String()[len(row.Cells[11].String()) - 1] == '%' {
						monthly.Sep, err = strconv.ParseFloat(row.Cells[17].Value[:len(row.Cells[17].Value)-1], 64)
					}
					if err != nil {monthly.Sep = 0; err = nil}

					if row.Cells[18].Type() == xlsx.CellTypeNumeric {
						monthly.Oct, err = strconv.ParseFloat(row.Cells[18].Value, 64)
					}else if row.Cells[18].Value != "" && row.Cells[11].String()[len(row.Cells[11].String()) - 1] == '%' {
						monthly.Oct, err = strconv.ParseFloat(row.Cells[18].Value[:len(row.Cells[18].Value)-1], 64)
					}
					if err != nil {monthly.Oct = 0; err = nil}

					if row.Cells[19].Type() == xlsx.CellTypeNumeric {
						monthly.Nov, err = strconv.ParseFloat(row.Cells[19].Value, 64)
					}else if row.Cells[19].Value != "" && row.Cells[11].String()[len(row.Cells[11].String()) - 1] == '%' {
						monthly.Nov, err = strconv.ParseFloat(row.Cells[19].Value[:len(row.Cells[19].Value)-1], 64)
					}
					if err != nil {monthly.Nov = 0; err = nil}

					if row.Cells[20].Type() == xlsx.CellTypeNumeric {
						monthly.Dec, err = strconv.ParseFloat(row.Cells[20].Value, 64)
					}else if row.Cells[20].Value != "" && row.Cells[11].String()[len(row.Cells[11].String()) - 1] == '%' {
						monthly.Dec, err = strconv.ParseFloat(row.Cells[20].Value[:len(row.Cells[20].Value)-1], 64)
					}
					if err != nil {monthly.Dec = 0; err = nil}
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
						tempMasalah.FolDate = row.Cells[10].String()
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

func (ps *parseService) ParseSummary (input multipart.File) (*model.Summary, error){
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
			Summary := model.Summary{}
			date := sheet.Rows[1].Cells[1].String()
			parts := strings.Split(date, ": ")
			if len(parts) < 2 {return nil, errors.New("invalid format")}
			datePart := parts[1]
			layout := "02 January 2006"
			issuedDate, err := time.Parse(layout, datePart)
			if err != nil {return nil, err}
			Summary.IssuedDate = &issuedDate
			
			for i := 4; i <= len(sheet.Rows); i++{
				if sheet.Rows[i].Cells[2].String() != "" && sheet.Rows[i].Cells[2].String() != "PERSENTAGE FINISH" {
					Summary.Status = append(Summary.Status, sheet.Rows[i].Cells[2].String()) 
				}else {break}
			}
			for i := 3; i <= len(sheet.Rows[2].Cells); i+=2{
				if sheet.Rows[2].Cells[i].String() != "REMARKS" && !strings.Contains(sheet.Rows[2].Cells[i].String(), "vs"){
					tempProject := model.Project{}
					item := []int32{}
					qty := []int32{}
					tempProject.Name = sheet.Rows[2].Cells[i].String()
					for j := 0; j < len(Summary.Status); j++{
						data, err := sheet.Rows[j+4].Cells[i].Int()
						if err != nil{return nil, err}
						item = append(item, int32(data))
						data, err = sheet.Rows[j+4].Cells[i+1].Int()
						if err != nil{return nil, err}
						qty = append(qty, int32(data))
					}
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