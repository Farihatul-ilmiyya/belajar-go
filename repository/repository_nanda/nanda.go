package repository_nanda

import (
	"echo-golang/model/product"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)
//go:generate mockery --name RepositoryProduct
type RepositoryProduct interface {
	Create(product product.Product) (product.Product, error)
}

type repository_product struct {
	db      *gorm.DB
	logging *logrus.Logger
}

func NewRepositoryProduct(db *gorm.DB, logging *logrus.Logger) *repository_product {
	return &repository_product{db: db, logging: logging}
}

func (r *repository_product) Create(newProduct product.Product) (product.Product, error) {
	time.Sleep(5 * time.Second)
	r.logging.Info("menyimpan data ke database")

	err := r.db.Create(&newProduct).Error
	if err != nil {
		return product.Product{}, err
	}
	time.Sleep(5 * time.Second)
	r.logging.Info("data di kembalikan ke service")
	return newProduct, nil
}
