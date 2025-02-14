package product

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type ProductTypeRepository interface {
	GetAll() ([]models.ProductTypeDocResponse, error)
	GetById(int) (*models.ProductTypeDocResponse, error)
	Create(*models.ProductTypeDocResponse) error
	ExistInDb(string) (bool, error)
}
