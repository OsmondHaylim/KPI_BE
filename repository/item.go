package repository

import (
	"goreact/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ItemRepo interface {
	Store(Item *model.Item) error
	Update(id int, item model.Item) error
	UpdateNecessary(id int, item model.Item) error
	Saves(item model.Item) error
	Delete(id int) error
	GetByID(id int) (*model.Item, error)
	GetList() ([]model.Item, error)
}

type itemRepo struct {
	db *gorm.DB
}

func NewItemRepo(db *gorm.DB) *itemRepo {
	return &itemRepo{db}
}

func (is *itemRepo) Store(item *model.Item) error {
	return is.db.Create(item).Error
}

func (is *itemRepo) Update(id int, item model.Item) error {
	return is.db.Where(id).Updates(item).Error
}

func (is *itemRepo) UpdateNecessary(id int, item model.Item) error {
	return is.db.Where(id).Omit("Results").Updates(item).Error
}

func (is *itemRepo) Saves(item model.Item) error {
	return is.db.Save(item).Error
}

func (is *itemRepo) Delete(id int) error {	
	return is.db.Where(id).Delete(&model.Item{}).Error 
}

func (is *itemRepo) GetByID(id int) (*model.Item, error) {
	var Item model.Item
	err := is.db.
	Preload(clause.Associations).
	Preload("Results.Factors").
	Preload("Results.Factors.Plan").
	Preload("Results.Factors.Actual").
	Preload("Results.Factors.Plan.Monthly").
	Preload("Results.Factors.Actual.Monthly").
	Where("item_id = ?", id).
	First(&Item).Error
	if err != nil {
		return nil, err
	}
	return &Item, nil
}

func (is *itemRepo) GetList() ([]model.Item, error) {
	var item []model.Item
	err := is.db.
	Preload(clause.Associations).
	Preload("Results.Factors").
	Preload("Results.Factors.Plan").
	Preload("Results.Factors.Actual").
	Preload("Results.Factors.Plan.Monthly").
	Preload("Results.Factors.Actual.Monthly").
	Find(&item).Error
	if err != nil{
		return []model.Item{}, err
	}
	return item, nil 
}
