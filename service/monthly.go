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

func (ks *monthlyService) Store(monthly *model.Monthly) error {
	return ks.db.Create(monthly).Error
}

func (ks *monthlyService) Update(id int, monthly model.Monthly) error {
	return ks.db.Where(id).Updates(monthly).Error
}

func (ks *monthlyService) Delete(id int) error {	
	return ks.db.Where(id).Delete(&model.Monthly{}).Error 
}

func (ks *monthlyService) GetByID(id int) (*model.Monthly, error) {
	var Monthly model.Monthly
	err := ks.db.Where("id = ?", id).First(&Monthly).Error
	if err != nil {
		return nil, err
	}
	return &Monthly, nil
}

func (ks *monthlyService) GetList() ([]model.Monthly, error) {
	var result []model.Monthly
	err := ks.db.Find(&result).Error
	if err != nil{
		return []model.Monthly{}, err
	}
	return result, nil 
}
