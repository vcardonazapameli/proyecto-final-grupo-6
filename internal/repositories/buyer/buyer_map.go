package buyer

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

func NewBuyerMap(db map[int]models.Buyer) BuyerRepository {
	defaultDb := make(map[int]models.Buyer)
	if db != nil {
		defaultDb = db
	}
	return &BuyerMap{db: defaultDb}
}

type BuyerMap struct {
	db map[int]models.Buyer
}

// IncrementId implements BuyerRepository.
func (b *BuyerMap) IncrementId() int {
	maxId:= 0
	for _,value := range b.db{
		if maxId < value.Id{
			maxId = value.Id
		}
	}
	return maxId + 1
}

// UpdateBuyer implements BuyerRepository.
func (b *BuyerMap) UpdateBuyer(id int, buyer models.Buyer) models.Buyer {
	b.db[id] = buyer
	return buyer
}

// ValidateIfExistsById implements BuyerRepository.
func (b *BuyerMap) ValidateIfExistsById(id int) (exists bool) {
	_, exists = b.db[id]
	return
}

// DeleteBuyer implements BuyerRepository.
func (b *BuyerMap) DeleteBuyer(buyerId int) {
	delete(b.db, buyerId)
}

// ValidateCardNumberId implements BuyerRepository.
func (b *BuyerMap) ValidateCardNumberIdToUpdate(cardNumber, id int) (exists bool) {
	for _, value := range b.db {
		if value.Id != id && cardNumber == value.CardNumberId {
			return true
		}
	}
	return
}
func (b *BuyerMap) ValidateCardNumberId(cardNumber int) (exists bool) {
	for _, value := range b.db {
		if cardNumber == value.CardNumberId {
			return true
		}
	}
	return
}

// CreateBuyer implements BuyerRepository.
func (b *BuyerMap) CreateBuyer(buyer models.BuyerAttributes) {
	id := b.IncrementId()
	newBuyer := models.Buyer{
		Id: id,
		BuyerAttributes: buyer,
	}
	b.db[id] = newBuyer
}

// GetById implements BuyerRepository.
func (b *BuyerMap) GetById(id int) (buyer models.Buyer, exists bool) {
	buyer, exists = b.db[id]
	return
}

// GetAll implements BuyerRepository.
func (b *BuyerMap) GetAll() map[int]models.Buyer {
	return b.db
}
