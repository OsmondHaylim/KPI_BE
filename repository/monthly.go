package repository

import (
	"goreact/model"
	"gorm.io/gorm"
)

type MonthlyRepo interface {
	Store(Monthly *model.Monthly) error
	Update(id int, monthly model.Monthly) error
	Saves(monthly model.Monthly) error
	Delete(id int) error
	GetByID(id int) (*model.Monthly, error)
	GetList() ([]model.Monthly, error)
}

type monthlyRepo struct {
	db *gorm.DB
}

func NewMonthlyRepo(db *gorm.DB) *monthlyRepo {
	return &monthlyRepo{db}
}

func (ms *monthlyRepo) Store(monthly *model.Monthly) error {
	return ms.db.Create(monthly).Error
}

func (ms *monthlyRepo) Update(id int, monthly model.Monthly) error {
	return ms.db.Where(id).Updates(monthly).Error
}

func (ms *monthlyRepo) Saves(monthly model.Monthly) error{
	return ms.db.Save(monthly).Error
}

func (ms *monthlyRepo) Delete(id int) error {	
	return ms.db.Where(id).Delete(&model.Monthly{}).Error 
}

func (ms *monthlyRepo) GetByID(id int) (*model.Monthly, error) {
	var Monthly model.Monthly
	err := ms.db.Where("monthly_id = ?", id).First(&Monthly).Error
	if err != nil {
		return nil, err
	}
	return &Monthly, nil
}

func (ms *monthlyRepo) GetList() ([]model.Monthly, error) {
	var result []model.Monthly
	err := ms.db.Find(&result).Error
	if err != nil{
		return []model.Monthly{}, err
	}
	return result, nil 
}
