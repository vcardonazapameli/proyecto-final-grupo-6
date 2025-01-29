package buyer

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type BuyerRepository interface{
	GetAll()map[int]models.Buyer
	GetById(id int)(models.Buyer, bool)
	CreateBuyer(buyer models.Buyer)
	ValidateCardNumberId(cardNumber int)(exists bool)
}