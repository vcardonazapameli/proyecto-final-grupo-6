package seller

import (
	loader "github.com/arieleon_meli/proyecto-final-grupo-6/internal/loader/seller"
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
