package product

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type ProductRepository interface {
	GetAll() ([]models.ProductDocResponse, error)
	GetById(int) (*models.ProductDocResponse, error)
	Delete(int) error
	Create(*models.ProductDocResponse) error
	ExistInDb(string) (bool, error)
	Update(int, *models.ProductDocResponse) error
	MatchWithTheSameProductCode(int, string) (bool, error)
}
