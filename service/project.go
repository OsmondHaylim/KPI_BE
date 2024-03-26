package service

import (
	"goreact/model"
	"gorm.io/gorm"

)

type ProjectService interface {
	Store(Project *model.Project) error
	Update(id int, project model.Project) error
	Saves(project model.Project) error
	Delete(id int) error
	GetByID(id int) (*model.Project, error)
	GetList() ([]model.Project, error)
}
type projectService struct {
	db *gorm.DB
}

func NewProjectService(db *gorm.DB) *projectService {
	return &projectService{db}
}

func (ms *projectService) Store(project *model.Project) error {
	return ms.db.Create(project).Error
}
func (ms *projectService) Update(id int, project model.Project) error {
	return ms.db.Where(id).Updates(project).Error
}
func (ms *projectService) Saves(project model.Project) error{
	return ms.db.Save(project).Error
}
func (ms *projectService) Delete(id int) error {	
	return ms.db.Where(id).Delete(&model.Project{}).Error 
}
func (ms *projectService) GetByID(id int) (*model.Project, error) {
	var Project model.Project
	err := ms.db.Where(id).First(&Project).Error
	if err != nil {
		return nil, err
	}
	return &Project, nil
}
func (ms *projectService) GetList() ([]model.Project, error) {
	var result []model.Project
	err := ms.db.Find(&result).Error
	if err != nil{
		return []model.Project{}, err
	}
	return result, nil 
}
