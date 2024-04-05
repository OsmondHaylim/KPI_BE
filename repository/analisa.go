package repository

import (
	"goreact/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

)

type AnalisaRepo interface {
	Store(Analisa *model.Analisa) error
	Update(id int, analisa model.Analisa) error
	Saves(analisa model.Analisa) error
	Delete(id int) error
	GetByID(id int) (*model.Analisa, error)
	GetList() ([]model.Analisa, error)
}

type analisaRepo struct {
	db *gorm.DB
}

func NewAnalisaRepo(db *gorm.DB) *analisaRepo {
	return &analisaRepo{db}
}

func (ms *analisaRepo) Store(analisa *model.Analisa) error {
	return ms.db.Create(analisa).Error
}
func (ms *analisaRepo) Update(id int, analisa model.Analisa) error {
	return ms.db.Where(id).Updates(analisa).Error
}
func (ms *analisaRepo) Saves(analisa model.Analisa) error{
	return ms.db.Save(analisa).Error
}
func (ms *analisaRepo) Delete(id int) error {	
	return ms.db.Where(id).Delete(&model.Analisa{}).Error 
}
func (ms *analisaRepo) GetByID(id int) (*model.Analisa, error) {
	var Analisa model.Analisa
	err := ms.db.Preload(clause.Associations).Where(id).First(&Analisa).Error
	if err != nil {
		return nil, err
	}
	return &Analisa, nil
}
func (ms *analisaRepo) GetList() ([]model.Analisa, error) {
	var result []model.Analisa
	err := ms.db.Preload(clause.Associations).Find(&result).Error
	if err != nil{
		return []model.Analisa{}, err
	}
	return result, nil 
}