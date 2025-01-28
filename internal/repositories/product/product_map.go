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

func (r *ProductMap) GetAll() (map[int]models.Product, error) {
	return r.db, nil
}

func (r *ProductMap) GetById(id int) (*models.Product, error) {
	product, exist := r.db[id]
	if !exist {
		return nil, nil
	}
	return &product, nil
}
