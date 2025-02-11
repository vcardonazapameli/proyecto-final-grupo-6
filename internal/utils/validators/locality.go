package validators

import (
	e "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func ValidateLocality(loc models.Locality) error {
	messages := make([]string, 0)
	if loc.Id < 0 {
		messages = append(messages, "Wrong ID format")
	}
	if loc.CountryName == "" {
		messages = append(messages, "Wrong ID format")
	}
	if loc.LocalityName == "" {
		messages = append(messages, "Wrong ID format")
	}
	if loc.ProvinceName == "" {
		messages = append(messages, "Wrong ID format")
	}

	if len(messages) > 0 {
		return e.ValidationError{Messages: messages}
	}
	return nil
}
