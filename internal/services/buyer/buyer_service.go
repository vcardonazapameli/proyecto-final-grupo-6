package buyer

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"


type BuyerService interface{
	GetAll()(map[int]models.Buyer,error)
	GetById(id int)(models.Buyer, error)
	CreateBuyer(buyer models.Buyer)error
}