package models

type Product struct {
	Id int
	ProductAttributes
}

type ProductAttributes struct {
	ProductCode                    string
	Description                    string
	ExpirationRate                 float64
	RecommendedFreezingTemperature float64
	FreezingRate                   float64
	Dimensions
	ProductType int
	Seller      int
}

type Dimensions struct {
	Width     float64
	Height    float64
	Length    float64
	NetWeight float64
}

type ProductDocResponse struct {
	Id                             int     `json:"id"`
	ProductCode                    string  `json:"product_code"`
	Description                    string  `json:"description"`
	ExpirationRate                 float64 `json:"expiration_rate"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature"`
	FreezingRate                   float64 `json:"freezing_rate"`
	Width                          float64 `json:"width"`
	Height                         float64 `json:"height"`
	Length                         float64 `json:"length"`
	NetWeight                      float64 `json:"net_weight"`
	ProductType                    int     `json:"product_type_id"`
	Seller                         int     `json:"seller_id"`
}

type ProductDocRequest struct {
	ProductCode                    string  `json:"product_code"`
	Description                    string  `json:"description"`
	ExpirationRate                 float64 `json:"expiration_rate"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature"`
	FreezingRate                   float64 `json:"freezing_rate"`
	Width                          float64 `json:"width"`
	Height                         float64 `json:"height"`
	Length                         float64 `json:"length"`
	NetWeight                      float64 `json:"net_weight"`
	ProductType                    int     `json:"product_type_id"`
	Seller                         int     `json:"seller_id"`
}

type ProductUpdateDocRequest struct {
	ProductCode                    *string  `json:"product_code"`
	Description                    *string  `json:"description"`
	ExpirationRate                 *float64 `json:"expiration_rate"`
	RecommendedFreezingTemperature *float64 `json:"recommended_freezing_temperature"`
	FreezingRate                   *float64 `json:"freezing_rate"`
	Width                          *float64 `json:"width"`
	Height                         *float64 `json:"height"`
	Length                         *float64 `json:"length"`
	NetWeight                      *float64 `json:"net_weight"`
	ProductType                    *int     `json:"product_type_id"`
	Seller                         *int     `json:"seller_id"`
}
