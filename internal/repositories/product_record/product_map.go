package product

import (
	"database/sql"
	"time"

	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

type productRecordRepository struct {
	db *sql.DB
}

func NewProductRecordRepository(db *sql.DB) ProductRecordRepository {
	repository := &productRecordRepository{
		db: db,
	}
	return repository
}

func (productRecordRepository *productRecordRepository) Create(productRecord *models.ProductRecordDocResponse) error {
	formattedLastUpdateDate, _ := time.Parse("2006-01-02", productRecord.LastUpdateDate)
	result, err := productRecordRepository.db.Exec("insert into product_records (last_update_date, purchase_price, sale_price, product_id) values (?, ?, ?, ?)",
		formattedLastUpdateDate,
		productRecord.PurchasePrice,
		productRecord.SalePrice,
		productRecord.ProductId,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	productRecord.Id = int(id)
	return nil
}
