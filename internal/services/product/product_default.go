package product

import (
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product"
	errorCustom "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/errors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func NewProductDefault(rp repository.ProductRepository) ProductService {
	return &ProductDefault{rp: rp}
}

type ProductDefault struct {
	rp repository.ProductRepository
}

func (s *ProductDefault) GetAll() (map[int]models.ProductDoc, error) {
	products, _ := s.rp.GetAll()
	productsDoc := mappers.ProductsToProductsDoc(products)
	return productsDoc, nil
}

func (s *ProductDefault) GetById(id int) (*models.ProductDoc, error) {
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

func (s *ProductDefault) Create(productDoc models.ProductDoc) (*models.ProductDoc, error) {
	existInDb := s.rp.ExistInDb(productDoc.ProductCode)
	if existInDb {
		return nil, errorCustom.ErrorConflict
	}
	product := mappers.ProductDocToProduct(productDoc)
	product.Id = s.rp.GenerateId()
	err := s.rp.Create(product)
	if err != nil {
		return nil, nil
	}
	productDoc = mappers.ProductToProductDoc(product)
	return &productDoc, nil
}
