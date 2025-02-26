package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	handler "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers/product"
	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/product"
	customErrors "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_Product_Create(t *testing.T) {
	t.Run("create_ok", func(t *testing.T) {
		// Arrange
		productRequest := models.ProductDocRequest{
			ProductCode:                    "PROD004",
			Description:                    "Producto D",
			ExpirationRate:                 0.7,
			RecommendedFreezingTemperature: -20,
			FreezingRate:                   -20,
			Width:                          2,
			Height:                         5,
			Length:                         10,
			NetWeight:                      0.5,
			ProductType:                    2,
			Seller:                         2,
		}

		productResponse := &models.ProductDocResponse{
			Id:                             8,
			ProductCode:                    "PROD004",
			Description:                    "Producto D",
			ExpirationRate:                 0.7,
			RecommendedFreezingTemperature: -20,
			FreezingRate:                   -20,
			Width:                          2,
			Height:                         5,
			Length:                         10,
			NetWeight:                      0.5,
			ProductType:                    2,
			Seller:                         2,
		}

		sv := new(service.ProductServiceMock)
		sv.On("Create", productRequest).Return(productResponse, nil)

		hd := handler.NewProductHandler(sv)
		rt := chi.NewRouter()
		rt.Post("/products", hd.Create())

		productCreated, _ := json.Marshal(productRequest)

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader(productCreated)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedCode := http.StatusCreated
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedResponse := map[string]interface{}{
			"message": "Success",
			"data":    productResponse,
		}
		expectedBody, _ := json.Marshal(expectedResponse)
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, string(expectedBody), res.Body.String())
		sv.AssertExpectations(t)
	})
	t.Run("create_fail", func(t *testing.T) {
		// Arrange
		productRequest := models.ProductDocRequest{
			ProductCode:                    "P004",
			Description:                    "Producto D",
			ExpirationRate:                 365,
			RecommendedFreezingTemperature: -18,
			FreezingRate:                   0.5,
			Width:                          10,
			Height:                         20,
			Length:                         30,
			NetWeight:                      1,
			ProductType:                    1,
			Seller:                         1,
		}

		sv := new(service.ProductServiceMock)
		sv.On("Create", productRequest).Return(&models.ProductDocResponse{}, customErrors.ErrorBadRequest)

		hd := handler.NewProductHandler(sv)
		rt := chi.NewRouter()
		rt.Post("/products", hd.Create())

		productCreated, _ := json.Marshal(productRequest)

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(productCreated)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedCode := http.StatusBadRequest
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{"status_code": 400, "message": "bad request"}`
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})
	t.Run("create_conflict", func(t *testing.T) {
		// Arrange
		productRequest := models.ProductDocRequest{
			ProductCode:                    "P004",
			Description:                    "Producto D",
			ExpirationRate:                 365,
			RecommendedFreezingTemperature: -20,
			FreezingRate:                   -20,
			Width:                          10,
			Height:                         20,
			Length:                         30,
			NetWeight:                      1,
			ProductType:                    1,
			Seller:                         1,
		}

		sv := new(service.ProductServiceMock)
		sv.On("Create", productRequest).Return(&models.ProductDocResponse{}, customErrors.ErrorConflict)
		hd := handler.NewProductHandler(sv)
		rt := chi.NewRouter()
		rt.Post("/products", hd.Create())

		productCreated, _ := json.Marshal(productRequest)

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(productCreated)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedCode := http.StatusConflict
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{"status_code": 409, "message": "conflict occurred"}`
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})
	t.Run("create_invalid_json", func(t *testing.T) {
		// Arrange
		sv := new(service.ProductServiceMock)
		hd := handler.NewProductHandler(sv)
		rt := chi.NewRouter()
		rt.Post("/products", hd.Create())

		req, res := httptest.NewRequest(http.MethodPost, "/products", nil), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedCode := http.StatusInternalServerError
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{"status_code": 500, "message": "Internal Server Error"}`
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})
}

func Test_Product_Read(t *testing.T) {
	t.Run("find_all_ok", func(t *testing.T) {
		//Arrange
		sv := new(service.ProductServiceMock)
		sv.On("GetAll").Return([]models.ProductDocResponse{
			{
				Id:                             1,
				ProductCode:                    "P001",
				Description:                    "Producto A",
				ExpirationRate:                 365,
				RecommendedFreezingTemperature: -18,
				FreezingRate:                   0.5,
				Width:                          10,
				Height:                         20,
				Length:                         30,
				NetWeight:                      1,
				ProductType:                    1,
				Seller:                         1,
			},
			{
				Id:                             2,
				ProductCode:                    "P002",
				Description:                    "Producto B",
				ExpirationRate:                 180,
				RecommendedFreezingTemperature: -15,
				FreezingRate:                   0.3,
				Width:                          5,
				Height:                         5,
				Length:                         10,
				NetWeight:                      0.5,
				ProductType:                    2,
				Seller:                         2,
			},
		}, error(nil))
		hd := handler.NewProductHandler(sv)
		rt := chi.NewRouter()
		rt.Get("/products", hd.GetAll())

		//Act
		req, res := httptest.NewRequest(http.MethodGet, "/products", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
							"message": "Success",
							"data": [
								{
									"id": 1,
									"product_code": "P001",
									"description": "Producto A",
									"expiration_rate": 365,
									"recommended_freezing_temperature": -18,
									"freezing_rate": 0.5,
									"width": 10,
									"height": 20,
									"length": 30,
									"net_weight": 1,
									"product_type_id": 1,
									"seller_id": 1
								},
								{
									"id": 2,
									"product_code": "P002",
									"description": "Producto B",
									"expiration_rate": 180,
									"recommended_freezing_temperature": -15,
									"freezing_rate": 0.3,
									"width": 5,
									"height": 5,
									"length": 10,
									"net_weight": 0.5,
									"product_type_id": 2,
									"seller_id": 2
								}
							]
						}`
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})
	t.Run("find_all_fail", func(t *testing.T) {
		//Arrange
		sv := new(service.ProductServiceMock)
		sv.On("GetAll").Return(nil, customErrors.ErrorInternalServerError)
		hd := handler.NewProductHandler(sv)
		rt := chi.NewRouter()
		rt.Get("/products", hd.GetAll())

		//Act
		req, res := httptest.NewRequest(http.MethodGet, "/products", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		expectedCode := http.StatusInternalServerError
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{"status_code": 500, "message": "Internal Server Error"}`
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})
	t.Run("find_by_id_existent", func(t *testing.T) {
		//Arrange
		sv := new(service.ProductServiceMock)
		sv.On("GetById", 1).Return(&models.ProductDocResponse{
			Id:                             1,
			ProductCode:                    "P001",
			Description:                    "Producto A",
			ExpirationRate:                 365,
			RecommendedFreezingTemperature: -18,
			FreezingRate:                   0.5,
			Width:                          10,
			Height:                         20,
			Length:                         30,
			NetWeight:                      1,
			ProductType:                    1,
			Seller:                         1,
		}, error(nil))
		hd := handler.NewProductHandler(sv)
		rt := chi.NewRouter()
		rt.Get("/products/{id}", hd.GetById())

		//Act
		req, res := httptest.NewRequest(http.MethodGet, "/products/1", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
							"message": "Success",
							"data": {
								"id": 1,
								"product_code": "P001",
								"description": "Producto A",
								"expiration_rate": 365,
								"recommended_freezing_temperature": -18,
								"freezing_rate": 0.5,
								"width": 10,
								"height": 20,
								"length": 30,
								"net_weight": 1,
								"product_type_id": 1,
								"seller_id": 1
							}
						}`
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})
	t.Run("find_by_id_non_existent", func(t *testing.T) {
		//Arrange
		sv := new(service.ProductServiceMock)
		sv.On("GetById", 2).Return(nil, customErrors.ErrorNotFound)
		hd := handler.NewProductHandler(sv)
		rt := chi.NewRouter()
		rt.Get("/products/{id}", hd.GetById())

		//Act
		req, res := httptest.NewRequest(http.MethodGet, "/products/2", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		expectedCode := http.StatusNotFound
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{"status_code": 404, "message": "resource not found"}`
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})
}

func Test_Product_Update(t *testing.T) {
	t.Run("update_ok", func(t *testing.T) {
		// Arrange
		productRequest := models.ProductUpdateDocRequest{
			ProductCode:    new(string),
			Description:    new(string),
			ExpirationRate: new(float64),
		}
		*productRequest.ProductCode = "P001B"
		*productRequest.Description = "Producto B"
		*productRequest.ExpirationRate = 200

		productResponse := &models.ProductDocResponse{
			Id:                             1,
			ProductCode:                    "P001B",
			Description:                    "Producto B",
			ExpirationRate:                 200,
			RecommendedFreezingTemperature: -20,
			FreezingRate:                   -20,
			Width:                          10,
			Height:                         20,
			Length:                         30,
			NetWeight:                      1,
			ProductType:                    1,
			Seller:                         1,
		}

		sv := new(service.ProductServiceMock)
		sv.On("Update", 1, productRequest).Return(productResponse, nil)

		hd := handler.NewProductHandler(sv)
		rt := chi.NewRouter()
		rt.Patch("/products/{id}", hd.Update())

		productUpdated, _ := json.Marshal(productRequest)

		// Act
		req, res := httptest.NewRequest(http.MethodPatch, "/products/1", bytes.NewReader(productUpdated)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedResponse := map[string]interface{}{
			"message": "Success",
			"data":    productResponse,
		}
		expectedBody, _ := json.Marshal(expectedResponse)
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, string(expectedBody), res.Body.String())
		sv.AssertExpectations(t)
	})
	t.Run("update_fail", func(t *testing.T) {
		// Arrange
		sv := new(service.ProductServiceMock)
		hd := handler.NewProductHandler(sv)
		rt := chi.NewRouter()
		rt.Patch("/products/{id}", hd.Update())

		req, res := httptest.NewRequest(http.MethodPatch, "/products/1", nil), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedCode := http.StatusInternalServerError
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{"status_code": 500, "message": "Internal Server Error"}`
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})
	t.Run("update_non_existent", func(t *testing.T) {
		// Arrange
		productRequest := models.ProductUpdateDocRequest{
			ProductCode:    new(string),
			Description:    new(string),
			ExpirationRate: new(float64),
		}
		*productRequest.ProductCode = "P001B"
		*productRequest.Description = "Producto B"
		*productRequest.ExpirationRate = 200

		sv := new(service.ProductServiceMock)
		sv.On("Update", 3, productRequest).Return(nil, customErrors.ErrorNotFound)

		hd := handler.NewProductHandler(sv)
		rt := chi.NewRouter()
		rt.Patch("/products/{id}", hd.Update())

		productUpdated, _ := json.Marshal(productRequest)

		// Act
		req, res := httptest.NewRequest(http.MethodPatch, "/products/3", bytes.NewReader(productUpdated)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedCode := http.StatusNotFound
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{"status_code": 404, "message": "resource not found"}`
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, string(expectedBody), res.Body.String())
		sv.AssertExpectations(t)
	})
}

func Test_Product_Delete(t *testing.T) {
	t.Run("delete_ok", func(t *testing.T) {
		//Arrange
		sv := new(service.ProductServiceMock)
		sv.On("Delete", 2).Return(nil)
		hd := handler.NewProductHandler(sv)
		rt := chi.NewRouter()
		rt.Delete("/products/{id}", hd.Delete())

		//Act
		req, res := httptest.NewRequest(http.MethodDelete, "/products/2", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		expectedCode := http.StatusNoContent
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		sv.AssertExpectations(t)
	})
	t.Run("delete_non_existent", func(t *testing.T) {
		//Arrange
		sv := new(service.ProductServiceMock)
		sv.On("Delete", 3).Return(customErrors.ErrorNotFound)
		hd := handler.NewProductHandler(sv)
		rt := chi.NewRouter()
		rt.Delete("/products/{id}", hd.Delete())

		//Act
		req, res := httptest.NewRequest(http.MethodDelete, "/products/3", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		expectedCode := http.StatusNotFound
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{"status_code": 404, "message": "resource not found"}`
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})
}

func Test_Product_Record_Read(t *testing.T) {
	t.Run("find_all_product_records_ok", func(t *testing.T) {
		//Arrange
		sv := new(service.ProductServiceMock)
		sv.On("GetProductRecords", mock.Anything, mock.Anything, mock.Anything).Return([]models.ProductRecordByProductResponse{
			{
				ProductId:    1,
				Description:  "Producto A",
				RecordsCount: 1,
			},
			{
				ProductId:    2,
				Description:  "Producto B",
				RecordsCount: 2,
			},
			{
				ProductId:    3,
				Description:  "Producto C",
				RecordsCount: 3,
			},
		}, error(nil))
		hd := handler.NewProductHandler(sv)
		rt := chi.NewRouter()
		rt.Get("/products/productRecords", hd.GetProductRecords())

		//Act
		req, res := httptest.NewRequest(http.MethodGet, "/products/productRecords", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
							"message": "Success",
							"data": [
								{
									"product_id": 1,
									"description": "Producto A",
									"records_count": 1
								},
								{
									"product_id": 2,
									"description": "Producto B",
									"records_count": 2
								},
								{
									"product_id": 3,
									"description": "Producto C",
									"records_count": 3
								}
							]
						}`
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})
	t.Run("find_all_product_records_fail", func(t *testing.T) {
		//Arrange
		sv := new(service.ProductServiceMock)
		sv.On("GetProductRecords", mock.Anything, mock.Anything, mock.Anything).Return(nil, customErrors.ErrorInternalServerError)
		hd := handler.NewProductHandler(sv)
		rt := chi.NewRouter()
		rt.Get("/products/productRecords", hd.GetProductRecords())

		//Act
		req, res := httptest.NewRequest(http.MethodGet, "/products/productRecords", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		expectedCode := http.StatusInternalServerError
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{"status_code": 500, "message": "Internal Server Error"}`
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})
	t.Run("find_all_product_records_by_id", func(t *testing.T) {
		//Arrange
		sv := new(service.ProductServiceMock)
		sv.On("GetProductRecords", mock.Anything, mock.Anything, mock.Anything).Return([]models.ProductRecordByProductResponse{
			{
				ProductId:    1,
				Description:  "Producto A",
				RecordsCount: 1,
			},
		}, error(nil))
		hd := handler.NewProductHandler(sv)
		rt := chi.NewRouter()
		rt.Get("/products/productRecords", hd.GetProductRecords())

		//Act
		req, res := httptest.NewRequest(http.MethodGet, "/products/productRecords?id=1", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
							"message": "Success",
							"data": [
								{
									"product_id": 1,
									"description": "Producto A",
									"records_count": 1
								}
							]
						}`
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})
	t.Run("find_all_product_records_by_product_code", func(t *testing.T) {
		//Arrange
		sv := new(service.ProductServiceMock)
		sv.On("GetProductRecords", mock.Anything, mock.Anything, mock.Anything).Return([]models.ProductRecordByProductResponse{
			{
				ProductId:    1,
				Description:  "Producto A",
				RecordsCount: 1,
			},
		}, error(nil))
		hd := handler.NewProductHandler(sv)
		rt := chi.NewRouter()
		rt.Get("/products/productRecords", hd.GetProductRecords())

		//Act
		req, res := httptest.NewRequest(http.MethodGet, "/products/productRecords?product_code=P001B", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
							"message": "Success",
							"data": [
								{
									"product_id": 1,
									"description": "Producto A",
									"records_count": 1
								}
							]
						}`
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})
	t.Run("find_all_product_records_by_product_type_id", func(t *testing.T) {
		//Arrange
		sv := new(service.ProductServiceMock)
		sv.On("GetProductRecords", mock.Anything, mock.Anything, mock.Anything).Return([]models.ProductRecordByProductResponse{
			{
				ProductId:    1,
				Description:  "Producto A",
				RecordsCount: 1,
			},
		}, error(nil))
		hd := handler.NewProductHandler(sv)
		rt := chi.NewRouter()
		rt.Get("/products/productRecords", hd.GetProductRecords())

		//Act
		req, res := httptest.NewRequest(http.MethodGet, "/products/productRecords?product_type_id=1", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
							"message": "Success",
							"data": [
								{
									"product_id": 1,
									"description": "Producto A",
									"records_count": 1
								}
							]
						}`
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})
}
