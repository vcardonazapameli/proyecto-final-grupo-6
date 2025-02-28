package handlers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	handlers "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers/warehouse"
	services "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/warehouse"
	errorsCustom "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestWarehouseCreate(t *testing.T) {
	t.Run("case create_ok", func(t *testing.T) {
		// Arrange
		sv := new(services.WarehouseServiceMock)
		localityID := uint64(3)
		newWarehouse := models.WarehouseDocResponse{
			ID:                  10,
			Warehouse_code:      "WH0010",
			Address:             "Calle nueva 456",
			Telephone:           "1100223344",
			Minimun_capacity:    100,
			Minimun_temperature: -5.0,
			Locality_id:         &localityID,
		}
		sv.On("CreateWarehouse", mock.Anything).Return(&newWarehouse, nil)

		hd := handlers.NewWarehouseHandler(sv)
		rt := chi.NewRouter()
		rt.Post("/warehouse", hd.CreateWarehouse())

		requestBody := `{
            "warehouse_code": "WH0010",
            "address": "Calle nueva 456",
            "telephone": "1100223344",
            "minimun_capacity": 100,
            "minimun_temperature": -5.0,
            "locality_id": 3
        }`
		//Act
		req, res := httptest.NewRequest(http.MethodPost, "/warehouse", strings.NewReader(requestBody)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedCode := http.StatusCreated
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
            "message": "Success",
            "data": {
                "id": 10,
                "warehouse_code": "WH0010",
                "address": "Calle nueva 456",
                "telephone": "1100223344",
                "minimun_capacity": 100,
                "minimun_temperature": -5.0,
                "locality_id": 3
            }
        }`

		require.Equal(t, expectedCode, res.Code, "expected 201 created")
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})

	t.Run("case create_fail", func(t *testing.T) {
		// Arrange
		sv := new(services.WarehouseServiceMock)
		sv.On("CreateWarehouse", mock.Anything).Return(&models.WarehouseDocResponse{}, errorsCustom.ErrorBadRequest)

		hd := handlers.NewWarehouseHandler(sv)
		rt := chi.NewRouter()
		rt.Post("/warehouse", hd.CreateWarehouse())

		requestBody := `{
							"warehouse_code": "WHQ10",
							"addres": "",
							"telephone": "1100223344",
							"minimun_capacity": 100,
							"minimun_temperature": -5.0,
							"locality_id": 3
		           		}`
		//Act
		req, res := httptest.NewRequest(http.MethodPost, "/warehouse", strings.NewReader(requestBody)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedCode := http.StatusBadRequest
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
		       				"status_code": 400,
		       				"message": "bad request"
		           		}`

		require.Equal(t, expectedCode, res.Code, "expected 400 bad request")
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})

	t.Run("case create_conflic", func(t *testing.T) {
		// Arrange
		sv := new(services.WarehouseServiceMock)
		sv.On("CreateWarehouse", mock.Anything).Return(nil, errorsCustom.ErrorConflict)

		hd := handlers.NewWarehouseHandler(sv)
		rt := chi.NewRouter()
		rt.Post("/warehouse", hd.CreateWarehouse())

		requestBody := `{
            "warehouse_code": "WH0010",
            "address": "Calle nueva 456",
            "telephone": "1100223344",
            "minimun_capacity": 100,
            "minimun_temperature": -5.0,
            "locality_id": 3
        }`
		//Act
		req, res := httptest.NewRequest(http.MethodPost, "/warehouse", strings.NewReader(requestBody)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedCode := http.StatusConflict
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
							"status_code": 409,
							"message": "conflict occurred"
        				}`

		require.Equal(t, expectedCode, res.Code, "expected 409 conflict")
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})

	t.Run("case create_bad_request", func(t *testing.T) {
		// Arrange
		sv := new(services.WarehouseServiceMock)
		sv.On("CreateWarehouse", mock.Anything).Return(nil, errorsCustom.ErrorBadRequest)

		hd := handlers.NewWarehouseHandler(sv)
		rt := chi.NewRouter()
		rt.Post("/warehouse", hd.CreateWarehouse())

		requestBody := `{
            "warehouse_code": "WH0010",
            "address": "Calle nueva 456",
            "telephone": "1100223344",
            "minimun_capacity": 100,
            "minimun_temperature": -5.0,
            "locality_id": 3
        }`
		//Act
		req, res := httptest.NewRequest(http.MethodPost, "/warehouse", strings.NewReader(requestBody)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedCode := http.StatusBadRequest
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
							"status_code": 400,
							"message": "bad request"
        				}`

		require.Equal(t, expectedCode, res.Code, "expected 400 bad request")
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})
}

