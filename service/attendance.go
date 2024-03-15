package service

import (
	"goreact/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AttendanceService interface {
	Store(Attendance *model.Attendance) error
	Update(id int, attendance model.Attendance) error
	Delete(id int) error
	GetByID(id int) (*model.Attendance, error)
	GetList() ([]model.Attendance, error)
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

func (as *attendanceService) Delete(id int) error {	
	return as.db.Where(id).Delete(&model.Attendance{}).Error 
}

func (as *attendanceService) GetByID(id int) (*model.Attendance, error) {
	var Attendance model.Attendance
	err := as.db.
	Preload(clause.Associations).
	Where("id = ?", id).First(&Attendance).Error
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
