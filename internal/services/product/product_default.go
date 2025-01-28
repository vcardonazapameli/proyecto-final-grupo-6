package product

import (
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product"
)

func NewProductDefault(rp repository.ProductRepository) ProductService {
	return &ProductDefault{rp: rp}
}

type ProductDefault struct {
	rp repository.ProductRepository
}
