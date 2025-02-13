package product_record

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type ProductRecordService interface {
	Create(models.ProductRecordDocRequest) (*models.ProductRecordDocResponse, error)
}
