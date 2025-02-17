package models




type PurchaseOrder struct {

	Id	 				uint
	OrderNumber			uint
	OrderDate			string
	TrackingCode		string
	BuyerId				uint
	CarrierId			uint
	OrderStatus			uint
	WarehouseId			uint

}

type PurchaseOrderResponse struct {

	Id	 				uint  	`json:"id"`
	OrderNumber			uint		`json:"order_number"`
	OrderDate			string		`json:"order_date"`
	TrackingCode		string		`json:"tracking_code"`
	BuyerId				uint		`json:"buyer_id"`
	CarrierId			uint		`json:"carrier_id"`
	OrderStatusId		uint		`json:"order_status_id"`
	WarehouseId			uint		`json:"wareHouse_id"`
}

type PurchaseOrderRequest struct {

	OrderNumber			uint		`json:"order_number"`
	OrderDate			string		`json:"order_date"`
	TrackingCode		string		`json:"tracking_code"`
	BuyerId				uint		`json:"buyer_id"`
	CarrierId			uint		`json:"carrier_id"`
	OrderStatusId		uint		`json:"order_status_id"`
	WarehouseId			uint		`json:"wareHouse_id"`
}
type PurchaseOrderReport struct {
	BuyerDocResponse
	PurchaseOrdersCount uint `json:"purchase_orders_count"`
}