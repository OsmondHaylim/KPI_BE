package service

import (
	"goreact/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FactorService interface {
	Store(Factor *model.Factor) error
	Update(id int, factor model.Factor) error
	Saves(factor model.Factor) error
	Delete(id int) error
	GetByID(id int) (*model.Factor, error)
	GetList() ([]model.Factor, error)
}

type factorService struct {
	db *gorm.DB
}

func NewFactorService(db *gorm.DB) *factorService {
	return &factorService{db}
}

func (fs *factorService) Store(factor *model.Factor) error {
	return fs.db.Create(factor).Error
}

func (fs *factorService) Update(id int, factor model.Factor) error {
	return fs.db.Where(id).Updates(factor).Error
}

func (fs *factorService) Saves(factor model.Factor) error {
	return fs.db.Save(factor).Error
}

func (fs *factorService) Delete(id int) error {	
	return fs.db.Where(id).Delete(&model.Factor{}).Error 
}

func (fs *factorService) GetByID(id int) (*model.Factor, error) {
	var Factor model.Factor
	err := fs.db.
	Preload(clause.Associations).
	Preload("Plan.Monthly").
	Preload("Actual.Monthly").
	Where("factor_id = ?", id).First(&Factor).Error
	if err != nil {
		return nil, err
	}
	return &Factor, nil
}

func (fs *factorService) GetList() ([]model.Factor, error) {
	var result []model.Factor
	err := fs.db.
	Preload(clause.Associations).
	Preload("Plan.Monthly").
	Preload("Actual.Monthly").
	Find(&result).Error
	if err != nil{
		return []model.Factor{}, err
	}
	return result, nil 
}
