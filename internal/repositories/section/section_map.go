package section

import (
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

func (r *SectionMap) GetAll() (v map[int]models.Section, err error) {
	return
}
