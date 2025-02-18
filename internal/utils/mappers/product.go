package mappers

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

func ProductDocRequestToProductDocResponse(productDoc models.ProductDocRequest) models.ProductDocResponse {
	return models.ProductDocResponse{
		ProductCode:                    productDoc.ProductCode,
		Description:                    productDoc.Description,
		ExpirationRate:                 productDoc.ExpirationRate,
		RecommendedFreezingTemperature: productDoc.RecommendedFreezingTemperature,
		FreezingRate:                   productDoc.FreezingRate,
		Width:                          productDoc.Width,
		Height:                         productDoc.Height,
		Length:                         productDoc.Length,
		NetWeight:                      productDoc.NetWeight,
		ProductType:                    productDoc.ProductType,
		Seller:                         productDoc.Seller,
	}
}

func ProductDocResponseToProductDocRequest(productDoc *models.ProductDocResponse) models.ProductDocRequest {
	return models.ProductDocRequest{
		ProductCode:                    productDoc.ProductCode,
		Description:                    productDoc.Description,
		ExpirationRate:                 productDoc.ExpirationRate,
		RecommendedFreezingTemperature: productDoc.RecommendedFreezingTemperature,
		FreezingRate:                   productDoc.FreezingRate,
		Width:                          productDoc.Width,
		Height:                         productDoc.Height,
		Length:                         productDoc.Length,
		NetWeight:                      productDoc.NetWeight,
		ProductType:                    productDoc.ProductType,
		Seller:                         productDoc.Seller,
	}
}
