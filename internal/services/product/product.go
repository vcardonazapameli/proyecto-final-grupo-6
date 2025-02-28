package product

import (
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product"
	productTypeRepository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product_type"
	sellerRepository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/seller"
	errorCustom "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/validators"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func NewProductService(repository repository.ProductRepository, productTypeRepository productTypeRepository.ProductTypeRepository, sellerRepository sellerRepository.SellerRepository) ProductService {
	return &productService{repository: repository, productTypeRespository: productTypeRepository, sellerRepository: sellerRepository}
}

type productService struct {
	repository             repository.ProductRepository
	productTypeRespository productTypeRepository.ProductTypeRepository
	sellerRepository       sellerRepository.SellerRepository
}

func (productService *productService) GetAll() ([]models.ProductDocResponse, error) {
	products, err := productService.repository.GetAll()
	if err != nil {
		return nil, err
	}
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
	err := productService.repository.Delete(product.Id)
	if err != nil {
		return err
	}
	return nil
}

func (productService *productService) Create(productDocRequest models.ProductDocRequest) (*models.ProductDocResponse, error) {
	if errorValidateFields := validators.ValidateFieldsProduct(productDocRequest); errorValidateFields != nil {
		return nil, errorValidateFields
	}
	existInDb, _ := productService.repository.ExistInDb(productDocRequest.ProductCode)
	if existInDb {
		return nil, errorCustom.ErrorConflict
	}
	productType, _ := productService.productTypeRespository.GetById(productDocRequest.ProductType)
	if productType == nil {
		return nil, errorCustom.ErrorConflict
	}
	_, err := productService.sellerRepository.GetByID(productDocRequest.Seller)
	if err != nil {
		return nil, errorCustom.ErrorConflict
	}
	product := mappers.ProductDocRequestToProductDocResponse(productDocRequest)
	if err := productService.repository.Create(&product); err != nil {
		return nil, err
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
	productType, _ := productService.productTypeRespository.GetById(productUpdate.ProductType)
	if productType == nil {
		return nil, errorCustom.ErrorNotFound
	}
	productDoc := mappers.ProductDocResponseToProductDocRequest(productUpdate)
	if errorValidateFields := validators.ValidateFieldsProduct(productDoc); errorValidateFields != nil {
		return nil, errorValidateFields
	}
	productService.repository.Update(id, productUpdate)
	return productUpdate, nil
}
