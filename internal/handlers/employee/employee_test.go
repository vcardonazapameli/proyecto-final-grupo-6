package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/go-chi/chi/v5"

	handlers "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers/employee"
	sv "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/employee"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
)

func TestEmployeeHandler(t *testing.T) {

	t.Run("Create_ok", func(t *testing.T) {
		// Arrange
		mockService := new(sv.EmployeeServiceMock)
		handler := handlers.NewEmployeeHandler(mockService)
		rt := chi.NewRouter()
		rt.Post("/employees", handler.Create())

		newEmployee := models.RequestEmployee{
			CardNumberID: "123456",
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  1,
		}

		expectedEmployee := models.EmployeeDoc{
			Id:           1,
			CardNumberID: "123456",
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  1,
		}

		expectedResponse := map[string]interface{}{
			"message": "Success",
			"data":    expectedEmployee,
		}

		expectedResponseJSON, _ := json.Marshal(expectedResponse)
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		mockService.On("Create", newEmployee).Return(&expectedEmployee, nil)

		// Convertir el objeto newEmployee a JSON
		newEmployeeJSON, _ := json.Marshal(newEmployee)

		// Crear la petición con el body
		req := httptest.NewRequest(http.MethodPost, "/employees", bytes.NewBuffer(newEmployeeJSON))
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()

		// Act
		rt.ServeHTTP(res, req)

		// Assert
		require.Equal(t, http.StatusCreated, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
		assert.JSONEq(t, string(expectedResponseJSON), res.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("Create_fail", func(t *testing.T) {
		// Arrange
		mockService := new(sv.EmployeeServiceMock)
		handler := handlers.NewEmployeeHandler(mockService)
		rt := chi.NewRouter()
		rt.Post("/employees", handler.Create())

		newEmployee := models.RequestEmployee{
			CardNumberID: "",
			FirstName:    "",
			LastName:     "",
			WarehouseID:  0,
		}

		expectedResponse := map[string]interface{}{
			"status_code": http.StatusUnprocessableEntity,
			"message":     customErrors.ErrorUnprocessableContent.Error(),
		}

		expectedResponseJSON, _ := json.Marshal(expectedResponse)
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		mockService.On("Create", newEmployee).Return(nil, customErrors.ErrorUnprocessableContent)

		// Convertir el objeto newEmployee a JSON
		newEmployeeJSON, _ := json.Marshal(newEmployee)

		// Crear la petición con el body
		req := httptest.NewRequest(http.MethodPost, "/employees", bytes.NewBuffer(newEmployeeJSON))
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()

		// Act
		rt.ServeHTTP(res, req)

		// Assert
		require.Equal(t, http.StatusUnprocessableEntity, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
		assert.JSONEq(t, string(expectedResponseJSON), res.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("Create_conflict", func(t *testing.T) {
		// Arrange
		mockService := new(sv.EmployeeServiceMock)
		handler := handlers.NewEmployeeHandler(mockService)
		rt := chi.NewRouter()
		rt.Post("/employees", handler.Create())

		newEmployee := models.RequestEmployee{
			CardNumberID: "123456",
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  1,
		}

		expectedResponse := map[string]interface{}{
			"status_code": http.StatusConflict,
			"message":     customErrors.ErrorConflict.Error(),
		}

		expectedResponseJSON, _ := json.Marshal(expectedResponse)
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		mockService.On("Create", newEmployee).Return(nil, customErrors.ErrorConflict)

		// Convertir el objeto newEmployee a JSON
		newEmployeeJSON, _ := json.Marshal(newEmployee)

		// Crear la petición con el body
		req := httptest.NewRequest(http.MethodPost, "/employees", bytes.NewBuffer(newEmployeeJSON))
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()

		// Act
		rt.ServeHTTP(res, req)

		// Assert
		require.Equal(t, http.StatusConflict, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
		assert.JSONEq(t, string(expectedResponseJSON), res.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("Create_decode_error", func(t *testing.T) {
		// Arrange
		mockService := new(sv.EmployeeServiceMock)
		handler := handlers.NewEmployeeHandler(mockService)
		rt := chi.NewRouter()
		rt.Post("/employees", handler.Create())

		expectedResponse := map[string]interface{}{
			"status_code": http.StatusBadRequest,
			"message":     customErrors.ErrorBadRequest.Error(),
		}

		expectedResponseJSON, _ := json.Marshal(expectedResponse)
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		// Crear la petición con el body invalido
		req := httptest.NewRequest(http.MethodPost, "/employees", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()

		// Act
		rt.ServeHTTP(res, req)

		// Assert
		require.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
		assert.JSONEq(t, string(expectedResponseJSON), res.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("Find_all_ok", func(t *testing.T) {

		mockService := new(sv.EmployeeServiceMock)
		handler := handlers.NewEmployeeHandler(mockService)
		rt := chi.NewRouter()
		rt.Get("/employees", handler.GetAll())

		// Arrange
		expectedEmployees := []models.EmployeeDoc{
			{Id: 1, CardNumberID: "123456", FirstName: "John", LastName: "Doe", WarehouseID: 1},
			{Id: 2, CardNumberID: "654321", FirstName: "Sara", LastName: "Ramirez", WarehouseID: 2},
		}

		expectedResponse := map[string]interface{}{
			"message": "Success",
			"data":    expectedEmployees,
		}

		expectedResponseJSON, _ := json.Marshal(expectedResponse)
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		mockService.On("GetAll").Return(expectedEmployees, nil)

		req, res := httptest.NewRequest(http.MethodGet, "/employees", nil), httptest.NewRecorder()

		//Act
		rt.ServeHTTP(res, req)

		// Assert
		require.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
		assert.JSONEq(t, string(expectedResponseJSON), res.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("Find_all_fail", func(t *testing.T) {

		mockService := new(sv.EmployeeServiceMock)
		handler := handlers.NewEmployeeHandler(mockService)
		rt := chi.NewRouter()
		rt.Get("/employees", handler.GetAll())

		expectedResponse := map[string]interface{}{
			"status_code": http.StatusInternalServerError,
			"message":     "Internal Server Error",
		}

		expectedResponseJSON, _ := json.Marshal(expectedResponse)
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		mockService.On("GetAll").Return(nil, customErrors.ErrorInternalServerError)

		req, res := httptest.NewRequest(http.MethodGet, "/employees", nil), httptest.NewRecorder()

		//Act
		rt.ServeHTTP(res, req)

		// Assert
		require.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
		assert.JSONEq(t, string(expectedResponseJSON), res.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("Find_by_id_existent", func(t *testing.T) {

		mockService := new(sv.EmployeeServiceMock)
		handler := handlers.NewEmployeeHandler(mockService)
		rt := chi.NewRouter()
		rt.Get("/employees/{id}", handler.GetById())

		// Arrange
		expectedEmployee := models.EmployeeDoc{Id: 1, CardNumberID: "123456", FirstName: "John", LastName: "Doe", WarehouseID: 1}

		expectedResponse := map[string]interface{}{
			"message": "Success",
			"data":    expectedEmployee,
		}

		expectedResponseJSON, _ := json.Marshal(expectedResponse)
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		mockService.On("GetById", 1).Return(&expectedEmployee, nil)
		req, res := httptest.NewRequest(http.MethodGet, "/employees/1", nil), httptest.NewRecorder()

		//Act
		rt.ServeHTTP(res, req)

		// Assert
		require.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
		assert.JSONEq(t, string(expectedResponseJSON), res.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("Find_by_id_non_existent", func(t *testing.T) {

		mockService := new(sv.EmployeeServiceMock)
		handler := handlers.NewEmployeeHandler(mockService)
		rt := chi.NewRouter()
		rt.Get("/employees/{id}", handler.GetById())

		expectedResponse := map[string]interface{}{
			"status_code": http.StatusNotFound,
			"message":     customErrors.ErrorNotFound.Error(),
		}
		expectedResponseJSON, _ := json.Marshal(expectedResponse)
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		mockService.On("GetById", 12312).Return(nil, customErrors.ErrorNotFound)
		req, res := httptest.NewRequest(http.MethodGet, "/employees/12312", nil), httptest.NewRecorder()

		//Act
		rt.ServeHTTP(res, req)

		// Assert
		require.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
		assert.JSONEq(t, string(expectedResponseJSON), res.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("Find_by_id_bad_request", func(t *testing.T) {

		mockService := new(sv.EmployeeServiceMock)
		handler := handlers.NewEmployeeHandler(mockService)
		rt := chi.NewRouter()
		rt.Get("/employees/{id}", handler.GetById())

		expectedResponse := map[string]interface{}{
			"status_code": http.StatusBadRequest,
			"message":     customErrors.ErrorBadRequest.Error(),
		}
		expectedResponseJSON, _ := json.Marshal(expectedResponse)
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		req, res := httptest.NewRequest(http.MethodGet, "/employees/asd", nil), httptest.NewRecorder()

		//Act
		rt.ServeHTTP(res, req)

		// Assert
		require.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
		assert.JSONEq(t, string(expectedResponseJSON), res.Body.String())
	})

	t.Run("Update_ok", func(t *testing.T) {
		// Arrange
		mockService := new(sv.EmployeeServiceMock)
		handler := handlers.NewEmployeeHandler(mockService)
		rt := chi.NewRouter()
		rt.Put("/employees/{id}", handler.Update())

		var idEmployee int = 1
		var cardNumberID string = "123456"
		var firstname string = "John"
		var lastname string = "Doe"
		var warehouseID int = 1

		updEmployee := models.UpdateEmployee{
			CardNumberID: &cardNumberID,
			FirstName:    &firstname,
			LastName:     &lastname,
			WarehouseID:  &warehouseID,
		}

		expectedEmployee := models.EmployeeDoc{
			Id:           idEmployee,
			CardNumberID: cardNumberID,
			FirstName:    firstname,
			LastName:     lastname,
			WarehouseID:  warehouseID,
		}

		expectedResponse := map[string]interface{}{
			"message": "Success",
			"data":    expectedEmployee,
		}

		expectedResponseJSON, _ := json.Marshal(expectedResponse)
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		mockService.On("Update", idEmployee, updEmployee).Return(&expectedEmployee, nil)

		// Convertir el objeto updEmployee a JSON
		updEmployeeJSON, _ := json.Marshal(updEmployee)

		// Crear la petición con el body
		req := httptest.NewRequest(http.MethodPut, "/employees/1", bytes.NewBuffer(updEmployeeJSON))
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()

		// Act
		rt.ServeHTTP(res, req)

		// Assert
		require.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
		assert.JSONEq(t, string(expectedResponseJSON), res.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("Update_non_existent", func(t *testing.T) {
		// Arrange
		mockService := new(sv.EmployeeServiceMock)
		handler := handlers.NewEmployeeHandler(mockService)
		rt := chi.NewRouter()
		rt.Put("/employees/{id}", handler.Update())

		var idEmployee int = 1234
		var cardNumberID string = "123456"
		var firstname string = "John"
		var lastname string = "Doe"
		var warehouseID int = 1

		updEmployee := models.UpdateEmployee{
			CardNumberID: &cardNumberID,
			FirstName:    &firstname,
			LastName:     &lastname,
			WarehouseID:  &warehouseID,
		}

		expectedResponse := map[string]interface{}{
			"status_code": http.StatusNotFound,
			"message":     customErrors.ErrorNotFound.Error(),
		}

		expectedResponseJSON, _ := json.Marshal(expectedResponse)
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		mockService.On("Update", idEmployee, updEmployee).Return(nil, customErrors.ErrorNotFound)

		// Convertir el objeto updEmployee a JSON
		updEmployeeJSON, _ := json.Marshal(updEmployee)

		// Crear la petición con el body
		req := httptest.NewRequest(http.MethodPut, "/employees/1234", bytes.NewBuffer(updEmployeeJSON))
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()

		// Act
		rt.ServeHTTP(res, req)

		// Assert
		require.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
		assert.JSONEq(t, string(expectedResponseJSON), res.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("Update_invalid_id", func(t *testing.T) {
		// Arrange
		mockService := new(sv.EmployeeServiceMock)
		handler := handlers.NewEmployeeHandler(mockService)
		rt := chi.NewRouter()
		rt.Put("/employees/{id}", handler.Update())

		var cardNumberID string = "123456"
		var firstname string = "John"
		var lastname string = "Doe"
		var warehouseID int = 1

		updEmployee := models.UpdateEmployee{
			CardNumberID: &cardNumberID,
			FirstName:    &firstname,
			LastName:     &lastname,
			WarehouseID:  &warehouseID,
		}

		expectedResponse := map[string]interface{}{
			"status_code": http.StatusBadRequest,
			"message":     customErrors.ErrorBadRequest.Error(),
		}

		expectedResponseJSON, _ := json.Marshal(expectedResponse)
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		// Convertir el objeto updEmployee a JSON
		updEmployeeJSON, _ := json.Marshal(updEmployee)

		// Crear la petición con el body y un id invalido
		req := httptest.NewRequest(http.MethodPut, "/employees/asd", bytes.NewBuffer(updEmployeeJSON))
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()

		// Act
		rt.ServeHTTP(res, req)

		// Assert
		require.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
		assert.JSONEq(t, string(expectedResponseJSON), res.Body.String())
	})

	t.Run("Update_invalid_body", func(t *testing.T) {
		// Arrange
		mockService := new(sv.EmployeeServiceMock)
		handler := handlers.NewEmployeeHandler(mockService)
		rt := chi.NewRouter()
		rt.Put("/employees/{id}", handler.Update())

		expectedResponse := map[string]interface{}{
			"status_code": http.StatusBadRequest,
			"message":     customErrors.ErrorBadRequest.Error(),
		}

		expectedResponseJSON, _ := json.Marshal(expectedResponse)
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		// Crear la petición con el body invalido
		req := httptest.NewRequest(http.MethodPut, "/employees/1", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()

		// Act
		rt.ServeHTTP(res, req)

		// Assert
		require.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
		assert.JSONEq(t, string(expectedResponseJSON), res.Body.String())
	})

	t.Run("Delete_ok", func(t *testing.T) {
		// Arrange
		mockService := new(sv.EmployeeServiceMock)
		handler := handlers.NewEmployeeHandler(mockService)
		rt := chi.NewRouter()
		rt.Delete("/employees/{id}", handler.Delete())

		mockService.On("Delete", 1).Return(nil)

		// Crear la petición con el body
		req := httptest.NewRequest(http.MethodDelete, "/employees/1", nil)
		res := httptest.NewRecorder()

		// Act
		rt.ServeHTTP(res, req)

		// Assert
		require.Equal(t, http.StatusNoContent, res.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("Delete_non_existent", func(t *testing.T) {
		// Arrange
		mockService := new(sv.EmployeeServiceMock)
		handler := handlers.NewEmployeeHandler(mockService)
		rt := chi.NewRouter()
		rt.Delete("/employees/{id}", handler.Delete())

		expectedResponse := map[string]interface{}{
			"status_code": http.StatusNotFound,
			"message":     customErrors.ErrorNotFound.Error(),
		}

		expectedResponseJSON, _ := json.Marshal(expectedResponse)
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		mockService.On("Delete", 123).Return(customErrors.ErrorNotFound)

		// Crear la petición con el body
		req := httptest.NewRequest(http.MethodDelete, "/employees/123", nil)
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()

		// Act
		rt.ServeHTTP(res, req)

		// Assert
		require.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
		assert.JSONEq(t, string(expectedResponseJSON), res.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("Delete_invalid_id", func(t *testing.T) {
		// Arrange
		mockService := new(sv.EmployeeServiceMock)
		handler := handlers.NewEmployeeHandler(mockService)
		rt := chi.NewRouter()
		rt.Delete("/employees/{id}", handler.Delete())

		expectedResponse := map[string]interface{}{
			"status_code": http.StatusBadRequest,
			"message":     customErrors.ErrorBadRequest.Error(),
		}

		expectedResponseJSON, _ := json.Marshal(expectedResponse)
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		// Crear la petición con el body y un id invalido
		req := httptest.NewRequest(http.MethodDelete, "/employees/asd", nil)
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()

		// Act
		rt.ServeHTTP(res, req)

		// Assert
		require.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
		assert.JSONEq(t, string(expectedResponseJSON), res.Body.String())
	})

	t.Run("Get_report_inboundOrders_ok", func(t *testing.T) {
		// Arrange
		mockService := new(sv.EmployeeServiceMock)
		handler := handlers.NewEmployeeHandler(mockService)
		rt := chi.NewRouter()
		rt.Get("/employees/reportInboundOrders", handler.GetReportInboundOrders())

		expectedReport := []models.EmployeeWithOrders{
			{Id: 1, CardNumberID: "123456", FirstName: "John", LastName: "Doe", WarehouseID: 1, InboundOrdersCount: 2},
			{Id: 2, CardNumberID: "654321", FirstName: "Sara", LastName: "Ramirez", WarehouseID: 2, InboundOrdersCount: 3},
		}

		expectedResponse := map[string]interface{}{
			"message": "Success",
			"data":    expectedReport,
		}

		expectedResponseJSON, _ := json.Marshal(expectedResponse)
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		mockService.On("GetReportInboundOrders", (*int)(nil)).Return(&expectedReport, nil)

		// Crear la petición con el body
		req := httptest.NewRequest(http.MethodGet, "/employees/reportInboundOrders", nil)
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()

		// Act
		rt.ServeHTTP(res, req)

		// Assert
		require.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
		assert.JSONEq(t, string(expectedResponseJSON), res.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("Get_report_inboundOrders_fail", func(t *testing.T) {
		// Arrange
		mockService := new(sv.EmployeeServiceMock)
		handler := handlers.NewEmployeeHandler(mockService)
		rt := chi.NewRouter()
		rt.Get("/employees/reportInboundOrders", handler.GetReportInboundOrders())

		expectedResponse := map[string]interface{}{
			"status_code": http.StatusInternalServerError,
			"message":     "Internal Server Error",
		}

		expectedResponseJSON, _ := json.Marshal(expectedResponse)
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		mockService.On("GetReportInboundOrders", (*int)(nil)).Return(nil, customErrors.ErrorInternalServerError)

		// Crear la petición con el body
		req := httptest.NewRequest(http.MethodGet, "/employees/reportInboundOrders", nil)
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()

		// Act
		rt.ServeHTTP(res, req)

		// Assert
		require.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
		assert.JSONEq(t, string(expectedResponseJSON), res.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("GetReportInboundOrders_by_id", func(t *testing.T) {
		// Arrange
		mockService := new(sv.EmployeeServiceMock)
		handler := handlers.NewEmployeeHandler(mockService)
		rt := chi.NewRouter()
		rt.Get("/employees/reportInboundOrders", handler.GetReportInboundOrders())

		var idEmployee int = 1

		expectedReport := models.EmployeeWithOrders{
			Id:                 idEmployee,
			CardNumberID:       "123456",
			FirstName:          "John",
			LastName:           "Doe",
			WarehouseID:        1,
			InboundOrdersCount: 2,
		}

		expectedResponse := map[string]interface{}{
			"message": "Success",
			"data":    expectedReport,
		}

		expectedResponseJSON, _ := json.Marshal(expectedResponse)
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		mockService.On("GetReportInboundOrders", &idEmployee).Return(&expectedReport, nil)

		req := httptest.NewRequest(http.MethodGet, "/employees/reportInboundOrders?id=1", nil)
		res := httptest.NewRecorder()

		// Act
		rt.ServeHTTP(res, req)

		// Assert
		require.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
		assert.JSONEq(t, string(expectedResponseJSON), res.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("GetReportInboundOrders_by_id_not_found", func(t *testing.T) {
		// Arrange
		mockService := new(sv.EmployeeServiceMock)
		handler := handlers.NewEmployeeHandler(mockService)
		rt := chi.NewRouter()
		rt.Get("/employees/reportInboundOrders", handler.GetReportInboundOrders())

		var idEmployee int = 12312

		expectedResponse := map[string]interface{}{
			"status_code": http.StatusNotFound,
			"message":     customErrors.ErrorNotFound.Error(),
		}

		expectedResponseJSON, _ := json.Marshal(expectedResponse)
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		mockService.On("GetReportInboundOrders", &idEmployee).Return(nil, customErrors.ErrorNotFound)

		req := httptest.NewRequest(http.MethodGet, "/employees/reportInboundOrders?id=12312", nil)
		res := httptest.NewRecorder()

		// Act
		rt.ServeHTTP(res, req)

		// Assert
		require.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
		assert.JSONEq(t, string(expectedResponseJSON), res.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("GetReportInboundOrders_with_invalid_ID", func(t *testing.T) {
		// Arrange
		mockService := new(sv.EmployeeServiceMock)
		handler := handlers.NewEmployeeHandler(mockService)
		rt := chi.NewRouter()
		rt.Get("/employees/reportInboundOrders", handler.GetReportInboundOrders())

		expectedResponse := map[string]interface{}{
			"status_code": http.StatusBadRequest,
			"message":     customErrors.ErrorBadRequest.Error(),
		}

		expectedResponseJSON, _ := json.Marshal(expectedResponse)
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		req := httptest.NewRequest(http.MethodGet, "/employees/reportInboundOrders?id=invalid", nil)
		res := httptest.NewRecorder()

		// Act
		rt.ServeHTTP(res, req)

		// Assert
		require.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
		assert.JSONEq(t, string(expectedResponseJSON), res.Body.String())
		mockService.AssertExpectations(t)
	})
}
