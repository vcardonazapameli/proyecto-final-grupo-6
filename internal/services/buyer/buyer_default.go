package buyer

import (
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/buyer"
	customErrors "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/errors"
	validators "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/validators"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func NewBuyerDefault(rp repository.BuyerRepository) BuyerService {
	return &BuyerDefault{rp: rp}
}

type BuyerDefault struct {
	rp repository.BuyerRepository
}

// CreateBuyer implements BuyerService.
func (b *BuyerDefault) CreateBuyer(buyer models.Buyer)error {
	
	if b.rp.ValidateCardNumberId(buyer.CardNumberId){
		return customErrors.ErrorConflict
	}
	if !validators.ValidateBuyer(buyer){
		return customErrors.ErrorBadRequest
	}
	b.rp.CreateBuyer(buyer)
	return nil
}

// GetById implements BuyerService.
func (b *BuyerDefault) GetById(id int) (models.Buyer, error) {
	buyer, exists := b.rp.GetById(id)
	if !exists {
		return models.Buyer{}, customErrors.ErrorNotFound
	}
	return buyer, nil
}

// GetAll implements BuyerService.
func (b *BuyerDefault) GetAll() (map[int]models.Buyer, error) {
	buyers := b.rp.GetAll()
	if len(buyers) == 0 {
		return buyers, customErrors.ErrorNotFound
	}
	return buyers, nil
}
