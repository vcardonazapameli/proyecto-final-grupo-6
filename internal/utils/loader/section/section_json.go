package section

import (
	"encoding/json"
	"os"

	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func NewSectionJSONFile(path string) *SectionJSONFile {
	return &SectionJSONFile{
		path: path,
	}
}

type SectionJSONFile struct {
	path string
}

func (l *SectionJSONFile) Load() (sections map[int]models.Section, err error) {
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	var sectionsJSON []models.SectionDoc
	err = json.NewDecoder(file).Decode(&sectionsJSON)
	if err != nil {
		return
	}

	sections = make(map[int]models.Section)
	for _, section := range sectionsJSON {
		sections[section.Id] = models.Section{
			Id: section.Id,
			SectionAttributes: models.SectionAttributes{
				SectionNumber:      section.SectionNumber,
				CurrentCapacity:    section.CurrentCapacity,
				CurrentTemperature: section.CurrentTemperature,
				MaximumCapacity:    section.MaximumCapacity,
				MinimumCapacity:    section.MinimumCapacity,
				MinimumTemperature: section.MinimumTemperature,
				ProductTypeId:      section.ProductTypeId,
				WarehouseId:        section.WarehouseId,
			},
		}
	}
	return
}
