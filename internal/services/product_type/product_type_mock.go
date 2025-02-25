// Code generated by mockery v2.52.3. DO NOT EDIT.

package product_type

import (
	models "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// ProductTypeServiceMock is an autogenerated mock type for the ProductTypeServiceMock type
type ProductTypeServiceMock struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *ProductTypeServiceMock) Create(_a0 models.ProductTypeDocRequest) (*models.ProductTypeDocResponse, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *models.ProductTypeDocResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(models.ProductTypeDocRequest) (*models.ProductTypeDocResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(models.ProductTypeDocRequest) *models.ProductTypeDocResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ProductTypeDocResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(models.ProductTypeDocRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with no fields
func (_m *ProductTypeServiceMock) GetAll() ([]models.ProductTypeDocResponse, error) {
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
func (_m *ProductTypeServiceMock) GetById(_a0 int) (*models.ProductTypeDocResponse, error) {
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

// NewProductTypeService creates a new instance of ProductTypeServiceMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProductTypeServiceMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProductTypeServiceMock {
	mock := &ProductTypeServiceMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
