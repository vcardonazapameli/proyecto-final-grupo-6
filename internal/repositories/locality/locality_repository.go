package locality

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type LocalityRepository interface {
	Save(loc *models.Locality) error
}
