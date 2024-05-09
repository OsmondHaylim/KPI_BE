package repository

import (
	"goreact/model"
	"gorm.io/gorm"
)

type UserRepo interface {
	Store(User *model.User) error
	Update(id int, user model.User) error
	Delete(id int) error
	GetByID(id int) (*model.User, error)
	GetList() ([]model.User, error)
	GetByEmail(Email string) (model.User, bool)
	GetByName(Name string) (model.User, bool)
	// GetPrivileged() ([]model.User, error)
	SearchName(name string) ([]model.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{db}
}

func (us *userRepo) Store(User *model.User) error {
	return us.db.Create(User).Error
}

func (us *userRepo) Update(id int, user model.User) error {
	return us.db.Where(id).Updates(user).Error
}

func (us *userRepo) Delete(id int) error {	
	return us.db.Where(id).Delete(&model.User{}).Error 
}

func (us *userRepo) GetByID(id int) (*model.User, error) {
	var User model.User
	err := us.db.Where("id = ?", id).First(&User).Error
	if err != nil {
		return nil, err
	}
	return &User, nil
}

func (us *userRepo) GetList() ([]model.User, error) {
	var result []model.User
	err := us.db.Find(&result).Error
	if err != nil {
		return []model.User{}, err
	}
	return result, nil 
}

// func (us *userRepo) GetPrivileged() ([]model.User, error) {
// 	var result []model.User
// 	rows, err := us.db.Where("role = ?", "admin").Or("role = ?", "maintainer").Table("users").Rows()
// 	if err != nil || rows == nil{
// 		return []model.User{}, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() { 
// 		us.db.ScanRows(rows, &result)
// 	}
// 	return result, nil 
// }

func (us *userRepo) GetByEmail(email string) (model.User, bool) {
	var result model.User
	err := us.db.Where("email = ?", email).First(&result).Error
	if err != nil {
		return model.User{}, false
	}
	return result, true
}

func (us *userRepo) GetByName(name string) (model.User, bool) {
	var result model.User
	err := us.db.Where("username = ?", name).First(&result).Error
	if err != nil {
		return model.User{}, false
	}
	return result, true
}

func (us *userRepo) SearchName(name string) ([]model.User, error){
	var result []model.User
	rows, err := us.db.Where("username LIKE ?", "%" + name + "%").Table("users").Rows()
	if err != nil || rows == nil{
		return []model.User{}, err
	}
	defer rows.Close()
	for rows.Next() { 
		us.db.ScanRows(rows, &result)
	}
	return result, nil 
}