package product

import (
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product"
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
