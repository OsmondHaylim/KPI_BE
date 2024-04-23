package repository

import (
	"goreact/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AttendanceRepo interface {
	Store(Attendance *model.Attendance) error
	Update(id int, attendance model.Attendance) error
	Saves(attendance model.Attendance) error
	Delete(id int) error
	GetByID(id int) (*model.Attendance, error)
	GetList() ([]model.Attendance, error)
	GetAttendanceFromMonthly(id int) (*model.Attendance, *string, error)
}

type attendanceRepo struct {
	db *gorm.DB
}

func NewAttendanceRepo(db *gorm.DB) *attendanceRepo {
	return &attendanceRepo{db}
}

func (as *attendanceRepo) Store(attendance *model.Attendance) error {
	return as.db.Create(attendance).Error
}

func (as *attendanceRepo) Update(id int, attendance model.Attendance) error {
	return as.db.Where(id).Updates(attendance).Error
}

func (as *attendanceRepo) Saves(attendance model.Attendance) error {
	return as.db.Save(attendance).Error
}

func (as *attendanceRepo) Delete(id int) error {	
	return as.db.Where(id).Select(clause.Associations).Delete(&model.Attendance{}).Error
}

func (as *attendanceRepo) GetByID(id int) (*model.Attendance, error) {
	var Attendance model.Attendance
	err := as.db.
	Preload(clause.Associations).
	Where("year = ?", id).First(&Attendance).Error
	if err != nil {return nil, err}
	return &Attendance, nil
}

func (as *attendanceRepo) GetList() ([]model.Attendance, error) {
	var result []model.Attendance
	err := as.db.
	Preload(clause.Associations).
	Find(&result).Error
	if err != nil{
		return []model.Attendance{}, err
	}
	return result, nil 
}

func (as *attendanceRepo) GetAttendanceFromMonthly(id_monthly int)(*model.Attendance, *string, error){
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
