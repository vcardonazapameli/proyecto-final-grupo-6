package section

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type SectionLoader interface {
	Load() (v map[int]models.Section, err error)
}
