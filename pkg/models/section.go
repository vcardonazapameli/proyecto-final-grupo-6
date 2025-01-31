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
	MinimumTemperature float64
	ProductTypeId      int
	WarehouseId        int
	ProductBatchId     []int
}

type SectionDoc struct {
	Id                 int     `json:"id"`
	SectionNumber      string  `json:"section_number"`
	CurrentCapacity    int     `json:"current_capacity"`
	CurrentTemperature float64 `json:"current_temperature"`
	MaximumCapacity    int     `json:"maximum_capacity"`
	MinimumCapacity    int     `json:"minimum_capacity"`
	MinimumTemperature float64 `json:"minimum_temperature"`
	ProductTypeId      int     `json:"product_type_id"`
	WarehouseId        int     `json:"warehouse_id"`
	ProductBatchId     []int   `json:"product_batch_id"`
}

type SectionValidation struct {
	SectionNumber      string  `json:"section_number"`
	CurrentCapacity    int     `json:"current_capacity"`
	CurrentTemperature float64 `json:"current_temperature"`
	MaximumCapacity    int     `json:"maximum_capacity"`
	MinimumCapacity    int     `json:"minimum_capacity"`
	MinimumTemperature float64 `json:"minimum_temperature"`
	ProductTypeId      int     `json:"product_type_id"`
	WarehouseId        int     `json:"warehouse_id"`
	ProductBatchId     []int   `json:"product_batch_id"`
}

type UpdateSectionDto struct {
	SectionNumber      *string  `json:"section_number"`
	CurrentCapacity    *int     `json:"current_capacity"`
	CurrentTemperature *float64 `json:"current_temperature"`
	MaximumCapacity    *int     `json:"maximum_capacity"`
	MinimumCapacity    *int     `json:"minimum_capacity"`
	MinimumTemperature *float64 `json:"minimum_temperature"`
	ProductTypeId      *int     `json:"product_type_id"`
	WarehouseId        *int     `json:"warehouse_id"`
	ProductBatchId     *[]int   `json:"product_batch_id"`
}
