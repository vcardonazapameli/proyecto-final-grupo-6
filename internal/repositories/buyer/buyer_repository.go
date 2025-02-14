package buyer

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type BuyerRepository interface{
	GetAll()([]models.BuyerDocResponse, error)
	GetById(id int)(*models.BuyerDocResponse, error)
	CreateBuyer(buyer *models.BuyerDocResponse)error
	DeleteBuyer(buyerId int)error
	UpdateBuyer(id int , buyer *models.BuyerDocRequest)error
	ValidateCardNumberId(cardNumber int)(bool)
	ValidateCardNumberIdToUpdate(cardNumber, id int)(bool)
	ValidateIfExistsById(id int)(bool)
	GetPurchasesReports(cardNumberId int)([]models.PurchaseOrderReport,error)
}