package section

import (
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/section"
)

func NewSectionDefault(rp repository.SectionRepository) SectionService {
	return &SectionDefault{rp: rp}
}

type SectionDefault struct {
	rp repository.SectionRepository
}
