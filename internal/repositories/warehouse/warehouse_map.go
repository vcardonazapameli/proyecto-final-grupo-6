package warehouse

import (
	errorsCustom "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/errors"

	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func NewWarehouseMap(db map[int]models.Warehouse) WarehouseRepository {
	defaultDb := make(map[int]models.Warehouse)
	if db != nil {
		defaultDb = db
	}
	return &WarehouseMap{db: defaultDb}
}

type WarehouseMap struct {
	db map[int]models.Warehouse
}

func (r *WarehouseMap) GetAll() (v map[int]models.Warehouse, err error) {
	v = make(map[int]models.Warehouse)
	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

func (r *WarehouseMap) GetById(idWarehouse int) (models.Warehouse, error) {
	warehouse, exists := r.db[idWarehouse]
	if !exists {
		return models.Warehouse{}, errorsCustom.ErrorNotFound
	}
	return warehouse, nil
}

func (r *WarehouseMap) CreateWarehouse(warehouse models.Warehouse) (models.Warehouse, error) {
	warehouses, err := r.GetAll()
	if err != nil {
		return models.Warehouse{}, err
	}
	for _, warehouseData := range warehouses {
		if warehouseData.Warehouse_code == warehouse.Warehouse_code {
			return models.Warehouse{}, errorsCustom.ErrorWarehouseCoreRepeat
		}
	}
	warehouse.Id = len(warehouses) + 1
	warehouses[warehouse.Id] = warehouse
	return warehouse, nil
}

func (r *WarehouseMap) DeleteWarehouse(idWarehouse int) error {
	warehouses, err := r.GetAll()
	if err != nil {
		return err
	}
	delete(warehouses, idWarehouse)
	return nil

}
