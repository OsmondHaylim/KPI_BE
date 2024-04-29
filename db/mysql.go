package db

import (
	// "fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MySQL struct{}

type Config struct{
	Host		string
	Port		string
	Password	string
	User		string
	DBName		string
	SSLMode		string
}

func (m *MySQL) Connect(creds *Config) (*gorm.DB, error) {
	dsn := "root@/kpiv"

	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Error)})
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

func NewDB() *MySQL {
	return &MySQL{}
}

func (m *MySQL) Reset(db *gorm.DB, table string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("TRUNCATE " + table).Error; err != nil {
			return err
		}

		if err := tx.Exec("ALTER SEQUENCE " + table + "_id_seq RESTART WITH 1").Error; err != nil {
			return err
		}

		return nil
	})
}