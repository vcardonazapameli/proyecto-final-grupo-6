package mappers

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

func ProductToProductDoc(product models.Product) models.ProductDocResponse {
	return models.ProductDocResponse{
		Id:                             product.Id,
		ProductCode:                    product.ProductCode,
		Description:                    product.Description,
		ExpirationRate:                 product.ExpirationRate,
		RecommendedFreezingTemperature: product.RecommendedFreezingTemperature,
		FreezingRate:                   product.FreezingRate,
		Width:                          product.Width,
		Height:                         product.Height,
		Length:                         product.Length,
		NetWeight:                      product.NetWeight,
		ProductType:                    product.ProductType,
		Seller:                         product.Seller,
	}
}

func ProductDocRequestToProduct(productDoc models.ProductDocRequest) models.Product {
	return models.Product{
		ProductAttributes: models.ProductAttributes{
			ProductCode:                    productDoc.ProductCode,
			Description:                    productDoc.Description,
			ExpirationRate:                 productDoc.ExpirationRate,
			RecommendedFreezingTemperature: productDoc.RecommendedFreezingTemperature,
			FreezingRate:                   productDoc.FreezingRate,
			Dimensions: models.Dimensions{
				Width:     productDoc.Width,
				Height:    productDoc.Height,
				Length:    productDoc.Length,
				NetWeight: productDoc.NetWeight,
			},
			ProductType: productDoc.ProductType,
			Seller:      productDoc.Seller,
		},
	}
}

func ProductsToProductsDoc(products map[int]models.Product) map[int]models.ProductDocResponse {
	productsDoc := map[int]models.ProductDocResponse{}
	for _, product := range products {
		productsDoc[product.Id] = models.ProductDocResponse{
			Id:                             product.Id,
			ProductCode:                    product.ProductCode,
			Description:                    product.Description,
			ExpirationRate:                 product.ExpirationRate,
			RecommendedFreezingTemperature: product.RecommendedFreezingTemperature,
			FreezingRate:                   product.FreezingRate,
			Width:                          product.Width,
			Height:                         product.Height,
			Length:                         product.Length,
			NetWeight:                      product.NetWeight,
			ProductType:                    product.ProductType,
			Seller:                         product.Seller,
		}
	}
	return productsDoc
}

func ProductUpdateDocRequestToProductDocRequest(product *models.Product) models.ProductDocRequest {
	return models.ProductDocRequest{
		ProductCode:                    product.ProductCode,
		Description:                    product.Description,
		ExpirationRate:                 product.ExpirationRate,
		RecommendedFreezingTemperature: product.RecommendedFreezingTemperature,
		FreezingRate:                   product.FreezingRate,
		Width:                          product.Width,
		Height:                         product.Height,
		Length:                         product.Length,
		NetWeight:                      product.NetWeight,
		ProductType:                    product.ProductType,
		Seller:                         product.Seller,
	}
}
