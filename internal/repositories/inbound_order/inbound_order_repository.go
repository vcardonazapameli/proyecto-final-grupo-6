package inbound_order

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type InboundOrderRepository interface {
	Create(request models.InboundOrder) (*models.InboundOrder, error)
	GetReportByEmployeeID(id int) (*models.EmployeeWithOrders, error)
	GetAllReport() ([]models.EmployeeWithOrders, error)
}
