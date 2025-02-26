package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	handlers "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers/buyer"
	services "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/buyer"
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
}