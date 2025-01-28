package models

type Section struct {
	Id int
	SectionAttributes
}

type SectionAttributes struct {
	SectionNumber      string
	CurrentCapacity    int
	CurrentTemperature float64
	MaximumCapacity    int
	MinimumCapacity    int
	MinimunTemperature float64
	ProductType        int
	WarehouseId        int
}

type SectionDoc struct {
	Id                 int     `json:"id"`
	SectionNumber      string  `json:"section_number"`
	CurrentCapacity    int     `json:"current_capacity"`
	CurrentTemperature float64 `json:"current_temperature"`
	MaximumCapacity    int     `json:"maximum_capacity"`
	MinimumCapacity    int     `json:"minimum_capacity"`
	MinimunTemperature float64 `json:"minimum_temperature"`
	ProductType        int     `json:"product_type_id"`
	WarehouseId        int     `json:"warehouse_id"`
}
