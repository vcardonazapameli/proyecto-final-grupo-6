package section

import (
	"fmt"

	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/errors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/validators"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func NewSectionMap(db map[int]models.Section) SectionRepository {
	defaultDb := make(map[int]models.Section)
	if db != nil {
		defaultDb = db
	}
	return &SectionMap{db: defaultDb}
}

type SectionMap struct {
	db map[int]models.Section
}

// List all sections
func (s *SectionMap) GetAll() (map[int]models.Section, error) {
	return s.db, nil
}

// Get a section by ID
func (s *SectionMap) GetByID(id int) (models.Section, error) {
	section, ok := s.db[id]
	if !ok {
		return models.Section{}, errors.ErrorNotFound
	}
	return section, nil
}

// Create a new section
func (s *SectionMap) Create(section models.Section) (models.Section, error) {
	sectionDoc := mappers.SectionToSectionValidation(section)
	if err := validators.ValidateNoEmptyFields(sectionDoc); err != nil {
		fmt.Println("Validation error:", err)
		return models.Section{}, errors.ErrorUnprocessableContent
	}
	for _, sec := range s.db {
		if sec.SectionNumber == section.SectionNumber {
			return models.Section{}, errors.ErrorConflict
		}
	}
	newID := len(s.db) + 1
	section.Id = newID
	s.db[newID] = section
	return section, nil
}

func (s *SectionMap) Update(id int, sectionDTO models.UpdateSectionDto) (models.Section, error) {
	existSection, ok := s.db[id]
	if !ok {
		return models.Section{}, errors.ErrorNotFound
	}

	updatedSection := validators.UpdateEntity(sectionDTO, &existSection)
	s.db[id] = *updatedSection
	return *updatedSection, nil
}

func (s *SectionMap) Delete(id int) error {
	_, ok := s.db[id]
	if !ok {
		return errors.ErrorNotFound
	}
	delete(s.db, id)
	return nil
}
