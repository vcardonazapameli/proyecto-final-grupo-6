package mappers

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

func BuyerDocToBuyer(buyerDoc models.BuyerDoc) models.Buyer {
	return models.Buyer{
		Id: buyerDoc.Id,
		BuyerAttributes: models.BuyerAttributes{
			CardNumberId: buyerDoc.CardNumberId,
			FirstName:    buyerDoc.FirstName,
			LastName:     buyerDoc.LastName,
		},
	}
}
func BuyerDocToBuyerAttributes(buyerDoc models.CreateBuyerDto) models.BuyerAttributes {
	return models.BuyerAttributes{
		CardNumberId: buyerDoc.CardNumberId,
		FirstName:    buyerDoc.FirstName,
		LastName:     buyerDoc.LastName,
	}
}
func BuyerToBuyerDoc(buyer models.Buyer) models.BuyerDoc {
	return models.BuyerDoc{
		Id:           buyer.Id,
		CardNumberId: buyer.CardNumberId,
		FirstName:    buyer.FirstName,
		LastName:     buyer.LastName,
	}
}
