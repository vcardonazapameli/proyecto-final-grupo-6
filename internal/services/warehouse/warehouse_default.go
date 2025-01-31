package warehouse

import (
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/warehouse"
	errorsCustom "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	validators "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/validators"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func NewWarehouseDefault(rp repository.WarehouseRepository) WarehouseService {
	return &WarehouseDefault{rp: rp}
}

type WarehouseDefault struct {
	rp repository.WarehouseRepository
}

func (s *WarehouseDefault) GetAll() (map[int]models.Warehouse, error) {
	data, err := s.rp.GetAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *WarehouseDefault) GetById(idWarehouse int) (models.Warehouse, error) {
	warehouse, err := s.rp.GetById(idWarehouse)
	if err != nil {
		return models.Warehouse{}, err
	}
	return warehouse, nil

}

func (s *WarehouseDefault) CreateWarehouse(warehouse models.Warehouse) (models.Warehouse, error) {
	if err := validators.ValidateFieldsWarehouse(warehouse); err != nil {
		return models.Warehouse{}, err
	}
	return s.rp.CreateWarehouse(warehouse)
}

func (s *WarehouseDefault) DeleteWarehouse(idWarehouse int) error {
	return s.rp.DeleteWarehouse(idWarehouse)
}

func (s *WarehouseDefault) UpdateWarehouse(id int, warehouseData models.WarehouseDocUpdate) (models.Warehouse, error) {
	existingWarehouse, err := s.rp.GetById(id)
	if err != nil {
		return models.Warehouse{}, errorsCustom.ErrorNotFound
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
}
