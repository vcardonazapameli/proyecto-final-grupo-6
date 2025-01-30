package validators

import e "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/errors"

func ValidateSellerAttrs(cid int, companyName string, address string, telephone int) error {
	messages := make([]string, 0)
	if cid <= 0 {
		messages = append(messages, "CID must not be negative nor zero")
	}
	if companyName == "" {
		messages = append(messages, "Company Name cannot be empty")
	}
	if address == "" {
		messages = append(messages, "Company Address cannot be empty")
	}
	if telephone < 10000000 || telephone > 99999999 {
		messages = append(messages, "Wrong telephone format. Must have between 8 and 10 digits")
	}

	if len(messages) > 0 {
		return e.ValidationError{Messages: messages}
	}
	return nil
}

func ValidateSellerAttrPointers(cid *int, companyName *string, address *string, telephone *int) error {
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
		if *telephone < 10000000 || *telephone > 99999999 {
			messages = append(messages, "Wrong telephone format. Must have between 8 and 10 digits")
		}
	}
	if len(messages) > 0 {
		return e.ValidationError{Messages: messages}
	}
	return nil
}
