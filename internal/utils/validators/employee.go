package validators

import (
	"strings"

	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func validateField(field *string, fieldName string, invalidFields *[]string) {
	if field != nil && strings.TrimSpace(*field) == "" {
		*invalidFields = append(*invalidFields, fieldName+" cannot be empty")
	}
}

func validateIntField(field *int, fieldName string, invalidFields *[]string) {
	if field != nil && *field <= 0 {
		*invalidFields = append(*invalidFields, fieldName+" cannot be less than or equal to 0")
	}
}

func ValidateUpdateEmployee(request models.UpdateEmployee) error {
	var invalidFields []string

	validateField(request.CardNumberID, "CardNumberID", &invalidFields)
	validateField(request.FirstName, "FirstName", &invalidFields)
	validateField(request.LastName, "LastName", &invalidFields)
	validateIntField(request.WarehouseID, "WarehouseID", &invalidFields)

	if len(invalidFields) > 0 {
		return customErrors.ValidationError{Messages: invalidFields}
	}
	return nil
}

func ValidateCreateEmployee(request models.RequestEmployee) error {
	var invalidFields []string

	cardNumberID := &request.CardNumberID
	firstName := &request.FirstName
	lastName := &request.LastName
	warehouseID := &request.WarehouseID

	validateField(cardNumberID, "CardNumberID", &invalidFields)
	validateField(firstName, "FirstName", &invalidFields)
	validateField(lastName, "LastName", &invalidFields)
	validateIntField(warehouseID, "WarehouseID", &invalidFields)

	if len(invalidFields) > 0 {
		return customErrors.ValidationError{Messages: invalidFields}
	}
	return nil
}
