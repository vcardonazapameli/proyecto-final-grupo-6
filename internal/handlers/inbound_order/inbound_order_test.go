package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/go-chi/chi/v5"

	handlers "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers/inbound_order"
	sv "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/inbound_order"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func TestInboundOrderHandler(t *testing.T) {
	t.Run("Create_ok", func(t *testing.T) {
		// Arrange
		mockService := new(sv.InboundOrderServiceMock)
		handler := handlers.NewInboundOrderHandler(mockService)
		rt := chi.NewRouter()
		rt.Post("/inbound_orders", handler.Create())

		newInboundOrder := models.RequestInboundOrder{
			OrderDate:      "2021-10-10",
			OrderNumber:    "123456",
			EmployeeID:     1,
			ProductBatchID: 2,
			WarehouseID:    1,
		}

		expectedInboundOrder := models.InboundOrder{
			ID:             1,
			OrderDate:      "2021-10-10",
			OrderNumber:    "123456",
			EmployeeID:     1,
			ProductBatchID: 2,
			WarehouseID:    1,
		}

		expectedResponse := map[string]interface{}{
			"message": "Success",
			"data":    expectedInboundOrder,
		}

		expectedResponseJSON, _ := json.Marshal(expectedResponse)
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		mockService.On("Create", newInboundOrder).Return(&expectedInboundOrder, nil)

		// Convertir el objeto newInboundOrder a JSON
		newInboundOrderJSON, _ := json.Marshal(newInboundOrder)

		// Crear la petición con el body
		req := httptest.NewRequest(http.MethodPost, "/inbound_orders", bytes.NewBuffer(newInboundOrderJSON))
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

	t.Run("Create_decode_error", func(t *testing.T) {
		// Arrange
		mockService := new(sv.InboundOrderServiceMock)
		handler := handlers.NewInboundOrderHandler(mockService)
		rt := chi.NewRouter()
		rt.Post("/inbound_orders", handler.Create())

		expectedResponse := map[string]interface{}{
			"status_code": http.StatusBadRequest,
			"message":     customErrors.ErrorBadRequest.Error(),
		}

		expectedResponseJSON, _ := json.Marshal(expectedResponse)
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		// Crear la petición con el body invalido
		req := httptest.NewRequest(http.MethodPost, "/inbound_orders", bytes.NewBuffer([]byte("invalid json")))
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

	t.Run("Create_service_error", func(t *testing.T) {
		// Arrange
		mockService := new(sv.InboundOrderServiceMock)
		handler := handlers.NewInboundOrderHandler(mockService)
		rt := chi.NewRouter()
		rt.Post("/inbound_orders", handler.Create())

		newInboundOrder := models.RequestInboundOrder{
			OrderDate:      "2021-10-10",
			OrderNumber:    "123456",
			EmployeeID:     1,
			ProductBatchID: 2,
			WarehouseID:    1,
		}

		expectedResponse := map[string]interface{}{
			"status_code": http.StatusInternalServerError,
			"message":     "Internal Server Error",
		}

		expectedResponseJSON, _ := json.Marshal(expectedResponse)
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		mockService.On("Create", newInboundOrder).Return(nil, customErrors.ErrorInternalServerError)

		// Convertir el objeto newInboundOrder a JSON
		newInboundOrderJSON, _ := json.Marshal(newInboundOrder)

		// Crear la petición con el body
		req := httptest.NewRequest(http.MethodPost, "/inbound_orders", bytes.NewBuffer(newInboundOrderJSON))
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
}
