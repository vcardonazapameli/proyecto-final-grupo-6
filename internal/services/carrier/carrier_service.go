package carrier

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type CarrierService interface {
	Create(models.CarrierDocRequest) (*models.CarrierDocResponse, error)
}
