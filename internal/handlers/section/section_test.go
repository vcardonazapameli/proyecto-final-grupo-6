package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	handlers "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers/section"
	services "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/section"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestSectionCreate(t *testing.T) {
	t.Run("create_ok", func(t *testing.T) {
		//Arrange
		serviceMock := new(services.SectionServiceMock)
		handler := handlers.NewSectionHandler(serviceMock)
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
		sectionResponse := &models.SectionDoc{
			SectionNumber:      "A1",
			CurrentCapacity:    20,
			CurrentTemperature: 10,
			MaximumCapacity:    30,
			MinimumCapacity:    10,
			MinimumTemperature: 5,
			ProductTypeId:      1,
			WarehouseId:        1,
		}
		serviceMock.On("Create", sectionRequest).Return(sectionResponse, nil)
		body, err := json.Marshal(sectionRequest)

		router := chi.NewRouter()
		router.Post("/section", handler.Create())

		//Act
		req, res := httptest.NewRequest(http.MethodPost, "/section", bytes.NewReader(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(res, req)
		//Assert
		expectedResponse := map[string]interface{}{
			"message": "Success",
			"data":    sectionResponse,
		}
		expectedBody, _ := json.Marshal(expectedResponse)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, res.Code)
		assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
		assert.JSONEq(t, string(expectedBody), res.Body.String())
		serviceMock.AssertExpectations(t)
	})
	t.Run("create_conflict", func(t *testing.T) {
		//arrnge
		serviceMock := new(services.SectionServiceMock)
		handler := handlers.NewSectionHandler(serviceMock)
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
		serviceMock.On("Create", sectionRequest).Return(nil, customErrors.ErrorConflict)
		router := chi.NewRouter()
		router.Post("/section", handler.Create())
		body, _ := json.Marshal(sectionRequest)
		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/section", bytes.NewReader(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(res, req)
		// Assert
		assert.Equal(t, http.StatusConflict, res.Code)
		serviceMock.AssertExpectations(t)
	})
	t.Run("create_bad_request", func(t *testing.T) {
		// Arrange
		serviceMock := new(services.SectionServiceMock)
		handler := handlers.NewSectionHandler(serviceMock)
		router := chi.NewRouter()
		router.Post("/section", handler.Create())
		body := []byte{}

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/section", bytes.NewReader(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(res, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, res.Code)
		serviceMock.AssertExpectations(t)
	})

}

func TestSectionRead(t *testing.T) {
	t.Run("find_all", func(t *testing.T) {
		// Arrange
		serviceMock := new(services.SectionServiceMock)
		handler := handlers.NewSectionHandler(serviceMock)
		sectionResponse := []models.SectionDoc{
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
		serviceMock.On("GetAll").Return(sectionResponse, nil)
		router := chi.NewRouter()
		router.Get("/section", handler.GetAll())
		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/section", nil), httptest.NewRecorder()
		router.ServeHTTP(res, req)
		// Assert
		expectedResponse := map[string]interface{}{
			"message": "Success",
			"data":    sectionResponse,
		}
		expectedBody, _ := json.Marshal(expectedResponse)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
		assert.JSONEq(t, string(expectedBody), res.Body.String())
		serviceMock.AssertExpectations(t)

	})
	t.Run("find_all_not_found", func(t *testing.T) {
		// Arrange
		serviceMock := new(services.SectionServiceMock)
		handler := handlers.NewSectionHandler(serviceMock)
		serviceMock.On("GetAll").Return(nil, customErrors.ErrorNotFound)
		router := chi.NewRouter()
		router.Get("/section", handler.GetAll())
		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/section", nil), httptest.NewRecorder()
		router.ServeHTTP(res, req)
		// Assert
		assert.Equal(t, http.StatusNotFound, res.Code)
		serviceMock.AssertExpectations(t)
	})
	t.Run("find_by_id", func(t *testing.T) {
		// Arrange
		serviceMock := new(services.SectionServiceMock)
		handler := handlers.NewSectionHandler(serviceMock)
		sectionResponse := &models.SectionDoc{
			SectionNumber:      "A1",
			CurrentCapacity:    20,
			CurrentTemperature: 10,
			MaximumCapacity:    30,
			MinimumCapacity:    10,
			MinimumTemperature: 5,
			ProductTypeId:      1,
			WarehouseId:        1,
		}
		serviceMock.On("GetByID", 1).Return(sectionResponse, nil)
		router := chi.NewRouter()
		router.Get("/section/{id}", handler.GetByID())
		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/section/1", nil), httptest.NewRecorder()
		router.ServeHTTP(res, req)
		// Assert
		expectedResponse := map[string]interface{}{
			"message": "Success",
			"data":    sectionResponse,
		}
		expectedBody, _ := json.Marshal(expectedResponse)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
		assert.JSONEq(t, string(expectedBody), res.Body.String())
		serviceMock.AssertExpectations(t)
	})
	t.Run("find_by_id_bad_request", func(t *testing.T) {
		// Arrange
		serviceMock := new(services.SectionServiceMock)
		handler := handlers.NewSectionHandler(serviceMock)
		router := chi.NewRouter()
		router.Get("/section/{id}", handler.GetByID())
		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/section/abc", nil), httptest.NewRecorder()
		router.ServeHTTP(res, req)
		// Assert
		assert.Equal(t, http.StatusBadRequest, res.Code)
		serviceMock.AssertExpectations(t)
	})
	t.Run("find_by_id_not_found", func(t *testing.T) {
		// Arrange
		serviceMock := new(services.SectionServiceMock)
		handler := handlers.NewSectionHandler(serviceMock)
		serviceMock.On("GetByID", 1).Return(nil, customErrors.ErrorNotFound)
		router := chi.NewRouter()
		router.Get("/section/{id}", handler.GetByID())
		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/section/1", nil), httptest.NewRecorder()
		router.ServeHTTP(res, req)
		// Assert
		assert.Equal(t, http.StatusNotFound, res.Code)
		serviceMock.AssertExpectations(t)
	})

}

func TestSectionUpdate(t *testing.T) {
	t.Run("update_ok", func(t *testing.T) {
		// Arrange
		serviceMock := new(services.SectionServiceMock)
		handler := handlers.NewSectionHandler(serviceMock)
		sectionDoc := models.SectionDoc{
			CurrentCapacity:    20,
			CurrentTemperature: 10,
		}
		sectionRequest := models.UpdateSectionDto{
			SectionNumber:      nil,
			CurrentCapacity:    &sectionDoc.CurrentCapacity,
			CurrentTemperature: &sectionDoc.CurrentTemperature,
			MaximumCapacity:    nil,
			MinimumCapacity:    nil,
			MinimumTemperature: nil,
			ProductTypeId:      nil,
			WarehouseId:        nil,
		}
		sectionResponse := &models.SectionDoc{
			SectionNumber:      "A1",
			CurrentCapacity:    20,
			CurrentTemperature: 10,
			MaximumCapacity:    30,
			MinimumCapacity:    10,
			MinimumTemperature: 5,
			ProductTypeId:      1,
			WarehouseId:        1,
		}
		serviceMock.On("Update", 1, sectionRequest).Return(sectionResponse, nil)
		body, _ := json.Marshal(sectionRequest)
		router := chi.NewRouter()
		router.Patch("/section/{id}", handler.Update())
		// Act
		req, res := httptest.NewRequest(http.MethodPatch, "/section/1", bytes.NewReader(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(res, req)
		// Assert
		expectedResponse := map[string]interface{}{
			"message": "Success",
			"data":    sectionResponse,
		}
		expectedBody, _ := json.Marshal(expectedResponse)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
		assert.JSONEq(t, string(expectedBody), res.Body.String())
		serviceMock.AssertExpectations(t)
	})
	t.Run("update_bad_request_invalid_id", func(t *testing.T) {
		// Arrange
		serviceMock := new(services.SectionServiceMock)
		handler := handlers.NewSectionHandler(serviceMock)
		router := chi.NewRouter()
		router.Patch("/section/{id}", handler.Update())
		// Act
		req, res := httptest.NewRequest(http.MethodPatch, "/section/abc", nil), httptest.NewRecorder()
		router.ServeHTTP(res, req)
		// Assert
		assert.Equal(t, http.StatusBadRequest, res.Code)
		serviceMock.AssertExpectations(t)
	})
	t.Run("update_bad_request_invalid_body", func(t *testing.T) {
		// Arrange
		serviceMock := new(services.SectionServiceMock)
		handler := handlers.NewSectionHandler(serviceMock)
		router := chi.NewRouter()
		router.Patch("/section/{id}", handler.Update())
		// Act
		req, res := httptest.NewRequest(http.MethodPatch, "/section/1", nil), httptest.NewRecorder()
		router.ServeHTTP(res, req)
		// Assert
		assert.Equal(t, http.StatusBadRequest, res.Code)
		serviceMock.AssertExpectations(t)

	})
	t.Run("update_not_found", func(t *testing.T) {
		// Arrange
		serviceMock := new(services.SectionServiceMock)
		handler := handlers.NewSectionHandler(serviceMock)
		sectionRequest := models.UpdateSectionDto{
			SectionNumber:      nil,
			CurrentCapacity:    nil,
			CurrentTemperature: nil,
			MaximumCapacity:    nil,
			MinimumCapacity:    nil,
			MinimumTemperature: nil,
			ProductTypeId:      nil,
			WarehouseId:        nil,
		}
		serviceMock.On("Update", 1, sectionRequest).Return(nil, customErrors.ErrorNotFound)
		body, _ := json.Marshal(sectionRequest)
		router := chi.NewRouter()
		router.Patch("/section/{id}", handler.Update())
		// Act
		req, res := httptest.NewRequest(http.MethodPatch, "/section/1", bytes.NewReader(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(res, req)
		// Assert
		assert.Equal(t, http.StatusNotFound, res.Code)
		serviceMock.AssertExpectations(t)
	})
}

func TestSectionDelete(t *testing.T) {
	t.Run("delete_ok", func(t *testing.T) {
		// Arrange
		serviceMock := new(services.SectionServiceMock)
		handler := handlers.NewSectionHandler(serviceMock)
		serviceMock.On("Delete", 1).Return(nil)
		router := chi.NewRouter()
		router.Delete("/section/{id}", handler.Delete())
		// Act
		req, res := httptest.NewRequest(http.MethodDelete, "/section/1", nil), httptest.NewRecorder()
		router.ServeHTTP(res, req)
		// Assert
		assert.Equal(t, http.StatusNoContent, res.Code)
		serviceMock.AssertExpectations(t)
	})
	t.Run("delete_bad_request_invalid_id", func(t *testing.T) {
		// Arrange
		serviceMock := new(services.SectionServiceMock)
		handler := handlers.NewSectionHandler(serviceMock)
		router := chi.NewRouter()
		router.Delete("/section/{id}", handler.Delete())
		// Act
		req, res := httptest.NewRequest(http.MethodDelete, "/section/abc", nil), httptest.NewRecorder()
		router.ServeHTTP(res, req)
		// Assert
		assert.Equal(t, http.StatusBadRequest, res.Code)
		serviceMock.AssertExpectations(t)
	})
	t.Run("delete_not_found", func(t *testing.T) {
		// Arrange
		serviceMock := new(services.SectionServiceMock)
		handler := handlers.NewSectionHandler(serviceMock)
		serviceMock.On("Delete", 1).Return(customErrors.ErrorNotFound)
		router := chi.NewRouter()
		router.Delete("/section/{id}", handler.Delete())
		// Act
		req, res := httptest.NewRequest(http.MethodDelete, "/section/1", nil), httptest.NewRecorder()
		router.ServeHTTP(res, req)
		// Assert
		assert.Equal(t, http.StatusNotFound, res.Code)
		serviceMock.AssertExpectations(t)
	})
}

func TestSectionReports(t *testing.T) {
	t.Run("get_section_reports_ok", func(t *testing.T) {
		serviceMock := new(services.SectionServiceMock)
		handler := handlers.NewSectionHandler(serviceMock)
		sectionReports := []models.SectionReport{
			{
				SectionId:     1,
				SectionNumber: "A1",
				ProductsCount: 10,
			},
			{
				SectionId:     1,
				SectionNumber: "A2",
				ProductsCount: 20,
			},
		}
		serviceMock.On("GetSectionReports", 0).Return(sectionReports, nil)
		router := chi.NewRouter()
		router.Get("/sections/reportProducts", handler.GetSectionReports())
		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/sections/reportProducts", nil), httptest.NewRecorder()
		router.ServeHTTP(res, req)
		// Assert
		expectedResponse := map[string]interface{}{
			"message": "Success",
			"data":    sectionReports,
		}
		expectedBody, _ := json.Marshal(expectedResponse)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
		assert.JSONEq(t, string(expectedBody), res.Body.String())
		serviceMock.AssertExpectations(t)
	})

	t.Run("get_section_reports_with_id_ok", func(t *testing.T) {
		// Arrange
		serviceMock := new(services.SectionServiceMock)
		handler := handlers.NewSectionHandler(serviceMock)
		sectionReports := []models.SectionReport{
			{
				SectionId:     1,
				SectionNumber: "A1",
				ProductsCount: 10,
			},
		}
		serviceMock.On("GetSectionReports", 1).Return(sectionReports, nil)
		router := chi.NewRouter()
		router.Get("/sections/reportProducts", handler.GetSectionReports())
		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/sections/reportProducts?id=1", nil), httptest.NewRecorder()
		router.ServeHTTP(res, req)
		// Assert
		expectedResponse := map[string]interface{}{
			"message": "Success",
			"data":    sectionReports,
		}
		expectedBody, _ := json.Marshal(expectedResponse)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
		assert.JSONEq(t, string(expectedBody), res.Body.String())
		serviceMock.AssertExpectations(t)
	})
	t.Run("get_section_reports_bad_request_invalid_id", func(t *testing.T) {
		// Arrange
		serviceMock := new(services.SectionServiceMock)
		handler := handlers.NewSectionHandler(serviceMock)
		router := chi.NewRouter()
		router.Get("/sections/reportProducts", handler.GetSectionReports())
		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/sections/reportProducts?id=abc", nil), httptest.NewRecorder()
		router.ServeHTTP(res, req)
		// Assert
		assert.Equal(t, http.StatusBadRequest, res.Code)
		serviceMock.AssertExpectations(t)
	})
	t.Run("get_section_reports_not_found", func(t *testing.T) {
		// Arrange
		serviceMock := new(services.SectionServiceMock)
		handler := handlers.NewSectionHandler(serviceMock)
		serviceMock.On("GetSectionReports", 1).Return(nil, customErrors.ErrorNotFound)
		router := chi.NewRouter()
		router.Get("/sections/reportProducts", handler.GetSectionReports())
		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/sections/reportProducts?id=1", nil), httptest.NewRecorder()
		router.ServeHTTP(res, req)
		// Assert
		assert.Equal(t, http.StatusNotFound, res.Code)
		serviceMock.AssertExpectations(t)
	})

}
