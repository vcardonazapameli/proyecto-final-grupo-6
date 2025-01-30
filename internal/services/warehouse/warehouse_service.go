package warehouse

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type WarehouseService interface {
	GetAll() (map[int]models.Warehouse, error)
	GetById(int) (models.Warehouse, error)
	CreateWarehouse(models.Warehouse) (models.Warehouse, error)
	DeleteWarehouse(int) error
	UpdateWarehouse(int, models.WarehouseDocUpdate) (models.Warehouse, error)
}
