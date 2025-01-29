package section

import (
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/section"
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
	section, err := s.rp.Create(section)
	if err != nil {
		return models.Section{}, err
	}
	return section, nil
}

func (s *SectionDefault) Update(id int, section models.Section) (models.Section, error) {
	section, err := s.rp.Update(id, section)
	if err != nil {
		return models.Section{}, err
	}
	return section, nil
}

func (s *SectionDefault) Delete(id int) error {
	err := s.rp.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
