package repository

import (
	"goreact/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MiniPAPRepo interface {
	Store(MiniPAP *model.MiniPAP) error
	Update(id int, minipap model.MiniPAP) error
	Saves(minipap model.MiniPAP) error
	Delete(id int) error
	GetByID(id int) (*model.MiniPAP, error)
	GetList() ([]model.MiniPAP, error)
}

type minipapRepo struct {
	db *gorm.DB
}

func NewMiniPAPRepo(db *gorm.DB) *minipapRepo {
	return &minipapRepo{db}
}

func (ms *minipapRepo) Store(minipap *model.MiniPAP) error {
	return ms.db.Create(minipap).Error
}

func (ms *minipapRepo) Update(id int, minipap model.MiniPAP) error {
	return ms.db.Where(id).Updates(minipap).Error
}

func (ms *minipapRepo) Saves(minipap model.MiniPAP) error {
	return ms.db.Save(minipap).Error
}

func (ms *minipapRepo) Delete(id int) error {	
	return ms.db.Where(id).Delete(&model.MiniPAP{}).Error 
}

func (ms *minipapRepo) GetByID(id int) (*model.MiniPAP, error) {
	var MiniPAP model.MiniPAP
	err := ms.db.
	Preload(clause.Associations).
	Where("minipap_id = ?", id).First(&MiniPAP).Error
	if err != nil {
		return nil, err
	}
	return &MiniPAP, nil
}

func (ms *minipapRepo) GetList() ([]model.MiniPAP, error) {
	var result []model.MiniPAP
	err := ms.db.
	Preload(clause.Associations).
	Find(&result).Error
	if err != nil{
		return []model.MiniPAP{}, err
	}
	return result, nil 
}
