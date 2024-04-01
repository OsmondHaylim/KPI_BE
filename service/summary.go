package service

import (
	"goreact/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SummaryService interface {
	Store(Summary *model.Summary) error
	Update(id int, summary model.Summary) error
	Saves(summary model.Summary) error
	Delete(id int) error
	GetByID(id int) (*model.Summary, error)
	GetList() ([]model.Summary, error)
}

type summaryService struct {
	db *gorm.DB
}

func NewSummaryService(db *gorm.DB) *summaryService {
	return &summaryService{db}
}

func (ms *summaryService) Store(summary *model.Summary) error {
	return ms.db.Create(summary).Error
}

func (ms *summaryService) Update(id int, summary model.Summary) error {
	return ms.db.Where(id).Updates(summary).Error
}

func (ms *summaryService) Saves(summary model.Summary) error {
	return ms.db.Save(summary).Error
}

func (ms *summaryService) Delete(id int) error {	
	return ms.db.Where(id).Delete(&model.Summary{}).Error 
}

func (ms *summaryService) GetByID(id int) (*model.Summary, error) {
	var Summary model.Summary
	err := ms.db.
	Preload(clause.Associations).
	Where("summary_id = ?", id).First(&Summary).Error
	if err != nil {
		return nil, err
	}
	return &Summary, nil
}

func (ms *summaryService) GetList() ([]model.Summary, error) {
	var result []model.Summary
	err := ms.db.
	Preload(clause.Associations).
	Find(&result).Error
	if err != nil{
		return []model.Summary{}, err
	}
	return result, nil 
}
