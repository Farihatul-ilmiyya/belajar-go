package injectorService

import (
	"echo-golang/injector/injectorRepository"
	"echo-golang/model/product"
	"echo-golang/pkg/logs"
	servicenanda "echo-golang/service/service_nanda"
	"echo-golang/service/service_product"
)

type InjectorService struct {
	Product      service_product.ServiceProduct
	ProductNanda servicenanda.ServiceProductNanda
}

func NewInjectorService(repository injectorRepository.RepositoryInjector) InjectorService {
	return InjectorService{
		Product:      service_product.NewServiceProduct(repository, logs.NewLogger()),
		ProductNanda: servicenanda.NewServiceProduct(repository, logs.NewLogger()),
	}
}
func NewService(repository *injectorRepository.RepositoryMock) InjectorService {
	// repository := injectorRepository.NewRepositoryMock()
	productModel := product.Product{
		Name:   "product 1",
		Price:  120,
		Weight: 20,
		Brand:  "Brand 1",
	}

	injector := InjectorService{
		Product: service_product.NewServiceProduct(injectorRepository.RepositoryInjector{
			RepositoryProduct: &repository.RepositoryProduct,
		}, logs.NewLogger()),
		ProductNanda: servicenanda.NewServiceProduct(injectorRepository.RepositoryInjector{
			RepositoryProductNanda: &repository.RepositoryProductNanda,
		}, logs.NewLogger()),
	}

	repository.RepositoryProduct.On("Create", productModel).Return(productModel, nil)
	repository.RepositoryProductNanda.On("Create", productModel).Return(productModel, nil)
	return injector
}
