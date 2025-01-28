package product

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

func NewProductJSONFile(path string) *ProductJSONFile {
	return &ProductJSONFile{
		path: path,
	}
}

type ProductJSONFile struct {
	path string
}

func (l *ProductJSONFile) Load() (v map[int]models.Product, err error) {
	return
}
