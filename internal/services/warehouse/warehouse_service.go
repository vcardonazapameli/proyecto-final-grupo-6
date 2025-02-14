package warehouse

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type WarehouseService interface {
	GetAll() ([]models.WarehouseDocResponse, error)
	GetById(int) (*models.WarehouseDocResponse, error)
	CreateWarehouse(models.WarehouseDocRequest) (*models.WarehouseDocResponse, error)
	DeleteWarehouse(int) error
	UpdateWarehouse(int, models.WarehouseUpdateDocRequest) (*models.WarehouseDocResponse, error)
}
