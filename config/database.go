package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() (*gorm.DB, error) {
	dsn := "db-user:db-password@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
