package validators

import (
	"strings"

	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func ValidateFieldsWarehouse(warehouse models.Warehouse) error {
	messages := make([]string, 0)

	if strings.TrimSpace(warehouse.Warehouse_code) == "" {
		messages = append(messages, "warehouse_code is required")
	}
	if strings.TrimSpace(warehouse.Address) == "" {
		messages = append(messages, "address is required")
	}
	if strings.TrimSpace(warehouse.Telephone) == "" {
		messages = append(messages, "telephone is required")
	}
	if warehouse.Minimun_capacity <= 0 {
		messages = append(messages, "minimun_capacity must be greater than to 0")
	}
	if warehouse.Minimun_temperature <= -30 {
		messages = append(messages, "The minimum temperature must be greater than or equal to -30")
	}
	if warehouse.Minimun_temperature > 30 {
		messages = append(messages, "The minimum temperature must be less than or equal to 30")
	}
	if len(warehouse.Telephone) < 8 || len(warehouse.Telephone) > 10 {
		messages = append(messages, "La longitud del teléfono debe ser mayor o igual a 8 y menor o igual a 10")
	}
	if len(messages) > 0 {
		return customErrors.ValidationError{Messages: messages}
	}
	return nil
}

func ValidateFieldsUpdate(warehouse models.Warehouse) error {
	messages := make([]string, 0)

	if warehouse.Minimun_capacity <= 0 {
		messages = append(messages, "minimun_capacity must be greater than to 0")
	}
	if warehouse.Minimun_temperature <= -30 {
		messages = append(messages, "The minimum temperature must be greater than or equal to -30")
	}
	if warehouse.Minimun_temperature > 30 {
		messages = append(messages, "The minimum temperature must be less than or equal to 30")
	}
	if len(warehouse.Telephone) < 8 || len(warehouse.Telephone) > 10 {
		messages = append(messages, "La longitud del teléfono debe ser mayor o igual a 8 y menor o igual a 10")
	}
	if len(messages) > 0 {
		return customErrors.ValidationError{Messages: messages}
	}
	return nil
}
