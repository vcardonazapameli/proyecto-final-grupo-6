package product_type

import (
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product_type"
	errorCustom "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/validators"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func NewProductTypeService(repository repository.ProductTypeRepository) ProductTypeService {
	return &productTypeService{repository: repository}
}

type productTypeService struct {
	repository repository.ProductTypeRepository
}

func (productTypeService *productTypeService) GetAll() ([]models.ProductTypeDocResponse, error) {
	productTypes, _ := productTypeService.repository.GetAll()
	return productTypes, nil
}

func (productTypeService *productTypeService) GetById(id int) (*models.ProductTypeDocResponse, error) {
	productType, _ := productTypeService.repository.GetById(id)
	if productType == nil {
		return nil, errorCustom.ErrorNotFound
	}
	return productType, nil
}

func (productTypeService *productTypeService) Create(productTypeDocRequest models.ProductTypeDocRequest) (*models.ProductTypeDocResponse, error) {
	if errorValidateFields := validators.ValidateFieldsProductType(productTypeDocRequest); errorValidateFields != nil {
		return nil, errorValidateFields
	}
	existInDb, err := productTypeService.repository.ExistInDb(productTypeDocRequest.Description)
	if err != nil {
		return nil, err
	}
	if existInDb {
		return nil, errorCustom.ErrorConflict
	}
	product := mappers.ProductTypeDocRequestToProductTypeDocResponse(productTypeDocRequest)
	if err := productTypeService.repository.Create(&product); err != nil {
		return nil, nil
	}
	return &product, nil
}
