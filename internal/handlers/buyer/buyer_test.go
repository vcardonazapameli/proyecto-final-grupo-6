package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	handlers "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers/buyer"
	services "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/buyer"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestBuyerCreate(t *testing.T){
	t.Run("create_ok", func(t *testing.T) {
		//Arrange
		serviceMock := new(services.BuyerServiceMock)
		handler := handlers.NewBuyerHandler(serviceMock)
		router := chi.NewRouter()
		router.Post("/buyer", handler.CreateBuyer())
		jsonInput := `{
			"card_number_id": 123456,
			"first_name": "Albert",
			"last_name": "Camus"
		}`
		buyerToCreate := models.BuyerDocRequest{
			
			CardNumberId: 123456,
			FirstName: "Albert",
			LastName: "Camus",
		}
		createdBuyer := &models.BuyerDocResponse{
			Id: 0,
			CardNumberId: 123456,
			FirstName: "Albert",
			LastName: "Camus",
		}
		serviceMock.On("CreateBuyer", buyerToCreate).Return(createdBuyer, nil)
		//Act
		request, response := httptest.NewRequest(http.MethodPost, "/buyer", bytes.NewReader([]byte(jsonInput)) ), httptest.NewRecorder()
		router.ServeHTTP(response, request) 
		//assert
		expectedCode := http.StatusCreated
		expectedHeader := http.Header{"Content-Type":[]string{"application/json"}}
		expectedBody := `{
						"message": "Success",
						"data": {
							"id": 0,
							"card_number_id": 123456,
							"first_name": "Albert",
							"last_name": "Camus"
						}
					}`
		require.Equal(t, expectedCode, response.Code)
		require.Equal(t, expectedHeader, response.Header())
		require.JSONEq(t, expectedBody, response.Body.String())
		serviceMock.AssertExpectations(t)
	})
	t.Run("create_fail", func(t *testing.T) {
		//Arrange
		serviceMock := new(services.BuyerServiceMock)
		handler := handlers.NewBuyerHandler(serviceMock)
		router := chi.NewRouter()
		router.Post("/buyer", handler.CreateBuyer())
		jsonInput := `{
			"card_number_id": 0,
			"first_name": "",
			"last_name": ""
		}`
		buyerToCreate := models.BuyerDocRequest{
			
			CardNumberId: 0,
			FirstName: "",
			LastName: "",
		}
		
		serviceMock.On("CreateBuyer", buyerToCreate).Return(nil, customErrors.ErrorUnprocessableContent)
		//Act
		request, response := httptest.NewRequest(http.MethodPost, "/buyer", bytes.NewReader([]byte(jsonInput)) ), httptest.NewRecorder()
		router.ServeHTTP(response, request) 
		//assert
		expectedCode := http.StatusUnprocessableEntity
		expectedHeader := http.Header{"Content-Type":[]string{"application/json"}}
		expectedBody := 
		`{			
			"status_code": 422,
			"message": "unprocessable content"
		}`
		require.Equal(t, expectedCode, response.Code)
		require.Equal(t, expectedHeader, response.Header())
		require.JSONEq(t, expectedBody, response.Body.String())
		serviceMock.AssertExpectations(t)
	})
	t.Run("create_conflict", func(t *testing.T) {
		//Arrange
		serviceMock := new(services.BuyerServiceMock)
		handler := handlers.NewBuyerHandler(serviceMock)
		router := chi.NewRouter()
		router.Post("/buyer", handler.CreateBuyer())
		jsonInput := `{
			"card_number_id": 1234,
			"first_name": "Camilo",
			"last_name": "Cuarto"
		}`
		buyerToCreate := models.BuyerDocRequest{
			
			CardNumberId: 1234,
			FirstName: "Camilo",
			LastName: "Cuarto",
		}
		
		serviceMock.On("CreateBuyer", buyerToCreate).Return(nil, customErrors.ErrorConflict)
		//Act
		request, response := httptest.NewRequest(http.MethodPost, "/buyer", bytes.NewReader([]byte(jsonInput)) ), httptest.NewRecorder()
		router.ServeHTTP(response, request) 
		//assert
		expectedCode := http.StatusConflict
		expectedHeader := http.Header{"Content-Type":[]string{"application/json"}}
		expectedBody := 
		`{			
			"status_code": 409,
			"message": "conflict occurred"
		}`
		require.Equal(t, expectedCode, response.Code)
		require.Equal(t, expectedHeader, response.Header())
		require.JSONEq(t, expectedBody, response.Body.String())
		serviceMock.AssertExpectations(t)
	})
	t.Run("create_bad_request", func(t *testing.T) {
		//Arrange
		serviceMock := new(services.BuyerServiceMock)
		handler := handlers.NewBuyerHandler(serviceMock)
		router := chi.NewRouter()
		router.Post("/buyer", handler.CreateBuyer())
		jsonInput := `{
			"card_number_id": 1234,
			"first_name": 42,
		}`
		//Act
		request, response := httptest.NewRequest(http.MethodPost, "/buyer", bytes.NewReader([]byte(jsonInput)) ), httptest.NewRecorder()
		router.ServeHTTP(response, request) 
		//assert
		expectedCode := http.StatusBadRequest
		expectedHeader := http.Header{"Content-Type":[]string{"application/json"}}
		expectedBody := 
		`{			
			"status_code": 400,
			"message": "bad request"
		}`
		require.Equal(t, expectedCode, response.Code)
		require.Equal(t, expectedHeader, response.Header())
		require.JSONEq(t, expectedBody, response.Body.String())
		serviceMock.AssertExpectations(t)
	})
}

