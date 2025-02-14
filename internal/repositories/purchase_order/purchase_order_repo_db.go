package purchaseorder

import (
	"database/sql"

	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	_ "github.com/go-sql-driver/mysql"
)

type purchaseOrderRepository struct {
	db *sql.DB
}



func NewPurchaseOrderRepository(db *sql.DB) PurchaseOrderRepository {
	repository := &purchaseOrderRepository{
		db: db,
	}
	return repository
}

// CreatePurchaseOrder implements PurchaseOrderRepository.
func (p *purchaseOrderRepository) CreatePurchaseOrder(purchaseOrder *models.PurchaseOrderResponse) error {
	query := "INSERT INTO purchase_orders(order_number, order_date, tracking_code, buyer_id, carrier_id, order_status_id, wareHouse_id) VALUES (?,?,?,?,?,?,?)"
	result, err := p.db.Exec(query,
		purchaseOrder.OrderNumber,
		purchaseOrder.OrderDate,
		purchaseOrder.TrackingCode,
		purchaseOrder.BuyerId,
		purchaseOrder.CarrierId,
		purchaseOrder.OrderStatusId,
		purchaseOrder.WarehouseId)
	if err != nil {
		return customErrors.HandleSqlError(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	purchaseOrder.Id = uint(id)
	return nil
}



// ValidateIfProductRecordExist implements PurchaseOrderRepository.
func (p *purchaseOrderRepository) ValidateIfOrderStatusExist(orderStatusId int) bool {
	var exist bool
	query := "SELECT EXISTS (SELECT 1 FROM order_status os WHERE os.id = ?)"
	err := p.db.QueryRow(query, orderStatusId).Scan(&exist)
	if err != nil {
		return false
	}
	return exist
}


func (p *purchaseOrderRepository) ValidateIfOrderNumberExist(orderNumber uint) ( bool) {
	var exist bool
		query:= "SELECT EXISTS (SELECT 1 FROM purchase_orders po WHERE b.order_number = ?)"
		err := p.db.QueryRow(query, orderNumber).Scan(&exist)
		if err != nil {
			return false
		}
		return exist
	}