package seller

import repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/seller"

type SellerServiceDefault struct {
	rp repository.SellerRepository
}

func NewSellerServiceDefault(rp repository.SellerRepository) *SellerServiceDefault {
	return &SellerServiceDefault{rp}
}
