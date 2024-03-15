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

func (rs *resultService) Store(result *model.Result) error {
	return rs.db.Create(result).Error
}

func (rs *resultService) Update(id int, result model.Result) error {
	return rs.db.Where(id).Updates(result).Error
}

func (rs *resultService) Delete(id int) error {	
	return rs.db.Where(id).Delete(&model.Result{}).Error 
}

func (rs *resultService) GetByID(id int) (*model.Result, error) {
	var Result model.Result
	err := rs.db.
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

func (rs *resultService) GetList() ([]model.Result, error) {
	var result []model.Result
	err := rs.db.
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
