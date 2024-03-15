package service

import (
	"goreact/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type YearlyService interface {
	Store(Yearly *model.Yearly) error
	Update(id int, yearly model.Yearly) error
	Delete(id int) error
	GetByID(id int) (*model.Yearly, error)
	GetList() ([]model.Yearly, error)
}

type yearlyService struct {
	db *gorm.DB
}

func NewYearlyService(db *gorm.DB) *yearlyService {
	return &yearlyService{db}
}

func (ys *yearlyService) Store(yearly *model.Yearly) error {
	return ys.db.Create(yearly).Error
}

func (ys *yearlyService) Update(id int, yearly model.Yearly) error {
	return ys.db.Where(id).Updates(yearly).Error
}

func (ys *yearlyService) Delete(id int) error {	
	return ys.db.Where(id).Delete(&model.Yearly{}).Error 
}

func (ys *yearlyService) GetByID(id int) (*model.Yearly, error) {
	var Yearly model.Yearly
	err := ys.db.
	Preload(clause.Associations).
	Preload("Item.Result").
	Preload("Item.Result.Factor").
	Preload("Item.Result.Factor.PAP").
	Preload("Item.Result.Factor.PAP.MiniPAP").
	Preload("Item.Result.Factor.PAP.MiniPAP.Monthly").
	Where("year = ?", id).
	First(&Yearly).Error
	if err != nil {
		return nil, err
	}
	return &Yearly, nil
}

func (ys *yearlyService) GetList() ([]model.Yearly, error) {
	var yearly []model.Yearly
	err := ys.db.
	Preload(clause.Associations).
	Preload("Item.Result").
	Preload("Item.Result.Factor").
	Preload("Item.Result.Factor.PAP").
	Preload("Item.Result.Factor.PAP.MiniPAP").
	Preload("Item.Result.Factor.PAP.MiniPAP.Monthly").
	Find(&yearly).Error
	if err != nil{
		return []model.Yearly{}, err
	}
	return yearly, nil 
}
