package seller

import (
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

type SellerService interface {
	GetAll() (map[int]models.SellerDoc, error)
	Create(cid int, companyName string, address string, telephone int) (models.SellerDoc, error)
	GetByID(id int) (models.SellerDoc, error)
	Delete(id int) error
	Update(id int, cid *int, companyName *string, address *string, telephone *int) (models.SellerDoc, error)
}
