package repository_nanda

import "gorm.io/gorm"

type product struct {
	db *gorm.DB
}

func NewProduct(db *gorm.DB) *product {
	return &product{db: db}
}
