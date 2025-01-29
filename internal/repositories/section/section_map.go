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

// Update a section
func (s *SectionMap) Update(id int, section models.Section) (st models.Section, err error) {
	existSection, ok := s.db[id]
	if !ok {
		return models.Section{}, errors.ErrorNotFound
	}
	if section.SectionNumber != "" {
		existSection.SectionNumber = section.SectionNumber
	}
	if section.CurrentCapacity != 0 {
		existSection.CurrentCapacity = section.CurrentCapacity
	}
	if section.CurrentTemperature != 0 {
		existSection.CurrentTemperature = section.CurrentTemperature
	}
	if section.MaximumCapacity != 0 {
		existSection.MaximumCapacity = section.MaximumCapacity
	}
	if section.MinimumCapacity != 0 {
		existSection.MinimumCapacity = section.MinimumCapacity
	}
	if section.MinimumTemperature != 0 {
		existSection.MinimumTemperature = section.MinimumTemperature
	}
	if section.ProductTypeId != 0 {
		existSection.ProductTypeId = section.ProductTypeId
	}
	if section.WarehouseId != 0 {
		existSection.WarehouseId = section.WarehouseId
	}
	if len(section.ProductBatchId) != 0 {
		existSection.ProductBatchId = section.ProductBatchId
	}
	s.db[id] = existSection
	return existSection, nil
}

func (s *SectionMap) Delete(id int) error {
	_, ok := s.db[id]
	if !ok {
		return errors.ErrorNotFound
	}
	delete(s.db, id)
	return nil
}
