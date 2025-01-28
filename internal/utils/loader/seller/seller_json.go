package seller

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

// SellerJSONFile is a Loader
type SellerJSONFile struct {
	path string
}

func NewSellerJSONFile(path string) *SellerJSONFile {
	return &SellerJSONFile{path}
}

func (l *SellerJSONFile) Load() (v map[int]models.Seller, err error) {
	return
}
