package inbound_order

import (
	productBatchRepo "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/ProductBatch"
	warehouseRepo "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/Warehouse"
	employeeRepo "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/employee"
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/inbound_order"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func NewInboundOrderDefault(
	rp repository.InboundOrderRepository,
	empRepo employeeRepo.EmployeeRepository,
	pbRepo productBatchRepo.ProductBatchRepository,
	whRepo warehouseRepo.WarehouseRepository) InboundOrderService {

	return &InboundOrderDefault{rp: rp, empRepo: empRepo, pbRepo: pbRepo, whRepo: whRepo}
}

// InboundOrderDefault is a struct that represents the default service for vehicles
type InboundOrderDefault struct {
	// rp is the repository that will be used by the service
	rp      repository.InboundOrderRepository
	empRepo employeeRepo.EmployeeRepository
	pbRepo  productBatchRepo.ProductBatchRepository
	whRepo  warehouseRepo.WarehouseRepository
}

// Create implements InboundOrderService.
func (i *InboundOrderDefault) Create(request models.RequestInboundOrder) (*models.InboundOrder, error) {

	// Validate if OrderNumber exists
	existingOrder, err := i.rp.FindByOrderNumber(request.OrderNumber)
	if existingOrder != nil && err == nil {
		return nil, customErrors.ErrorConflict
	}

	// Validate if EmployeeID exists
	employee, err := i.empRepo.GetById(request.EmployeeID)
	if employee != nil && err == nil {
		return nil, customErrors.ErrorConflict
	}

	// Validate if ProductBatchID exists
	productBatch, err := i.pbRepo.GetById(request.ProductBatchID)
	if productBatch != nil && err == nil {
		return nil, customErrors.ErrorConflict
	}

	// Validate if WarehouseID exists
	warehouse, err := i.whRepo.GetById(request.WarehouseID)
	if warehouse != nil && err == nil {
		return nil, customErrors.ErrorNotFound
	}

	//Map request to model
	inboundOrderMap := mappers.RequestInboundOrderToInboundOrder(request)
	result, err := i.rp.Create(inboundOrderMap)
	if err != nil {
		return nil, err
	}

	return result, nil
}
