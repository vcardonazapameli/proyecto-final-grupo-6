package warehouse

import (
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/warehouse"
	errorCustom "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/validators"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func NewWarehouseService(rp repository.WarehouseRepository) WarehouseService {
	return &warehouseService{rp: rp}
}

type warehouseService struct {
	rp repository.WarehouseRepository
}

func (s *warehouseService) GetAll() ([]models.WarehouseDocResponse, error) {
	data, err := s.rp.GetAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *warehouseService) GetById(idWarehouse int) (*models.WarehouseDocResponse, error) {
	warehouse, _ := s.rp.GetById(idWarehouse)
	if warehouse == nil {
		return nil, errorCustom.ErrorNotFound
	}
	return warehouse, nil

}

func (s *warehouseService) CreateWarehouse(warehouseDocRequest models.WarehouseDocRequest) (*models.WarehouseDocResponse, error) {
	if err := validators.ValidateFieldsWarehouseCreate(warehouseDocRequest); err != nil {
		return nil, err
	}
	existInDb, err := s.rp.ExistInDbWarehouseCode(warehouseDocRequest.Warehouse_code)
	if err != nil {
		return nil, err
	}
	if existInDb {
		return nil, errorCustom.ErrorConflict
	}
	warehouse := mappers.WarehouseDocRequestToWarehouseDocResponse(warehouseDocRequest)
	if err := s.rp.CreateWarehouse(&warehouse); err != nil {
		return nil, nil
	}
	return &warehouse, nil
}

func (s *warehouseService) DeleteWarehouse(id int) error {
	warehouse, _ := s.rp.GetById(id)
	if warehouse == nil {
		return errorCustom.ErrorNotFound
	}
	s.rp.DeleteWarehouse(warehouse.ID)
	return nil
}

func (s *warehouseService) UpdateWarehouse(id int, warehouseDocRequest models.WarehouseUpdateDocRequest) (*models.WarehouseDocResponse, error) {
	warehouse, _ := s.rp.GetById(id)
	if warehouse == nil {
		return nil, errorCustom.ErrorNotFound
	}
	warehouseUpdate := validators.UpdateEntity(warehouseDocRequest, warehouse)
	warehouseCodeExists, err := s.rp.MatchWarehouseCode(warehouseUpdate.ID, warehouseUpdate.Warehouse_code)
	if err != nil {
		return nil, err
	}
	if warehouseCodeExists {
		return nil, errorCustom.ErrorConflict
	}

	warehouseDoc := mappers.WarehouseDocResponseToWarehouseDocRequest(warehouseUpdate)
	if errorValidateFields := validators.ValidateFieldsWarehouseUpdate(warehouseDoc); errorValidateFields != nil {
		return nil, errorValidateFields
	}
	if err := s.rp.UpdateWarehouse(id, warehouseUpdate); err != nil {
		return nil, nil
	}
	return warehouseUpdate, nil
	/*
		existingWarehouse, err := s.rp.GetById(id)
		if err != nil {
			return models.WarehouseDocResponse{}, errorsCustom.ErrorNotFound
		}
		updatedWarehouse := validators.UpdateEntity(warehouseData, &existingWarehouse)
		if err := validators.ValidateFieldsUpdate(*updatedWarehouse); err != nil {
			return models.Warehouse{}, err
		}
		if warehouseData.Warehouse_code != nil {
			existingWarehouse.Warehouse_code = *warehouseData.Warehouse_code
		}
		if warehouseData.Address != nil {
			existingWarehouse.Address = *warehouseData.Address
		}
		if warehouseData.Telephone != nil {
			existingWarehouse.Telephone = *warehouseData.Telephone
		}
		if warehouseData.Minimun_capacity != nil {
			existingWarehouse.Minimun_capacity = *warehouseData.Minimun_capacity
		}
		if warehouseData.Minimun_temperature != nil {
			existingWarehouse.Minimun_temperature = *warehouseData.Minimun_temperature
		}
		if warehouseData.Locality_id != nil {
			existingWarehouse.Locality_id = *warehouseData.Locality_id
		}
		return s.rp.UpdateWarehouse(id, existingWarehouse)
	*/
}
