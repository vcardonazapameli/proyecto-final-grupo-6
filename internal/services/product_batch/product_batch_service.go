package productbatch

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type ProductBatchService interface {
	Create(pb models.ProductBatchRequest) (*models.ProductBatchResponse, error)
}
