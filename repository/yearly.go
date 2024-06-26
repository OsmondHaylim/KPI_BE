package repository

import (
	"goreact/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type YearlyRepo interface {
	Store(Yearly *model.Yearly) error
	Update(id int, yearly model.Yearly) error
	Saves(yearly model.Yearly) error
	Delete(id int) error
	GetByID(id int) (*model.Yearly, error)
	GetList() ([]model.Yearly, error)
}

type yearlyRepo struct {
	db *gorm.DB
}

func NewYearlyRepo(db *gorm.DB) *yearlyRepo {
	return &yearlyRepo{db}
}

func (ys *yearlyRepo) Store(yearly *model.Yearly) error {
	return ys.db.Create(yearly).Error
}

func (ys *yearlyRepo) Update(id int, yearly model.Yearly) error {
	return ys.db.Where(id).Updates(yearly).Error
}

func (ys *yearlyRepo) Saves(yearly model.Yearly) error {
	return ys.db.Save(yearly).Error
}

func (ys *yearlyRepo) Delete(id int) error {	
	return ys.db.
	Select(clause.Associations).
	Select("Attendance.Plan").
	Select("Attendance.Actual").
	Select("Attendance.Cuti").
	Select("Attendance.Izin").
	Select("Attendance.Lain").
	Select("Items.Results").
	Select("Items.Results.Factors").
	Select("Items.Results.Factors.Plan").
	Select("Items.Results.Factors.Actual").
	Select("Items.Results.Factors.Plan.Monthly").
	Select("Items.Results.Factors.Actual.Monthly").
	Delete(&model.Yearly{Year: id}).Error 
}

func (ys *yearlyRepo) GetByID(id int) (*model.Yearly, error) {
	var Yearly model.Yearly
	err := ys.db.
	Preload(clause.Associations).
	Preload("Attendance.Plan").
	Preload("Attendance.Actual").
	Preload("Attendance.Cuti").
	Preload("Attendance.Izin").
	Preload("Attendance.Lain").
	Preload("Items.Results").
	Preload("Items.Results.Factors").
	Preload("Items.Results.Factors.Plan").
	Preload("Items.Results.Factors.Actual").
	Preload("Items.Results.Factors.Plan.Monthly").
	Preload("Items.Results.Factors.Actual.Monthly").
	Where("year = ?", id).
	First(&Yearly).Error
	if err != nil {
		return nil, err
	}
	return &Yearly, nil
}

func (ys *yearlyRepo) GetList() ([]model.Yearly, error) {
	var yearly []model.Yearly
	err := ys.db.
	Preload(clause.Associations).
	Preload("Attendance.Plan").
	Preload("Attendance.Actual").
	Preload("Attendance.Cuti").
	Preload("Attendance.Izin").
	Preload("Attendance.Lain").
	Preload("Items.Results").
	Preload("Items.Results.Factors").
	Preload("Items.Results.Factors.Plan").
	Preload("Items.Results.Factors.Actual").
	Preload("Items.Results.Factors.Plan.Monthly").
	Preload("Items.Results.Factors.Actual.Monthly").
	Find(&yearly).Error
	if err != nil{
		return []model.Yearly{}, err
	}
	return yearly, nil 
}
