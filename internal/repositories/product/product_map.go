package product

import (
	"strings"

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

func (r *ProductMap) Delete(id int) error {
	delete(r.db, id)
	return nil
}

func (r *ProductMap) Create(product models.Product) error {
	r.db[product.Id] = product
	return nil
}

func (r *ProductMap) ExistInDb(productCode string) bool {
	for _, product := range r.db {
		if strings.EqualFold(product.ProductCode, productCode) {
			return true
		}
	}
	return false
}

func (r *ProductMap) GenerateId() int {
	var assignedId int = 1
	var firstIteration bool = true
	for _, product := range r.db {
		if firstIteration || product.Id >= assignedId {
			firstIteration = false
			assignedId = product.Id + 1
		}
	}
	return assignedId
}
