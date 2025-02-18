package mappers

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

func BuyerDocToBuyer(buyerDoc models.BuyerDocResponse) models.Buyer {
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
func BuyerDocRequestToBuyerDocResponse(buyer models.BuyerDocRequest) models.BuyerDocResponse {
	return models.BuyerDocResponse{
		CardNumberId: buyer.CardNumberId,
		FirstName:    buyer.FirstName,
		LastName:     buyer.LastName,
	}
}
func BuyerDocResponseToBuyerDocRequest(buyer models.BuyerDocResponse) models.BuyerDocRequest {
	return models.BuyerDocRequest{
		CardNumberId: buyer.CardNumberId,
		FirstName:    buyer.FirstName,
		LastName:     buyer.LastName,
	}
}
