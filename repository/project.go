package repository

import (
	"goreact/model"
	"gorm.io/gorm"

)

type ProjectRepo interface {
	Store(Project *model.Project) error
	Update(id int, project model.Project) error
	Saves(project model.Project) error
	Delete(id int) error
	GetByID(id int) (*model.Project, error)
	GetList() ([]model.Project, error)
}
type projectRepo struct {
	db *gorm.DB
}

func NewProjectRepo(db *gorm.DB) *projectRepo {
	return &projectRepo{db}
}

func (ms *projectRepo) Store(project *model.Project) error {
	return ms.db.Create(project).Error
}
func (ms *projectRepo) Update(id int, project model.Project) error {
	return ms.db.Where(id).Updates(project).Error
}
func (ms *projectRepo) Saves(project model.Project) error{
	return ms.db.Save(project).Error
}
func (ms *projectRepo) Delete(id int) error {	
	return ms.db.Where(id).Delete(&model.Project{}).Error 
}
func (ms *projectRepo) GetByID(id int) (*model.Project, error) {
	var Project model.Project
	err := ms.db.Where(id).First(&Project).Error
	if err != nil {
		return nil, err
	}
	return &Project, nil
}
func (ms *projectRepo) GetList() ([]model.Project, error) {
	var result []model.Project
	err := ms.db.Find(&result).Error
	if err != nil{
		return []model.Project{}, err
	}
	return result, nil 
}
