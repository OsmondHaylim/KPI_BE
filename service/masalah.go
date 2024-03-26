package service

import (
	"goreact/model"
	"gorm.io/gorm"

)

type MasalahService interface {
	Store(Masalah *model.Masalah) error
	Update(id int, masalah model.Masalah) error
	Saves(masalah model.Masalah) error
	Delete(id int) error
	GetByID(id int) (*model.Masalah, error)
	GetList() ([]model.Masalah, error)
}

type masalahService struct {
	db *gorm.DB
}

func NewMasalahService(db *gorm.DB) *masalahService {
	return &masalahService{db}
}

func (ms *masalahService) Store(masalah *model.Masalah) error {
	return ms.db.Create(masalah).Error
}
func (ms *masalahService) Update(id int, masalah model.Masalah) error {
	return ms.db.Where(id).Updates(masalah).Error
}
func (ms *masalahService) Saves(masalah model.Masalah) error{
	return ms.db.Save(masalah).Error
}
func (ms *masalahService) Delete(id int) error {	
	return ms.db.Where(id).Delete(&model.Masalah{}).Error 
}
func (ms *masalahService) GetByID(id int) (*model.Masalah, error) {
	var Masalah model.Masalah
	err := ms.db.Where(id).First(&Masalah).Error
	if err != nil {
		return nil, err
	}
	return &Masalah, nil
}
func (ms *masalahService) GetList() ([]model.Masalah, error) {
	var result []model.Masalah
	err := ms.db.Find(&result).Error
	if err != nil{
		return []model.Masalah{}, err
	}
	return result, nil 
}
