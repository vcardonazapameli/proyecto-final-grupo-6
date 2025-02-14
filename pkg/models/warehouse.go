package models

type Warehouse struct {
	Id int
	WarehouseAttributes
}

type WarehouseAttributes struct {
	Warehouse_code      string
	Address             string
	Telephone           string
	Minimun_capacity    uint64
	Minimun_temperature float64
	Locality_id         uint64
}

type WarehouseDocRequest struct {
	Warehouse_code      string  `json:"warehouse_code"`
	Address             string  `json:"address"`
	Telephone           string  `json:"telephone"`
	Minimun_capacity    uint64  `json:"minimun_capacity"`
	Minimun_temperature float64 `json:"minimun_temperature"`
	Locality_id         uint64  `json:"locality_id"`
}

type WarehouseDocResponse struct {
	ID                  int     `json:"id"`
	Warehouse_code      string  `json:"warehouse_code"`
	Address             string  `json:"address"`
	Telephone           string  `json:"telephone"`
	Minimun_capacity    uint64  `json:"minimun_capacity"`
	Minimun_temperature float64 `json:"minimun_temperature"`
	Locality_id         uint64  `json:"locality_id"`
}

type WarehouseUpdateDocRequest struct {
	ID                  *int     `json:"id"`
	Warehouse_code      *string  `json:"warehouse_code"`
	Address             *string  `json:"address"`
	Telephone           *string  `json:"telephone"`
	Minimun_capacity    *uint64  `json:"minimun_capacity"`
	Minimun_temperature *float64 `json:"minimun_temperature"`
	Locality_id         *uint64  `json:"locality_id"`
}
