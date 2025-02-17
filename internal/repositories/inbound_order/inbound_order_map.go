package inbound_order

import (
	"database/sql"
	"errors"
	"log"

	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func NewInboundOrderMap(db *sql.DB) InboundOrderRepository {
	return &InboundOrderMap{db: db}
}

type InboundOrderMap struct {
	db *sql.DB
}

// ExistOrderNumber implements InboundOrderRepository.
func (i *InboundOrderMap) ExistOrderNumber(orderNumber string) (bool, error) {
	row := i.db.QueryRow("SELECT EXISTS(SELECT 1 FROM inbound_orders WHERE order_number = ?)", orderNumber)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

// GetAllReport implements InboundOrderRepository.
func (i *InboundOrderMap) GetAllReport() ([]models.EmployeeWithOrders, error) {
	rows, err := i.db.Query(`SELECT  e.id, 
									e.first_name, 
									e.last_name, 
									e.id_card_number, 
									e.warehouse_id,
									count(i.id) inbound_orders_count
							FROM employees e
							LEFT JOIN inbound_orders i ON e.id = i.employe_id
							GROUP BY e.id`)
	if err != nil {
		log.Print("Error in GetAllReport InboundOrder: ", err)
		return nil, err
	}
	defer rows.Close()

	var employees []models.EmployeeWithOrders
	for rows.Next() {
		var employee models.EmployeeWithOrders
		if err := rows.Scan(&employee.Id, &employee.FirstName, &employee.LastName, &employee.CardNumberID, &employee.WarehouseID, &employee.InboundOrdersCount); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

// GetReportByEmployeeID implements InboundOrderRepository.
func (i *InboundOrderMap) GetReportByEmployeeID(id int) (*models.EmployeeWithOrders, error) {
	row := i.db.QueryRow(`SELECT  e.id, 
									e.first_name, 
									e.last_name, 
									e.id_card_number, 
									e.warehouse_id,
									count(i.id) inbound_orders_count
							FROM employees e
							LEFT JOIN inbound_orders i ON e.id = i.employe_id
                            WHERE e.id = ?
							GROUP BY e.id`, id)
	var employee models.EmployeeWithOrders
	if err := row.Scan(&employee.Id, &employee.FirstName, &employee.LastName, &employee.CardNumberID, &employee.WarehouseID, &employee.InboundOrdersCount); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, customErrors.ErrorNotFound
		}
		return nil, err
	}

	return &employee, nil
}

// Create implements InboundOrderRepository.
func (i *InboundOrderMap) Create(request models.InboundOrder) (*models.InboundOrder, error) {
	result, err := i.db.Exec("INSERT INTO inbound_orders (order_date, order_number, employe_id, product_batch_id, wareHouse_id) VALUES (?,?,?,?,?)",
		request.OrderDate, request.OrderNumber, request.EmployeeID, request.ProductBatchID, request.WarehouseID)
	if err != nil {
		log.Print("Error in Create InboundOrder: ", err)
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Print("Error in GetLastInsertID InboundOrder: ", err)
		return nil, err
	}
	request.ID = int(id)

	return &request, nil
}
