package mappers

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

func WarehouseDocRequestToWarehouseDocResponse(warehouseDocRequest models.WarehouseDocRequest) models.WarehouseDocResponse {
	return models.WarehouseDocResponse{
		Warehouse_code:      warehouseDocRequest.Warehouse_code,
		Address:             warehouseDocRequest.Address,
		Telephone:           warehouseDocRequest.Telephone,
		Minimun_capacity:    warehouseDocRequest.Minimun_capacity,
		Minimun_temperature: warehouseDocRequest.Minimun_temperature,
		Locality_id:         warehouseDocRequest.Locality_id,
	}
}

func WarehouseDocResponseToWarehouseDocRequest(warehouseDocResponse *models.WarehouseDocResponse) models.WarehouseDocRequest {
	return models.WarehouseDocRequest{
		Warehouse_code:      warehouseDocResponse.Warehouse_code,
		Address:             warehouseDocResponse.Address,
		Telephone:           warehouseDocResponse.Telephone,
		Minimun_capacity:    warehouseDocResponse.Minimun_capacity,
		Minimun_temperature: warehouseDocResponse.Minimun_temperature,
		Locality_id:         warehouseDocResponse.Locality_id,
	}
}
