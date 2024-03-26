package service

import (
	"goreact/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

)

type AnalisaService interface {
	Store(Analisa *model.Analisa) error
	Update(id int, analisa model.Analisa) error
	Saves(analisa model.Analisa) error
	Delete(id int) error
	GetByID(id int) (*model.Analisa, error)
	GetList() ([]model.Analisa, error)
}

type analisaService struct {
	db *gorm.DB
}

func NewAnalisaService(db *gorm.DB) *analisaService {
	return &analisaService{db}
}

func (ms *analisaService) Store(analisa *model.Analisa) error {
	return ms.db.Create(analisa).Error
}
func (ms *analisaService) Update(id int, analisa model.Analisa) error {
	return ms.db.Where(id).Updates(analisa).Error
}
func (ms *analisaService) Saves(analisa model.Analisa) error{
	return ms.db.Save(analisa).Error
}
func (ms *analisaService) Delete(id int) error {	
	return ms.db.Where(id).Delete(&model.Analisa{}).Error 
}
func (ms *analisaService) GetByID(id int) (*model.Analisa, error) {
	var Analisa model.Analisa
	err := ms.db.Preload(clause.Associations).Where(id).First(&Analisa).Error
	if err != nil {
		return nil, err
	}
	return &Analisa, nil
}
func (ms *analisaService) GetList() ([]model.Analisa, error) {
	var result []model.Analisa
	err := ms.db.Preload(clause.Associations).Find(&result).Error
	if err != nil{
		return []model.Analisa{}, err
	}
	return result, nil 
}