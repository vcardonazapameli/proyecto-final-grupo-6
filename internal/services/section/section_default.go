package section

import (
	productTypeRepository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product_type"
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/section"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
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

func (s *SectionDefault) GetAll() ([]models.SectionDoc, error) {
	sections, err := s.rp.GetAll()
	if err != nil {
		return nil, err
	}
	return sections, nil
}

func (s *SectionDefault) GetByID(id int) (*models.SectionDoc, error) {
	section, err := s.rp.GetByID(id)
	if err != nil {
		return nil, customErrors.ErrorNotFound
	}
	return section, nil
}

func (s *SectionDefault) Create(section models.SectionDocRequest) (*models.SectionDoc, error) {
	if err := validators.ValidateNoEmptyFields(section); err != nil {
		return nil, customErrors.ErrorUnprocessableContent
	}
	sectionDoc := mappers.SectionDocRequestToSectioDocResponse(section)
	if err := validators.ValidateCapacity(sectionDoc); err != nil {
		return nil, customErrors.ErrorConflict
	}
	if err := validators.ValidateTemperature(sectionDoc); err != nil {
		return nil, customErrors.ErrorConflict
	}
	productType, err := s.productTypeRespository.GetById(section.ProductTypeId)
	if err != nil {
		return nil, customErrors.ErrorNotFound
	}
	if productType == nil {
		return nil, customErrors.ErrorNotFound
	}

	sectionResponse := mappers.SectionDocRequestToSectioDocResponse(section)

	err = s.rp.Create(&sectionResponse)
	if err != nil {
		return nil, err
	}
	return &sectionResponse, nil
}

func (s *SectionDefault) Update(id int, sectionDto models.UpdateSectionDto) (*models.SectionDoc, error) {
	sectionToUpdate, err := s.rp.GetByID(id)
	if err != nil {
		return nil, err
	}
	updatedSection := validators.UpdateEntity(sectionDto, sectionToUpdate)
	if err := validators.ValidateCapacity(*updatedSection); err != nil {
		return nil, customErrors.ErrorConflict
	}
	if err := validators.ValidateTemperature(*updatedSection); err != nil {
		return nil, customErrors.ErrorConflict
	}
	productType, _ := s.productTypeRespository.GetById(updatedSection.ProductTypeId)
	if productType == nil {
		return nil, customErrors.ErrorNotFound
	}
	section := mappers.SectionDocResponseToSectionDocRequest(*updatedSection)
	if err := s.rp.Update(id, &section); err != nil {
		return nil, err
	}
	return updatedSection, nil
}

func (s *SectionDefault) Recover(id int) (models.SectionDoc, error) {
	err := s.rp.Recover(id)
	if err != nil {
		return models.SectionDoc{}, err
	}
	recoveredSection, err := s.rp.GetByID(id)
	if err != nil {
		return models.SectionDoc{}, err
	}

	return *recoveredSection, nil
}

func (s *SectionDefault) Delete(id int) error {
	section, err := s.rp.GetByID(id)
	if err != nil {
		return customErrors.ErrorNotFound
	}
	s.rp.Delete(section.Id)
	return nil
}

func (s *SectionDefault) GetSectionReports(sectionId int) ([]models.SectionReport, error) {
	sectionReports, err := s.rp.GetSectionReports(sectionId)
	if err != nil {
		return sectionReports, customErrors.ErrorNotFound
	}
	return sectionReports, nil
}
