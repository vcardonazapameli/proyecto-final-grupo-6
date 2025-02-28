package warehouse_test

import (
	"errors"
	"testing"

	mockRepo "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/warehouse"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/warehouse"
	errorCustom "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestWarehouseServiceCreate(t *testing.T) {
	t.Run("case create_ok", func(t *testing.T) {
		mockRepo := new(mockRepo.WarehouseRepositoryMock)
		service := warehouse.NewWarehouseService(mockRepo)
		localityID := uint64(3)

		warehouseReq := models.WarehouseDocRequest{Warehouse_code: "WH123", Address: "Calle Falsa 123", Telephone: "1122334455", Minimun_capacity: 200, Minimun_temperature: -12.2, Locality_id: &localityID}
		warehouseRes := models.WarehouseDocResponse{Warehouse_code: "WH123", Address: "Calle Falsa 123", Telephone: "1122334455", Minimun_capacity: 200, Minimun_temperature: -12.2, Locality_id: &localityID}

		mockRepo.On("ExistInDbWarehouseCode", "WH123").Return(false, nil)
		mockRepo.On("CreateWarehouse", mock.Anything).Return(nil)

		result, err := service.CreateWarehouse(warehouseReq)

		assert.NoError(t, err)
		assert.Equal(t, warehouseRes, *result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("case create_conflict", func(t *testing.T) {
		mockRepo := new(mockRepo.WarehouseRepositoryMock)
		service := warehouse.NewWarehouseService(mockRepo)

		warehouseReq := models.WarehouseDocRequest{Warehouse_code: "WH123", Address: "Calle Falsa 123", Telephone: "1122334455", Minimun_capacity: 200, Minimun_temperature: -12.2}

		mockRepo.On("ExistInDbWarehouseCode", "WH123").Return(true, nil)

		result, err := service.CreateWarehouse(warehouseReq)

		assert.ErrorIs(t, err, errorCustom.ErrorConflict)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("case create_err_warehouse_code", func(t *testing.T) {
		mockRepo := new(mockRepo.WarehouseRepositoryMock)
		service := warehouse.NewWarehouseService(mockRepo)

		warehouseReq := models.WarehouseDocRequest{Warehouse_code: "WH123", Address: "Calle Falsa 123", Telephone: "1122334455", Minimun_capacity: 200, Minimun_temperature: -12.2}

		mockRepo.On("ExistInDbWarehouseCode", "WH123").Return(false, errors.New("Database error"))
		result, err := service.CreateWarehouse(warehouseReq)

		assert.Nil(t, result)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("case create_err_create_warehouse", func(t *testing.T) {
		mockRepo := new(mockRepo.WarehouseRepositoryMock)
		service := warehouse.NewWarehouseService(mockRepo)

		warehouseReq := models.WarehouseDocRequest{
			Warehouse_code:      "WH123",
			Address:             "Calle Falsa 123",
			Telephone:           "1122334455",
			Minimun_capacity:    200,
			Minimun_temperature: -12.2,
		}

		warehouseRes := mappers.WarehouseDocRequestToWarehouseDocResponse(warehouseReq)

		mockRepo.On("ExistInDbWarehouseCode", "WH123").Return(false, nil).Once()
		mockRepo.On("CreateWarehouse", &warehouseRes).Return(errors.New("Database error")).Once()

		result, err := service.CreateWarehouse(warehouseReq)

		assert.Nil(t, result)
		assert.Error(t, err)
		assert.Equal(t, "Database error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestWarehouseServiceRead(t *testing.T) {
	t.Run("case find_all", func(t *testing.T) {
		mockRepo := new(mockRepo.WarehouseRepositoryMock)
		service := warehouse.NewWarehouseService(mockRepo)

		warehouses := []models.WarehouseDocResponse{
			{Warehouse_code: "WH123", Address: "Calle Falsa 123", Telephone: "1122334455", Minimun_capacity: 200, Minimun_temperature: -12.2},
			{Warehouse_code: "WH124", Address: "Rivadavia 221", Telephone: "1551555555", Minimun_capacity: 160, Minimun_temperature: -2.3},
		}

		mockRepo.On("GetAll").Return(warehouses, nil)

		result, err := service.GetAll()

		assert.NoError(t, err)
		assert.Equal(t, warehouses, result)
		assert.Len(t, result, len(warehouses))
		mockRepo.AssertExpectations(t)
	})

	t.Run("case find_by_id_non_existent", func(t *testing.T) {
		mockRepo := new(mockRepo.WarehouseRepositoryMock)
		service := warehouse.NewWarehouseService(mockRepo)

		mockRepo.On("GetById", 1).Return(nil, nil)

		result, err := service.GetById(1)

		assert.ErrorIs(t, err, errorCustom.ErrorNotFound)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("case find_by_id_existent", func(t *testing.T) {
		mockRepo := new(mockRepo.WarehouseRepositoryMock)
		service := warehouse.NewWarehouseService(mockRepo)

		warehouseRes := &models.WarehouseDocResponse{Warehouse_code: "WH123", Address: "Calle Falsa 123", Telephone: "1122334455", Minimun_capacity: 200, Minimun_temperature: -12.2}
		mockRepo.On("GetById", 1).Return(warehouseRes, nil)

		result, err := service.GetById(1)

		assert.NoError(t, err)
		assert.Equal(t, warehouseRes, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("case find_all_err_repo", func(t *testing.T) {
		mockRepo := new(mockRepo.WarehouseRepositoryMock)
		service := warehouse.NewWarehouseService(mockRepo)

		var warehouses []models.WarehouseDocResponse
		errorResponse := errors.New("Database error")
		mockRepo.On("GetAll").Return(warehouses, errorResponse)

		warehousesRes, err := service.GetAll()

		assert.Equal(t, warehouses, warehousesRes)
		assert.Equal(t, errorResponse, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("case find_by_id_err_repo", func(t *testing.T) {
		mockRepo := new(mockRepo.WarehouseRepositoryMock)
		service := warehouse.NewWarehouseService(mockRepo)

		warehouseExpected := &models.WarehouseDocResponse{}
		errorResponse := errors.New("Database error")
		mockRepo.On("GetById", 1).Return(warehouseExpected, errorResponse)

		warehouseRes, err := service.GetById(1)

		assert.Nil(t, warehouseRes)
		assert.Equal(t, errorResponse, err)
		mockRepo.AssertExpectations(t)
	})

}

func TestWarehouseServiceUpdate(t *testing.T) {
	t.Run("case update_existent", func(t *testing.T) {
		mockRepo := new(mockRepo.WarehouseRepositoryMock)
		service := warehouse.NewWarehouseService(mockRepo)

		warehouse_code := "WH123"
		address := "Calle Falsa 123"
		telephone := "1155000000"
		minimun_capacity := uint64(200)
		minimun_temperature := -12.2
		localityID := uint64(3)

		warehouseDataID := models.WarehouseDocResponse{
			ID:                  1,
			Warehouse_code:      "WH123",
			Address:             "Calle Falsa 123",
			Telephone:           "1122334455",
			Minimun_capacity:    200,
			Minimun_temperature: -12.2,
			Locality_id:         &localityID,
		}

		warehouseReq := models.WarehouseUpdateDocRequest{
			Warehouse_code:      &warehouse_code,
			Address:             &address,
			Telephone:           &telephone,
			Minimun_capacity:    &minimun_capacity,
			Minimun_temperature: &minimun_temperature,
			Locality_id:         &localityID,
		}

		warehouseRes := models.WarehouseUpdateDocResponse{
			ID:                  &warehouseDataID.ID,
			Warehouse_code:      &warehouseDataID.Warehouse_code,
			Address:             &warehouseDataID.Address,
			Telephone:           &telephone,
			Minimun_capacity:    &warehouseDataID.Minimun_capacity,
			Minimun_temperature: &warehouseDataID.Minimun_temperature,
			Locality_id:         &localityID,
		}

		mockRepo.On("GetById", 1).Return(&warehouseDataID, nil)
		mockRepo.On("MatchWarehouseCode", 1, warehouse_code).Return(false, nil)
		mockRepo.On("UpdateWarehouse", 1, mock.Anything).Return(nil)

		result, err := service.UpdateWarehouse(1, warehouseReq)
		warehouseResService := models.WarehouseUpdateDocResponse{
			ID:                  result.ID,
			Warehouse_code:      result.Warehouse_code,
			Address:             result.Address,
			Telephone:           result.Telephone,
			Minimun_capacity:    result.Minimun_capacity,
			Minimun_temperature: result.Minimun_temperature,
			Locality_id:         result.Locality_id,
		}

		assert.NoError(t, err)
		assert.Equal(t, warehouseRes, warehouseResService)
		mockRepo.AssertExpectations(t)
	})

	t.Run("case update_non_existent", func(t *testing.T) {
		mockRepo := new(mockRepo.WarehouseRepositoryMock)
		service := warehouse.NewWarehouseService(mockRepo)

		warehouseReq := models.WarehouseUpdateDocRequest{}

		mockRepo.On("GetById", 1).Return(nil, nil)

		result, err := service.UpdateWarehouse(1, warehouseReq)

		assert.ErrorIs(t, err, errorCustom.ErrorNotFound)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("case update_err_get_by_id", func(t *testing.T) {
		mockRepo := new(mockRepo.WarehouseRepositoryMock)
		service := warehouse.NewWarehouseService(mockRepo)

		warehouse_code := "WH123"
		address := "Calle Falsa 123"
		telephone := "1155000000"
		minimun_capacity := uint64(200)
		minimun_temperature := -12.2
		localityID := uint64(3)

		warehouseReq := models.WarehouseUpdateDocRequest{
			Warehouse_code:      &warehouse_code,
			Address:             &address,
			Telephone:           &telephone,
			Minimun_capacity:    &minimun_capacity,
			Minimun_temperature: &minimun_temperature,
			Locality_id:         &localityID,
		}

		mockRepo.On("GetById", 1).Return(nil, errors.New("Database error"))

		result, err := service.UpdateWarehouse(1, warehouseReq)

		assert.Nil(t, result)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

}

func TestWarehouseServiceDelete(t *testing.T) {
	t.Run("case delete_non_existent", func(t *testing.T) {
		mockRepo := new(mockRepo.WarehouseRepositoryMock)
		service := warehouse.NewWarehouseService(mockRepo)

		mockRepo.On("GetById", 1).Return(nil, nil)

		err := service.DeleteWarehouse(1)

		assert.ErrorIs(t, err, errorCustom.ErrorNotFound)
		mockRepo.AssertExpectations(t)
	})

	t.Run("case delete_ok", func(t *testing.T) {
		mockRepo := new(mockRepo.WarehouseRepositoryMock)
		service := warehouse.NewWarehouseService(mockRepo)

		mockRepo.On("GetById", 1).Return(&models.WarehouseDocResponse{ID: 1, Warehouse_code: "WH123", Address: "Calle Falsa 123", Telephone: "1122334455", Minimun_capacity: 200, Minimun_temperature: -12.2}, nil)
		mockRepo.On("DeleteWarehouse", 1).Return(nil)

		err := service.DeleteWarehouse(1)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("case delete_err_repo", func(t *testing.T) {
		mockRepo := new(mockRepo.WarehouseRepositoryMock)
		service := warehouse.NewWarehouseService(mockRepo)

		mockRepo.On("GetById", 1).Return(nil, errors.New("Database error"))

		err := service.DeleteWarehouse(1)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("case delete_err_serv", func(t *testing.T) {
		mockRepo := new(mockRepo.WarehouseRepositoryMock)
		service := warehouse.NewWarehouseService(mockRepo)

		warehouse := models.WarehouseDocResponse{ID: 1, Warehouse_code: "WH123", Address: "Calle Falsa 123", Telephone: "1122334455", Minimun_capacity: 200, Minimun_temperature: -12.2}
		mockRepo.On("GetById", 1).Return(&warehouse, nil)
		mockRepo.On("DeleteWarehouse", 1).Return(errors.New("Database error"))

		err := service.DeleteWarehouse(1)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}
