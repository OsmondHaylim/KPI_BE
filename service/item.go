package service

import (
	"goreact/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ItemService interface {
	Store(Item *model.Item) error
	Update(id int, item model.Item) error
	Saves(item model.Item) error
	Delete(id int) error
	GetByID(id int) (*model.Item, error)
	GetList() ([]model.Item, error)
}

type itemService struct {
	db *gorm.DB
}

func NewItemService(db *gorm.DB) *itemService {
	return &itemService{db}
}

func (is *itemService) Store(item *model.Item) error {
	return is.db.Create(item).Error
}

func (is *itemService) Update(id int, item model.Item) error {
	return is.db.Where(id).Updates(item).Error
}

func (is *itemService) Saves(item model.Item) error {
	return is.db.Save(item).Error
}

func (is *itemService) Delete(id int) error {	
	return is.db.Where(id).Delete(&model.Item{}).Error 
}

func (is *itemService) GetByID(id int) (*model.Item, error) {
	var Item model.Item
	err := is.db.
	Preload(clause.Associations).
	Preload("Results.Factors").
	Preload("Results.Factors.Statistic").
	Preload("Results.Factors.Statistic.Plan").
	Preload("Results.Factors.Statistic.Actual").
	Preload("Results.Factors.Statistic.Plan.Monthly").
	Preload("Results.Factors.Statistic.Actual.Monthly").
	Where("item_id = ?", id).
	First(&Item).Error
	if err != nil {
		return nil, err
	}
	return &Item, nil
}

func (is *itemService) GetList() ([]model.Item, error) {
	var item []model.Item
	err := is.db.
	Preload(clause.Associations).
	Preload("Results.Factors").
	Preload("Results.Factors.Statistic").
	Preload("Results.Factors.Statistic.Plan").
	Preload("Results.Factors.Statistic.Actual").
	Preload("Results.Factors.Statistic.Plan.Monthly").
	Preload("Results.Factors.Statistic.Actual.Monthly").
	Find(&item).Error
	if err != nil{
		return []model.Item{}, err
	}
	return item, nil 
}
