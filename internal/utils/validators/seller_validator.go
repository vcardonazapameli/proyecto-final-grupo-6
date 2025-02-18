package validators

import (
	e "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func ValidateSellerAttrs(s models.SellerDoc) error {
	messages := make([]string, 0)
	if s.Cid <= 0 {
		messages = append(messages, "CID must not be negative nor zero")
	}
	if s.CompanyName == "" {
		messages = append(messages, "Company Name cannot be empty")
	}
	if s.Address == "" {
		messages = append(messages, "Company Address cannot be empty")
	}
	if len(s.Telephone) < 8 || len(s.Telephone) > 10 {
		messages = append(messages, "Wrong telephone format. Must have between 8 and 10 digits")
	}

	if len(messages) > 0 {
		return e.ValidationError{Messages: messages}
	}
	return nil
}

func ValidateSellerAttrPointers(cid *int, companyName *string, address *string, telephone *string, localityId *int) error {
	messages := make([]string, 0)
	if cid != nil {
		if *cid <= 0 {
			messages = append(messages, "CID must not be negative nor zero")
		}
	}
	if companyName != nil {
		if *companyName == "" {
			messages = append(messages, "Company Name cannot be empty")
		}
	}
	if address != nil {
		if *address == "" {
			messages = append(messages, "Company Address cannot be empty")
		}
	}
	if telephone != nil {
		if len(*telephone) < 8 || len(*telephone) > 10 {
			messages = append(messages, "Wrong telephone format. Must have between 8 and 10 digits")
		}
	}
	if localityId != nil {
		if *localityId <= 0 {
			messages = append(messages, "Locality ID must not be negative nor zero")
		}
	}
	if len(messages) > 0 {
		return e.ValidationError{Messages: messages}
	}
	return nil
}
