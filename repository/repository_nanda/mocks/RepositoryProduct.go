// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	product "echo-golang/model/product"

	mock "github.com/stretchr/testify/mock"
)

// RepositoryProduct is an autogenerated mock type for the RepositoryProduct type
type RepositoryProduct struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *RepositoryProduct) Create(_a0 product.Product) (product.Product, error) {
	ret := _m.Called(_a0)

	var r0 product.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(product.Product) (product.Product, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(product.Product) product.Product); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(product.Product)
	}

	if rf, ok := ret.Get(1).(func(product.Product) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepositoryProduct interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepositoryProduct creates a new instance of RepositoryProduct. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepositoryProduct(t mockConstructorTestingTNewRepositoryProduct) *RepositoryProduct {
	mock := &RepositoryProduct{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
