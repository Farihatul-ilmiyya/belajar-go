package repository_rahmat

import "gorm.io/gorm"

type product struct {
	db *gorm.DB
}

func New(db *gorm.DB) *product{
	return &product{db: db}
}
