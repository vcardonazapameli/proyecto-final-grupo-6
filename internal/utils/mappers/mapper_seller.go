package mappers

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

func SellerDocToSeller(sellerDoc models.SellerDoc) models.Seller {
	return *models.NewSeller(
		sellerDoc.Id,
		sellerDoc.Cid,
		sellerDoc.CompanyName,
		sellerDoc.Address,
		sellerDoc.Telephone,
		sellerDoc.LocalityID,
	)
}

func SellerToSellerDoc(seller models.Seller) models.SellerDoc {
	return *models.NewSellerDoc(
		seller.Id,
		seller.Cid,
		seller.CompanyName,
		seller.Address,
		seller.Telephone,
		seller.LocalityID,
	)
}
