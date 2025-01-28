package section

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

func NewSectionJSONFile(path string) *SectionJSONFile {
	return &SectionJSONFile{
		path: path,
	}
}

type SectionJSONFile struct {
	path string
}

func (l *SectionJSONFile) Load() (v map[int]models.Section, err error) {
	return
}
