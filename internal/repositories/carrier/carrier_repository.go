package carrier

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type CarrierRepository interface {
	Create(*models.CarrierDocResponse) error
	ExistInDb(string) (bool, error)
	ExistLocalityInDb(int) (bool, error)
	ExistCarrierInDb(int) (bool, error)
}
