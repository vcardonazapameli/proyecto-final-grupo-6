package seller

import (
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/seller"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/errors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

type SellerServiceDefault struct {
	rp repository.SellerRepository
}

func NewSellerServiceDefault(rp repository.SellerRepository) *SellerServiceDefault {
	return &SellerServiceDefault{rp}
}

func (sv *SellerServiceDefault) GetAll() (map[int]models.SellerDoc, error) {
	s := make(map[int]models.SellerDoc)
	sellers, err := sv.rp.GetAll()

	if err != nil {
		return s, errors.ErrorInternalServerError
	}

	for _, sel := range sellers {
		sDoc := mappers.SellerToSellerDoc(sel)
		s[sDoc.Id] = sDoc
	}
	return s, nil
}

func (sv *SellerServiceDefault) Create(cid int, companyName string, address string, telephone int) (models.SellerDoc, error) {
	// Validate Seller
	if err := ValidateSeller(cid, companyName, address, telephone); err != nil {
		return models.SellerDoc{}, err
	}

	new, err := sv.rp.Save(cid, companyName, address, telephone)
	if err != nil {
		return models.SellerDoc{}, err
	}

	newDoc := mappers.SellerToSellerDoc(new)
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

func (sv *SellerServiceDefault) Update(id int, cid *int, companyName *string, address *string, telephone *int) (models.SellerDoc, error) {
	seller, err := sv.rp.GetByID(id)
	if err != nil {
		return models.SellerDoc{}, err
	}

	valError := make([]string, 0)
	if cid != nil {
		if *cid <= 0 {
			valError = append(valError, "CID must not be negative nor zero")
		} else {
			seller.Cid = *cid
		}
	}
	if companyName != nil {
		if *companyName == "" {
			valError = append(valError, "Company Name cannot be empty")
		} else {
			seller.CompanyName = *companyName
		}
	}
	if address != nil {
		if *address == "" {
			valError = append(valError, "Company Address cannot be empty")
		} else {
			seller.Address = *address
		}
	}
	if telephone != nil {
		if *telephone < 10000000 || *telephone > 99999999 {
			valError = append(valError, "Wrong telephone format. Must have between 8 and 10 digits")
		} else {
			seller.Telephone = *telephone
		}
	}
	if len(valError) > 0 {
		return models.SellerDoc{}, ValidationError{valError}
	}

	sv.rp.Update(seller)
	return mappers.SellerToSellerDoc(seller), nil
}
