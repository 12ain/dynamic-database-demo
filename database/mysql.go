package database

import (
	"dynamic-database-demo/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLInstance() (*gorm.DB, error) {
	// read from config file or environment variables
	dsn := viper.GetString("database.mysql.dsn")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.User{})

	return db, nil
}
