package models

type Product struct {
	Id uint64
	ProductAttributes
}

type ProductAttributes struct {
	ProductCode                    string
	Description                    string
	ExpirationRate                 float64
	recommendedFreezingTemperature float64
	freezingRate                   float64
	Dimensions
	ProductType uint64
	Seller      uint64
}

type Dimensions struct {
	Width     float64
	Height    float64
	Length    float64
	NetWeight float64
}

type ProductDoc struct {
	Id                             uint64  `json:"id"`
	ProductCode                    string  `json:"product_code"`
	Description                    string  `json:"description"`
	ExpirationRate                 float64 `json:"expiration_rate"`
	recommendedFreezingTemperature float64 `json:"recommended_freezing_temperature"`
	freezingRate                   float64 `json:"freezing_rate"`
	Width                          float64 `json:"width"`
	Height                         float64 `json:"height"`
	Length                         float64 `json:"length"`
	NetWeight                      float64 `json:"net_weight"`
	ProductType                    uint64  `json:"product_type_id"`
	Seller                         uint64  `json:"seller_id"`
}
