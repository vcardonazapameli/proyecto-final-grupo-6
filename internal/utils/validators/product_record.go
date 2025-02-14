package validators

import (
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func ValidateFieldsProductRecord(productRecordDocRequest models.ProductRecordDocRequest) error {

	messages := make([]string, 0)
	if len(productRecordDocRequest.LastUpdateDate) == 0 {
		messages = append(messages, "last_update_date must not be empty")
	}
	if productRecordDocRequest.PurchasePrice < 0 {
		messages = append(messages, "purchase_price cannot be negative")
	}
	if productRecordDocRequest.SalePrice < 0 {
		messages = append(messages, "sale_price cannot be negative")
	}
	if productRecordDocRequest.ProductId < 0 {
		messages = append(messages, "product_id cannot be negative")
	}
	if len(messages) > 0 {
		return customErrors.ValidationError{Messages: messages}
	}
	return nil
}
