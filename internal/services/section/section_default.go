package section

import (
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/section"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/validators"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func NewSectionDefault(rp repository.SectionRepository) SectionService {
	return &SectionDefault{rp: rp}
}

type SectionDefault struct {
	rp repository.SectionRepository
}

func (s *SectionDefault) GetAll() (map[int]models.Section, error) {
	sections, err := s.rp.GetAll()
	if err != nil {
		return nil, err
	}
	return sections, nil
}

func (s *SectionDefault) GetByID(id int) (models.Section, error) {
	section, err := s.rp.GetByID(id)
	if err != nil {
		return models.Section{}, err
	}
	return section, nil
}

func (s *SectionDefault) Create(section models.Section) (models.Section, error) {
	if err := validators.ValidateNoEmptyFields(section.SectionAttributes); err != nil {
		return models.Section{}, customErrors.ErrorUnprocessableContent
	}
	if err := validators.ValidateCapacity(section); err != nil {
		return models.Section{}, err
	}
	if err := validators.ValidateTemperature(section); err != nil {
		return models.Section{}, err
	}
	createdSection, err := s.rp.Create(section.SectionAttributes)
	if err != nil {
		return models.Section{}, err
	}
	return createdSection, nil
}

func (s *SectionDefault) Update(id int, sectionDto models.UpdateSectionDto) (models.Section, error) {
	sectionToUpdate, err := s.rp.GetByID(id)
	if err != nil {
		return models.Section{}, err
	}
	updatedSection := validators.UpdateEntity(sectionDto, &sectionToUpdate)
	if err := validators.ValidateCapacity(*updatedSection); err != nil {
		return models.Section{}, err
	}
	if err := validators.ValidateTemperature(*updatedSection); err != nil {
		return models.Section{}, err
	}
	s.rp.Update(id, *updatedSection)
	return *updatedSection, nil
}

func (s *SectionDefault) Delete(id int) error {
	section, _ := s.rp.GetByID(id)
	s.rp.Delete(section.Id)
	return nil
}
