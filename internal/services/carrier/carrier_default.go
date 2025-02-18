package carrier

import (
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/carrier"
	errorCustom "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/validators"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func NewCarrierService(repository repository.CarrierRepository) CarrierService {
	return &carrierService{repository: repository}
}

type carrierService struct {
	repository repository.CarrierRepository
}

func (carrierService *carrierService) Create(carrierDocRequest models.CarrierDocRequest) (*models.CarrierDocResponse, error) {
	if errorValidateFields := validators.ValidateFieldsCarrier(carrierDocRequest); errorValidateFields != nil {
		return nil, errorValidateFields
	}
	existInDb, err := carrierService.repository.ExistInDb(carrierDocRequest.Cid)
	if err != nil {
		return nil, err
	}
	if existInDb {
		return nil, errorCustom.ErrorConflict
	}
	idLocalityExist, err := carrierService.repository.ExistLocalityInDb(carrierDocRequest.Locality_id)
	if err != nil {
		return nil, err
	}
	if !idLocalityExist {
		return nil, errorCustom.ErrorConflict
	}
	carrier := mappers.CarrierDocRequestToCarrierDocResponse(carrierDocRequest)
	if err := carrierService.repository.Create(&carrier); err != nil {
		return nil, nil
	}
	return &carrier, nil
}
