package buyer

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

func ValidateBuyer(buyer models.Buyer)(valid bool){
	switch{
		case buyer.CardNumberId == 0 :
			return 
		case buyer.FirstName == "" :
			return 
		case buyer.LastName == "" :
			return 
	}
	return true
}