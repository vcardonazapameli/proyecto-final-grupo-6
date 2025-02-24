/*
package warehouse_test

import (

	"testing"

	mockRepo "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/warehouse"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/warehouse"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/stretchr/testify/assert"

)

	func TestGetAllWarehouses(t *testing.T) {
		mockRepo := new(mockRepo.WarehouseRepositoryMock)
		service := warehouse.NewWarehouseService(mockRepo)

		expectedWarehouses := []models.WarehouseDocResponse{{ID: 1}, {ID: 2}, {ID: 3}}
		mockRepo.On("GetAll").Return([]models.WarehouseDocResponse{{ID: 1}, {ID: 2}, {ID: 3}}, nil)

		warehouses, err := service.GetAll()
		assert.NoError(t, err)
		assert.Equal(t, expectedWarehouses, warehouses)
		mockRepo.AssertExpectations(t)
	}

	func TestGetWarehouseById_NotFound(t *testing.T) {
		mockRepo := new(mockRepo.WarehouseRepositoryMock)
		service := warehouse.NewWarehouseService(mockRepo)

		mockRepo.On("GetById", 1).Return(nil, nil)

		warehouse, err := service.GetById(1)
		assert.Error(t, err)
		assert.Nil(t, warehouse)
		mockRepo.AssertExpectations(t)
	}

/*

	func TestCreateWarehouse_Success(t *testing.T) {
		mockRepo := new(mockRepo.WarehouseRepositoryMock)
		service := warehouse.NewWarehouseService(mockRepo)

		warehouseReq := models.WarehouseDocRequest{Warehouse_code: "W123"}
		expectedWarehouse := models.WarehouseDocResponse{ID: 1, Warehouse_code: "W123"}

		mockRepo.On("ExistInDbWarehouseCode", "W123").Return(false, nil)
		mockRepo.On("CreateWarehouse", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*models.WarehouseDocResponse)
			arg.ID = expectedWarehouse.ID
		})

		warehouse, err := service.CreateWarehouse(warehouseReq)
		assert.NoError(t, err)
		assert.Equal(t, expectedWarehouse.ID, warehouse.ID)
		mockRepo.AssertExpectations(t)
	}

	func TestDeleteWarehouse_NotFound(t *testing.T) {
		mockRepo := new(mockRepo.WarehouseRepositoryMock)
		service := warehouse.NewWarehouseService(mockRepo)

		mockRepo.On("GetById", 1).Return(nil, nil)

		err := service.DeleteWarehouse(1)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	}
*/
package warehouse_test

import (
	"testing"

	mockRepo "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/warehouse"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/warehouse"
	errorCustom "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/stretchr/testify/assert"
)

/*
func TestCreateWarehouse_OK(t *testing.T) {
	mockRepo := new(mockRepo.WarehouseRepositoryMock)
	service := warehouse.NewWarehouseService(mockRepo)
	var idLocality *uint64
	*idLocality = 1
	warehouseReq := models.WarehouseDocRequest{Warehouse_code: "WH123", Address: "Calle Falsa 123", Telephone: "1122334455", Minimun_capacity: 200, Minimun_temperature: -12.2, Locality_id: 1}
	warehouseRes := models.WarehouseDocResponse{ID: 0, Warehouse_code: "WH123", Address: "Calle Falsa 123", Telephone: "1122334455", Minimun_capacity: 200, Minimun_temperature: -12.2, Locality_id: 1}

	mockRepo.On("ExistInDbWarehouseCode", "WH123").Return(false, nil)
	mockRepo.On("CreateWarehouse", mock.Anything).Return(nil)

	result, err := service.CreateWarehouse(warehouseReq)

	assert.NoError(t, err)
	assert.Equal(t, warehouseRes, result)
	mockRepo.AssertExpectations(t)
}
*/

