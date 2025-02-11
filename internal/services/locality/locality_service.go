package locality

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type LocalityService interface {
	Create(l *models.Locality) error
}
