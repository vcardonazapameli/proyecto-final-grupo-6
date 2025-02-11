package validators

import (
	defaultErrors "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func ValidateCapacity(section models.Section) error {
	var messages []string

	if section.MaximumCapacity <= section.MinimumCapacity {
		messages = append(messages, "maximum capacity must be greater than minimum capacity")
	}
	if section.CurrentCapacity < section.MinimumCapacity || section.CurrentCapacity > section.MaximumCapacity {
		messages = append(messages, "current capacity must be between minimum and maximum capacity")
	}
	if section.MinimumCapacity < 0 {
		messages = append(messages, "minimum capacity cannot be negative")
	}

	if len(messages) > 0 {
		return defaultErrors.ValidationError{Messages: messages}
	}
	return nil
}

func ValidateTemperature(section models.Section) error {
	if section.CurrentTemperature < section.MinimumTemperature {
		return defaultErrors.ValidationError{Messages: []string{"current temperature cannot be below minimum temperature"}}
	}
	return nil
}
