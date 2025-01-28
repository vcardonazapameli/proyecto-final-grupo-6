package seller

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type SellerLoader interface {
	Load() (map[int]models.Seller, error)
}
