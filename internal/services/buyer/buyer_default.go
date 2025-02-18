package buyer

import (
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/buyer"
	customErrors "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
	validators "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/validators"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func NewBuyerDefault(rp repository.BuyerRepository) BuyerService {
	return &BuyerDefault{rp: rp}
}

type BuyerDefault struct {
	rp repository.BuyerRepository
}

// UpdateBuyer implements BuyerService.
func (b *BuyerDefault) UpdateBuyer(id int, buyerRequest models.UpdateBuyerDto) (*models.BuyerDocResponse, error) {
	buyerToUpdate, err := b.GetById(id)
	if err != nil {
		return nil, err
	}

	updatedBuyer := validators.UpdateEntity(buyerRequest, buyerToUpdate)
	if b.rp.ValidateCardNumberIdToUpdate(updatedBuyer.CardNumberId, id) {
		return nil, customErrors.ErrorConflict
	}
	buyerDoc := mappers.BuyerDocResponseToBuyerDocRequest(*updatedBuyer)
	if err := b.rp.UpdateBuyer(id, &buyerDoc); err != nil {
		return nil, err
	}
	return updatedBuyer, nil
}

// DeleteBuyer implements BuyerService.
func (b *BuyerDefault) DeleteBuyer(buyerId int) error {
	if !b.rp.ValidateIfExistsById(buyerId) {
		return customErrors.ErrorNotFound
	}
	b.rp.DeleteBuyer(buyerId)
	return nil
}

// CreateBuyer implements BuyerService.
func (b *BuyerDefault) CreateBuyer(buyer models.BuyerDocRequest) (*models.BuyerDocResponse, error) {

	if err := validators.ValidateNoEmptyFields(buyer); err != nil {
		return nil, customErrors.ErrorUnprocessableContent
	}

	if b.rp.ValidateCardNumberId(buyer.CardNumberId) {
		return nil, customErrors.ErrorConflict
	}
	buyerDocResponse := mappers.BuyerDocRequestToBuyerDocResponse(buyer)

	err := b.rp.CreateBuyer(&buyerDocResponse)
	if err != nil {
		return nil, err
	}

	return &buyerDocResponse, nil
}

// GetById implements BuyerService.
func (b *BuyerDefault) GetById(id int) (*models.BuyerDocResponse, error) {
	buyer, err := b.rp.GetById(id)
	if err != nil {
		return nil, customErrors.ErrorNotFound
	}
	return buyer, nil
}

// GetAll implements BuyerService.
func (b *BuyerDefault) GetAll() ([]models.BuyerDocResponse, error) {
	buyers, err := b.rp.GetAll()
	if err != nil {
		return nil, err
	}
	return buyers, nil
}
func (p *BuyerDefault) GetPurchasesReports(cardNumberId int) ([]models.PurchaseOrderReport, error) {
	if cardNumberId > 0 && !p.rp.ValidateCardNumberId(cardNumberId) {
		return nil, customErrors.ErrorNotFound
	}
	purchaseReports, err := p.rp.GetPurchasesReports(cardNumberId)
	if err != nil {
		return nil, err
	}
	return purchaseReports, nil
}
