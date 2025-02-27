package mappers

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

func SectionDocRequestToSectioDocResponse(section models.SectionDocRequest) models.SectionDoc {
	return models.SectionDoc{
		SectionNumber:      section.SectionNumber,
		CurrentTemperature: section.CurrentTemperature,
		MinimumTemperature: section.MinimumTemperature,
		CurrentCapacity:    section.CurrentCapacity,
		MinimumCapacity:    section.MinimumCapacity,
		MaximumCapacity:    section.MaximumCapacity,
		WarehouseId:        section.WarehouseId,
		ProductTypeId:      section.ProductTypeId,
	}
}
func SectionDocResponseToSectionDocRequest(section models.SectionDoc) models.SectionDocRequest {
	return models.SectionDocRequest{
		SectionNumber:      section.SectionNumber,
		CurrentTemperature: section.CurrentTemperature,
		MinimumTemperature: section.MinimumTemperature,
		CurrentCapacity:    section.CurrentCapacity,
		MinimumCapacity:    section.MinimumCapacity,
		MaximumCapacity:    section.MaximumCapacity,
		WarehouseId:        section.WarehouseId,
		ProductTypeId:      section.ProductTypeId,
	}
}
