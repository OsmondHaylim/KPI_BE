package service

import (
	"goreact/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AttendanceService interface {
	Store(Attendance *model.Attendance) error
	Update(id int, attendance model.Attendance) error
	Saves(attendance model.Attendance) error
	Delete(id int) error
	GetByID(id int) (*model.Attendance, error)
	GetList() ([]model.Attendance, error)
	GetAttendanceFromMonthly(id int) (*model.Attendance, *string, error)
}

type attendanceService struct {
	db *gorm.DB
}

func NewAttendanceService(db *gorm.DB) *attendanceService {
	return &attendanceService{db}
}

func (as *attendanceService) Store(attendance *model.Attendance) error {
	return as.db.Create(attendance).Error
}

func (as *attendanceService) Update(id int, attendance model.Attendance) error {
	return as.db.Where(id).Updates(attendance).Error
}

func (as *attendanceService) Saves(attendance model.Attendance) error {
	return as.db.Save(attendance).Error
}

func (as *attendanceService) Delete(id int) error {	
	return as.db.Where(id).Delete(&model.Attendance{}).Error 
}

func (as *attendanceService) GetByID(id int) (*model.Attendance, error) {
	var Attendance model.Attendance
	err := as.db.
	Preload(clause.Associations).
	Where("Year = ?", id).First(&Attendance).Error
	if err != nil {
		return nil, err
	}
	return &Attendance, nil
}

func (as *attendanceService) GetList() ([]model.Attendance, error) {
	var result []model.Attendance
	err := as.db.
	Preload(clause.Associations).
	Find(&result).Error
	if err != nil{
		return []model.Attendance{}, err
	}
	return result, nil 
}

func (as *attendanceService) GetAttendanceFromMonthly(id_monthly int)(*model.Attendance, *string, error){
	var Attendance model.Attendance
	err := as.db.
	Preload(clause.Associations).
	Where(
		`plan_id = ? OR 
		actual_id = ? OR 
		cuti_id = ? OR 
		izin_id = ? OR 
		lain_id = ?`, 
		id_monthly, 
		id_monthly, 
		id_monthly, 
		id_monthly, 
		id_monthly).
		First(&Attendance).Error
	if err != nil {
		return nil, nil, err
	}else{
		var where string
		if *Attendance.PlanID == id_monthly{
			where = "plan_id"
		}else if *Attendance.ActualID == id_monthly{
			where = "actual_id"
		}else if *Attendance.CutiID == id_monthly{
			where = "cuti_id"
		}else if *Attendance.IzinID == id_monthly{
			where = "izin_id"
		}else{
			where = "lain_id"
		}
		return &Attendance, &where, nil
	}
}