func TestCreateWarehouse_Conflict(t *testing.T) {
	mockRepo := new(mockRepo.WarehouseRepositoryMock)
	service := warehouse.NewWarehouseService(mockRepo)

	warehouseReq := models.WarehouseDocRequest{Warehouse_code: "WH123", Address: "Calle Falsa 123", Telephone: "1122334455", Minimun_capacity: 200, Minimun_temperature: -12.2}

	mockRepo.On("ExistInDbWarehouseCode", "WH123").Return(true, nil)

	result, err := service.CreateWarehouse(warehouseReq)

	assert.ErrorIs(t, err, errorCustom.ErrorConflict)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestFindAllWarehouses(t *testing.T) {
	mockRepo := new(mockRepo.WarehouseRepositoryMock)
	service := warehouse.NewWarehouseService(mockRepo)

	warehouses := []models.WarehouseDocResponse{{Warehouse_code: "WH123"}, {Warehouse_code: "WH456"}}

	mockRepo.On("GetAll").Return(warehouses, nil)

	result, err := service.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, warehouses, result)
	//assert.Len(t, result, len(warehouses))
	mockRepo.AssertExpectations(t)
}

/*
func TestFindById_NonExistent(t *testing.T) {
	mockRepo := new(mockRepo.WarehouseRepositoryMock)
	service := warehouse.NewWarehouseService(mockRepo)

	mockRepo.On("GetById", 1).Return(nil, nil)

	result, err := service.GetById(1)

	assert.ErrorIs(t, err, errorCustom.ErrorNotFound)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestFindById_Existent(t *testing.T) {
	mockRepo := new(mockRepo.WarehouseRepositoryMock)
	service := warehouse.NewWarehouseService(mockRepo)

	warehouseRes := &models.WarehouseDocResponse{Warehouse_code: "WH123"}
	mockRepo.On("GetById", 1).Return(warehouseRes, nil)

	result, err := service.GetById(1)

	assert.NoError(t, err)
	assert.Equal(t, warehouseRes, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateWarehouse_Existent(t *testing.T) {
	mockRepo := new(mockRepo.WarehouseRepositoryMock)
	service := warehouse.NewWarehouseService(mockRepo)

	warehouseReq := models.WarehouseUpdateDocRequest{Warehouse_code: new(string)}
	warehouseRes := &models.WarehouseUpdateDocResponse{}

	mockRepo.On("GetById", 1).Return(&models.WarehouseDocResponse{}, nil)
	mockRepo.On("UpdateWarehouse", 1, mock.Anything).Return(nil)

	result, err := service.UpdateWarehouse(1, warehouseReq)

	assert.NoError(t, err)
	assert.Equal(t, warehouseRes, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateWarehouse_NonExistent(t *testing.T) {
	mockRepo := new(mockRepo.WarehouseRepositoryMock)
	service := warehouse.NewWarehouseService(mockRepo)

	warehouseReq := models.WarehouseUpdateDocRequest{}

	mockRepo.On("GetById", 1).Return(nil, nil)

	result, err := service.UpdateWarehouse(1, warehouseReq)

	assert.ErrorIs(t, err, errorCustom.ErrorNotFound)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestDeleteWarehouse_NonExistent(t *testing.T) {
	mockRepo := new(mockRepo.WarehouseRepositoryMock)
	service := warehouse.NewWarehouseService(mockRepo)

	mockRepo.On("GetById", 1).Return(nil, nil)

	err := service.DeleteWarehouse(1)

	assert.ErrorIs(t, err, errorCustom.ErrorNotFound)
	mockRepo.AssertExpectations(t)
}

func TestDeleteWarehouse_OK(t *testing.T) {
	mockRepo := new(mockRepo.WarehouseRepositoryMock)
	service := warehouse.NewWarehouseService(mockRepo)

	mockRepo.On("GetById", 1).Return(&models.WarehouseDocResponse{}, nil)
	mockRepo.On("DeleteWarehouse", 1).Return(nil)

	err := service.DeleteWarehouse(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
*/
