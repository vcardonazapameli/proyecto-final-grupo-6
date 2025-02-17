package product_type

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type ProductTypeService interface {
	GetAll() ([]models.ProductTypeDocResponse, error)
	GetById(int) (*models.ProductTypeDocResponse, error)
	Create(models.ProductTypeDocRequest) (*models.ProductTypeDocResponse, error)
}
