// Code generated by mockery v2.52.3. DO NOT EDIT.

package product_type

import (
	models "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// ProductTypeRepositoryMock is an autogenerated mock type for the ProductTypeRepositoryMock type
type ProductTypeRepositoryMock struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *ProductTypeRepositoryMock) Create(_a0 *models.ProductTypeDocResponse) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.ProductTypeDocResponse) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExistInDb provides a mock function with given fields: _a0
func (_m *ProductTypeRepositoryMock) ExistInDb(_a0 string) (bool, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for ExistInDb")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with no fields
func (_m *ProductTypeRepositoryMock) GetAll() ([]models.ProductTypeDocResponse, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []models.ProductTypeDocResponse
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.ProductTypeDocResponse, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.ProductTypeDocResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ProductTypeDocResponse)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: _a0
func (_m *ProductTypeRepositoryMock) GetById(_a0 int) (*models.ProductTypeDocResponse, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *models.ProductTypeDocResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*models.ProductTypeDocResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(int) *models.ProductTypeDocResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ProductTypeDocResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewProductTypeRepository creates a new instance of ProductTypeRepositoryMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProductTypeRepositoryMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProductTypeRepositoryMock {
	mock := &ProductTypeRepositoryMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
