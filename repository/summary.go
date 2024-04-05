package repository

import (
	"goreact/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SummaryRepo interface {
	Store(Summary *model.Summary) error
	Update(id int, summary model.Summary) error
	Saves(summary model.Summary) error
	Delete(id int) error
	GetByID(id int) (*model.Summary, error)
	GetList() ([]model.Summary, error)
}

type summaryRepo struct {
	db *gorm.DB
}

func NewSummaryRepo(db *gorm.DB) *summaryRepo {
	return &summaryRepo{db}
}

func (ms *summaryRepo) Store(summary *model.Summary) error {
	return ms.db.Create(summary).Error
}

func (ms *summaryRepo) Update(id int, summary model.Summary) error {
	return ms.db.Where(id).Updates(summary).Error
}

func (ms *summaryRepo) Saves(summary model.Summary) error {
	return ms.db.Save(summary).Error
}

func (ms *summaryRepo) Delete(id int) error {	
	return ms.db.Where(id).Delete(&model.Summary{}).Error 
}

func (ms *summaryRepo) GetByID(id int) (*model.Summary, error) {
	var Summary model.Summary
	err := ms.db.
	Preload(clause.Associations).
	Where("summary_id = ?", id).First(&Summary).Error
	if err != nil {
		return nil, err
	}
	return &Summary, nil
}

func (ms *summaryRepo) GetList() ([]model.Summary, error) {
	var result []model.Summary
	err := ms.db.
	Preload(clause.Associations).
	Find(&result).Error
	if err != nil{
		return []model.Summary{}, err
	}
	return result, nil 
}
