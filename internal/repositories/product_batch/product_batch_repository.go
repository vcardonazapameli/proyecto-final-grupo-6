package productbatch

import (
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

type ProductBatchRepository interface {
	Save(pb *models.ProductBatch) error
}
