package product

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type ProductRepository interface {
	GetAll() (map[int]models.Product, error)
	GetById(int) (*models.Product, error)
}
