package repository

// import (
// 	"goreact/model"
// 	"gorm.io/gorm"

// )

// type MasalahRepo interface {
// 	Store(Masalah *model.Masalah) error
// 	Update(id int, masalah model.Masalah) error
// 	Saves(masalah model.Masalah) error
// 	Delete(id int) error
// 	GetByID(id int) (*model.Masalah, error)
// 	GetList() ([]model.Masalah, error)
// }

// type masalahRepo struct {
// 	db *gorm.DB
// }

// func NewMasalahRepo(db *gorm.DB) *masalahRepo {
// 	return &masalahRepo{db}
// }

// func (ms *masalahRepo) Store(masalah *model.Masalah) error {
// 	return ms.db.Create(masalah).Error
// }
// func (ms *masalahRepo) Update(id int, masalah model.Masalah) error {
// 	return ms.db.Where(id).Updates(masalah).Error
// }
// func (ms *masalahRepo) Saves(masalah model.Masalah) error{
// 	return ms.db.Save(masalah).Error
// }
// func (ms *masalahRepo) Delete(id int) error {	
// 	return ms.db.Where(id).Delete(&model.Masalah{}).Error 
// }
// func (ms *masalahRepo) GetByID(id int) (*model.Masalah, error) {
// 	var Masalah model.Masalah
// 	err := ms.db.Where(id).First(&Masalah).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Masalah, nil
// }
// func (ms *masalahRepo) GetList() ([]model.Masalah, error) {
// 	var result []model.Masalah
// 	err := ms.db.Find(&result).Error
// 	if err != nil{
// 		return []model.Masalah{}, err
// 	}
// 	return result, nil 
// }
