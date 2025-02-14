package models

type LocalityDoc struct {
	Id           int    `json:"id"`
	LocalityName string `json:"locality_name"`
	ProvinceName string `json:"province_name"`
	CountryName  string `json:"country_name"`
}

type LocalitySellerCountDoc struct {
	LocalityID   int    `json:"locality_id"`
	LocalityName string `json:"locality_name"`
	SellerCount  int    `json:"sellers_count"`
}

type Country struct {
	Id   int
	Name string
}

type Province struct {
	Id        int
	Name      string
	CountryId int
}

type LocalityCarriesCountDoc struct {
	LocalityID   int    `json:"locality_id"`
	LocalityName string `json:"locality_name"`
	CarriesCount int    `json:"carries_count"`
}