func TestBuyerRead(t *testing.T){
	t.Run("find_all", func(t *testing.T) {
		//Arrange
		serviceMock := new(services.BuyerServiceMock)
		handler := handlers.NewBuyerHandler(serviceMock)
		router := chi.NewRouter()
		router.Get("/buyer", handler.GetAll())
		expectedBuyers := []models.BuyerDocResponse{
			{
				Id:           2,
				CardNumberId: 234532,
				FirstName:    "juan",
				LastName:     "Correa",
			},
			{
				Id:           3,
				CardNumberId: 9842,
				FirstName:    "José",
				LastName:     "Correa",
			},
		}
		serviceMock.On("GetAll").Return(expectedBuyers, nil)
		//Act
		request, response := httptest.NewRequest(http.MethodGet, "/buyer", nil ), httptest.NewRecorder()
		router.ServeHTTP(response, request) 
		//assert
		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type":[]string{"application/json"}}
		expectedBody := `
		{
			"message": "Success",
			"data": [
				{
					"id": 2,
					"card_number_id": 234532,
					"first_name": "juan",
					"last_name": "Correa"
				},
				{
					"id": 3,
					"card_number_id": 9842,
					"first_name": "José",
					"last_name": "Correa"
				}
					]
		}`
		require.Equal(t, expectedCode, response.Code)
		require.Equal(t, expectedHeader, response.Header())
		require.JSONEq(t, expectedBody, response.Body.String())
		serviceMock.AssertExpectations(t)
	})
	t.Run("find_all_internal_error", func(t *testing.T) {
		//Arrange
		serviceMock := new(services.BuyerServiceMock)
		handler := handlers.NewBuyerHandler(serviceMock)
		router := chi.NewRouter()
		router.Get("/buyer", handler.GetAll())
		
		serviceMock.On("GetAll").Return(nil, customErrors.ErrorInternalServerError)
		//Act
		request, response := httptest.NewRequest(http.MethodGet, "/buyer", nil ), httptest.NewRecorder()
		router.ServeHTTP(response, request) 
		//assert
		expectedCode := http.StatusInternalServerError
		expectedHeader := http.Header{"Content-Type":[]string{"application/json"}}
		expectedBody := `
		{
			
			"status_code": 500,
		    "message": "Internal Server Error"

		}`
		require.Equal(t, expectedCode, response.Code)
		require.Equal(t, expectedHeader, response.Header())
		require.JSONEq(t, expectedBody, response.Body.String())
		serviceMock.AssertExpectations(t)
	})
	t.Run("find_by_id_existent", func(t *testing.T) {
		//Arrange
		serviceMock := new(services.BuyerServiceMock)
		handler := handlers.NewBuyerHandler(serviceMock)
		router := chi.NewRouter()
		router.Get("/buyer/{id}", handler.GetById())
		buyerToReturn := &models.BuyerDocResponse{
			Id: 2,
			CardNumberId: 234532,
			FirstName: "Argemiro",
			LastName: "Lopez",
		}
		serviceMock.On("GetById",2).Return(buyerToReturn, nil)
		//Act
		request, response := httptest.NewRequest(http.MethodGet, "/buyer/2", nil ), httptest.NewRecorder()
		router.ServeHTTP(response, request) 
		//assert
		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type":[]string{"application/json"}}
		expectedBody := `
		{
			"message": "Success",
			"data": {
				"id": 2,
				"card_number_id": 234532,
				"first_name": "Argemiro",
				"last_name": "Lopez"
			}
		}`
		require.Equal(t, expectedCode, response.Code)
		require.Equal(t, expectedHeader, response.Header())
		require.JSONEq(t, expectedBody, response.Body.String())
		serviceMock.AssertExpectations(t)
	})
	
	t.Run("find_by_id_non_existent", func(t *testing.T) {
		//Arrange
		serviceMock := new(services.BuyerServiceMock)
		handler := handlers.NewBuyerHandler(serviceMock)
		router := chi.NewRouter()
		router.Get("/buyer/{id}", handler.GetById())
		
		serviceMock.On("GetById",2123).Return(nil, customErrors.ErrorNotFound)
		//Act
		request, response := httptest.NewRequest(http.MethodGet, "/buyer/2123", nil ), httptest.NewRecorder()
		router.ServeHTTP(response, request) 
		//assert
		expectedCode := http.StatusNotFound
		expectedHeader := http.Header{"Content-Type":[]string{"application/json"}}
		expectedBody := `
		{
			"status_code": 404,
			"message": "resource not found"
		}`
		require.Equal(t, expectedCode, response.Code)
		require.Equal(t, expectedHeader, response.Header())
		require.JSONEq(t, expectedBody, response.Body.String())
		serviceMock.AssertExpectations(t)
	})
	t.Run("find_by_id_bad_request", func(t *testing.T) {
		//Arrange
		serviceMock := new(services.BuyerServiceMock)
		handler := handlers.NewBuyerHandler(serviceMock)
		router := chi.NewRouter()
		router.Get("/buyer/{id}", handler.GetById())
		
		//Act
		request, response := httptest.NewRequest(http.MethodGet, "/buyer/asd", nil ), httptest.NewRecorder()
		router.ServeHTTP(response, request) 
		//Assert
		expectedCode := http.StatusBadRequest
		expectedHeader := http.Header{"Content-Type":[]string{"application/json"}}
		expectedBody := `
		{
			"status_code": 400,
			"message": "bad request"
		}`
		require.Equal(t, expectedCode, response.Code)
		require.Equal(t, expectedHeader, response.Header())
		require.JSONEq(t, expectedBody, response.Body.String())
		serviceMock.AssertExpectations(t)
	})

	t.Run("find_report_by_buyer", func(t *testing.T) {
		//Arrange
		serviceMock := new(services.BuyerServiceMock)
		handler := handlers.NewBuyerHandler(serviceMock)
		router := chi.NewRouter()
		router.Get("/reportPurchaseOrders", handler.GetPurchaseOrderReports())
		expectedReports := []models.PurchaseOrderReport{
			{
				BuyerDocResponse: models.BuyerDocResponse {
					Id:           2,
					CardNumberId: 234532,
					FirstName:    "Jacinto",
					LastName:     "Orejuela",
				},
				PurchaseOrdersCount: 2,
			},
			{
				BuyerDocResponse: models.BuyerDocResponse {
					Id:           3,
					CardNumberId: 5633,
					FirstName:    "Pensilvanio",
					LastName:     "Buendia",
				},
				PurchaseOrdersCount: 3,
			},
		}
		serviceMock.On("GetPurchasesReports",0).Return(expectedReports, nil)
		//Act
		request, response := httptest.NewRequest(http.MethodGet, "/reportPurchaseOrders?card_number_id=0", nil ), httptest.NewRecorder()
		router.ServeHTTP(response, request) 
		//assert
		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type":[]string{"application/json"}}
		expectedBody := `{
				"message": "Success",
				"data": [
					{
						"id": 2,
						"card_number_id": 234532,
						"first_name": "Jacinto",
						"last_name": "Orejuela",
						"purchase_orders_count": 2
					},
					{
						"id": 3,
						"card_number_id": 5633,
						"first_name": "Pensilvanio",
						"last_name": "Buendia",
						"purchase_orders_count": 3
					}
				]
			}`
		require.Equal(t, expectedCode, response.Code)
		require.Equal(t, expectedHeader, response.Header())
		require.JSONEq(t, expectedBody, response.Body.String())
		serviceMock.AssertExpectations(t)
	})
	t.Run("find_report_by_buyer_bad_request", func(t *testing.T) {
		//Arrange
		serviceMock := new(services.BuyerServiceMock)
		handler := handlers.NewBuyerHandler(serviceMock)
		router := chi.NewRouter()
		router.Get("/reportPurchaseOrders", handler.GetPurchaseOrderReports())
	
		//Act
		request, response := httptest.NewRequest(http.MethodGet, "/reportPurchaseOrders?card_number_id=asd", nil ), httptest.NewRecorder()
		router.ServeHTTP(response, request) 
		//assert
		expectedCode := http.StatusBadRequest
		expectedHeader := http.Header{"Content-Type":[]string{"application/json"}}
		expectedBody := `
		{
			"status_code": 400,
			"message": "bad request"
		}`
		require.Equal(t, expectedCode, response.Code)
		require.Equal(t, expectedHeader, response.Header())
		require.JSONEq(t, expectedBody, response.Body.String())
		serviceMock.AssertExpectations(t)
	})
	t.Run("find_report_by_buyer_not_found", func(t *testing.T) {
		//Arrange
		serviceMock := new(services.BuyerServiceMock)
		handler := handlers.NewBuyerHandler(serviceMock)
		router := chi.NewRouter()
		router.Get("/reportPurchaseOrders", handler.GetPurchaseOrderReports())
		serviceMock.On("GetPurchasesReports",1245).Return(nil, customErrors.ErrorNotFound)
		//Act
		request, response := httptest.NewRequest(http.MethodGet, "/reportPurchaseOrders?card_number_id=1245", nil ), httptest.NewRecorder()
		router.ServeHTTP(response, request) 
		//assert
		expectedCode := http.StatusNotFound
		expectedHeader := http.Header{"Content-Type":[]string{"application/json"}}
		expectedBody := `
		{
			"status_code": 404,
			"message": "resource not found"
		}`
		require.Equal(t, expectedCode, response.Code)
		require.Equal(t, expectedHeader, response.Header())
		require.JSONEq(t, expectedBody, response.Body.String())
		serviceMock.AssertExpectations(t)
	})
}

