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
	warehouse, err := s.rp.GetById(idWarehouse)
	if err != nil {
		return nil, err
	}
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
		return nil, err
	}
	return &warehouse, nil
}

func (s *warehouseService) DeleteWarehouse(id int) error {
	warehouse, err := s.rp.GetById(id)
	if err != nil {
		return err
	}
	if warehouse == nil {
		return errorCustom.ErrorNotFound
	}
	if err := s.rp.DeleteWarehouse(warehouse.ID); err != nil {
		return err
	}
	return nil
}

func (s *warehouseService) UpdateWarehouse(id int, warehouseDocRequest models.WarehouseUpdateDocRequest) (*models.WarehouseUpdateDocResponse, error) {
	warehouse, err := s.rp.GetById(id)
	if err != nil {
		return nil, err
	}
	if warehouse == nil {
		return nil, errorCustom.ErrorNotFound
	}
	//warehouseUpdate := validators.UpdateEntity(warehouseDocRequest, warehouse)
	if warehouseDocRequest.Warehouse_code != nil {
		newCode := *warehouseDocRequest.Warehouse_code
		warehouseCodeExists, err := s.rp.MatchWarehouseCode(warehouse.ID, newCode)
		if err != nil {
			return nil, err
		}
		if warehouseCodeExists {
			return nil, errorCustom.ErrorConflict
		}
	}
	warehouseDocUpdate := mappers.WarehouseUpdateDocRequestToWarehouseUpdateDocResponse(warehouseDocRequest)
	if err := s.rp.UpdateWarehouse(id, warehouseDocUpdate); err != nil {
		return nil, err
	}
	warehouseDocUpdate.ID = &id
	return warehouseDocUpdate, nil
}
