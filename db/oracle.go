package db

// import (
// 	"fmt"
// 	"gorm.io/driver/oracle"
// 	"gorm.io/gorm"

// 	// "agatra/model"
// )

// type Oracle struct{}

// type Config struct{
// 	Host		string
// 	Port		string
// 	Password	string
// 	User		string
// 	DBName		string
// }

// func (o *Oracle) Connect(creds *Config) (*gorm.DB, error) {
// 	dsn := fmt.Sprintf("%s/%s@%s:%s/%s", creds.User, creds.Password, creds.Host, creds.Port, creds.DBName)

// 	dbConn, err := gorm.Open(oracle.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return dbConn, nil
// }

// func NewDB() *Oracle {
// 	return &Oracle{}
// }

// func (o *Oracle) Reset(db *gorm.DB, table string) error {
// 	return db.Transaction(func(tx *gorm.DB) error {
// 		if err := tx.Exec("TRUNCATE " + table).Error; err != nil {
// 			return err
// 		}

// 		if err := tx.Exec("ALTER SEQUENCE " + table + "_id_seq RESTART WITH 1").Error; err != nil {
// 			return err
// 		}

// 		return nil
// 	})
// }