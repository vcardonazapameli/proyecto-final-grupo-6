package product

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type ProductService interface {
	GetAll() ([]models.ProductDocResponse, error)
	GetById(int) (*models.ProductDocResponse, error)
	Delete(int) error
	Create(models.ProductDocRequest) (*models.ProductDocResponse, error)
	Update(int, models.ProductUpdateDocRequest) (*models.ProductDocResponse, error)
}
