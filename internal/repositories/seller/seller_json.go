package seller

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type SellerRepositoryJSON struct {
	db       map[int]models.Seller
	jsonPath string
}

func NewSellerRepositoryJSON(db map[int]models.Seller, path string) *SellerRepositoryJSON {
	// default db
	defaultDb := make(map[int]models.Seller)
	if db != nil {
		defaultDb = db
	}
	return &SellerRepositoryJSON{db: defaultDb, jsonPath: path}
}
