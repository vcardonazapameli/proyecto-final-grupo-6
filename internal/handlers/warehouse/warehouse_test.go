package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	handlers "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers/warehouse"
	services "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/warehouse"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestWarehouseById(t *testing.T) {
	t.Run("case existent id", func(t *testing.T) {
		//Arrange
		sv := new(services.WarehouseServiceMock)
		sv.On("GetById", 7).Return(&models.WarehouseDocResponse{ID: 7, Warehouse_code: "qweqqwewe", Address: "Calle falsa 123", Telephone: "123441122", Minimun_capacity: 20, Minimun_temperature: -19.2, Locality_id: nil}, error(nil))
		hd := handlers.NewWarehouseHandler(sv)
		rt := chi.NewRouter()
		rt.Get("/warehouse/{id}", hd.GetById())

		//Act
		req, res := httptest.NewRequest(http.MethodGet, "/warehouse/7", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
						"message": "Success",
						"data": {
							"id": 7,
							"warehouse_code": "qweqqwewe",
							"address": "Calle falsa 123",
							"telephone": "123441122",
							"minimun_capacity": 20,
							"minimun_temperature": -19.2,
							"locality_id": null
						}
					}`
		require.Equal(t, expectedCode, res.Code, "an error has warehouse")
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})
}
