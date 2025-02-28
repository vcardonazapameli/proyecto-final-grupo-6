// Code generated by mockery v2.52.3. DO NOT EDIT.

package buyer

import (
	models "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// BuyerRepositoryMock is an autogenerated mock type for the BuyerRepository type
type BuyerRepositoryMock struct {
	mock.Mock
}

// CreateBuyer provides a mock function with given fields: _a0
func (_m *BuyerRepositoryMock) CreateBuyer(_a0 *models.BuyerDocResponse) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for CreateBuyer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.BuyerDocResponse) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBuyer provides a mock function with given fields: buyerId
func (_m *BuyerRepositoryMock) DeleteBuyer(buyerId int) error {
	ret := _m.Called(buyerId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteBuyer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(buyerId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with no fields
func (_m *BuyerRepositoryMock) GetAll() ([]models.BuyerDocResponse, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []models.BuyerDocResponse
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.BuyerDocResponse, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.BuyerDocResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.BuyerDocResponse)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *BuyerRepositoryMock) GetById(id int) (*models.BuyerDocResponse, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *models.BuyerDocResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*models.BuyerDocResponse, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *models.BuyerDocResponse); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.BuyerDocResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPurchasesReports provides a mock function with given fields: cardNumberId
func (_m *BuyerRepositoryMock) GetPurchasesReports(cardNumberId int) ([]models.PurchaseOrderReport, error) {
	ret := _m.Called(cardNumberId)

	if len(ret) == 0 {
		panic("no return value specified for GetPurchasesReports")
	}

	var r0 []models.PurchaseOrderReport
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]models.PurchaseOrderReport, error)); ok {
		return rf(cardNumberId)
	}
	if rf, ok := ret.Get(0).(func(int) []models.PurchaseOrderReport); ok {
		r0 = rf(cardNumberId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.PurchaseOrderReport)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(cardNumberId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBuyer provides a mock function with given fields: id, _a1
func (_m *BuyerRepositoryMock) UpdateBuyer(id int, _a1 *models.BuyerDocRequest) error {
	ret := _m.Called(id, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateBuyer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, *models.BuyerDocRequest) error); ok {
		r0 = rf(id, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateCardNumberId provides a mock function with given fields: cardNumber
func (_m *BuyerRepositoryMock) ValidateCardNumberId(cardNumber int) bool {
	ret := _m.Called(cardNumber)

	if len(ret) == 0 {
		panic("no return value specified for ValidateCardNumberId")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(cardNumber)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ValidateCardNumberIdToUpdate provides a mock function with given fields: cardNumber, id
func (_m *BuyerRepositoryMock) ValidateCardNumberIdToUpdate(cardNumber int, id int) bool {
	ret := _m.Called(cardNumber, id)

	if len(ret) == 0 {
		panic("no return value specified for ValidateCardNumberIdToUpdate")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(int, int) bool); ok {
		r0 = rf(cardNumber, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ValidateIfExistsById provides a mock function with given fields: id
func (_m *BuyerRepositoryMock) ValidateIfExistsById(id int) bool {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for ValidateIfExistsById")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewBuyerRepositoryMock creates a new instance of BuyerRepositoryMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBuyerRepositoryMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *BuyerRepositoryMock {
	mock := &BuyerRepositoryMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
