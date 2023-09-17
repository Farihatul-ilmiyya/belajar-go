package injectorRepository

import (
	"echo-golang/repository/repository_nanda"
	repositoryNandaMock "echo-golang/repository/repository_nanda/mocks"
	"echo-golang/repository/repository_product"
	"echo-golang/repository/repository_product/mocks"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RepositoryInjector struct {
	RepositoryProduct      repository_product.RepositoryProduct
	RepositoryProductNanda repository_nanda.RepositoryProduct
}

func NewInjectorRepository(db *gorm.DB, logging *logrus.Logger) RepositoryInjector {
	return RepositoryInjector{
		RepositoryProduct: repository_product.NewRepositoryProduct(db, logging),
	}
}

type RepositoryMock struct {
	RepositoryProduct      mocks.RepositoryProduct
	RepositoryProductNanda repositoryNandaMock.RepositoryProduct
}

func NewRepositoryMock() RepositoryMock {
	return RepositoryMock{
		RepositoryProduct:      mocks.RepositoryProduct{},
		RepositoryProductNanda: repositoryNandaMock.RepositoryProduct{},
	}
}
