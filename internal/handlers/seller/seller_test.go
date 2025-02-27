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

func TestHandlerSeller_Create(t *testing.T) {
	t.Run("Create Seller Ok: Successfully", func(t *testing.T) {

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

	t.Run("Create Seller Error: Conflict", func(t *testing.T) {

		//Arrange
		svMock := new(service.SellerServiceMock)

		svIn := models.SellerDoc{
			Id: -1, Cid: 1000, CompanyName: "Company 1", Address: "Address 1", Telephone: "11222333", LocalityID: 1,
		}

		svMock.On("Create", svIn).Return(models.SellerDoc{}, customErrors.ErrorConflict)
		hd := handler.NewSellerHandler(svMock)

		rt := chi.NewRouter()
		rt.Post("/seller", hd.Create())

		requestBody := map[string]any{
			"cid":          1000,
			"company_name": "Company 1",
			"address":      "Address 1",
			"telephone":    "11222333",
			"locality_id":  1,
		}

		jsonBody, _ := json.Marshal(requestBody)

		// -- Expected Values --
		expectedBody := `{
			"status_code": 409,
			"message": "conflict occurred"
			}`

		expectedCode := http.StatusConflict
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

func TestHandler_Read(t *testing.T) {

	t.Run("Read By ID: Ok", func(t *testing.T) {

		// Arrange
		svMock := new(service.SellerServiceMock)

		svOut := models.SellerDoc{
			Id: 1, Cid: 1000, CompanyName: "Company 1", Address: "Address 1", Telephone: "11222333", LocalityID: 1,
		}

		svMock.On("GetByID", 1).Return(svOut, nil)

		hd := handler.NewSellerHandler(svMock)

		rt := chi.NewRouter()
		rt.Get("/seller/{id}", hd.GetByID())

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

		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		//Act
		req, res := httptest.NewRequest(http.MethodGet, "/seller/1", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("Read Failed: Non-Existing ID", func(t *testing.T) {

		// Arrange
		svMock := new(service.SellerServiceMock)

		svMock.On("GetByID", 9999).Return(models.SellerDoc{}, customErrors.ErrorNotFound)
		hd := handler.NewSellerHandler(svMock)

		rt := chi.NewRouter()
		rt.Get("/seller/{id}", hd.GetByID())

		// -- Expected Values --
		expectedBody := `{
			"status_code": 404,
			"message": "resource not found"
			}`

		expectedCode := http.StatusNotFound
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		//Act
		req, res := httptest.NewRequest(http.MethodGet, "/seller/9999", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("Read All: Ok", func(t *testing.T) {

		// Arrange
		svMock := new(service.SellerServiceMock)

		svOut := map[int]models.SellerDoc{
			1: {Id: 1, Cid: 1000, CompanyName: "Company 1", Address: "Address 1", Telephone: "11222333", LocalityID: 1},
			2: {Id: 2, Cid: 2000, CompanyName: "Company 2", Address: "Address 2", Telephone: "11333444", LocalityID: 2},
		}

		svMock.On("GetAll").Return(svOut, nil)

		hd := handler.NewSellerHandler(svMock)

		rt := chi.NewRouter()
		rt.Get("/seller", hd.GetAll())

		// -- Expected Values --
		expectedBody := map[string]any{
			"data": map[string]any{
				"1": map[string]any{
					"address":      "Address 1",
					"cid":          1000,
					"company_name": "Company 1",
					"id":           1,
					"locality_id":  1,
					"telephone":    "11222333",
				},
				"2": map[string]interface{}{
					"address":      "Address 2",
					"cid":          2000,
					"company_name": "Company 2",
					"id":           2,
					"locality_id":  2,
					"telephone":    "11333444",
				},
			},
			"message": "Success",
		}
		expectedJSON, _ := json.Marshal(expectedBody)

		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		//Act
		req, res := httptest.NewRequest(http.MethodGet, "/seller", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert

		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, string(expectedJSON), res.Body.String())
	})
}

func TestHandler_Update(t *testing.T) {
	t.Run("Update Ok", func(t *testing.T) {

		// Arrange
		var (
			newName    string = "New Name"
			newAddress string = "New Address"
		)

		svMock := new(service.SellerServiceMock)

		svOut := models.SellerDoc{Id: 1, Cid: 1000, CompanyName: "New Name", Address: "New Address", Telephone: "11222333", LocalityID: 1}

		svMock.On("Update", 1, (*int)(nil), &newName, &newAddress, (*string)(nil), (*int)(nil)).Return(svOut, nil)

		hd := handler.NewSellerHandler(svMock)

		rt := chi.NewRouter()
		rt.Patch("/seller/{id}", hd.Update())

		requestObj := handler.UpdateSellerRequest{
			Cid:         nil,
			CompanyName: &newName,
			Address:     &newAddress,
			Telephone:   nil,
			LocalityID:  nil,
		}
		requestBody, _ := json.Marshal(requestObj)

		// -- Expected Values --
		expectedBody := `
			{
			"message": "Success",
			"data": {
				"id": 1,
				"cid": 1000,
				"company_name": "New Name",
				"address": "New Address",
				"telephone": "11222333",
				"locality_id": 1
			}
		}
		`

		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		//Act
		req, res := httptest.NewRequest(http.MethodPatch, "/seller/1", bytes.NewReader(requestBody)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("Update Failed: Non Existing ID", func(t *testing.T) {

		// Arrange
		var (
			newName    string = "New Name"
			newAddress string = "New Address"
		)

		svMock := new(service.SellerServiceMock)

		svMock.On("Update", 9999, (*int)(nil), &newName, &newAddress, (*string)(nil), (*int)(nil)).Return(models.SellerDoc{}, customErrors.ErrorNotFound)

		hd := handler.NewSellerHandler(svMock)

		rt := chi.NewRouter()
		rt.Patch("/seller/{id}", hd.Update())

		requestObj := handler.UpdateSellerRequest{
			Cid:         nil,
			CompanyName: &newName,
			Address:     &newAddress,
			Telephone:   nil,
			LocalityID:  nil,
		}
		requestBody, _ := json.Marshal(requestObj)

		// -- Expected Values --
		expectedBody := `
			{
			"status_code": 404,
			"message": "resource not found"
			}
		`

		expectedCode := http.StatusNotFound
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		//Act
		req, res := httptest.NewRequest(http.MethodPatch, "/seller/9999", bytes.NewReader(requestBody)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
	})

}
