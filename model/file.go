package model

import "gorm.io/gorm"

type UploadFile struct {
	gorm.Model
	FileName 	string	`gorm:"notNull"`
	File 		[]byte	`gorm:"type:bytea"`
}

