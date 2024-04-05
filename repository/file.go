package repository

import (
	"goreact/model"
	"gorm.io/gorm"
)

type FileRepo interface {
	Store(UploadFile *model.UploadFile) error
	Saves(file model.UploadFile) error
	Delete(id int) error
	GetByID(id int) (*model.UploadFile, error)
	GetList() ([]model.UploadFile, error)
}

type fileRepo struct {
	db *gorm.DB
}

func NewFileRepo(db *gorm.DB) *fileRepo {
	return &fileRepo{db}
}

func (fs *fileRepo) Store(file *model.UploadFile) error {
	return fs.db.Create(file).Error
}

func (fs *fileRepo) Saves(file model.UploadFile) error {
	return fs.db.Save(file).Error
}

func (fs *fileRepo) Delete(id int) error {	
	return fs.db.Where(id).Delete(&model.UploadFile{}).Error 
}

func (fs *fileRepo) GetByID(id int) (*model.UploadFile, error) {
	var UploadFile model.UploadFile
	err := fs.db.Where("monthly_id = ?", id).First(&UploadFile).Error
	if err != nil {
		return nil, err
	}
	return &UploadFile, nil
}

func (fs *fileRepo) GetList() ([]model.UploadFile, error) {
	var result []model.UploadFile
	err := fs.db.Find(&result).Error
	if err != nil{
		return []model.UploadFile{}, err
	}
	return result, nil 
}
