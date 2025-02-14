package productbatch

import (
	productbatch "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product_batch"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

type productBatchService struct {
	rp *productbatch.ProductBatchRepositoryDB
}

func NewProductBatchService(rp *productbatch.ProductBatchRepositoryDB) ProductBatchService {
	return &productBatchService{rp}
}

func (s *productBatchService) Create(pb *models.ProductBatch) error {
	err := s.rp.Save(pb)
	return err
}
