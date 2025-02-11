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

func (s *LocalityServiceDefault) Create(loc *models.Locality) error {
	err := validators.ValidateLocality(*loc)
	if err != nil {
		return err
	}

	err = s.rp.Save(loc)
	return err
}
