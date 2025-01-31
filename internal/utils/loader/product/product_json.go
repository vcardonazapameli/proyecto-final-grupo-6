package product

import (
	"encoding/json"
	"os"

	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func NewProductJSONFile(path string) *ProductJSONFile {
	return &ProductJSONFile{
		path: path,
	}
}

type ProductJSONFile struct {
	path string
}

func (l *ProductJSONFile) Load() (products map[int]models.Product, err error) {
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	var productsJSON []models.ProductDocResponse
	err = json.NewDecoder(file).Decode(&productsJSON)
	if err != nil {
		return
	}

	products = make(map[int]models.Product)
	for key, product := range productsJSON {
		products[key] = models.Product{
			Id: product.Id,
			ProductAttributes: models.ProductAttributes{
				ProductCode:                    product.ProductCode,
				Description:                    product.Description,
				ExpirationRate:                 product.ExpirationRate,
				RecommendedFreezingTemperature: product.RecommendedFreezingTemperature,
				FreezingRate:                   product.FreezingRate,
				Dimensions: models.Dimensions{
					Width:     product.Width,
					Height:    product.Height,
					Length:    product.Length,
					NetWeight: product.NetWeight,
				},
				ProductType: product.ProductType,
				Seller:      product.Seller,
			},
		}
	}
	return
}
