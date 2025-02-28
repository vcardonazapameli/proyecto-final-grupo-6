package section

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type SectionRepository interface {
	GetAll() ([]models.SectionDoc, error)
	GetByID(id int) (*models.SectionDoc, error)
	Create(section *models.SectionDoc) error
	Update(id int, section *models.SectionDocRequest) error
	Delete(id int) error
	Recover(id int) error
	SectionExists(id int) bool
	GetSectionReports(sectionId int) ([]models.SectionReport, error)
}
