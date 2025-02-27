package section_test

import (
	"testing"

	pt "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product_type"
	st "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/section"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/section"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
		sectionRequest := models.SectionDocRequest{
			SectionNumber:      "A1",
			CurrentCapacity:    20,
			CurrentTemperature: 10,
			MaximumCapacity:    30,
			MinimumCapacity:    10,
			MinimumTemperature: 5,
			ProductTypeId:      1,
			WarehouseId:        1,
		}
		expectedSection := &models.SectionDoc{
			SectionNumber:      "A1",
			CurrentCapacity:    20,
			CurrentTemperature: 10,
			MaximumCapacity:    30,
			MinimumCapacity:    10,
			MinimumTemperature: 5,
			ProductTypeId:      1,
			WarehouseId:        1,
		}
		mockRepo.On("Create", mock.Anything).Return(nil)
		mockProductTypeRepo.On("GetById", sectionRequest.ProductTypeId).Return(&models.ProductTypeDocResponse{}, nil)
		// Act
		res, err := sv.Create(sectionRequest)
		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedSection, res)
		mockRepo.AssertExpectations(t)
		mockProductTypeRepo.AssertExpectations(t)
	})
	t.Run("create_conflict: section_number already exists", func(t *testing.T) {
		// Arrange
		mockRepo, mockProductTypeRepo, sv := setupMocks()
		sectionRequest := models.SectionDocRequest{
			SectionNumber:      "A1",
			CurrentCapacity:    20,
			CurrentTemperature: 10,
			MaximumCapacity:    30,
			MinimumCapacity:    10,
			MinimumTemperature: 5,
			ProductTypeId:      1,
			WarehouseId:        1,
		}
		mockProductTypeRepo.On("GetById", sectionRequest.ProductTypeId).Return(&models.ProductTypeDocResponse{}, nil)
		mockRepo.On("Create", mock.Anything).Return(customErrors.ErrorConflict)
		expected := customErrors.ErrorConflict
		// Act
		result, err := sv.Create(sectionRequest)
		// Assert
		assert.ErrorIs(t, err, expected)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
		mockProductTypeRepo.AssertExpectations(t)
	})
	t.Run("create_unprocessable_content: validation error", func(t *testing.T) {
		// Arrange
		mockRepo, mockProductTypeRepo, sv := setupMocks()
		sectionRequest := models.SectionDocRequest{
			SectionNumber:      "",
			CurrentCapacity:    20,
			CurrentTemperature: 10,
			MaximumCapacity:    30,
			MinimumCapacity:    10,
			MinimumTemperature: 5,
			ProductTypeId:      1,
			WarehouseId:        1,
		}
		expected := customErrors.ErrorUnprocessableContent
		// Act
		result, err := sv.Create(sectionRequest)
		// Assert
		assert.ErrorIs(t, err, expected)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
		mockProductTypeRepo.AssertExpectations(t)
	})
	t.Run("create_not_found: product type not found", func(t *testing.T) {
		// Arrange
		mockRepo, mockProductTypeRepo, sv := setupMocks()
		sectionRequest := models.SectionDocRequest{
			SectionNumber:      "A1",
			CurrentCapacity:    20,
			CurrentTemperature: 10,
			MaximumCapacity:    30,
			MinimumCapacity:    10,
			MinimumTemperature: 5,
			ProductTypeId:      1,
			WarehouseId:        1,
		}
		mockProductTypeRepo.On("GetById", sectionRequest.ProductTypeId).Return(nil, customErrors.ErrorNotFound)
		expected := customErrors.ErrorNotFound
		// Act
		result, err := sv.Create(sectionRequest)
		// Assert
		assert.ErrorIs(t, err, expected)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
		mockProductTypeRepo.AssertExpectations(t)
	})
	t.Run("create_conflict: capacity exceeds limit", func(t *testing.T) {
		// Arrange
		mockRepo, mockProductTypeRepo, sv := setupMocks()
		sectionRequest := models.SectionDocRequest{
			SectionNumber:      "A1",
			CurrentCapacity:    40,
			CurrentTemperature: 10,
			MaximumCapacity:    30,
			MinimumCapacity:    10,
			MinimumTemperature: 5,
			ProductTypeId:      1,
			WarehouseId:        1,
		}
		expected := customErrors.ErrorConflict

		// Act
		result, err := sv.Create(sectionRequest)

		// Assert
		assert.ErrorIs(t, err, expected)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
		mockProductTypeRepo.AssertExpectations(t)
	})

	t.Run("create_conflict: temperature below minimum", func(t *testing.T) {
		// Arrange
		mockRepo, mockProductTypeRepo, sv := setupMocks()
		sectionRequest := models.SectionDocRequest{
			SectionNumber:      "A1",
			CurrentCapacity:    20,
			CurrentTemperature: 2,
			MaximumCapacity:    30,
			MinimumCapacity:    10,
			MinimumTemperature: 5,
			ProductTypeId:      1,
			WarehouseId:        1,
		}
		expected := customErrors.ErrorConflict

		// Act
		result, err := sv.Create(sectionRequest)

		// Assert
		assert.ErrorIs(t, err, expected)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
		mockProductTypeRepo.AssertExpectations(t)
	})

	t.Run("create_not_found: productType is nil without error", func(t *testing.T) {
		// Arrange
		mockRepo, mockProductTypeRepo, sv := setupMocks()
		sectionRequest := models.SectionDocRequest{
			SectionNumber:      "A1",
			CurrentCapacity:    20,
			CurrentTemperature: 10,
			MaximumCapacity:    30,
			MinimumCapacity:    10,
			MinimumTemperature: 5,
			ProductTypeId:      1,
			WarehouseId:        1,
		}
		mockProductTypeRepo.On("GetById", sectionRequest.ProductTypeId).Return(nil, nil)
		expected := customErrors.ErrorNotFound

		// Act
		result, err := sv.Create(sectionRequest)

		// Assert
		assert.ErrorIs(t, err, expected)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
		mockProductTypeRepo.AssertExpectations(t)
	})

}

