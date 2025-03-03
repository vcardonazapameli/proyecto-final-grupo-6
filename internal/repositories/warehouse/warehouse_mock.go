// Code generated by mockery v2.52.3. DO NOT EDIT.

package warehouse

import (
	models "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// WarehouseRepositoryMock is an autogenerated mock type for the WarehouseRepositoryMock type
type WarehouseRepositoryMock struct {
	mock.Mock
}

// CreateWarehouse provides a mock function with given fields: _a0
func (_m *WarehouseRepositoryMock) CreateWarehouse(_a0 *models.WarehouseDocResponse) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for CreateWarehouse")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.WarehouseDocResponse) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteWarehouse provides a mock function with given fields: _a0
func (_m *WarehouseRepositoryMock) DeleteWarehouse(_a0 int) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for DeleteWarehouse")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExistInDbWarehouseCode provides a mock function with given fields: _a0
func (_m *WarehouseRepositoryMock) ExistInDbWarehouseCode(_a0 string) (bool, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for ExistInDbWarehouseCode")
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
func (_m *WarehouseRepositoryMock) GetAll() ([]models.WarehouseDocResponse, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []models.WarehouseDocResponse
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.WarehouseDocResponse, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.WarehouseDocResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.WarehouseDocResponse)
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
func (_m *WarehouseRepositoryMock) GetById(_a0 int) (*models.WarehouseDocResponse, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *models.WarehouseDocResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*models.WarehouseDocResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(int) *models.WarehouseDocResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.WarehouseDocResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MatchWarehouseCode provides a mock function with given fields: _a0, _a1
func (_m *WarehouseRepositoryMock) MatchWarehouseCode(_a0 int, _a1 string) (bool, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for MatchWarehouseCode")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(int, string) (bool, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(int, string) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(int, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateWarehouse provides a mock function with given fields: _a0, _a1
func (_m *WarehouseRepositoryMock) UpdateWarehouse(_a0 int, _a1 *models.WarehouseUpdateDocResponse) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateWarehouse")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, *models.WarehouseUpdateDocResponse) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewWarehouseRepository creates a new instance of WarehouseRepositoryMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWarehouseRepositoryMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *WarehouseRepositoryMock {
	mock := &WarehouseRepositoryMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
