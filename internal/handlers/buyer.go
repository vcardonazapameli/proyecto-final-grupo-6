package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/buyer"
	customErrors "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/errors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

func NewBuyerHandler(sv service.BuyerService) *BuyerHandler {
	return &BuyerHandler{sv: sv}
}

// EmployeeDefault is a struct with methods that represent handlers for Employees
type BuyerHandler struct {
	// sv is the service that will be used by the handler
	sv service.BuyerService
}

func (handler *BuyerHandler)GetAll()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		buyers, err := handler.sv.GetAll()
		if err != nil {
			response.Error(w,http.StatusNotFound,err.Error())
			return
		}
		response.JSON(w,http.StatusOK, buyers)
		return
	}
}

func (handler *BuyerHandler)GetById()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil{
			response.Error(w,http.StatusBadRequest, customErrors.ErrorBadRequest.Error() )
		}
		buyer,err := handler.sv.GetById(id)
		if  err!= nil {
			response.Error(w,http.StatusNotFound, customErrors.ErrorNotFound.Error() )
		}
		response.JSON(w,http.StatusOK,buyer)
	}
}

func (handler *BuyerHandler)CreateBuyer()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var buyerDoc models.BuyerDoc
		if err := json.NewDecoder(r.Body).Decode(&buyerDoc); err != nil{
			response.Error(w, http.StatusBadRequest, customErrors.ErrorBadRequest.Error())
			return	
		}
		newBuyer := mappers.BuyerDocToBuyer(buyerDoc)
		err := handler.sv.CreateBuyer(newBuyer)
		if err != nil {
			
		}
	}
}