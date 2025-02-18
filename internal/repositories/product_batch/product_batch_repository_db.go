package productbatch

import (
	"database/sql"

	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

type ProductBatchRepositoryDB struct {
	db *sql.DB
}

func NewProductBatchRepository(db *sql.DB) ProductBatchRepository {
	return &ProductBatchRepositoryDB{db}
}

func (r *ProductBatchRepositoryDB) BatchNumberExists(batchNumber string) bool {
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM product_batches WHERE batch_number = ?)", batchNumber).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func (r *ProductBatchRepositoryDB) Save(pb *models.ProductBatchResponse) error {
	query := "INSERT INTO product_batches (batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := r.db.Exec(query,
		pb.BatchNumber,
		pb.CurrentQuantity,
		pb.CurrentTemperature,
		pb.DueDate.Format("2006-01-02 15:04:05"),
		pb.InitialQuantity,
		pb.ManufacturingDate.Format("2006-01-02 15:04:05"),
		pb.ManufacturingHour.Format("2006-01-02 15:04:05"),
		pb.MinimumTemperature,
		pb.ProductId,
		pb.SectionId,
	)
	if err != nil {
		return customErrors.HandleSqlError(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	pb.Id = int(id)
	return nil

}

// ExistsByID implements ProductBatchRepository.
func (r *ProductBatchRepositoryDB) ExistsByID(id int) (bool, error) {
	row := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM product_batches WHERE id = ?)", id)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, err
	}
	return exists, nil
}
