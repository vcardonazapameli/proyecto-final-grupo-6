package productbatch

import (
	p "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product"
	productbatch "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product_batch"
	s "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/section"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/validators"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

type productBatchService struct {
	productbatch      productbatch.ProductBatchRepository
	sectionRepository s.SectionRepository
	productRepository p.ProductRepository
}

func NewProductBatchService(productbatch productbatch.ProductBatchRepository,
	sectionRepo s.SectionRepository, productRepo p.ProductRepository) ProductBatchService {
	return &productBatchService{productbatch: productbatch, sectionRepository: sectionRepo, productRepository: productRepo}
}

func (s *productBatchService) Create(pb models.ProductBatchRequest) (*models.ProductBatchResponse, error) {
	if err := validators.ValidateNoEmptyFields(pb); err != nil {
		return nil, customErrors.ErrorUnprocessableContent
	}
	if s.productbatch.BatchNumberExists(pb.BatchNumber) {
		return nil, customErrors.ErrorConflict
	}
	if !s.sectionRepository.SectionExists(pb.SectionId) {
		return nil, customErrors.ErrorConflict
	}
	product, err := s.productRepository.GetById(pb.ProductId)
	if err != nil || product == nil {
		return nil, customErrors.ErrorConflict
	}
	productBatchResponse, err := mappers.ProductBatchRequestToProductBatch(pb)
	if err != nil {
		return nil, err
	}
	s.productbatch.Save(productBatchResponse)
	return productBatchResponse, nil
}
