package service

import (
	"goreact/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PAPService interface {
	Store(PAP *model.PAP) error
	Update(id int, pap model.PAP) error
	Delete(id int) error
	GetByID(id int) (*model.PAP, error)
	GetList() ([]model.PAP, error)
}

type papService struct {
	db *gorm.DB
}

func NewPAPService(db *gorm.DB) *papService {
	return &papService{db}
}

func (as *papService) Store(pap *model.PAP) error {
	return as.db.Create(pap).Error
}

func (as *papService) Update(id int, pap model.PAP) error {
	return as.db.Where(id).Updates(pap).Error
}

func (as *papService) Delete(id int) error {	
	return as.db.Where(id).Delete(&model.PAP{}).Error 
}

func (as *papService) GetByID(id int) (*model.PAP, error) {
	var PAP model.PAP
	err := as.db.
	Preload(clause.Associations).
	Preload("MiniPAP.monthly").
	Where("id = ?", id).First(&PAP).Error
	if err != nil {
		return nil, err
	}
	return &PAP, nil
}

func (as *papService) GetList() ([]model.PAP, error) {
	var result []model.PAP
	err := as.db.
	Preload(clause.Associations).
	Preload("MiniPAP.monthly").
	Find(&result).Error
	if err != nil{
		return []model.PAP{}, err
	}
	return result, nil 
}
