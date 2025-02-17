package purchaseorder

import (


	br "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/buyer"
	cr "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/carrier"
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/purchase_order"
	wr "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/warehouse"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/validators"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func NewPurchaseOrderService(repository repository.PurchaseOrderRepository,
	 						 warehouseRepo wr.WarehouseRepository,
							 buyerRepo br.BuyerRepository,
							 carrierRepo cr.CarrierRepository) PurchaseOrderService {
	return &purchaseOrderService{repository: repository, warehouseRepository: warehouseRepo, buyerRepository: buyerRepo, carrierRepository: carrierRepo}
}

type purchaseOrderService struct {
	repository          repository.PurchaseOrderRepository
	warehouseRepository wr.WarehouseRepository
	buyerRepository     br.BuyerRepository
	carrierRepository cr.CarrierRepository
}

// CreatePurchaseOrder implements PurchaseOrderService.
func (p *purchaseOrderService) CreatePurchaseOrder(purchaseOrder models.PurchaseOrderRequest) (*models.PurchaseOrderResponse, error) {
	//validar que no hayan campos vacios
	if err:= validators.ValidateNoEmptyFields(purchaseOrder); err!= nil{ 
		return nil, customErrors.ErrorUnprocessableContent
	}
	//validar que no haya otro purchaseOrder con dicho numero
	if p.repository.ValidateIfOrderNumberExist(purchaseOrder.OrderNumber){ 

		return nil, customErrors.ErrorConflict
	}
	//Validar si orderStatus existe
	if !p.repository.ValidateIfOrderStatusExist(int(purchaseOrder.OrderStatusId)){
			return nil, customErrors.ErrorConflict
	}
	//validar que exista el warehouseId
	if _,err :=p.warehouseRepository.GetById(int(purchaseOrder.WarehouseId)); err != nil{ 

		return nil, customErrors.ErrorConflict
	}
	//validar que exista el buyer
	if !p.buyerRepository.ValidateIfExistsById(int(purchaseOrder.BuyerId)){ 

		return nil, customErrors.ErrorConflict
	}
	//validar que exista el carrier
	if exist, _:=p.carrierRepository.ExistCarrierInDb(int(purchaseOrder.CarrierId)); !exist{

		return nil, customErrors.ErrorConflict
	}
	purchaseOrderResponse := mappers.PurchaseOrderRequestToResponse(purchaseOrder)
	p.repository.CreatePurchaseOrder(&purchaseOrderResponse)
	return &purchaseOrderResponse, nil
}



