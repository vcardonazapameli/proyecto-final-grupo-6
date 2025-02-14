package product

import (
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product"
	errorCustom "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/validators"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func NewProductService(repository repository.ProductRepository) ProductService {
	return &productService{repository: repository}
}

type productService struct {
	repository repository.ProductRepository
}

func (productService *productService) GetAll() ([]models.ProductDocResponse, error) {
	products, _ := productService.repository.GetAll()
	return products, nil
}

func (productService *productService) GetById(id int) (*models.ProductDocResponse, error) {
	product, _ := productService.repository.GetById(id)
	if product == nil {
		return nil, errorCustom.ErrorNotFound
	}
	return product, nil
}

func (productService *productService) GetProductRecords(id *int, productTypeId *int, productCode *string) ([]models.ProductRecordByProductResponse, error) {
	products, err := productService.repository.GetProductRecords(id, productTypeId, productCode)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (productService *productService) Delete(id int) error {
	product, _ := productService.repository.GetById(id)
	if product == nil {
		return errorCustom.ErrorNotFound
	}
	productService.repository.Delete(product.Id)
	return nil
}

func (productService *productService) Create(productDocRequest models.ProductDocRequest) (*models.ProductDocResponse, error) {
	if errorValidateFields := validators.ValidateFieldsProduct(productDocRequest); errorValidateFields != nil {
		return nil, errorValidateFields
	}
	existInDb, err := productService.repository.ExistInDb(productDocRequest.ProductCode)
	if err != nil {
		return nil, err
	}
	if existInDb {
		return nil, errorCustom.ErrorConflict
	}
	product := mappers.ProductDocRequestToProductDocResponse(productDocRequest)
	if err := productService.repository.Create(&product); err != nil {
		return nil, nil
	}
	return &product, nil
}

func (productService *productService) Update(id int, productDocRequest models.ProductUpdateDocRequest) (*models.ProductDocResponse, error) {
	product, _ := productService.repository.GetById(id)
	if product == nil {
		return nil, errorCustom.ErrorNotFound
	}
	productUpdate := validators.UpdateEntity(productDocRequest, product)
	productCodeAlreadyAssociated, err := productService.repository.MatchWithTheSameProductCode(productUpdate.Id, productUpdate.ProductCode)
	if err != nil {
		return nil, err
	}
	if productCodeAlreadyAssociated {
		return nil, errorCustom.ErrorConflict
	}
	productDoc := mappers.ProductDocResponseToProductDocRequest(productUpdate)
	if errorValidateFields := validators.ValidateFieldsProduct(productDoc); errorValidateFields != nil {
		return nil, errorValidateFields
	}
	if err := productService.repository.Update(id, productUpdate); err != nil {
		return nil, nil
	}
	return productUpdate, nil
}
