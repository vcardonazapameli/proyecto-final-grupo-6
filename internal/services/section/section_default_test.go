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

	currentCapacity    = 15
	currentTemperature = 22.0

	updateSectionDto = models.UpdateSectionDto{
		CurrentCapacity:    &currentCapacity,
		CurrentTemperature: &currentTemperature,
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

func TestReadSection(t *testing.T) {
	t.Run("find_all: read sections successfully", func(t *testing.T) {
		// Arrange
		mockRepo, _, sv := setupMocks()
		sections := map[int]models.Section{
			1: newSection,
		}
		mockRepo.On("GetAll").Return(sections, nil)
		// Act
		res, err := sv.GetAll()
		// Assert
		assert.NoError(t, err)
		assert.Equal(t, sections, res)
		assert.Len(t, res, 1)
	})

	t.Run("find_by_id: read section successfully", func(t *testing.T) {
		// Arrange
		mockRepo, _, sv := setupMocks()
		mockRepo.On("GetByID", newSection.Id).Return(newSection, nil)
		// Act
		res, err := sv.GetByID(newSection.Id)
		// Assert
		assert.NoError(t, err)
		assert.Equal(t, newSection, res)
	})

	t.Run("find_by_id_non_existent: section not found", func(t *testing.T) {
		// Arrange
		mockRepo, _, sv := setupMocks()
		err := customErrors.ErrorNotFound
		notExistentId := 9999
		mockRepo.On("GetByID", notExistentId).Return(models.Section{}, err)
		// Act
		res, err := sv.GetByID(notExistentId)
		// Assert
		assert.ErrorIs(t, err, customErrors.ErrorNotFound)
		assert.Empty(t, res)
	})
}

func TestUpdateSection(t *testing.T) {
	t.Run("update_existent: update section successfully", func(t *testing.T) {
		// Arrange
		mockRepo, mockProductTypeRepo, sv := setupMocks()
		newSection.CurrentCapacity = currentCapacity
		newSection.CurrentTemperature = currentTemperature
		mockRepo.On("GetByID", newSection.Id).Return(newSection, nil)
		mockProductTypeRepo.On("GetById", newSection.ProductTypeId).Return(&models.ProductTypeDocResponse{}, nil)
		mockRepo.On("Update", newSection.Id, newSection).Return(newSection, nil)
		// Act
		res, err := sv.Update(newSection.Id, updateSectionDto)
		// Assert
		assert.NoError(t, err)
		assert.Equal(t, newSection, res)
	})

	t.Run("update_non_existent: section not found", func(t *testing.T) {
		// Arrange
		mockRepo, _, sv := setupMocks()
		nonExistentID := 9999
		err := customErrors.ErrorNotFound
		mockRepo.On("GetByID", nonExistentID).Return(models.Section{}, err)
		// Act
		res, err := sv.Update(nonExistentID, updateSectionDto)
		// Assert
		assert.ErrorIs(t, err, customErrors.ErrorNotFound)
		assert.Empty(t, res)
	})
}

func TestDeleteSection(t *testing.T) {
	t.Run("delete_ok: delete section successfully", func(t *testing.T) {
		// Arrange
		mockRepo, _, sv := setupMocks()
		mockRepo.On("GetByID", newSection.Id).Return(newSection, nil)
		mockRepo.On("Delete", newSection.Id).Return(nil)
		// Act
		err := sv.Delete(newSection.Id)
		// Assert
		assert.NoError(t, err)
	})

	t.Run("delete_non_existent: section not found", func(t *testing.T) {
		// Arrange
		mockRepo, _, sv := setupMocks()
		nonExistentID := 9999
		err := customErrors.ErrorNotFound
		mockRepo.On("GetByID", nonExistentID).Return(models.Section{}, err)
		// Act
		err = sv.Delete(nonExistentID)
		// Assert
		assert.ErrorIs(t, err, customErrors.ErrorNotFound)
	})
}
