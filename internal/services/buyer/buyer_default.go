package buyer

import (
	"fmt"

	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/buyer"
	customErrors "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
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
func (b *BuyerDefault) UpdateBuyer(id int, buyerDto models.UpdateBuyerDto) (models.Buyer, error) {
	buyerToUpdate, err := b.GetById(id)
	if err != nil {
		return models.Buyer{}, err
	}
	updatedBuyer := validators.UpdateEntity(buyerDto, buyerToUpdate)
	fmt.Println("Updated buyer: ", *updatedBuyer)
	b.rp.UpdateBuyer(id, *updatedBuyer)
	return *updatedBuyer, nil
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
func (b *BuyerDefault) CreateBuyer(buyer models.Buyer) error {

	if err := validators.ValidateNoEmptyFields(buyer); err != nil {
		return customErrors.ErrorConflict
	}
	if !validators.ValidateBuyer(buyer) {
		return customErrors.ErrorBadRequest
	}
	b.rp.CreateBuyer(buyer)
	return nil
}

// GetById implements BuyerService.
func (b *BuyerDefault) GetById(id int) (*models.Buyer, error) {
	buyer, exists := b.rp.GetById(id)
	if !exists {
		return nil, customErrors.ErrorNotFound
	}
	return &buyer, nil
}

// GetAll implements BuyerService.
func (b *BuyerDefault) GetAll() (map[int]models.Buyer, error) {
	buyers := b.rp.GetAll()
	if len(buyers) == 0 {
		return buyers, customErrors.ErrorNotFound
	}
	return buyers, nil
}
