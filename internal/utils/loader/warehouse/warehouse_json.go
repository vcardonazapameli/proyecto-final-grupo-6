package warehouse

import (
	"encoding/json"
	"os"

	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func NewWarehouseJSONFile(path string) *WarehouseJSONFile {
	return &WarehouseJSONFile{
		path: path,
	}
}

type WarehouseJSONFile struct {
	path string
}

func (l *WarehouseJSONFile) Load() (v map[int]models.Warehouse, err error) {
	// open file
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	// decode file
	var warehouseJSON []models.WarehouseDocResponse
	err = json.NewDecoder(file).Decode(&warehouseJSON)
	if err != nil {
		return
	}

	// serialize warehouses
	v = make(map[int]models.Warehouse)
	for _, vh := range warehouseJSON {
		v[vh.ID] = models.Warehouse{
			Id: vh.ID,
			WarehouseAttributes: models.WarehouseAttributes{
				Warehouse_code:      vh.Warehouse_code,
				Address:             vh.Address,
				Telephone:           vh.Telephone,
				Minimun_capacity:    vh.Minimun_capacity,
				Minimun_temperature: vh.Minimun_temperature,
				Locality_id:         *vh.Locality_id,
			},
		}
	}
	return
}