func TestBuyerUpdate(t *testing.T){
	t.Run("update_ok", func(t *testing.T) {
		//Arrange
		serviceMock := new(services.BuyerServiceMock)
		handler := handlers.NewBuyerHandler(serviceMock)
		router := chi.NewRouter()
		router.Patch("/buyer/{id}", handler.UpdateBuyer())
		requestBody := `{
			"card_number_id": 1234,
			"first_name": "Albert",
			"last_name": "Camus"
		}`
		cardNumberId:= 1234
		first_name:= "Albert"
		last_name := "Camus"
		requestBuyerDoc := models.UpdateBuyerDto{
			CardNumberId: &cardNumberId,
			FirstName: &first_name,
			LastName:  &last_name,
		}
		updatedBuyer:= &models.BuyerDocResponse{
			Id: 1,
			CardNumberId: 1234,
			FirstName: "Albert",
			LastName: "Camus",
		}
		serviceMock.On("UpdateBuyer",1,requestBuyerDoc).Return(updatedBuyer, nil)
		//Act
		request, response := httptest.NewRequest(http.MethodPatch, "/buyer/1", bytes.NewReader([]byte(requestBody))), httptest.NewRecorder()
		router.ServeHTTP(response, request)
		//Assert
		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type":[]string{"application/json"}}
		expectedBody := `{
						"message": "Success",
						"data": {
							"id": 1,
							"card_number_id": 1234,
							"first_name": "Albert",
							"last_name": "Camus"
						}
					}`
		require.Equal(t, expectedCode, response.Code)
		require.Equal(t, expectedHeader, response.Header())
		require.JSONEq(t, expectedBody, response.Body.String())
		serviceMock.AssertExpectations(t)


	})
	t.Run("update_non_existent", func(t *testing.T) {
		//Arrange
		serviceMock := new(services.BuyerServiceMock)
		handler := handlers.NewBuyerHandler(serviceMock)
		router := chi.NewRouter()
		router.Patch("/buyer/{id}", handler.UpdateBuyer())
		requestBody := `{
			"card_number_id": 123421,
			"first_name": "Albert",
			"last_name": "Camus"
		}`
		cardNumberId:= 123421
		first_name:= "Albert"
		last_name := "Camus"
		requestBuyerDoc := models.UpdateBuyerDto{
			CardNumberId: &cardNumberId,
			FirstName: &first_name,
			LastName:  &last_name,
		}
		serviceMock.On("UpdateBuyer",123421,requestBuyerDoc).Return(nil, customErrors.ErrorNotFound)
		//Act
		request, response := httptest.NewRequest(http.MethodPatch, "/buyer/123421", bytes.NewReader([]byte(requestBody))), httptest.NewRecorder()
		router.ServeHTTP(response, request)
		//Assert
		expectedCode := http.StatusNotFound
		expectedHeader := http.Header{"Content-Type":[]string{"application/json"}}
		expectedBody := 
		`{
    		"status_code": 404,
    		"message": "resource not found"
		}`
		require.Equal(t, expectedCode, response.Code)
		require.Equal(t, expectedHeader, response.Header())
		require.JSONEq(t, expectedBody, response.Body.String())
		serviceMock.AssertExpectations(t)
	})
	t.Run("update_bad_path_param", func(t *testing.T) {
		//Arrange
		serviceMock := new(services.BuyerServiceMock)
		handler := handlers.NewBuyerHandler(serviceMock)
		router := chi.NewRouter()
		router.Patch("/buyer/{id}", handler.UpdateBuyer())
		requestBody := `{
			"card_number_id": 123421,
			"first_name": "Albert",
			"last_name": "Camus"
		}`
		//Act
		request, response := httptest.NewRequest(http.MethodPatch, "/buyer/asd", bytes.NewReader([]byte(requestBody))), httptest.NewRecorder()
		router.ServeHTTP(response, request)
		//Assert
		expectedCode := http.StatusBadRequest
		expectedHeader := http.Header{"Content-Type":[]string{"application/json"}}
		expectedBody := 
		`{
    		"status_code": 400,
    		"message": "bad request"
		}`
		require.Equal(t, expectedCode, response.Code)
		require.Equal(t, expectedHeader, response.Header())
		require.JSONEq(t, expectedBody, response.Body.String())
		serviceMock.AssertExpectations(t)


	})
	t.Run("update_bad_request", func(t *testing.T) {
		//Arrange
		serviceMock := new(services.BuyerServiceMock)
		handler := handlers.NewBuyerHandler(serviceMock)
		router := chi.NewRouter()
		router.Patch("/buyer/{id}", handler.UpdateBuyer())
		//Act
		request, response := httptest.NewRequest(http.MethodPatch, "/buyer/1", nil), httptest.NewRecorder()
		router.ServeHTTP(response, request)
		//Assert
		expectedCode := http.StatusBadRequest
		expectedHeader := http.Header{"Content-Type":[]string{"application/json"}}
		expectedBody := 
		`{
    		"status_code": 400,
    		"message": "bad request"
		}`
		require.Equal(t, expectedCode, response.Code)
		require.Equal(t, expectedHeader, response.Header())
		require.JSONEq(t, expectedBody, response.Body.String())
		serviceMock.AssertExpectations(t)


	})
}

