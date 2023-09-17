package service_product_test

import (
	"echo-golang/injector/injectorRepository"
	"echo-golang/injector/injectorService"
	"echo-golang/model/product"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	repository := injectorRepository.NewRepositoryMock()
	service := injectorService.NewService(&repository)
	requestProduct := product.RequestProduct{
		Name:   "product 1",
		Price:  120,
		Weight: 20,
		Brand:  "Brand 1",
	}
	response, status, err := service.ProductNanda.Create(requestProduct)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, status)
	assert.NotNil(t, response)
	assert.Equal(t, requestProduct.Name, response.Name)
	assert.Equal(t, requestProduct.Price, response.Price)
	assert.Equal(t, requestProduct.Weight, response.Weight)
	assert.Equal(t, requestProduct.Brand, response.Brand)
}
