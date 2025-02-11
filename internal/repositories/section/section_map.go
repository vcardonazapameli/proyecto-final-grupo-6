package section

import (
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
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
func (r *SectionMap) GetAll() (map[int]models.Section, error) {
	return r.db, nil
}

// Get a section by ID
func (r *SectionMap) GetByID(id int) (models.Section, error) {
	section, ok := r.db[id]
	if !ok {
		return models.Section{}, customErrors.ErrorNotFound
	}
	return section, nil
}

func (r *SectionMap) SearchBySectionNumber(sn string) (models.Section, bool) {
	for _, section := range r.db {
		if section.SectionNumber == sn {
			return section, true
		}
	}
	return models.Section{}, false
}

func (r *SectionMap) GetBiggestID() (max int) {
	for _, section := range r.db {
		if section.Id > max {
			max = section.Id
		}
	}
	return
}

// Create a new section
func (r *SectionMap) Create(section models.Section) (models.Section, error) {
	sectionDoc := mappers.SectionToSectionValidation(section)
	if err := validators.ValidateNoEmptyFields(sectionDoc); err != nil {
		return models.Section{}, customErrors.ErrorUnprocessableContent
	}

	if _, exists := r.SearchBySectionNumber(section.SectionNumber); exists {
		return models.Section{}, customErrors.ErrorConflict
	}
	section.Id = r.nextID
	r.nextID++

	r.db[section.Id] = section
	return section, nil
}

func (r *SectionMap) Update(id int, section models.Section) (models.Section, error) {
	r.db[id] = section
	return section, nil
}

func (r *SectionMap) Delete(id int) error {
	_, ok := r.db[id]
	if !ok {
		return customErrors.ErrorNotFound
	}
	delete(r.db, id)
	return nil
}
