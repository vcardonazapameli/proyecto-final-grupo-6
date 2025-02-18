package mappers

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

func ProductTypeDocRequestToProductTypeDocResponse(productTypeDoc models.ProductTypeDocRequest) models.ProductTypeDocResponse {
	return models.ProductTypeDocResponse{
		Description: productTypeDoc.Description,
	}
}

func ProductTypeDocResponseToProductTypeDocRequest(productTypeDoc *models.ProductTypeDocResponse) models.ProductTypeDocRequest {
	return models.ProductTypeDocRequest{
		Description: productTypeDoc.Description,
	}
}
