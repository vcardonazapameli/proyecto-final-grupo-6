package product

import (
	"database/sql"
	"log"

	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

type productTypeRepository struct {
	db *sql.DB
}

func NewProductTypeRepository(db *sql.DB) ProductTypeRepository {
	repository := &productTypeRepository{
		db: db,
	}
	return repository
}

func (productTypeRepository *productTypeRepository) GetAll() ([]models.ProductTypeDocResponse, error) {
	rows, err := productTypeRepository.db.Query("select id, description from products_types;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var productTypes []models.ProductTypeDocResponse
	for rows.Next() {
		var productType models.ProductTypeDocResponse
		err := rows.Scan(
			&productType.Id,
			&productType.Description,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		productTypes = append(productTypes, productType)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return productTypes, nil
}

func (productTypeRepository *productTypeRepository) GetById(id int) (*models.ProductTypeDocResponse, error) {
	row := productTypeRepository.db.QueryRow("select id, description from products_types where id = ?;", id)
	var productType models.ProductTypeDocResponse
	err := row.Scan(
		&productType.Id,
		&productType.Description,
	)
	if err != nil {
		return nil, err
	}
	return &productType, nil
}

func (productTypeRepository *productTypeRepository) Create(product *models.ProductTypeDocResponse) error {
	result, err := productTypeRepository.db.Exec("insert into products_types (description) values (?)",
		product.Description,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	product.Id = int(id)
	return nil
}

func (productTypeRepository *productTypeRepository) ExistInDb(description string) (bool, error) {
	var exist bool
	query := "select exists(select 1 from products_types pt where pt.description = ?)"
	err := productTypeRepository.db.QueryRow(query, description).Scan(&exist)
	if err != nil {
		return false, err
	}
	return exist, nil
}
