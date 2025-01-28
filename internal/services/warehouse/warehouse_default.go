package warehouse

import (
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/warehouse"
)

func NewWarehouseDefault(rp repository.WarehouseRepository) WarehouseService {
	return &WarehouseDefault{rp: rp}
}

type WarehouseDefault struct {
	rp repository.WarehouseRepository
}