func TestReadSection(t *testing.T) {
	t.Run("find_all: read sections successfully", func(t *testing.T) {
		// Arrange
		mockRepo, _, sv := setupMocks()
		sections := []models.SectionDoc{
			{
				SectionNumber:      "A1",
				CurrentCapacity:    20,
				CurrentTemperature: 10,
				MaximumCapacity:    30,
				MinimumCapacity:    10,
				MinimumTemperature: 5,
				ProductTypeId:      1,
				WarehouseId:        1,
			},
			{
				SectionNumber:      "A2",
				CurrentCapacity:    20,
				CurrentTemperature: 10,
				MaximumCapacity:    30,
				MinimumCapacity:    10,
				MinimumTemperature: 5,
				ProductTypeId:      1,
				WarehouseId:        1,
			},
		}
		mockRepo.On("GetAll").Return(sections, nil)
		// Act
		res, err := sv.GetAll()
		// Assert
		assert.NoError(t, err)
		assert.Equal(t, sections, res)
		mockRepo.AssertExpectations(t)
	})
	t.Run("find_all_with_error", func(t *testing.T) {
		// Arrange
		mockRepo, _, sv := setupMocks()
		mockRepo.On("GetAll").Return([]models.SectionDoc{}, customErrors.ErrorInternalServerError)
		// Act
		res, err := sv.GetAll()
		// Assert
		assert.ErrorIs(t, err, customErrors.ErrorInternalServerError)
		assert.Empty(t, res)
		mockRepo.AssertExpectations(t)
	})
	t.Run("find_by_id: read section successfully", func(t *testing.T) {
		// Arrange
		mockRepo, _, sv := setupMocks()
		section := &models.SectionDoc{
			SectionNumber:      "A1",
			CurrentCapacity:    20,
			CurrentTemperature: 10,
			MaximumCapacity:    30,
			MinimumCapacity:    10,
			MinimumTemperature: 5,
			ProductTypeId:      1,
			WarehouseId:        1,
		}
		mockRepo.On("GetByID", 1).Return(section, nil)
		// Act
		res, err := sv.GetByID(1)
		// Assert
		assert.NoError(t, err)
		assert.Equal(t, section, res)
		mockRepo.AssertExpectations(t)
	})
	t.Run("find_by_id_non_existent: section not found", func(t *testing.T) {
		// Arrange
		mockRepo, _, sv := setupMocks()
		expectedErr := customErrors.ErrorNotFound
		notExistentId := 9999
		mockRepo.On("GetByID", notExistentId).Return(nil, expectedErr)
		// Act
		res, err := sv.GetByID(notExistentId)
		// Assert
		assert.ErrorAs(t, err, &expectedErr)
		assert.Empty(t, res)
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateSection(t *testing.T) {
	t.Run("update_existent: update section successfully", func(t *testing.T) {
		// Arrange
		mockRepo, mockProductTypeRepo, sv := setupMocks()
		idSection := 1
		sectionNumber := "A1"
		currentCapacity := 30
		currentTemperature := 15.3
		maximumCapacity := 50
		minimumCapacity := 10
		minimumTemperature := 5.0
		productTypeId := 1
		warehouseId := 1

		sectionRequest := models.UpdateSectionDto{
			SectionNumber:      &sectionNumber,
			CurrentCapacity:    &currentCapacity,
			CurrentTemperature: &currentTemperature,
			MaximumCapacity:    &maximumCapacity,
			MinimumCapacity:    &minimumCapacity,
			MinimumTemperature: &minimumTemperature,
			ProductTypeId:      &productTypeId,
			WarehouseId:        &warehouseId,
		}
		sectionToUpdate := &models.SectionDoc{
			Id:                 idSection,
			SectionNumber:      "A1",
			CurrentCapacity:    25,
			CurrentTemperature: 10.2,
			MaximumCapacity:    40,
			MinimumCapacity:    5,
			MinimumTemperature: 2.0,
			ProductTypeId:      1,
			WarehouseId:        1,
		}
		updatedSection := &models.SectionDoc{
			Id:                 idSection,
			SectionNumber:      "A1",
			CurrentCapacity:    30,
			CurrentTemperature: 15.3,
			MaximumCapacity:    50,
			MinimumCapacity:    10,
			MinimumTemperature: 5.0,
			ProductTypeId:      1,
			WarehouseId:        1,
		}

		section := &models.SectionDocRequest{
			SectionNumber:      "A1",
			CurrentCapacity:    30,
			CurrentTemperature: 15.3,
			MaximumCapacity:    50,
			MinimumCapacity:    10,
			MinimumTemperature: 5.0,
			ProductTypeId:      1,
			WarehouseId:        1,
		}
		mockRepo.On("GetByID", idSection).Return(sectionToUpdate, nil)
		mockRepo.On("Update", idSection, section).Return(nil)
		mockProductTypeRepo.On("GetById", productTypeId).Return(&models.ProductTypeDocResponse{}, nil)

		// Act
		res, err := sv.Update(idSection, sectionRequest)
		// Assert
		assert.NoError(t, err)
		assert.Equal(t, updatedSection, res)
		mockRepo.AssertExpectations(t)
		mockProductTypeRepo.AssertExpectations(t)
	})
	t.Run("update_non_existent: section not found", func(t *testing.T) {
		// Arrange
		mockRepo, _, sv := setupMocks()
		idSection := 9999
		sectionNumber := "A1"
		currentCapacity := 30
		currentTemperature := 15.3
		maximumCapacity := 50
		minimumCapacity := 10
		minimumTemperature := 5.0
		productTypeId := 1
		warehouseId := 1
		sectionRequest := models.UpdateSectionDto{
			SectionNumber:      &sectionNumber,
			CurrentCapacity:    &currentCapacity,
			CurrentTemperature: &currentTemperature,
			MaximumCapacity:    &maximumCapacity,
			MinimumCapacity:    &minimumCapacity,
			MinimumTemperature: &minimumTemperature,
			ProductTypeId:      &productTypeId,
			WarehouseId:        &warehouseId,
		}
		mockRepo.On("GetByID", idSection).Return(nil, customErrors.ErrorNotFound)
		// Act
		res, err := sv.Update(idSection, sectionRequest)
		// Assert
		assert.ErrorIs(t, err, customErrors.ErrorNotFound)
		assert.Empty(t, res)
		mockRepo.AssertExpectations(t)
	})
	t.Run("update_conflict: capacity exceeds limit", func(t *testing.T) {
		// Arrange
		mockRepo, mockProductTypeRepo, sv := setupMocks()
		idSection := 1
		sectionNumber := "A1"
		currentCapacity := 60
		currentTemperature := 15.3
		maximumCapacity := 50
		minimumCapacity := 10
		minimumTemperature := 5.0
		productTypeId := 1
		warehouseId := 1
		sectionRequest := models.UpdateSectionDto{
			SectionNumber:      &sectionNumber,
			CurrentCapacity:    &currentCapacity,
			CurrentTemperature: &currentTemperature,
			MaximumCapacity:    &maximumCapacity,
			MinimumCapacity:    &minimumCapacity,
			MinimumTemperature: &minimumTemperature,
			ProductTypeId:      &productTypeId,
			WarehouseId:        &warehouseId,
		}
		sectionToUpdate := &models.SectionDoc{
			Id:                 idSection,
			SectionNumber:      "A1",
			CurrentCapacity:    25,
			CurrentTemperature: 10.2,
			MaximumCapacity:    40,
			MinimumCapacity:    5,
			MinimumTemperature: 2.0,
			ProductTypeId:      1,
			WarehouseId:        1,
		}
		mockRepo.On("GetByID", idSection).Return(sectionToUpdate, nil)
		// Act
		res, err := sv.Update(idSection, sectionRequest)
		// Assert
		assert.ErrorIs(t, err, customErrors.ErrorConflict)
		assert.Empty(t, res)
		mockRepo.AssertExpectations(t)
		mockProductTypeRepo.AssertExpectations(t)
	})
	t.Run("update_conflict: temperature below minimum", func(t *testing.T) {
		// Arrange
		mockRepo, mockProductTypeRepo, sv := setupMocks()
		idSection := 1
		sectionNumber := "A1"
		currentCapacity := 30
		currentTemperature := 2.0
		maximumCapacity := 50
		minimumCapacity := 10
		minimumTemperature := 5.0
		productTypeId := 1
		warehouseId := 1

		sectionRequest := models.UpdateSectionDto{
			SectionNumber:      &sectionNumber,
			CurrentCapacity:    &currentCapacity,
			CurrentTemperature: &currentTemperature,
			MaximumCapacity:    &maximumCapacity,
			MinimumCapacity:    &minimumCapacity,
			MinimumTemperature: &minimumTemperature,
			ProductTypeId:      &productTypeId,
			WarehouseId:        &warehouseId,
		}

		sectionToUpdate := &models.SectionDoc{
			Id:                 idSection,
			SectionNumber:      "A1",
			CurrentCapacity:    25,
			CurrentTemperature: 10.2,
			MaximumCapacity:    40,
			MinimumCapacity:    5,
			MinimumTemperature: 2.0,
			ProductTypeId:      1,
			WarehouseId:        1,
		}

		mockRepo.On("GetByID", idSection).Return(sectionToUpdate, nil)
		mockProductTypeRepo.On("GetById", productTypeId).Return(&models.ProductTypeDocResponse{}, nil)

		// Act
		res, err := sv.Update(idSection, sectionRequest)

		// Assert
		assert.ErrorIs(t, err, customErrors.ErrorConflict)
		assert.Nil(t, res)
		mockRepo.AssertExpectations(t)
	})
	t.Run("update_not_found: productType is nil without error", func(t *testing.T) {
		// Arrange
		mockRepo, mockProductTypeRepo, sv := setupMocks()
		idSection := 1
		sectionNumber := "A1"
		currentCapacity := 30
		currentTemperature := 15.3
		maximumCapacity := 50
		minimumCapacity := 10
		minimumTemperature := 5.0
		productTypeId := 1
		warehouseId := 1

		sectionRequest := models.UpdateSectionDto{
			SectionNumber:      &sectionNumber,
			CurrentCapacity:    &currentCapacity,
			CurrentTemperature: &currentTemperature,
			MaximumCapacity:    &maximumCapacity,
			MinimumCapacity:    &minimumCapacity,
			MinimumTemperature: &minimumTemperature,
			ProductTypeId:      &productTypeId,
			WarehouseId:        &warehouseId,
		}

		sectionToUpdate := &models.SectionDoc{
			Id:                 idSection,
			SectionNumber:      "A1",
			CurrentCapacity:    25,
			CurrentTemperature: 10.2,
			MaximumCapacity:    40,
			MinimumCapacity:    5,
			MinimumTemperature: 2.0,
			ProductTypeId:      1,
			WarehouseId:        1,
		}

		mockRepo.On("GetByID", idSection).Return(sectionToUpdate, nil)
		mockProductTypeRepo.On("GetById", productTypeId).Return(nil, nil)
		// Act
		res, err := sv.Update(idSection, sectionRequest)

		// Assert
		assert.ErrorIs(t, err, customErrors.ErrorNotFound)
		assert.Nil(t, res)
		mockRepo.AssertExpectations(t)
		mockProductTypeRepo.AssertExpectations(t)
	})
	t.Run("update_fail: repository update error", func(t *testing.T) {
		// Arrange
		mockRepo, mockProductTypeRepo, sv := setupMocks()
		idSection := 1
		sectionNumber := "A1"
		currentCapacity := 30
		currentTemperature := 15.3
		maximumCapacity := 50
		minimumCapacity := 10
		minimumTemperature := 5.0
		productTypeId := 1
		warehouseId := 1

		sectionRequest := models.UpdateSectionDto{
			SectionNumber:      &sectionNumber,
			CurrentCapacity:    &currentCapacity,
			CurrentTemperature: &currentTemperature,
			MaximumCapacity:    &maximumCapacity,
			MinimumCapacity:    &minimumCapacity,
			MinimumTemperature: &minimumTemperature,
			ProductTypeId:      &productTypeId,
			WarehouseId:        &warehouseId,
		}

		sectionToUpdate := &models.SectionDoc{
			Id:                 idSection,
			SectionNumber:      "A1",
			CurrentCapacity:    25,
			CurrentTemperature: 10.2,
			MaximumCapacity:    40,
			MinimumCapacity:    5,
			MinimumTemperature: 2.0,
			ProductTypeId:      1,
			WarehouseId:        1,
		}

		section := &models.SectionDocRequest{
			SectionNumber:      "A1",
			CurrentCapacity:    30,
			CurrentTemperature: 15.3,
			MaximumCapacity:    50,
			MinimumCapacity:    10,
			MinimumTemperature: 5.0,
			ProductTypeId:      1,
			WarehouseId:        1,
		}

		mockRepo.On("GetByID", idSection).Return(sectionToUpdate, nil)
		mockProductTypeRepo.On("GetById", productTypeId).Return(&models.ProductTypeDocResponse{}, nil)
		mockRepo.On("Update", idSection, section).Return(customErrors.ErrorInternalServerError)

		// Act
		res, err := sv.Update(idSection, sectionRequest)

		// Assert
		assert.Error(t, err)
		assert.ErrorIs(t, err, customErrors.ErrorInternalServerError)
		assert.Nil(t, res)
		mockRepo.AssertExpectations(t)
	})

}

func TestDeleteSection(t *testing.T) {
	t.Run("delete_ok: delete section successfully", func(t *testing.T) {
		// Arrange
		mockRepo, _, sv := setupMocks()
		section := &models.SectionDoc{
			Id:                 1,
			SectionNumber:      "A1",
			CurrentCapacity:    20,
			CurrentTemperature: 10,
			MaximumCapacity:    30,
			MinimumCapacity:    10,
			MinimumTemperature: 5,
			ProductTypeId:      1,
			WarehouseId:        1,
		}
		mockRepo.On("GetByID", 1).Return(section, nil)
		mockRepo.On("Delete", 1).Return(nil)
		// Act
		err := sv.Delete(1)
		// Assert
		assert.NoError(t, err)
		mockRepo.AssertCalled(t, "Delete", 1)
		mockRepo.AssertCalled(t, "GetByID", 1)
		mockRepo.AssertExpectations(t)
	})
	t.Run("delete_non_existent: section not found", func(t *testing.T) {
		// Arrange
		mockRepo, _, sv := setupMocks()
		nonExistentID := 9999
		mockRepo.On("GetByID", nonExistentID).Return(nil, customErrors.ErrorNotFound)
		// Act
		err := sv.Delete(nonExistentID)
		// Assert
		assert.ErrorIs(t, err, customErrors.ErrorNotFound)
		mockRepo.AssertExpectations(t)
	})
}

func TestRecoverSection(t *testing.T) {
	t.Run("recover_ok: recover section successfully", func(t *testing.T) {
		// Arrange
		mockRepo, _, sv := setupMocks()
		idSection := 1

		recoveredSection := &models.SectionDoc{
			Id:                 idSection,
			SectionNumber:      "A1",
			CurrentCapacity:    30,
			CurrentTemperature: 15.3,
			MaximumCapacity:    50,
			MinimumCapacity:    10,
			MinimumTemperature: 5.0,
			ProductTypeId:      1,
			WarehouseId:        1,
		}

		mockRepo.On("Recover", idSection).Return(nil)
		mockRepo.On("GetByID", idSection).Return(recoveredSection, nil)

		// Act
		res, err := sv.Recover(idSection)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, *recoveredSection, res)
		mockRepo.AssertExpectations(t)
	})
	t.Run("recover_fail: repository recover error", func(t *testing.T) {
		// Arrange
		mockRepo, _, sv := setupMocks()
		idSection := 1
		mockRepo.On("Recover", idSection).Return(customErrors.ErrorInternalServerError)
		// Act
		res, err := sv.Recover(idSection)
		// Assert
		assert.Error(t, err)
		assert.ErrorIs(t, err, customErrors.ErrorInternalServerError)
		assert.Equal(t, models.SectionDoc{}, res)
		mockRepo.AssertExpectations(t)

	})
	t.Run("recover_fail_id: error retrieving section", func(t *testing.T) {
		// Arrange
		mockRepo, _, sv := setupMocks()
		idSection := 9999
		mockRepo.On("Recover", idSection).Return(nil)
		mockRepo.On("GetByID", idSection).Return(nil, customErrors.ErrorNotFound)
		// Act
		res, err := sv.Recover(idSection)
		// Assert
		assert.Error(t, err)
		assert.ErrorIs(t, err, customErrors.ErrorNotFound)
		assert.Equal(t, models.SectionDoc{}, res)
		mockRepo.AssertExpectations(t)
	})
}

func TestReportsSection(t *testing.T) {
	t.Run("reports_ok: get reports successfully", func(t *testing.T) {
		// Arrange
		mockRepo, _, sv := setupMocks()
		sectionId := 1

		expectedReports := []models.SectionReport{
			{
				SectionId:     sectionId,
				SectionNumber: "A1",
				ProductsCount: 10,
			},
			{
				SectionId:     sectionId,
				SectionNumber: "A2",
				ProductsCount: 20,
			},
		}
		mockRepo.On("GetSectionReports", sectionId).Return(expectedReports, nil)

		// Act
		res, err := sv.GetSectionReports(sectionId)
		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedReports, res)
		mockRepo.AssertExpectations(t)
	})
	t.Run("reports_not_found: section reports not found", func(t *testing.T) {
		// Arrange
		mockRepo, _, sv := setupMocks()
		sectionId := 1
		mockRepo.On("GetSectionReports", sectionId).Return(nil, customErrors.ErrorNotFound)
		// Act
		res, err := sv.GetSectionReports(sectionId)
		// Assert
		assert.ErrorIs(t, err, customErrors.ErrorNotFound)
		assert.Empty(t, res)
		mockRepo.AssertExpectations(t)
	})
}
