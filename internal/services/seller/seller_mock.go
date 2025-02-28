// Code generated by mockery v2.52.3. DO NOT EDIT.

package seller

import (
	models "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// SellerServiceMock is an autogenerated mock type for the SellerService type
type SellerServiceMock struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *SellerServiceMock) Create(_a0 models.SellerDoc) (models.SellerDoc, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 models.SellerDoc
	var r1 error
	if rf, ok := ret.Get(0).(func(models.SellerDoc) (models.SellerDoc, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(models.SellerDoc) models.SellerDoc); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.SellerDoc)
	}

	if rf, ok := ret.Get(1).(func(models.SellerDoc) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *SellerServiceMock) Delete(id int) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with no fields
func (_m *SellerServiceMock) GetAll() (map[int]models.SellerDoc, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 map[int]models.SellerDoc
	var r1 error
	if rf, ok := ret.Get(0).(func() (map[int]models.SellerDoc, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() map[int]models.SellerDoc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[int]models.SellerDoc)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *SellerServiceMock) GetByID(id int) (models.SellerDoc, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 models.SellerDoc
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (models.SellerDoc, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) models.SellerDoc); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.SellerDoc)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, cid, companyName, address, telephone, localityId
func (_m *SellerServiceMock) Update(id int, cid *int, companyName *string, address *string, telephone *string, localityId *int) (models.SellerDoc, error) {
	ret := _m.Called(id, cid, companyName, address, telephone, localityId)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 models.SellerDoc
	var r1 error
	if rf, ok := ret.Get(0).(func(int, *int, *string, *string, *string, *int) (models.SellerDoc, error)); ok {
		return rf(id, cid, companyName, address, telephone, localityId)
	}
	if rf, ok := ret.Get(0).(func(int, *int, *string, *string, *string, *int) models.SellerDoc); ok {
		r0 = rf(id, cid, companyName, address, telephone, localityId)
	} else {
		r0 = ret.Get(0).(models.SellerDoc)
	}

	if rf, ok := ret.Get(1).(func(int, *int, *string, *string, *string, *int) error); ok {
		r1 = rf(id, cid, companyName, address, telephone, localityId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSellerService creates a new instance of SellerServiceMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSellerService(t interface {
	mock.TestingT
	Cleanup(func())
}) *SellerServiceMock {
	mock := &SellerServiceMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
