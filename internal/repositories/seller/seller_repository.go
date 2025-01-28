package seller

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type SellerRepository interface {
	GetAll() (map[int]models.Seller, error)
}
