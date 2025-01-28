package models

type Seller struct {
	Id          int
	Cid         int
	CompanyName string
	Address     string
	Telephone   int
}

type SellerDoc struct {
	Id          int    `json:"id"`
	Cid         int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   int    `json:"telephone"`
}
