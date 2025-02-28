package inbound_order_test

import (
	"testing"

	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	rpEmp "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/employee"
	rpIor "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/inbound_order"
	rpPb "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product_batch"
	rpWh "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/warehouse"
	sv "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/inbound_order"
)

func TestInboundOrderService(t *testing.T) {
	t.Run("create_fail_validation", func(t *testing.T) {
		// Arrange
		repoEmp := new(rpEmp.EmployeeRepositoryMock)
		repoIor := new(rpIor.InboundOrderRepositoryMock)
		repoPb := new(rpPb.ProductBatchRepositoryMock)
		repoWh := new(rpWh.WarehouseRepositoryMock)
		service := sv.NewInboundOrderDefault(repoIor, repoEmp, repoPb, repoWh)

		newOrder := models.RequestInboundOrder{
			OrderDate:      "",
			OrderNumber:    "",
			EmployeeID:     0,
			ProductBatchID: 0,
			WarehouseID:    0,
		}

		// Act
		createdOrder, err := service.Create(newOrder)

		// Assert
		require.Error(t, err)
		assert.Nil(t, createdOrder)
		assert.Equal(t, customErrors.ErrorUnprocessableContent, err)
	})

	t.Run("create_fail_order_number_exists", func(t *testing.T) {
		// Arrange
		repoEmp := new(rpEmp.EmployeeRepositoryMock)
		repoIor := new(rpIor.InboundOrderRepositoryMock)
		repoPb := new(rpPb.ProductBatchRepositoryMock)
		repoWh := new(rpWh.WarehouseRepositoryMock)
		service := sv.NewInboundOrderDefault(repoIor, repoEmp, repoPb, repoWh)

		newOrder := models.RequestInboundOrder{
			OrderDate:      "2025-02-27",
			OrderNumber:    "12345",
			EmployeeID:     1,
			ProductBatchID: 1,
			WarehouseID:    1,
		}

		repoIor.On("ExistOrderNumber", newOrder.OrderNumber).Return(true, nil)

		// Act
		createdOrder, err := service.Create(newOrder)

		// Assert
		require.Error(t, err)
		assert.Nil(t, createdOrder)
		assert.Equal(t, customErrors.ErrorConflict, err)
		repoIor.AssertExpectations(t)
	})

	t.Run("create_fail_employee_not_found", func(t *testing.T) {
		// Arrange
		repoEmp := new(rpEmp.EmployeeRepositoryMock)
		repoIor := new(rpIor.InboundOrderRepositoryMock)
		repoPb := new(rpPb.ProductBatchRepositoryMock)
		repoWh := new(rpWh.WarehouseRepositoryMock)
		service := sv.NewInboundOrderDefault(repoIor, repoEmp, repoPb, repoWh)

		newOrder := models.RequestInboundOrder{
			OrderDate:      "2025-02-27",
			OrderNumber:    "12345",
			EmployeeID:     1,
			ProductBatchID: 1,
			WarehouseID:    1,
		}

		repoIor.On("ExistOrderNumber", newOrder.OrderNumber).Return(false, nil)
		repoEmp.On("GetById", newOrder.EmployeeID).Return(nil, customErrors.ErrorNotFound)

		// Act
		createdOrder, err := service.Create(newOrder)

		// Assert
		require.Error(t, err)
		assert.Nil(t, createdOrder)
		assert.Equal(t, customErrors.ErrorConflict, err)
		repoIor.AssertExpectations(t)
		repoEmp.AssertExpectations(t)
	})

	t.Run("create_fail_product_batch_not_found", func(t *testing.T) {
		// Arrange
		repoEmp := new(rpEmp.EmployeeRepositoryMock)
		repoIor := new(rpIor.InboundOrderRepositoryMock)
		repoPb := new(rpPb.ProductBatchRepositoryMock)
		repoWh := new(rpWh.WarehouseRepositoryMock)
		service := sv.NewInboundOrderDefault(repoIor, repoEmp, repoPb, repoWh)

		newOrder := models.RequestInboundOrder{
			OrderNumber:    "12345",
			OrderDate:      "2025-02-27",
			EmployeeID:     1,
			ProductBatchID: 11231,
			WarehouseID:    1,
		}

		repoIor.On("ExistOrderNumber", newOrder.OrderNumber).Return(false, nil)
		repoEmp.On("GetById", newOrder.EmployeeID).Return(&models.Employee{}, nil)
		repoPb.On("ExistsByID", newOrder.ProductBatchID).Return(false, nil)

		// Act
		createdOrder, err := service.Create(newOrder)

		// Assert
		require.Error(t, err)
		assert.Nil(t, createdOrder)
		assert.Equal(t, customErrors.ErrorConflict, err)
		repoIor.AssertExpectations(t)
		repoEmp.AssertExpectations(t)
		repoPb.AssertExpectations(t)
	})

	t.Run("create_fail_warehouse_not_found", func(t *testing.T) {
		// Arrange
		repoEmp := new(rpEmp.EmployeeRepositoryMock)
		repoIor := new(rpIor.InboundOrderRepositoryMock)
		repoPb := new(rpPb.ProductBatchRepositoryMock)
		repoWh := new(rpWh.WarehouseRepositoryMock)
		service := sv.NewInboundOrderDefault(repoIor, repoEmp, repoPb, repoWh)

		newOrder := models.RequestInboundOrder{
			OrderNumber:    "12345",
			OrderDate:      "2025-02-27",
			EmployeeID:     1,
			ProductBatchID: 1,
			WarehouseID:    9127,
		}

		repoIor.On("ExistOrderNumber", newOrder.OrderNumber).Return(false, nil)
		repoEmp.On("GetById", newOrder.EmployeeID).Return(&models.Employee{}, nil)
		repoPb.On("ExistsByID", newOrder.ProductBatchID).Return(true, nil)
		repoWh.On("GetById", newOrder.WarehouseID).Return(nil, customErrors.ErrorNotFound)

		// Act
		createdOrder, err := service.Create(newOrder)

		// Assert
		require.Error(t, err)
		assert.Nil(t, createdOrder)
		assert.Equal(t, customErrors.ErrorConflict, err)
		repoIor.AssertExpectations(t)
		repoEmp.AssertExpectations(t)
		repoPb.AssertExpectations(t)
		repoWh.AssertExpectations(t)
	})

	t.Run("create_fail_repositoryError", func(t *testing.T) {
		// Arrange
		repoEmp := new(rpEmp.EmployeeRepositoryMock)
		repoIor := new(rpIor.InboundOrderRepositoryMock)
		repoPb := new(rpPb.ProductBatchRepositoryMock)
		repoWh := new(rpWh.WarehouseRepositoryMock)
		service := sv.NewInboundOrderDefault(repoIor, repoEmp, repoPb, repoWh)

		newOrder := models.RequestInboundOrder{
			OrderNumber:    "12345",
			OrderDate:      "2025-02-27",
			EmployeeID:     1,
			ProductBatchID: 1,
			WarehouseID:    1,
		}

		repoIor.On("ExistOrderNumber", newOrder.OrderNumber).Return(false, nil)
		repoEmp.On("GetById", newOrder.EmployeeID).Return(&models.Employee{}, nil)
		repoPb.On("ExistsByID", newOrder.ProductBatchID).Return(true, nil)
		repoWh.On("GetById", newOrder.WarehouseID).Return(&models.WarehouseDocResponse{}, nil)
		repoIor.On("Create", mock.AnythingOfType("models.InboundOrder")).Return(nil, customErrors.ErrorInternalServerError)

		// Act
		createdOrder, err := service.Create(newOrder)

		// Assert
		require.Error(t, err)
		assert.Nil(t, createdOrder)
		repoIor.AssertExpectations(t)
		repoEmp.AssertExpectations(t)
		repoPb.AssertExpectations(t)
		repoWh.AssertExpectations(t)
	})
	t.Run("create_success", func(t *testing.T) {
		// Arrange
		repoEmp := new(rpEmp.EmployeeRepositoryMock)
		repoIor := new(rpIor.InboundOrderRepositoryMock)
		repoPb := new(rpPb.ProductBatchRepositoryMock)
		repoWh := new(rpWh.WarehouseRepositoryMock)
		service := sv.NewInboundOrderDefault(repoIor, repoEmp, repoPb, repoWh)

		newOrder := models.RequestInboundOrder{
			OrderNumber:    "12345",
			OrderDate:      "2025-02-27",
			EmployeeID:     1,
			ProductBatchID: 1,
			WarehouseID:    1,
		}

		expectedOrder := models.InboundOrder{
			ID:             1,
			OrderNumber:    "12345",
			OrderDate:      "2025-02-27",
			EmployeeID:     1,
			ProductBatchID: 1,
			WarehouseID:    1,
		}

		repoIor.On("ExistOrderNumber", newOrder.OrderNumber).Return(false, nil)
		repoEmp.On("GetById", newOrder.EmployeeID).Return(&models.Employee{}, nil)
		repoPb.On("ExistsByID", newOrder.ProductBatchID).Return(true, nil)
		repoWh.On("GetById", newOrder.WarehouseID).Return(&models.WarehouseDocResponse{}, nil)
		repoIor.On("Create", mock.AnythingOfType("models.InboundOrder")).Return(&expectedOrder, nil)

		// Act
		createdOrder, err := service.Create(newOrder)

		// Assert
		require.NoError(t, err)
		assert.NotNil(t, createdOrder)
		assert.Equal(t, &expectedOrder, createdOrder)
		repoIor.AssertExpectations(t)
		repoEmp.AssertExpectations(t)
		repoPb.AssertExpectations(t)
		repoWh.AssertExpectations(t)
	})
}
