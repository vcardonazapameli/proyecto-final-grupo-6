package models

type LocalityDoc struct {
	Id           int    `json:"id"`
	LocalityName string `json:"locality_name"`
	ProvinceName string `json:"province_name"`
	CountryName  string `json:"country_name"`
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