func TestBuyerDelete(t *testing.T){
	t.Run("delete_ok", func(t *testing.T) {
		//Arrange
		serviceMock := new(services.BuyerServiceMock)
		handler := handlers.NewBuyerHandler(serviceMock)
		router := chi.NewRouter()
		router.Delete("/buyer/{id}", handler.DeleteBuyer())
		serviceMock.On("DeleteBuyer",2).Return(nil)
		//Act
		request, response := httptest.NewRequest(http.MethodDelete, "/buyer/2", nil ), httptest.NewRecorder()
		router.ServeHTTP(response, request) 
		//assert
		expectedCode := http.StatusNoContent
		expectedHeader := http.Header{"Content-Type":[]string{"application/json"}}
		
		require.Equal(t, expectedCode, response.Code)
		require.Equal(t, expectedHeader, response.Header())
		serviceMock.AssertExpectations(t)
	})
	t.Run("delete_non_existent", func(t *testing.T) {
		//Arrange
		serviceMock := new(services.BuyerServiceMock)
		handler := handlers.NewBuyerHandler(serviceMock)
		router := chi.NewRouter()
		router.Delete("/buyer/{id}", handler.DeleteBuyer())
		serviceMock.On("DeleteBuyer",2).Return(customErrors.ErrorNotFound)
		//Act
		request, response := httptest.NewRequest(http.MethodDelete, "/buyer/2", nil ), httptest.NewRecorder()
		router.ServeHTTP(response, request) 
		//assert
		expectedCode := http.StatusNotFound
		expectedHeader := http.Header{"Content-Type":[]string{"application/json"}}
		
		require.Equal(t, expectedCode, response.Code)
		require.Equal(t, expectedHeader, response.Header())
		serviceMock.AssertExpectations(t)
	})
	t.Run("delete_bad_request", func(t *testing.T) {
		//Arrange
		serviceMock := new(services.BuyerServiceMock)
		handler := handlers.NewBuyerHandler(serviceMock)
		router := chi.NewRouter()
		router.Delete("/buyer/{id}", handler.DeleteBuyer())
		
		//Act
		request, response := httptest.NewRequest(http.MethodDelete, "/buyer/as", nil ), httptest.NewRecorder()
		router.ServeHTTP(response, request) 
		//assert
		expectedCode := http.StatusBadRequest
		expectedHeader := http.Header{"Content-Type":[]string{"application/json"}}
		
		require.Equal(t, expectedCode, response.Code)
		require.Equal(t, expectedHeader, response.Header())
		serviceMock.AssertExpectations(t)
	})
	
}





