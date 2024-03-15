package service

import (
	"goreact/model"
	"gorm.io/gorm"
)

type MonthlyService interface {
	Store(Monthly *model.Monthly) error
	Update(id int, monthly model.Monthly) error
	Delete(id int) error
	GetByID(id int) (*model.Monthly, error)
	GetList() ([]model.Monthly, error)
}

type monthlyService struct {
	db *gorm.DB
}

func NewMonthlyService(db *gorm.DB) *monthlyService {
	return &monthlyService{db}
}

func (ms *monthlyService) Store(monthly *model.Monthly) error {
	return ms.db.Create(monthly).Error
}

func (ms *monthlyService) Update(id int, monthly model.Monthly) error {
	return ms.db.Where(id).Updates(monthly).Error
}

func (ms *monthlyService) Delete(id int) error {	
	return ms.db.Where(id).Delete(&model.Monthly{}).Error 
}

func (ms *monthlyService) GetByID(id int) (*model.Monthly, error) {
	var Monthly model.Monthly
	err := ms.db.Where("id = ?", id).First(&Monthly).Error
	if err != nil {
		return nil, err
	}
	return &Monthly, nil
}

func (ms *monthlyService) GetList() ([]model.Monthly, error) {
	var result []model.Monthly
	err := ms.db.Find(&result).Error
	if err != nil{
		return []model.Monthly{}, err
	}
	return result, nil 
}
