package section

import (
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/errors"
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
func (s *SectionMap) GetByID(id int) (st models.Section, err error) {
	section, ok := s.db[id]
	if !ok {
		return models.Section{}, errors.ErrorNotFound
	}
	return section, nil
}

// Create a new section
func (s *SectionMap) Create(section models.Section) (st models.Section, err error) {
	if section.SectionNumber == "" {
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
