package seller

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type SellerRepository interface {
	GetAll() (map[int]models.Seller, error)
	SearchByCID(int) (models.Seller, bool)
	Save(*models.Seller) error
	GetByID(id int) (models.Seller, error)
	Delete(id int) error
	Update(models.Seller)
}
