package models

type Employee struct {
	Id                 int
	EmployeeAttributes EmployeeAttributes
}

type EmployeeAttributes struct {
	CardNumberID string
	FirstName    string
	LastName     string
	WarehouseID  int
}

type EmployeeDoc struct {
	Id           int    `json:"id"`
	CardNumberID string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseID  int    `json:"warehouse_id"`
}

type RequestEmployee struct {
	CardNumberID string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseID  int    `json:"warehouse_id"`
}
