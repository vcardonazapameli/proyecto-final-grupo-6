package buyer

import (
	"database/sql"
	"log"

	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

type buyerRepository struct {
	db *sql.DB
}

func NewBuyerRepository(db *sql.DB) BuyerRepository {
	repository:= &buyerRepository{
		db: db,

	}
	return repository
}

// GetAll implements BuyerRepository.
func (buyerRepository *buyerRepository) GetAll() ([]models.BuyerDocResponse, error) {
	rows, err := buyerRepository.db.Query("SELECT id, id_card_number, first_name, last_name from buyers")
	if err != nil {
		return nil, customErrors.HandleSqlError(err)
	}
	defer rows.Close()
	var buyers []models.BuyerDocResponse
	for rows.Next() {
		var buyer models.BuyerDocResponse
		err := rows.Scan(
			&buyer.Id,
			&buyer.CardNumberId,
			&buyer.FirstName,
			&buyer.LastName,
		)
		if err != nil {
			log.Panicln(err)
			continue
		}
		buyers = append(buyers, buyer)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return buyers, nil
	
}

// GetById implements BuyerRepository.
func (buyerRepository *buyerRepository) GetById(id int) (*models.BuyerDocResponse, error) {
	row := buyerRepository.db.QueryRow("SELECT id, id_card_number, first_name, last_name from buyers WHERE id = ?", id)
	var buyer models.BuyerDocResponse
	err := row.Scan(
		&buyer.Id,
		&buyer.CardNumberId,
		&buyer.FirstName,
		&buyer.LastName,
	)
	if err != nil {
		return nil, customErrors.HandleSqlError(err)
	}
	return &buyer, nil
}


// CreateBuyer implements BuyerRepository.
func (buyerRepository *buyerRepository) CreateBuyer(buyer *models.BuyerDocResponse) error {
	result, err := buyerRepository.db.Exec("INSERT INTO buyers (id_card_number, first_name, last_name) VALUES (?, ?, ?)",
		buyer.CardNumberId,
		buyer.FirstName,
		buyer.LastName,
	)
	if err != nil {
		return customErrors.HandleSqlError(err)
		
	}
	id, err := result.LastInsertId()
	if err != nil {
		return customErrors.HandleSqlError(err)
	}

	buyer.Id = int(id)
	
	return nil

}

// DeleteBuyer implements BuyerRepository.
func (buyerRepository *buyerRepository) DeleteBuyer(buyerId int) error{
	_, err := buyerRepository.db.Exec("DELETE from buyers WHERE id = ?", buyerId)
	if err != nil {
		return customErrors.HandleSqlError(err)
	}
	return nil
}



// UpdateBuyer implements BuyerRepository.
func (buyerRepository *buyerRepository) UpdateBuyer(id int, buyer *models.BuyerDocRequest) error {
	_, err := buyerRepository.db.Exec("UPDATE buyers SET id_card_number = ?, first_name = ?, last_name = ? WHERE id = ?", buyer.CardNumberId, buyer.FirstName, buyer.LastName, id)
	if err != nil {
		return customErrors.HandleSqlError(err)
	}
	return nil
}

// ValidateCardNumberId implements BuyerRepository.
func (buyerRepository *buyerRepository) ValidateCardNumberId(cardNumber int) ( bool) {
var exist bool
	query:= "SELECT EXISTS (SELECT 1 FROM buyers b WHERE b.id_card_number = ?)"
	err := buyerRepository.db.QueryRow(query, cardNumber).Scan(&exist)
	if err != nil {
		return false
	}
	return exist
}

// ValidateCardNumberIdToUpdate implements BuyerRepository.
func (buyerRepository *buyerRepository) ValidateCardNumberIdToUpdate(cardNumber int, id int) (exists bool) {
	var exist bool
	query := "Select exists (select 1 FROM buyers b WHERE b.id != ? and b.id_card_number = ? );"
	err := buyerRepository.db.QueryRow(query, id, cardNumber).Scan(&exist)
	if err != nil {
		return false
	}
	return exist
}

// ValidateIfExistsById implements BuyerRepository.
func (buyerRepository *buyerRepository) ValidateIfExistsById(id int) (exists bool) {
	var exist bool
	query := "SELECT EXISTS (select 1 FROM buyers b WHERE b.id = ?)"
	err := buyerRepository.db.QueryRow(query, id).Scan(&exist)
	if err != nil {
		return false
	}
	return exist
}

func (buyerRepository *buyerRepository) GetPurchasesReports(cardNumberId int) ([]models.PurchaseOrderReport, error) {
	query, args := purchaseReportQuery(cardNumberId)
	rows, err := buyerRepository.db.Query(query, args...)
	if err != nil {
		return nil, customErrors.HandleSqlError(err)
	}
	defer rows.Close()
	var purchase_orders_reports []models.PurchaseOrderReport
	for rows.Next() {
		var purchase_report models.PurchaseOrderReport
		err := rows.Scan(
			&purchase_report.Id,
			&purchase_report.CardNumberId,
			&purchase_report.FirstName,
			&purchase_report.LastName,
			&purchase_report.PurchaseOrdersCount,
		)
		if err != nil {
			log.Panic(err)
			continue
		}
		purchase_orders_reports = append(purchase_orders_reports, purchase_report)
	}
	if err = rows.Err(); err != nil {
		return nil, customErrors.HandleSqlError(err)
	}
	return purchase_orders_reports, nil
}
func purchaseReportQuery(cardNumberId int) (string, []any) {
    var query string
    var args []interface{}
    
    if cardNumberId == 0 {
        query = "SELECT b.id, b.id_card_number, b.first_name, b.last_name, COUNT(pr.id) as purchase_orders_count FROM buyers b " +
                "LEFT JOIN purchase_orders pr ON pr.buyer_id = b.id " +
                "GROUP BY b.id;"
    } else {
        query = "SELECT b.id, b.id_card_number, b.first_name, b.last_name, COUNT(pr.id) as purchase_orders_count FROM buyers b " +
                "LEFT JOIN purchase_orders pr ON pr.buyer_id = b.id " +
                "WHERE b.id_card_number = ? " +
                "GROUP BY b.id;"
        args = append(args, cardNumberId)
    }
    
    return query, args
}