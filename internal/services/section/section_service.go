package section

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type SectionService interface {
	GetAll() ([]models.SectionDoc, error)
	GetByID(id int) (*models.SectionDoc, error)
	Create(section models.SectionDocRequest) (*models.SectionDoc, error)
	Update(id int, sectionDto models.UpdateSectionDto) (*models.SectionDoc, error)
	Delete(id int) error
	GetSectionReports(sectionId int) ([]models.SectionReport, error)
}
