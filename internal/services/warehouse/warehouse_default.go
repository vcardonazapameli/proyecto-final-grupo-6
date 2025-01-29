package warehouse

import (
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/warehouse"
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
	newWarehouse, err := s.rp.CreateWarehouse(warehouse)
	if err != nil {
		return models.Warehouse{}, err
	}
	return newWarehouse, nil
}
