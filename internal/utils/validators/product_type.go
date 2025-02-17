package validators

import (
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func ValidateFieldsProductType(productTypeDocRequest models.ProductTypeDocRequest) error {

	messages := make([]string, 0)
	if len(productTypeDocRequest.Description) == 0 {
		messages = append(messages, "description must not be empty")
	}
	if len(messages) > 0 {
		return customErrors.ValidationError{Messages: messages}
	}
	return nil
}
