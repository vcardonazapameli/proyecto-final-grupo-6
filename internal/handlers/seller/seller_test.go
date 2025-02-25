package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	handler "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers/seller"
	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/seller"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestHandlerSellerCreate(t *testing.T) {
	t.Run("Create Seller Successfully", func(t *testing.T) {

		//Arrange
		svMock := new(service.SellerServiceMock)

		svIn := models.SellerDoc{
			Id: -1, Cid: 1000, CompanyName: "Company 1", Address: "Address 1", Telephone: "11222333", LocalityID: 1,
		}

		svOut := models.SellerDoc{
			Id: 1, Cid: 1000, CompanyName: "Company 1", Address: "Address 1", Telephone: "11222333", LocalityID: 1,
		}

		svMock.On("Create", svIn).Return(svOut, nil)
		hd := handler.NewSellerHandler(svMock)

		rt := chi.NewRouter()
		rt.Post("/seller", hd.Create())

		requestBody := map[string]interface{}{
			"cid":          1000,
			"company_name": "Company 1",
			"address":      "Address 1",
			"telephone":    "11222333",
			"locality_id":  1,
		}

		jsonBody, _ := json.Marshal(requestBody)

		// -- Expected Values --
		expectedBody := `{
			"message": "Success",
			"data": {
				"id": 1,
				"cid": 1000,
				"company_name": "Company 1",
				"address": "Address 1",
				"telephone": "11222333",
				"locality_id": 1
				}
			}`

		expectedCode := http.StatusCreated
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		//Act
		req, res := httptest.NewRequest(http.MethodPost, "/seller", bytes.NewReader(jsonBody)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("Create Seller Error: Validation Error", func(t *testing.T) {

		//Arrange
		svMock := new(service.SellerServiceMock)

		svIn := models.SellerDoc{
			Id: -1, Cid: 0, CompanyName: "", Address: "", Telephone: "", LocalityID: 1,
		}

		svMock.On("Create", svIn).Return(models.SellerDoc{}, customErrors.ValidationError{})
		hd := handler.NewSellerHandler(svMock)

		rt := chi.NewRouter()
		rt.Post("/seller", hd.Create())

		requestBody := map[string]any{
			"cid":          0,
			"company_name": "",
			"address":      "",
			"telephone":    "",
			"locality_id":  1,
		}

		jsonBody, _ := json.Marshal(requestBody)

		// -- Expected Values --
		expectedBody := `{
			"status_code": 422,
			"message": "There were some errors validating:  "
			}`

		expectedCode := http.StatusUnprocessableEntity
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		//Act
		req, res := httptest.NewRequest(http.MethodPost, "/seller", bytes.NewReader(jsonBody)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("Create Seller Error: Bad Request", func(t *testing.T) {

		//Arrange
		svMock := new(service.SellerServiceMock)
		hd := handler.NewSellerHandler(svMock)

		rt := chi.NewRouter()
		rt.Post("/seller", hd.Create())

		requestBody := map[string]any{
			"cid":          1,
			"company_name": 9999, // Expect string
			"address":      "Address 1",
			"telephone":    "11222333",
			"locality_id":  1,
		}

		jsonBody, _ := json.Marshal(requestBody)

		// -- Expected Values --
		expectedBody := `{
			"status_code": 400,
			"message": "bad request"
			}`

		expectedCode := http.StatusBadRequest
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		//Act
		req, res := httptest.NewRequest(http.MethodPost, "/seller", bytes.NewReader(jsonBody)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
	})
}
