package mappers

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

func CarrierDocRequestToCarrierDocResponse(carrierDoc models.CarrierDocRequest) models.CarrierDocResponse {
	return models.CarrierDocResponse{
		Cid:          carrierDoc.Cid,
		Company_name: carrierDoc.Company_name,
		Address:      carrierDoc.Address,
		Telephone:    carrierDoc.Telephone,
		Locality_id:  carrierDoc.Locality_id,
	}
}

func CarrierDocResponseToCarrierDocRequest(carrier *models.CarrierDocResponse) models.CarrierDocRequest {
	return models.CarrierDocRequest{
		Cid:          carrier.Cid,
		Company_name: carrier.Company_name,
		Address:      carrier.Address,
		Telephone:    carrier.Telephone,
		Locality_id:  carrier.Locality_id,
	}

}