func TestWarehouseRead(t *testing.T) {
	t.Run("case find_all", func(t *testing.T) {
		//Arrange
		sv := new(services.WarehouseServiceMock)
		locID3 := uint64(3)
		locID4 := uint64(4)
		sv.On("GetAll").Return([]models.WarehouseDocResponse{
			{
				ID:                  2,
				Warehouse_code:      "AAAA",
				Address:             "PLIQ9212",
				Telephone:           "12222332",
				Minimun_capacity:    50,
				Minimun_temperature: -9.2,
				Locality_id:         &locID3,
			},
			{
				ID:                  4,
				Warehouse_code:      "BBBB",
				Address:             "Calle falsa 123",
				Telephone:           "1122334455",
				Minimun_capacity:    30,
				Minimun_temperature: -19.2,
				Locality_id:         &locID4,
			},
			{
				ID:                  5,
				Warehouse_code:      "BBBBDc",
				Address:             "Calle falsa 123",
				Telephone:           "123441122",
				Minimun_capacity:    30,
				Minimun_temperature: -19.2,
				Locality_id:         &locID4,
			},
		}, error(nil))
		hd := handlers.NewWarehouseHandler(sv)
		rt := chi.NewRouter()
		rt.Get("/warehouse", hd.GetAll())

		//Act
		req, res := httptest.NewRequest(http.MethodGet, "/warehouse", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
							"message": "Success",
							"data": [
								{
									"id": 2,
									"warehouse_code": "AAAA",
									"address": "PLIQ9212",
									"telephone": "12222332",
									"minimun_capacity": 50,
									"minimun_temperature": -9.2,
									"locality_id": 3
								},
								{
									"id": 4,
									"warehouse_code": "BBBB",
									"address": "Calle falsa 123",
									"telephone": "1122334455",
									"minimun_capacity": 30,
									"minimun_temperature": -19.2,
									"locality_id": 4
								},
								{
									"id": 5,
									"warehouse_code": "BBBBDc",
									"address": "Calle falsa 123",
									"telephone": "123441122",
									"minimun_capacity": 30,
									"minimun_temperature": -19.2,
									"locality_id": 4
								}
							]
						}`
		require.Equal(t, expectedCode, res.Code, "expected 200 ok")
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})

	t.Run("case find_by_id_non_existent", func(t *testing.T) {
		//Arrange
		sv := new(services.WarehouseServiceMock)
		sv.On("GetById", 7).Return(nil, errorsCustom.ErrorNotFound)
		hd := handlers.NewWarehouseHandler(sv)
		rt := chi.NewRouter()
		rt.Get("/warehouse/{id}", hd.GetById())

		//Act
		req, res := httptest.NewRequest(http.MethodGet, "/warehouse/7", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		expectedCode := http.StatusNotFound
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
							"status_code": 404,
							"message": "resource not found"
						}`
		require.Equal(t, expectedCode, res.Code, "expected 204 not found")
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})

	t.Run("case find_by_id_existent", func(t *testing.T) {
		//Arrange
		sv := new(services.WarehouseServiceMock)
		sv.On("GetById", 1).Return(&models.WarehouseDocResponse{ID: 1, Warehouse_code: "RWTF", Address: "Calle falsa 123", Telephone: "123441122", Minimun_capacity: 20, Minimun_temperature: -19.2, Locality_id: nil}, error(nil))
		hd := handlers.NewWarehouseHandler(sv)
		rt := chi.NewRouter()
		rt.Get("/warehouse/{id}", hd.GetById())

		//Act
		req, res := httptest.NewRequest(http.MethodGet, "/warehouse/1", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
						"message": "Success",
						"data": {
							"id": 1,
							"warehouse_code": "RWTF",
							"address": "Calle falsa 123",
							"telephone": "123441122",
							"minimun_capacity": 20,
							"minimun_temperature": -19.2,
							"locality_id": null
						}
					}`
		require.Equal(t, expectedCode, res.Code, "expected 200 ok")
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})

	t.Run("case find_all_err", func(t *testing.T) {
		//Arrange
		sv := new(services.WarehouseServiceMock)
		sv.On("GetAll").Return(nil, errors.New("Database error"))
		hd := handlers.NewWarehouseHandler(sv)
		rt := chi.NewRouter()
		rt.Get("/warehouse", hd.GetAll())

		//Act
		req, res := httptest.NewRequest(http.MethodGet, "/warehouse", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		expectedCode := http.StatusInternalServerError
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
							"status_code": 500,
							"message": "Internal Server Error"
						}`
		require.Equal(t, expectedCode, res.Code, "expected 500 internal server error")
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})

	t.Run("case find_by_id_bad_request", func(t *testing.T) {
		//Arrange
		sv := new(services.WarehouseServiceMock)
		sv.On("GetById", 1).Return(nil, errorsCustom.ErrorBadRequest)
		hd := handlers.NewWarehouseHandler(sv)
		rt := chi.NewRouter()
		rt.Get("/warehouse/{id}", hd.GetById())

		//Act
		req, res := httptest.NewRequest(http.MethodGet, "/warehouse/1", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		expectedCode := http.StatusBadRequest
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
							"status_code": 400,
							"message": "bad request"
						}`
		require.Equal(t, expectedCode, res.Code, "expected 400 bad request")
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})

}

func TestWarehouseUpdate(t *testing.T) {
	t.Run("case update_ok", func(t *testing.T) {
		// Arrange
		sv := new(services.WarehouseServiceMock)
		localityID := uint64(3)
		warehouseModel := models.WarehouseDocResponse{
			ID:                  1,
			Warehouse_code:      "WH0010",
			Address:             "Calle nueva 456",
			Telephone:           "1100223344",
			Minimun_capacity:    100,
			Minimun_temperature: -5.0,
		}
		newWarehouse := models.WarehouseUpdateDocResponse{
			ID:                  &warehouseModel.ID,
			Warehouse_code:      &warehouseModel.Warehouse_code,
			Address:             &warehouseModel.Address,
			Telephone:           &warehouseModel.Telephone,
			Minimun_capacity:    &warehouseModel.Minimun_capacity,
			Minimun_temperature: &warehouseModel.Minimun_temperature,
			Locality_id:         &localityID,
		}

		sv.On("UpdateWarehouse", 1, mock.Anything).Return(&newWarehouse, nil)

		hd := handlers.NewWarehouseHandler(sv)
		rt := chi.NewRouter()
		rt.Patch("/warehouse/{id}", hd.UpdateWarehouse())

		requestBody := `{
            "warehouse_code": "WH0010",
            "address": "Calle nueva 456",
            "telephone": "1100223344",
            "minimun_capacity": 100,
            "minimun_temperature": -5.0,
            "locality_id": 3
        }`
		//Act
		req, res := httptest.NewRequest(http.MethodPatch, "/warehouse/1", strings.NewReader(requestBody)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
            "message": "Success",
            "data": {
                "id": 1,
                "warehouse_code": "WH0010",
                "address": "Calle nueva 456",
                "telephone": "1100223344",
                "minimun_capacity": 100,
                "minimun_temperature": -5.0,
                "locality_id": 3
            }
        }`

		require.Equal(t, expectedCode, res.Code, "expected 200 ok")
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})

	t.Run("case update_ok", func(t *testing.T) {
		// Arrange
		sv := new(services.WarehouseServiceMock)
		sv.On("UpdateWarehouse", 1, mock.Anything).Return(&models.WarehouseUpdateDocResponse{}, errorsCustom.ErrorNotFound)

		hd := handlers.NewWarehouseHandler(sv)
		rt := chi.NewRouter()
		rt.Patch("/warehouse/{id}", hd.UpdateWarehouse())

		requestBody := `{
            "warehouse_code": "WH0010",
            "address": "Calle nueva 456",
            "telephone": "1100223344",
            "minimun_capacity": 100,
            "minimun_temperature": -5.0,
            "locality_id": 3
        }`
		//Act
		req, res := httptest.NewRequest(http.MethodPatch, "/warehouse/1", strings.NewReader(requestBody)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedCode := http.StatusNotFound
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
							"status_code": 404,
							"message": "resource not found"
						}`

		require.Equal(t, expectedCode, res.Code, "expected 404 not found")
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})
}

func TestWarehouseDelete(t *testing.T) {

	t.Run("case delete_non_existent", func(t *testing.T) {
		//Arrange
		sv := new(services.WarehouseServiceMock)
		sv.On("DeleteWarehouse", 1).Return(errorsCustom.ErrorNotFound)
		hd := handlers.NewWarehouseHandler(sv)
		rt := chi.NewRouter()
		rt.Delete("/warehouse/{id}", hd.DeleteWarehouse())

		//Act
		req, res := httptest.NewRequest(http.MethodDelete, "/warehouse/1", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		expectedCode := http.StatusNotFound
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
			"status_code": 404,
			"message": "resource not found"
		}`
		require.Equal(t, expectedCode, res.Code, "expected 204 not found")
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})

	t.Run("case delete_ok", func(t *testing.T) {
		//Arrange
		sv := new(services.WarehouseServiceMock)
		sv.On("DeleteWarehouse", 1).Return(error(nil))
		hd := handlers.NewWarehouseHandler(sv)
		rt := chi.NewRouter()
		rt.Delete("/warehouse/{id}", hd.DeleteWarehouse())

		//Act
		req, res := httptest.NewRequest(http.MethodDelete, "/warehouse/1", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		expectedCode := http.StatusNoContent
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code, "expected 204 No Content")
		require.Equal(t, expectedHeader, res.Header())
		sv.AssertExpectations(t)
	})
}
