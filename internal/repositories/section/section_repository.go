package section

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type SectionRepository interface {
	GetAll() (map[int]models.Section, error)
	GetByID(id int) (st models.Section, err error)
	Create(s models.Section) (st models.Section, err error)
}
