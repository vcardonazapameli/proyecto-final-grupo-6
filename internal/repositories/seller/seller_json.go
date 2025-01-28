package seller

import (
	loader "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/loader/seller"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

type SellerRepositoryJSON struct {
	db     map[int]models.Seller
	loader loader.SellerLoader
}

func NewSellerRepositoryJSON(db map[int]models.Seller, loader loader.SellerLoader) *SellerRepositoryJSON {
	// default db
	defaultDb := make(map[int]models.Seller)
	if db != nil {
		defaultDb = db
	}
	return &SellerRepositoryJSON{db: defaultDb, loader: loader}
}

func (r *SellerRepositoryJSON) GetAll() (s map[int]models.Seller, err error) {
	return r.db, err
}

func (r *SellerRepositoryJSON) searchByCID(cid int) (models.Seller, bool) {
	for _, s := range r.db {
		if s.Cid == cid {
			return s, true
		}
	}
	return models.Seller{}, false
}

func (r *SellerRepositoryJSON) GetBiggestID() (max int) {
	for _, s := range r.db {
		if s.Id > max {
			max = s.Id
		}
	}
	return
}

func (r *SellerRepositoryJSON) Save(cid int, companyName string, address string, telephone int) (models.Seller, error) {
	if _, exists := r.searchByCID(cid); exists {
		return models.Seller{}, ExistingCIdError
	}

	// If does not exist, save.
	id := r.GetBiggestID() + 1
	newSeller := *models.NewSeller(id, cid, companyName, address, telephone)
	r.db[id] = newSeller
	return newSeller, nil
}
