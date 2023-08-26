package repository_mia

import "gorm.io/gorm"

type mia struct {
	db *gorm.DB
}

func NewMia(db *gorm.DB) *mia {
	return &mia{db: db}
}