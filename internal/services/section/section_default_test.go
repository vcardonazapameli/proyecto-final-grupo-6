package section_test

import (
	"testing"

	pt "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product_type"
	st "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/section"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/section"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/stretchr/testify/assert"
)

var (
	sectionAttributes = models.SectionAttributes{
		SectionNumber:      "A1",
		CurrentCapacity:    10,
		CurrentTemperature: 20.0,
		MaximumCapacity:    20,
		MinimumCapacity:    5,
		MinimumTemperature: 10.0,
		ProductTypeId:      1,
		WarehouseId:        1,
	}

	newSection = models.Section{
		Id:                1,
		SectionAttributes: sectionAttributes,
	}
)

func setupMocks() (*st.SectionRepositoryMock, *pt.ProductTypeRepositoryMock, *section.SectionDefault) {
	mockRepo := new(st.SectionRepositoryMock)
	mockProductTypeRepo := new(pt.ProductTypeRepositoryMock)
	sv := section.NewSectionDefault(mockRepo, mockProductTypeRepo)
	return mockRepo, mockProductTypeRepo, sv.(*section.SectionDefault)
}

func TestCreateSection(t *testing.T) {
	t.Run("create_ok: create sections successfully", func(t *testing.T) {
		// Arrange
		mockRepo, mockProductTypeRepo, sv := setupMocks()
		mockProductTypeRepo.On("GetById", sectionAttributes.ProductTypeId).Return(&models.ProductTypeDocResponse{}, nil)
		mockRepo.On("Create", sectionAttributes).Return(newSection, nil)
		// Act
		res, err := sv.Create(newSection)
		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, newSection, res)
	})

	t.Run("create_conflict: section_number already exists", func(t *testing.T) {
		// Arrange
		mockRepo, mockProductTypeRepo, sv := setupMocks()
		err := customErrors.ErrorConflict
		mockProductTypeRepo.On("GetById", sectionAttributes.ProductTypeId).Return(&models.ProductTypeDocResponse{}, nil)
		mockRepo.On("Create", sectionAttributes).Return(models.Section{}, err)
		// Act
		res, err := sv.Create(newSection)
		// Assert
		assert.Error(t, err)
		assert.Equal(t, models.Section{}, res)

	})
}
