package warehouse

import (
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
