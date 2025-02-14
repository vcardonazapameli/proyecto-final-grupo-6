package mappers

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

func WarehouseDocRequestToWarehouseUpdateDocResponse(warehouse *models.WarehouseDocRequest) *models.WarehouseUpdateDocResponse {
	return &models.WarehouseUpdateDocResponse{
		Warehouse_code:      &warehouse.Warehouse_code,
		Address:             &warehouse.Address,
		Telephone:           &warehouse.Telephone,
		Minimun_capacity:    &warehouse.Minimun_capacity,
		Minimun_temperature: &warehouse.Minimun_temperature,
		Locality_id:         warehouse.Locality_id,
	}
}

func WarehouseDocResponseToWarehouseUpdateDocResponse(warehouse *models.WarehouseDocResponse) *models.WarehouseUpdateDocResponse {
	return &models.WarehouseUpdateDocResponse{
		Warehouse_code:      &warehouse.Warehouse_code,
		Address:             &warehouse.Address,
		Telephone:           &warehouse.Telephone,
		Minimun_capacity:    &warehouse.Minimun_capacity,
		Minimun_temperature: &warehouse.Minimun_temperature,
		Locality_id:         warehouse.Locality_id,
	}
}

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

func WarehouseDocResponseToWarehouseDocRequest(WarehouseUpdateDocResponse *models.WarehouseDocResponse) models.WarehouseDocRequest {
	return models.WarehouseDocRequest{
		Warehouse_code:      WarehouseUpdateDocResponse.Warehouse_code,
		Address:             WarehouseUpdateDocResponse.Address,
		Telephone:           WarehouseUpdateDocResponse.Telephone,
		Minimun_capacity:    WarehouseUpdateDocResponse.Minimun_capacity,
		Minimun_temperature: WarehouseUpdateDocResponse.Minimun_temperature,
		Locality_id:         WarehouseUpdateDocResponse.Locality_id,
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
