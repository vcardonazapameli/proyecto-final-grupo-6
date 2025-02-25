package section_test

import (
	"testing"

	pt "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product_type"
	st "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/section"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/section"
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

func TestCreateSection(t *testing.T) {
	t.Run("case 1: create sections successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(st.SectionRepositoryMock)
		mockRepo.On("Create", sectionAttributes).Return(newSection, nil)

		mockProductTypeRepo := new(pt.ProductTypeRepositoryMock)
		mockProductTypeRepo.On("GetById", sectionAttributes.ProductTypeId).Return(&models.ProductTypeDocResponse{}, nil)

		sv := section.NewSectionDefault(mockRepo, mockProductTypeRepo)
		// Act
		res, err := sv.Create(newSection)
		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, newSection)
		assert.Equal(t, newSection, res)
	})
}
