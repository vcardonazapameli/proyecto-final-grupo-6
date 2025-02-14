package models

type ProductType struct {
	Id          int
	Description string
}

type ProductTypeDocResponse struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}

type ProductTypeDocRequest struct {
	Description string `json:"description"`
}
