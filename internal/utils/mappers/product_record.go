package mappers

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

func ProductRecordDocRequestToProductRecordDocResponse(productRecordDoc models.ProductRecordDocRequest) models.ProductRecordDocResponse {
	return models.ProductRecordDocResponse{
		LastUpdateDate: productRecordDoc.LastUpdateDate,
		PurchasePrice:  productRecordDoc.PurchasePrice,
		SalePrice:      productRecordDoc.SalePrice,
		ProductId:      productRecordDoc.ProductId,
	}
}

func ProductRecordDocResponseToProductRecordDocRequest(productRecordDoc *models.ProductRecordDocResponse) models.ProductRecordDocRequest {
	return models.ProductRecordDocRequest{
		LastUpdateDate: productRecordDoc.LastUpdateDate,
		PurchasePrice:  productRecordDoc.PurchasePrice,
		SalePrice:      productRecordDoc.SalePrice,
		ProductId:      productRecordDoc.ProductId,
	}
}
