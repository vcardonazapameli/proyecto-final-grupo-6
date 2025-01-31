package validators

import (
	"reflect"

	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func ValidateBuyerEmpty(buyer models.BuyerAttributes)(valid bool){
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
func ValidateBuyerTypes(buyer models.BuyerAttributes)(valid bool){
	
	if reflect.TypeOf(buyer.CardNumberId) == reflect.TypeOf(int(0)){
		return 
	}
	if reflect.TypeOf(buyer.FirstName) == reflect.TypeOf(string("")){
		return 
	}
	if reflect.TypeOf(buyer.CardNumberId) == reflect.TypeOf(string("")){
		return 
	}
	return true
}