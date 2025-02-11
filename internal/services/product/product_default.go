package product

import (
	"sort"

	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product"
	errorCustom "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/validators"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func NewProductDefault(rp repository.ProductRepository) ProductService {
	return &ProductDefault{rp: rp}
}

type ProductDefault struct {
	rp repository.ProductRepository
}

func (s *ProductDefault) GetAll() ([]models.ProductDocResponse, error) {
	products, _ := s.rp.GetAll()
	productsDoc := mappers.ProductsToProductsDoc(products)
	productsSlice := []models.ProductDocResponse{}
	for _, product := range productsDoc {
		productsSlice = append(productsSlice, product)
	}
	sort.Slice(productsSlice, func(i, j int) bool {
		return productsSlice[i].Id < productsSlice[j].Id
	})
	return productsSlice, nil
}

func (s *ProductDefault) GetById(id int) (*models.ProductDocResponse, error) {
	product, _ := s.rp.GetById(id)
	if product == nil {
		return nil, errorCustom.ErrorNotFound
	}
	productDoc := mappers.ProductToProductDoc(*product)
	return &productDoc, nil
}

func (s *ProductDefault) Delete(id int) error {
	product, _ := s.rp.GetById(id)
	if product == nil {
		return errorCustom.ErrorNotFound
	}
	s.rp.Delete(product.Id)
	return nil
}

func (s *ProductDefault) Create(productDocRquest models.ProductDocRequest) (*models.ProductDocResponse, error) {

	if errorValidateFields := validators.ValidateFieldsProduct(productDocRquest); errorValidateFields != nil {

		return nil, errorValidateFields
	}
	existInDb := s.rp.ExistInDb(productDocRquest.ProductCode)
	if existInDb {
		return nil, errorCustom.ErrorConflict
	}
	product := mappers.ProductDocRequestToProduct(productDocRquest)
	product.Id = s.rp.GenerateId()
	err := s.rp.Create(product)
	if err != nil {
		return nil, nil
	}
	productDoc := mappers.ProductToProductDoc(product)
	return &productDoc, nil
}

func (s *ProductDefault) Update(id int, productDocRequest models.ProductUpdateDocRequest) (*models.ProductDocResponse, error) {
	product, _ := s.rp.GetById(id)
	if product == nil {
		return nil, errorCustom.ErrorNotFound
	}
	productUpdate := validators.UpdateEntity(productDocRequest, product)
	if s.rp.MatchWithTheSameProductCode(productUpdate.Id, productUpdate.ProductCode) {
		return nil, errorCustom.ErrorConflict
	}
	productUpdateDocRequest := mappers.ProductUpdateDocRequestToProductDocRequest(productUpdate)

	if errorValidateFields := validators.ValidateFieldsProduct(productUpdateDocRequest); errorValidateFields != nil {

		return nil, errorValidateFields
	}
	err := s.rp.Update(id, productUpdate)
	if err != nil {
		return nil, nil
	}
	productDoc := mappers.ProductToProductDoc(*productUpdate)
	return &productDoc, nil
}
