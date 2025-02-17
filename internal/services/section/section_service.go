package section

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type SectionService interface {
	GetAll() (map[int]models.Section, error)
	GetByID(id int) (models.Section, error)
	Create(section models.Section) (models.Section, error)
	Update(id int, sectionDto models.UpdateSectionDto) (models.Section, error)
	Delete(id int) error
	GetSectionReports(sectionId int) ([]models.SectionReport, error)
}
