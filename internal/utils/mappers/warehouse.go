package mappers

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

func WarehouseUpdateDocRequestToWarehouseUpdateDocResponse(warehouse models.WarehouseUpdateDocRequest) *models.WarehouseUpdateDocResponse {
	return &models.WarehouseUpdateDocResponse{
		Warehouse_code:      warehouse.Warehouse_code,
		Address:             warehouse.Address,
		Telephone:           warehouse.Telephone,
		Minimun_capacity:    warehouse.Minimun_capacity,
		Minimun_temperature: warehouse.Minimun_temperature,
		Locality_id:         warehouse.Locality_id,
	}
}

func WarehouseDocRequestToWarehouseDocResponse(WarehouseDocRequest models.WarehouseDocRequest) models.WarehouseDocResponse {
	var localityID *uint64
	if WarehouseDocRequest.Locality_id != nil {
		localityID = WarehouseDocRequest.Locality_id
	}
	return models.WarehouseDocResponse{
		Warehouse_code:      WarehouseDocRequest.Warehouse_code,
		Address:             WarehouseDocRequest.Address,
		Telephone:           WarehouseDocRequest.Telephone,
		Minimun_capacity:    WarehouseDocRequest.Minimun_capacity,
		Minimun_temperature: WarehouseDocRequest.Minimun_temperature,
		Locality_id:         localityID,
	}
}
