package mappers

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

func SectionToSectionDoc(sections map[int]models.Section) map[int]models.SectionDoc {
	data := make(map[int]models.SectionDoc)
	for _, value := range sections {
		data[value.Id] = models.SectionDoc{
			Id:                 value.Id,
			SectionNumber:      value.SectionNumber,
			CurrentTemperature: value.CurrentTemperature,
			MinimumTemperature: value.MinimumTemperature,
			CurrentCapacity:    value.CurrentCapacity,
			MinimumCapacity:    value.MinimumCapacity,
			MaximumCapacity:    value.MaximumCapacity,
			WarehouseId:        value.WarehouseId,
			ProductTypeId:      value.ProductTypeId,
			ProductBatchId:     value.ProductBatchId,
		}
	}
	return data
}
