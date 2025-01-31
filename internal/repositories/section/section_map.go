package section

import (
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
	sectionMap := &SectionMap{db: defaultDb}
	sectionMap.nextID = sectionMap.GetBiggestID() + 1

	return sectionMap
}

type SectionMap struct {
	db     map[int]models.Section
	nextID int
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

func (s *SectionMap) SearchBySectionNumber(sn string) (models.Section, bool) {
	for _, section := range s.db {
		if section.SectionNumber == sn {
			return section, true
		}
	}
	return models.Section{}, false
}

func (s *SectionMap) GetBiggestID() (max int) {
	for _, section := range s.db {
		if section.Id > max {
			max = section.Id
		}
	}
	return
}

// Create a new section
func (s *SectionMap) Create(section models.Section) (models.Section, error) {
	sectionDoc := mappers.SectionToSectionValidation(section)
	if err := validators.ValidateNoEmptyFields(sectionDoc); err != nil {
		return models.Section{}, errors.ErrorUnprocessableContent
	}

	if _, exists := s.SearchBySectionNumber(section.SectionNumber); exists {
		return models.Section{}, errors.ErrorConflict
	}
	section.Id = s.nextID
	s.nextID++

	s.db[section.Id] = section
	return section, nil
}

func (s *SectionMap) Update(id int, section models.Section) (models.Section, error) {
	// existSection, ok := s.db[id]
	// if !ok {
	// 	return models.Section{}, errors.ErrorNotFound
	// }

	// updatedSection := validators.UpdateEntity(sectionDTO, &existSection)
	// s.db[id] = *updatedSection
	// return *updatedSection, nil
	s.db[id] = section
	return section, nil
}

func (s *SectionMap) Delete(id int) error {
	_, ok := s.db[id]
	if !ok {
		return errors.ErrorNotFound
	}
	delete(s.db, id)
	return nil
}
