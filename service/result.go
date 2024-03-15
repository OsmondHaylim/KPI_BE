package service

import (
	"goreact/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ResultService interface {
	Store(Result *model.Result) error
	Update(id int, result model.Result) error
	Delete(id int) error
	GetByID(id int) (*model.Result, error)
	GetList() ([]model.Result, error)
}

type resultService struct {
	db *gorm.DB
}

func NewResultService(db *gorm.DB) *resultService {
	return &resultService{db}
}

func (ks *resultService) Store(result *model.Result) error {
	return ks.db.Create(result).Error
}

func (ks *resultService) Update(id int, result model.Result) error {
	return ks.db.Where(id).Updates(result).Error
}

func (ks *resultService) Delete(id int) error {	
	return ks.db.Where(id).Delete(&model.Result{}).Error 
}

func (ks *resultService) GetByID(id int) (*model.Result, error) {
	var Result model.Result
	err := ks.db.
	Preload(clause.Associations).
	Preload("Factor.PAP").
	Preload("Factor.PAP.MiniPAP").
	Preload("Factor.PAP.MiniPAP.Monthly").
	Where("id = ?", id).
	First(&Result).Error
	if err != nil {
		return nil, err
	}
	return &Result, nil
}

func (ks *resultService) GetList() ([]model.Result, error) {
	var result []model.Result
	err := ks.db.
	Preload(clause.Associations).
	Preload("Factor.PAP").
	Preload("Factor.PAP.MiniPAP").
	Preload("Factor.PAP.MiniPAP.Monthly").
	Find(&result).Error
	if err != nil{
		return []model.Result{}, err
	}
	return result, nil 
}
