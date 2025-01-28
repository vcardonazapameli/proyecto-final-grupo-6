package seller

import (
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/seller"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/errors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

type SellerServiceDefault struct {
	rp repository.SellerRepository
}

func NewSellerServiceDefault(rp repository.SellerRepository) *SellerServiceDefault {
	return &SellerServiceDefault{rp}
}

func (sv *SellerServiceDefault) GetAll() (map[int]models.SellerDoc, error) {
	s := make(map[int]models.SellerDoc)
	sellers, err := sv.rp.GetAll()

	if err != nil {
		return s, errors.ErrorInternalServerError
	}

	for _, sel := range sellers {
		sDoc := mappers.SellerToSellerDoc(sel)
		s[sDoc.Id] = sDoc
	}
	return s, nil
}
