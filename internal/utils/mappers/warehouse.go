package mappers

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

func WarehouseDocToWarehouse(warehouseDoc models.WarehouseDoc) models.WarehouseAttributes {
	return models.WarehouseAttributes{
		Warehouse_code:      warehouseDoc.Warehouse_code,
		Address:             warehouseDoc.Address,
		Telephone:           warehouseDoc.Telephone,
		Minimun_capacity:    warehouseDoc.Minimun_capacity,
		Minimun_temperature: warehouseDoc.Minimun_temperature,
		Locality_id:         warehouseDoc.Locality_id,
	}
}

func WarehouseToWarehouseDoc(warehouse models.Warehouse) models.WarehouseDoc {
	return models.WarehouseDoc{
		ID:                  warehouse.Id,
		Warehouse_code:      warehouse.Warehouse_code,
		Address:             warehouse.Address,
		Telephone:           warehouse.Telephone,
		Minimun_capacity:    warehouse.Minimun_capacity,
		Minimun_temperature: warehouse.Minimun_temperature,
		Locality_id:         warehouse.Locality_id,
	}
}
