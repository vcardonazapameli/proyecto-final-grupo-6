package buyer

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"


type BuyerService interface{
	GetAll()([]models.BuyerDocResponse,error)
	GetById(id int)(*models.BuyerDocResponse, error)
	CreateBuyer(buyer models.BuyerDocRequest)(*models.BuyerDocResponse, error)
	DeleteBuyer(buyerId int)(error)
	UpdateBuyer(id int , buyerRequest models.UpdateBuyerDto)(*models.BuyerDocResponse, error)
	GetPurchasesReports(CardNumberId int)([]models.PurchaseOrderReport, error)
}