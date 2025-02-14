package productbatch

import (
	"database/sql"

	defaultErrors "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/go-sql-driver/mysql"
)

type ProductBatchRepositoryDB struct {
	db *sql.DB
}

func NewProductBatchRepositoryDB(db *sql.DB) *ProductBatchRepositoryDB {
	return &ProductBatchRepositoryDB{db}
}

func (r *ProductBatchRepositoryDB) Save(pb *models.ProductBatch) error {
	result, err := r.db.Exec("INSERT INTO product_batches (batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		&pb.BatchNumber,
		&pb.CurrentQuantity,
		&pb.CurrentTemperature,
		&pb.DueDate,
		&pb.InitialQuantity,
		&pb.ManufacturingDate,
		&pb.ManufacturingHour,
		&pb.MinimumTemperature,
		&pb.ProductId,
		&pb.SectionId,
	)
	if err != nil {
		if v, ok := err.(*mysql.MySQLError); ok {
			if v.Number == 1062 {
				return defaultErrors.ErrorConflict
			}
		}
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	pb.Id = int(id)
	return nil

}
