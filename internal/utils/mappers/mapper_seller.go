package mappers

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

func SellerDocToSeller(sellerDoc models.SellerDoc) models.Seller {
	return models.Seller{
		Id:          sellerDoc.Id,
		Cid:         sellerDoc.Cid,
		CompanyName: sellerDoc.CompanyName,
		Address:     sellerDoc.Address,
		Telephone:   sellerDoc.Telephone,
		LocalityID:  sellerDoc.LocalityID,
	}
}

func SellerToSellerDoc(seller models.Seller) models.SellerDoc {
	return models.SellerDoc{
		Id:          seller.Id,
		Cid:         seller.Cid,
		CompanyName: seller.CompanyName,
		Address:     seller.Address,
		Telephone:   seller.Telephone,
		LocalityID:  seller.LocalityID,
	}
}
