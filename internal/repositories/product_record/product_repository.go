package product_record

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type ProductRecordRepository interface {
	Create(*models.ProductRecordDocResponse) error
}
