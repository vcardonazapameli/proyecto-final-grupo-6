package section

import (
	productTypeRepository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product_type"
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/section"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/validators"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func NewSectionDefault(rp repository.SectionRepository, productTypeRepository productTypeRepository.ProductTypeRepository) SectionService {
	return &SectionDefault{rp: rp, productTypeRespository: productTypeRepository}
}

type SectionDefault struct {
	rp                     repository.SectionRepository
	productTypeRespository productTypeRepository.ProductTypeRepository
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
	if err := validators.ValidateNoEmptyFields(section.SectionAttributes); err != nil {
		return models.Section{}, customErrors.ErrorUnprocessableContent
	}
	if err := validators.ValidateCapacity(section); err != nil {
		return models.Section{}, err
	}
	if err := validators.ValidateTemperature(section); err != nil {
		return models.Section{}, err
	}
	productType, err := s.productTypeRespository.GetById(section.ProductTypeId)
	if err != nil {
		return models.Section{}, err
	}
	if productType == nil {
		return models.Section{}, customErrors.ErrorNotFound
	}
	createdSection, err := s.rp.Create(section.SectionAttributes)
	if err != nil {
		return models.Section{}, err
	}
	return createdSection, nil
}

func (s *SectionDefault) Update(id int, sectionDto models.UpdateSectionDto) (models.Section, error) {
	sectionToUpdate, err := s.rp.GetByID(id)
	if err != nil {
		return models.Section{}, err
	}

	updatedSection := validators.UpdateEntity(sectionDto, &sectionToUpdate)
	if err := validators.ValidateCapacity(*updatedSection); err != nil {
		return models.Section{}, err
	}
	if err := validators.ValidateTemperature(*updatedSection); err != nil {
		return models.Section{}, err
	}

	productType, _ := s.productTypeRespository.GetById(updatedSection.ProductTypeId)
	if productType == nil {
		return models.Section{}, customErrors.ErrorNotFound
	}
	result, err := s.rp.Update(id, *updatedSection)
	if err != nil {
		return models.Section{}, err
	}

	return result, nil
}

func (s *SectionDefault) Recover(id int) (models.Section, error) {
	err := s.rp.Recover(id)
	if err != nil {
		return models.Section{}, err
	}
	recoveredSection, err := s.rp.GetByID(id)
	if err != nil {
		return models.Section{}, err
	}

	return recoveredSection, nil
}

func (s *SectionDefault) Delete(id int) error {
	section, err := s.rp.GetByID(id)
	if err != nil {
		return err
	}
	err = s.rp.Delete(section.Id)
	if err != nil {
		return err
	}

	return nil
}

func (s *SectionDefault) GetSectionReports(sectionId int) ([]models.SectionReport, error) {
	sectionReports, err := s.rp.GetSectionReports(sectionId)
	if err != nil {
		return sectionReports, customErrors.ErrorNotFound
	}
	return sectionReports, nil
}
