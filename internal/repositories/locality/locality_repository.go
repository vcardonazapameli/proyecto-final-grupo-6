package locality

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type LocalityRepository interface {
	Save(locId int, locName string, provId int) error
	// Searches if there exists a Province with matching provinceName and related with a Country with matching countryName
	GetProvinceWithCountryNames(provinceName string, countryName string) (models.Province, error)
}
