package service

import (
	"goreact/model"
	"gorm.io/gorm"
)

type FileService interface {
	Store(UploadFile *model.UploadFile) error
	Saves(file model.UploadFile) error
}

type fileService struct {
	db *gorm.DB
}

func NewFileService(db *gorm.DB) *fileService {
	return &fileService{db}
}

func (fs *fileService) Store(file *model.UploadFile) error {
	return fs.db.Create(file).Error
}

func (fs *fileService) Saves(file model.UploadFile) error {
	return fs.db.Save(file).Error
}

