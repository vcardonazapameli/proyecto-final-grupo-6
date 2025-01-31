package mappers

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

func SectionToSectionDoc(sections models.Section) models.SectionDoc {
	return models.SectionDoc{
		Id:                 sections.Id,
		SectionNumber:      sections.SectionNumber,
		CurrentTemperature: sections.CurrentTemperature,
		MinimumTemperature: sections.MinimumTemperature,
		CurrentCapacity:    sections.CurrentCapacity,
		MinimumCapacity:    sections.MinimumCapacity,
		MaximumCapacity:    sections.MaximumCapacity,
		WarehouseId:        sections.WarehouseId,
		ProductTypeId:      sections.ProductTypeId,
		ProductBatchId:     sections.ProductBatchId,
	}
}

func SectionDocToSection(sections models.SectionDoc) models.Section {
	return models.Section{
		SectionAttributes: models.SectionAttributes{
			SectionNumber:      sections.SectionNumber,
			CurrentTemperature: sections.CurrentTemperature,
			MinimumTemperature: sections.MinimumTemperature,
			CurrentCapacity:    sections.CurrentCapacity,
			MinimumCapacity:    sections.MinimumCapacity,
			MaximumCapacity:    sections.MaximumCapacity,
			WarehouseId:        sections.WarehouseId,
			ProductTypeId:      sections.ProductTypeId,
			ProductBatchId:     sections.ProductBatchId,
		},
	}
}

func SectionToSectionValidation(section models.Section) models.SectionValidation {
	return models.SectionValidation{
		SectionNumber:      section.SectionNumber,
		CurrentCapacity:    section.CurrentCapacity,
		CurrentTemperature: section.CurrentTemperature,
		MaximumCapacity:    section.MaximumCapacity,
		MinimumCapacity:    section.MinimumCapacity,
		MinimumTemperature: section.MinimumTemperature,
		ProductTypeId:      section.ProductTypeId,
		WarehouseId:        section.WarehouseId,
		ProductBatchId:     section.ProductBatchId,
	}
}
