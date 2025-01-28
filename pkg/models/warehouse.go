package models

type Warehouse struct {
	Id uint64
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

type WarehouseDoc struct {
	ID                  uint64  `json:"id"`
	Warehouse_code      string  `json:"warehouse_code"`
	Address             string  `json:"addres"`
	Telephone           string  `json:"telephone"`
	Minimun_capacity    uint64  `json:"minimun_capacity"`
	Minimun_temperature float64 `json:"minimun_temperature"`
	Locality_id         uint64  `json:"locality_id"`
}
