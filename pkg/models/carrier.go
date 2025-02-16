package models

type Carrier struct {
	Id int
	CarrierAttributes
}

type CarrierAttributes struct {
	Cid          string
	Company_name string
	Address      string
	Telephone    string
	Locality_id  string
}

type CarrierDocResponse struct {
	Id           int    `json:"id"`
	Cid          string `json:"cid"`
	Company_name string `json:"company_name"`
	Address      string `json:"address"`
	Telephone    string `json:"telephone"`
	Locality_id  int    `json:"locality_id"`
}

type CarrierDocRequest struct {
	Cid          string `json:"cid"`
	Company_name string `json:"company_name"`
	Address      string `json:"address"`
	Telephone    string `json:"telephone"`
	Locality_id  int    `json:"locality_id"`
}
