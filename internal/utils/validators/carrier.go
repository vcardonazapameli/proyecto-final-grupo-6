package validators

import (
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func ValidateFieldsCarrier(carrierDocRequest models.CarrierDocRequest) error {

	messages := make([]string, 0)
	if len(carrierDocRequest.Cid) < 5 {
		messages = append(messages, "cid must be at least 5 characters long")
	}
	if len(carrierDocRequest.Telephone) < 8 || len(carrierDocRequest.Telephone) > 10 {
		messages = append(messages, "The phone length must be greater than or equal to 8 and less than or equal to 10.")
	}
	if len(messages) > 0 {
		return customErrors.ValidationError{Messages: messages}
	}
	return nil
}
