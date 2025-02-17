package inbound_order

import (
	"log"

	employeeRepo "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/employee"
	inboundOrderRepo "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/inbound_order"

	// productBatchRepo "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product_batch"
	warehouseRepo "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/warehouse"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func NewInboundOrderDefault(
	rp inboundOrderRepo.InboundOrderRepository,
	empRepo employeeRepo.EmployeeRepository,
	// pbRepo productBatchRepo.ProductBatchRepository,
	whRepo warehouseRepo.WarehouseRepository) InboundOrderService {

	return &InboundOrderDefault{
		rp:      rp,
		empRepo: empRepo,
		// pbRepo: pbRepo,
		whRepo: whRepo}
}

// InboundOrderDefault is a struct that represents the default service for vehicles
type InboundOrderDefault struct {
	// rp is the repository that will be used by the service
	rp      inboundOrderRepo.InboundOrderRepository
	empRepo employeeRepo.EmployeeRepository
	// pbRepo  productBatchRepo.ProductBatchRepository
	whRepo warehouseRepo.WarehouseRepository
}

// Create implements InboundOrderService.
func (i *InboundOrderDefault) Create(request models.RequestInboundOrder) (*models.InboundOrder, error) {

	// Validate if OrderNumber exists
	existingOrder, err := i.rp.ExistOrderNumber(request.OrderNumber)
	if existingOrder && err == nil {
		log.Print("OrderNumber already exists")
		return nil, customErrors.ErrorConflict
	}

	// Validate if EmployeeID exists
	employee, err := i.empRepo.GetById(request.EmployeeID)
	if employee == nil && err != nil {
		log.Print("EmployeeID does not exist")
		return nil, customErrors.ErrorConflict
	}

	// Validate if ProductBatchID exists
	// productBatch, err := i.pbRepo.GetById(request.ProductBatchID)
	// if productBatch != nil && err == nil {
	// 	return nil, customErrors.ErrorConflict
	// }

	// Validate if WarehouseID exists
	warehouse, _ := i.whRepo.GetById(request.WarehouseID)
	if warehouse == nil {
		log.Print("WarehouseID does not exist")
		return nil, customErrors.ErrorConflict
	}

	//Map request to model
	inboundOrderMap := mappers.RequestInboundOrderToInboundOrder(request)
	result, err := i.rp.Create(inboundOrderMap)
	if err != nil {
		return nil, err
	}

	return result, nil
}
