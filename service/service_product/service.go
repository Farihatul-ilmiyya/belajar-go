package service_product

import (
	"echo-golang/injector/injectorRepository"
	"echo-golang/model/product"
	"errors"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type ServiceProduct interface {
	Create(request product.RequestProduct) (product.ResponseProduct, int, error)
}

type service_product struct {
	repository injectorRepository.RepositoryInjector
	logging    *logrus.Logger
}

func NewServiceProduct(repository injectorRepository.RepositoryInjector, logging *logrus.Logger) *service_product {
	return &service_product{repository: repository, logging: logging}
}

func (s *service_product) Create(request product.RequestProduct) (product.ResponseProduct, int, error) {
	s.logging.Info("User memasukkan input")
	time.Sleep(5 * time.Second)
	s.logging.Info("melakukan formatting")
	reqFormatter := product.Product{
		Name:   request.Name,
		Price:  request.Price,
		Weight: request.Weight,
		Brand:  request.Brand,
	}
	time.Sleep(5 * time.Second)
	s.logging.Info("mengirim data ke repository")
	result, err := s.repository.RepositoryProduct.Create(reqFormatter)
	if err != nil {
		s.logging.Error("error create product ", err)
		return product.ResponseProduct{}, http.StatusInternalServerError, errors.New("something want wrong")
	}
	time.Sleep(5 * time.Second)
	s.logging.Info("melakukan formatting untuk response")
	resultFormatter := product.ResponseProduct{
		ID:        result.ID,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
		Name:      result.Name,
		Price:     result.Price,
		Weight:    result.Weight,
		Brand:     result.Brand,
	}
	time.Sleep(3 * time.Second)
	s.logging.Info("mengembalikan data ke handler")
	return resultFormatter, http.StatusCreated, nil
}
