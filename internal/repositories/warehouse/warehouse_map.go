package warehouse

import (
	"database/sql"
	"log"

	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

type warehouseRepository struct {
	db *sql.DB
}

func NewWarehouseRepository(db *sql.DB) WarehouseRepository {
	repository := &warehouseRepository{
		db: db,
	}
	return repository
}

func (r *warehouseRepository) GetAll() ([]models.WarehouseDocResponse, error) {
	rows, err := r.db.Query("select id, warehouse_code, address, telephone, minimun_capacity, minimun_temperature, locality_id from warehouses;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var warehouses []models.WarehouseDocResponse
	for rows.Next() {
		var warehouse models.WarehouseDocResponse
		err := rows.Scan(
			&warehouse.ID,
			&warehouse.Warehouse_code,
			&warehouse.Address,
			&warehouse.Telephone,
			&warehouse.Minimun_capacity,
			&warehouse.Minimun_temperature,
			&warehouse.Locality_id,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		warehouses = append(warehouses, warehouse)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return warehouses, nil
}

func (r *warehouseRepository) GetById(id int) (*models.WarehouseDocResponse, error) {
	row := r.db.QueryRow("select id, warehouse_code, address, telephone, minimun_capacity, minimun_temperature, locality_id from warehouses where id = ?;", id)
	var warehouse models.WarehouseDocResponse
	err := row.Scan(
		&warehouse.ID,
		&warehouse.Warehouse_code,
		&warehouse.Address,
		&warehouse.Telephone,
		&warehouse.Minimun_capacity,
		&warehouse.Minimun_temperature,
		&warehouse.Locality_id,
	)
	if err != nil {
		return nil, err
	}
	return &warehouse, nil
}

func (r *warehouseRepository) CreateWarehouse(warehouse *models.WarehouseDocResponse) error {
	result, err := r.db.Exec("insert into warehouses (address, telephone, warehouse_code, minimun_capacity, minimun_temperature, locality_id) values (?, ?, ?, ?, ?, ?)",
		warehouse.Address,
		warehouse.Telephone,
		warehouse.Warehouse_code,
		warehouse.Minimun_capacity,
		warehouse.Minimun_temperature,
		warehouse.Locality_id,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	warehouse.ID = int(id)
	return nil
}

func (r *warehouseRepository) DeleteWarehouse(id int) error {
	_, err := r.db.Exec("delete from warehouses where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *warehouseRepository) UpdateWarehouse(id int, warehouseDoc *models.WarehouseDocResponse) error {
	_, err := r.db.Exec("update warehouses set address = ?, telephone = ?, warehouse_code = ?, minimun_capacity = ?, minimun_temperature = ?, locality_id = ? where id = ?",
		warehouseDoc.Address,
		warehouseDoc.Telephone,
		warehouseDoc.Warehouse_code,
		warehouseDoc.Minimun_capacity,
		warehouseDoc.Minimun_temperature,
		warehouseDoc.Locality_id,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *warehouseRepository) ExistInDbWarehouseCode(warehouse_code string) (bool, error) {
	var exist bool
	query := "select exists(select 1 from warehouses w where w.warehouse_code = ?)"
	err := r.db.QueryRow(query, warehouse_code).Scan(&exist)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (r *warehouseRepository) MatchWarehouseCode(id int, warehouse_code string) (bool, error) {
	var numberOfMatches int
	query := "select count(*) from warehouses where id != ? and warehouse_code = ?"
	err := r.db.QueryRow(query, id, warehouse_code).Scan(&numberOfMatches)
	if err != nil {
		return false, err
	}
	return numberOfMatches > 0, nil
}
