package repository

import (
	"goreact/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ResultRepo interface {
	Store(Result *model.Result) error
	Update(id int, result model.Result) error
	Saves(result model.Result) error
	Delete(id int) error
	GetByID(id int) (*model.Result, error)
	GetList() ([]model.Result, error)
}

type resultRepo struct {
	db *gorm.DB
}

func NewResultRepo(db *gorm.DB) *resultRepo {
	return &resultRepo{db}
}

func (rs *resultRepo) Store(result *model.Result) error {
	return rs.db.Create(result).Error
}

func (rs *resultRepo) Update(id int, result model.Result) error {
	return rs.db.Where(id).Updates(result).Error
}

func (rs *resultRepo) Saves(result model.Result) error {
	return rs.db.Save(result).Error
}

func (rs *resultRepo) Delete(id int) error {	
	return rs.db.Where(id).Delete(&model.Result{}).Error 
}

func (rs *resultRepo) GetByID(id int) (*model.Result, error) {
	var Result model.Result
	err := rs.db.
	Preload(clause.Associations).
	Preload("Factors.Plan").
	Preload("Factors.Actual").
	Preload("Factors.Plan.Monthly").
	Preload("Factors.Actual.Monthly").
	Where("result_id = ?", id).
	First(&Result).Error
	if err != nil {
		return nil, err
	}
	return &Result, nil
}

func (rs *resultRepo) GetList() ([]model.Result, error) {
	var result []model.Result
	err := rs.db.
	Preload(clause.Associations).
	Preload("Factors.Plan").
	Preload("Factors.Actual").
	Preload("Factors.Plan.Monthly").
	Preload("Factors.Actual.Monthly").
	Find(&result).Error
	if err != nil{
		return []model.Result{}, err
	}
	return result, nil 
}
