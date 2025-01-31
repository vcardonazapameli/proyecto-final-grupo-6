package validators

import (
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/errors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func ValidateFieldsProducts(productDocRequest models.ProductDocRequest) error {
	messages := make([]string, 0)
	if len(productDocRequest.ProductCode) < 4 {
		messages = append(messages, "product_code must be at least 5 characters long")
	}
	if productDocRequest.ExpirationRate < 0 {
		messages = append(messages, "expiration_rate cannot be negative")
	}
	if productDocRequest.RecommendedFreezingTemperature > -20 {
		messages = append(messages, "recommended_freezing_temperature must be less than or equal to -20.0")
	}
	if productDocRequest.FreezingRate > -20 {
		messages = append(messages, "freezing_rate must be less than or equal to -20.0")
	}
	if productDocRequest.Width < 0 {
		messages = append(messages, "width cannot be negative")
	}
	if productDocRequest.Height < 0 {
		messages = append(messages, "height cannot be negative")
	}
	if productDocRequest.Length < 0 {
		messages = append(messages, "length cannot be negative")
	}
	if productDocRequest.NetWeight < 0 {
		messages = append(messages, "net_weight cannot be negative")
	}
	if productDocRequest.ProductType < 0 {
		messages = append(messages, "product_type_id cannot be negative")
	}
	if productDocRequest.Seller < 0 {
		messages = append(messages, "seller_id cannot be negative")
	}
	if len(messages) > 0 {
		return errors.ValidationError{Messages: messages}
	}
	return nil
}
