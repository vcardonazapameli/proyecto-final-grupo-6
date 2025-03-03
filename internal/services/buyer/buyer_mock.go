// Code generated by mockery v2.52.3. DO NOT EDIT.

package buyer

import (
	models "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// BuyerServiceMock is an autogenerated mock type for the BuyerService type
type BuyerServiceMock struct {
	mock.Mock
}

// CreateBuyer provides a mock function with given fields: _a0
func (_m *BuyerServiceMock) CreateBuyer(_a0 models.BuyerDocRequest) (*models.BuyerDocResponse, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for CreateBuyer")
	}

	var r0 *models.BuyerDocResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(models.BuyerDocRequest) (*models.BuyerDocResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(models.BuyerDocRequest) *models.BuyerDocResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.BuyerDocResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(models.BuyerDocRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBuyer provides a mock function with given fields: buyerId
func (_m *BuyerServiceMock) DeleteBuyer(buyerId int) error {
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
func (_m *BuyerServiceMock) GetAll() ([]models.BuyerDocResponse, error) {
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
func (_m *BuyerServiceMock) GetById(id int) (*models.BuyerDocResponse, error) {
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

// GetPurchasesReports provides a mock function with given fields: CardNumberId
func (_m *BuyerServiceMock) GetPurchasesReports(CardNumberId int) ([]models.PurchaseOrderReport, error) {
	ret := _m.Called(CardNumberId)

	if len(ret) == 0 {
		panic("no return value specified for GetPurchasesReports")
	}

	var r0 []models.PurchaseOrderReport
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]models.PurchaseOrderReport, error)); ok {
		return rf(CardNumberId)
	}
	if rf, ok := ret.Get(0).(func(int) []models.PurchaseOrderReport); ok {
		r0 = rf(CardNumberId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.PurchaseOrderReport)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(CardNumberId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBuyer provides a mock function with given fields: id, buyerRequest
func (_m *BuyerServiceMock) UpdateBuyer(id int, buyerRequest models.UpdateBuyerDto) (*models.BuyerDocResponse, error) {
	ret := _m.Called(id, buyerRequest)

	if len(ret) == 0 {
		panic("no return value specified for UpdateBuyer")
	}

	var r0 *models.BuyerDocResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(int, models.UpdateBuyerDto) (*models.BuyerDocResponse, error)); ok {
		return rf(id, buyerRequest)
	}
	if rf, ok := ret.Get(0).(func(int, models.UpdateBuyerDto) *models.BuyerDocResponse); ok {
		r0 = rf(id, buyerRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.BuyerDocResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(int, models.UpdateBuyerDto) error); ok {
		r1 = rf(id, buyerRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBuyerServiceMock creates a new instance of BuyerServiceMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBuyerServiceMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *BuyerServiceMock {
	mock := &BuyerServiceMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
