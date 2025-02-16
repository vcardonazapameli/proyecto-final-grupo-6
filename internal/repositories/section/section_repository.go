package section

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type SectionRepository interface {
	GetAll() (map[int]models.Section, error)
	GetByID(id int) (models.Section, error)
	Create(section models.SectionAttributes) (models.Section, error)
	Update(id int, section models.Section) (models.Section, error)
	Delete(id int) error
	Recover(id int) error
	SectionExists(id int) bool
}
