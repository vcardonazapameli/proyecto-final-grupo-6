package locality

import (
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/locality"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/validators"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

type LocalityServiceDefault struct {
	rp repository.LocalityRepository
}

func NewLocalityServiceDefault(rp repository.LocalityRepository) *LocalityServiceDefault {
	return &LocalityServiceDefault{rp}
}

func (s *LocalityServiceDefault) Create(loc *models.LocalityDoc) error {
	err := validators.ValidateLocality(*loc)
	if err != nil {
		return err
	}

	// Searches if exists Province with same name and Same country name, if not 404.
	p, err := s.rp.GetProvinceWithCountryNames(loc.ProvinceName, loc.CountryName)
	if err != nil {
		return err
	}

	err = s.rp.Save(loc.Id, loc.LocalityName, p.Id)
	return err
}
