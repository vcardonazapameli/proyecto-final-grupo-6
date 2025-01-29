package product

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type ProductService interface {
	GetAll() (map[int]models.ProductDoc, error)
	GetById(int) (*models.ProductDoc, error)
	Delete(int) error
	Create(models.ProductDoc) (*models.ProductDoc, error)
}
