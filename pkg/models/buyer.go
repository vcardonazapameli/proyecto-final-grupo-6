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
type BuyerDocRequest struct {
	CardNumberId 	int			`json:"card_number_id"`
	FirstName		string		`json:"first_name"`
	LastName 		string		`json:"last_name"`
}
type BuyerDocResponse struct {
	Id 				int			`json:"id"`
	CardNumberId 	int			`json:"card_number_id"`
	FirstName		string		`json:"first_name"`
	LastName 		string		`json:"last_name"`
}
// Struct para 
type UpdateBuyerDto struct{
	CardNumberId 	*int			`json:"card_number_id"`
	FirstName		*string			`json:"first_name"`
	LastName 		*string			`json:"last_name"`
}
type CreateBuyerDto struct{
	CardNumberId 	int				`json:"card_number_id"`
	FirstName		string			`json:"first_name"`
	LastName 		string			`json:"last_name"`
}
