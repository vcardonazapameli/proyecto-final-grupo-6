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
	if len(r.db) == 1 && r.db[1].Id == 0 {
		return nil, errorsCustom.ErrorNotFound
	}
	v = make(map[int]models.Warehouse)
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
	for _, warehouseData := range r.db {
		if warehouseData.Warehouse_code == warehouse.Warehouse_code {
			return models.Warehouse{}, errorsCustom.ErrorWarehouseCoreRepeat
		}
	}

	newID := len(r.db) + 1
	warehouse.Id = newID

	r.db[newID] = warehouse
	return warehouse, nil
}

func (r *WarehouseMap) DeleteWarehouse(idWarehouse int) error {
	_, exists := r.db[idWarehouse]
	if !exists {
		return errorsCustom.ErrorNotFound
	}
	delete(r.db, idWarehouse)
	return nil
}

func (r *WarehouseMap) UpdateWarehouse(id int, warehouse models.Warehouse) (models.Warehouse, error) {
	existingWarehouse, exists := r.db[id]
	if !exists {
		return models.Warehouse{}, errorsCustom.ErrorNotFound
	}

	for key, wh := range r.db {
		if key != id && wh.Warehouse_code == warehouse.Warehouse_code {
			return models.Warehouse{}, errorsCustom.ErrorWarehouseCoreRepeat
		}
	}

	existingWarehouse.Warehouse_code = warehouse.Warehouse_code
	existingWarehouse.Address = warehouse.Address
	existingWarehouse.Telephone = warehouse.Telephone
	existingWarehouse.Minimun_capacity = warehouse.Minimun_capacity
	existingWarehouse.Minimun_temperature = warehouse.Minimun_temperature
	existingWarehouse.Locality_id = warehouse.Locality_id

	r.db[id] = existingWarehouse

	return existingWarehouse, nil
}
