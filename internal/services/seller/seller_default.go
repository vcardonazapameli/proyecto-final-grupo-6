package seller

import (
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/seller"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

type SellerServiceDefault struct {
	rp repository.SellerRepository
}

func NewSellerServiceDefault(rp repository.SellerRepository) *SellerServiceDefault {
	return &SellerServiceDefault{rp}
}

func (sv *SellerServiceDefault) GetAll() (s map[int]models.SellerDoc, err error) {
	return
}
