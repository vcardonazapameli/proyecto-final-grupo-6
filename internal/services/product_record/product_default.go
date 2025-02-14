package product_record

import (
	productRepository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product"
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product_record"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func NewProductRecordService(repository repository.ProductRecordRepository, productRepository productRepository.ProductRepository) ProductRecordService {
	return &productRecordService{repository: repository, productRepository: productRepository}
}

type productRecordService struct {
	repository        repository.ProductRecordRepository
	productRepository productRepository.ProductRepository
}

func (productRecordService *productRecordService) Create(productRecordDocRequest models.ProductRecordDocRequest) (*models.ProductRecordDocResponse, error) {
	// if errorValidateFields := validators.ValidateFieldsProduct(productDocRequest); errorValidateFields != nil {
	// 	return nil, errorValidateFields
	// }
	ExistProductInDb, err := productRecordService.productRepository.GetById(productRecordDocRequest.ProductId)
	if err != nil {
		return nil, err
	}
	if ExistProductInDb == nil {
		return nil, customErrors.ErrorNotFound
	}
	productRecord := mappers.ProductRecordDocRequestToProductRecordDocResponse(productRecordDocRequest)
	if err := productRecordService.repository.Create(&productRecord); err != nil {
		return nil, err
	}
	return &productRecord, nil
}
