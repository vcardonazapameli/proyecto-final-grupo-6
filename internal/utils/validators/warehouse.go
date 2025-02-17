package validators

import (
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func ValidateFieldsWarehouseCreate(warehouse models.WarehouseDocRequest) error {
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
		messages = append(messages, "The phone length must be greater than or equal to 8 and less than or equal to 10")
	}
	if len(messages) > 0 {
		return customErrors.ValidationError{Messages: messages}
	}
	return nil
}

func ValidateFieldsWarehouseUpdate(warehouse models.WarehouseDocRequest) error {
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
