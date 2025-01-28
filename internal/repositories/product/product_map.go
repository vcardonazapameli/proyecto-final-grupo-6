package product

import (
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func NewProductMap(db map[int]models.Product) ProductRepository {
	defaultDb := make(map[int]models.Product)
	if db != nil {
		defaultDb = db
	}
	return &ProductMap{db: defaultDb}
}

type ProductMap struct {
	db map[int]models.Product
}
