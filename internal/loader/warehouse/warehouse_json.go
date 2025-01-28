package warehouse

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

func NewWarehouseJSONFile(path string) *WarehouseJSONFile {
	return &WarehouseJSONFile{
		path: path,
	}
}

type WarehouseJSONFile struct {
	path string
}

func (l *WarehouseJSONFile) Load() (v map[int]models.Warehouse, err error) {
	return
}
