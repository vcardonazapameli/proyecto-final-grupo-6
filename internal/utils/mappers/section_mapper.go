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
		},
	}
}

func SectionAttributesToSection(attributes models.SectionAttributes, id int) models.Section {
	return models.Section{
		Id:                id,
		SectionAttributes: attributes,
	}
}

func SectionToSectionAttributes(section models.Section) models.SectionAttributes {
	return section.SectionAttributes
}
