package models


type  Buyer struct {
	Id 		int
	BuyerAttributes
}

type BuyerAttributes struct{
	CardNumberId 	int
	FirstName 		string
	LastName 		string
}

type BuyerDoc struct {
	Id 				int			`json:"id"`
	CardNumberId 	int		`json:"card_number_id"`
	FirstName		string		`json:"first_name"`
	LastName 		string		`json:"last_name"`
}