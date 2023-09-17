package service_product_test

// import (
// 	"echo-golang/model/product"
// 	"echo-golang/pkg/helpers"
// 	"echo-golang/repository/repository_product/mocks"
// 	"fmt"
// 	"net/http"
// 	"testing"

// 	"github.com/sirupsen/logrus"
// )

// func TestCreate1(t *testing.T) {
// 	repository := &mocks.RepositoryProduct{}
// 	service := NewServiceProduct(repository, &logrus.Logger{})

// 	productModel := product.Product{
// 		Name:   "product 2",
// 		Price:  120,
// 		Weight: 20,
// 		Brand:  "Brand 1",
// 	}
// 	repository.On("Create", productModel).Return(productModel, nil)

// 	requestProduct := product.RequestProduct{
// 		Name:   "product 1",
// 		Price:  120,
// 		Weight: 20,
// 		Brand:  "Brand 1",
// 	}
// 	response, status, err := service.Create(requestProduct)

// 	if err != nil {
// 		t.Errorf("Unexpected error: %v", err)
// 		return
// 	}
// 	if status != http.StatusCreated {
// 		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, status)
// 	}
// 	fmt.Println(helpers.PrettyPrint(response))
// }
