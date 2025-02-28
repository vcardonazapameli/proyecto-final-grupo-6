package mappers

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

func SellerDocToSeller(sellerDoc models.SellerDoc) models.Seller {

	return models.Seller(sellerDoc)
}

func SellerToSellerDoc(seller models.Seller) models.SellerDoc {
	return models.SellerDoc(seller)
}
