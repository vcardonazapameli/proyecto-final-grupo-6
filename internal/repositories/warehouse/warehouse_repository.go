package warehouse

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type WarehouseRepository interface {
	GetAll() ([]models.WarehouseDocResponse, error)
	GetById(int) (*models.WarehouseDocResponse, error)
	CreateWarehouse(*models.WarehouseDocResponse) error
	DeleteWarehouse(int) error
	UpdateWarehouse(int, *models.WarehouseUpdateDocResponse) error
	ExistInDbWarehouseCode(string) (bool, error)
	MatchWarehouseCode(int, string) (bool, error)
}
