package seller

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type SellerService interface {
	GetAll() (map[int]models.SellerDoc, error)
}
