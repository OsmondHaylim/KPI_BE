package repository

import (
	"goreact/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FactorRepo interface {
	Store(Factor *model.Factor) error
	Update(id int, factor model.Factor) error
	UpdateNecessary(id int, factor model.Factor) error
	Saves(factor model.Factor) error
	Delete(id int) error
	GetByID(id int) (*model.Factor, error)
	GetList() ([]model.Factor, error)
}

type factorRepo struct {
	db *gorm.DB
}

func NewFactorRepo(db *gorm.DB) *factorRepo {
	return &factorRepo{db}
}

func (fs *factorRepo) Store(factor *model.Factor) error {
	return fs.db.Create(factor).Error
}

func (fs *factorRepo) Update(id int, factor model.Factor) error {
	return fs.db.Where(id).Updates(factor).Error
}

func (fs *factorRepo) UpdateNecessary(id int, factor model.Factor) error {
	return fs.db.Where(id).Omit("Plan").Omit("Actual").Updates(factor).Error
}

func (fs *factorRepo) Saves(factor model.Factor) error {
	return fs.db.Save(factor).Error
}

func (fs *factorRepo) Delete(id int) error {	
	return fs.db.Where(id).Delete(&model.Factor{}).Error 
}

func (fs *factorRepo) GetByID(id int) (*model.Factor, error) {
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

func (fs *factorRepo) GetList() ([]model.Factor, error) {
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
