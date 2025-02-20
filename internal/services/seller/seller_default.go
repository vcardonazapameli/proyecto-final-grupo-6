package seller

import (
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/seller"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/validators"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

type SellerServiceDefault struct {
	rp repository.SellerRepository
}

func NewSellerServiceDefault(rp repository.SellerRepository) SellerService {
	return &SellerServiceDefault{rp}
}

func (sv *SellerServiceDefault) GetAll() (map[int]models.SellerDoc, error) {
	s := make(map[int]models.SellerDoc)
	sellers, err := sv.rp.GetAll()

	if err != nil {
		return s, customErrors.ErrorInternalServerError
	}

	for _, sel := range sellers {
		sDoc := mappers.SellerToSellerDoc(sel)
		s[sDoc.Id] = sDoc
	}
	return s, nil
}

func (sv *SellerServiceDefault) Create(sDoc models.SellerDoc) (models.SellerDoc, error) {
	// Validate Seller
	if err := validators.ValidateSellerAttrs(sDoc); err != nil {
		return models.SellerDoc{}, err
	}

	seller := mappers.SellerDocToSeller(sDoc)
	err := sv.rp.Save(&seller)
	if err != nil {
		return models.SellerDoc{}, err
	}

	newDoc := mappers.SellerToSellerDoc(seller)
	return newDoc, nil

}

func (sv *SellerServiceDefault) GetByID(id int) (models.SellerDoc, error) {
	s, err := sv.rp.GetByID(id)
	if err != nil {
		return models.SellerDoc{}, err
	}

	return mappers.SellerToSellerDoc(s), nil
}

func (sv *SellerServiceDefault) Delete(id int) error {
	return sv.rp.Delete(id)
}

func (sv *SellerServiceDefault) Update(id int, cid *int, companyName *string, address *string, telephone *string, localityId *int) (models.SellerDoc, error) {
	seller, err := sv.rp.GetByID(id)
	if err != nil {
		return models.SellerDoc{}, err
	}

	if err := validators.ValidateSellerAttrPointers(cid, companyName, address, telephone, localityId); err != nil {
		return models.SellerDoc{}, err
	}

	if cid != nil {
		// If exists a Seller with same CID that the one trying to update, raise conflcit.
		if s, exists := sv.rp.SearchByCID(*cid); exists && s.Id != id {
			return models.SellerDoc{}, customErrors.ErrorConflict
		}

		seller.Cid = *cid
	}
	if companyName != nil {
		seller.CompanyName = *companyName

	}
	if address != nil {
		seller.Address = *address

	}
	if telephone != nil {
		seller.Telephone = *telephone
	}

	if localityId != nil {
		seller.LocalityID = *localityId
	}

	err = sv.rp.Update(seller)
	return mappers.SellerToSellerDoc(seller), err
}
