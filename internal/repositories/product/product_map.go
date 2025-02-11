package product

import (
	"database/sql"
	"log"

	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	repository := &productRepository{
		db: db,
	}
	return repository
}

func (productRepository *productRepository) GetAll() ([]models.ProductDocResponse, error) {
	rows, err := productRepository.db.Query("select id, product_code, description, expiration_rate, recommended_freezing_temperature, freezing_rate, width, height, length, net_weight, product_type_id, seller_id from products;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []models.ProductDocResponse
	for rows.Next() {
		var product models.ProductDocResponse
		err := rows.Scan(
			&product.Id,
			&product.ProductCode,
			&product.Description,
			&product.ExpirationRate,
			&product.RecommendedFreezingTemperature,
			&product.FreezingRate,
			&product.Width,
			&product.Height,
			&product.Length,
			&product.NetWeight,
			&product.ProductType,
			&product.Seller,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (productRepository *productRepository) GetById(id int) (*models.ProductDocResponse, error) {
	row := productRepository.db.QueryRow("select id, product_code, description, expiration_rate, recommended_freezing_temperature, freezing_rate, width, height, length, net_weight, product_type_id, seller_id from products where id = ?;", id)
	var product models.ProductDocResponse
	err := row.Scan(
		&product.Id,
		&product.ProductCode,
		&product.Description,
		&product.ExpirationRate,
		&product.RecommendedFreezingTemperature,
		&product.FreezingRate,
		&product.Width,
		&product.Height,
		&product.Length,
		&product.NetWeight,
		&product.ProductType,
		&product.Seller,
	)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (productRepository *productRepository) Delete(id int) error {
	_, err := productRepository.db.Exec("delete from products where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (productRepository *productRepository) Create(product *models.ProductDocResponse) error {
	result, err := productRepository.db.Exec("insert into products (product_code, description, expiration_rate, recommended_freezing_temperature, freezing_rate, width, height, length, net_weight, product_type_id, seller_id) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		product.ProductCode,
		product.Description,
		product.ExpirationRate,
		product.RecommendedFreezingTemperature,
		product.FreezingRate,
		product.Width,
		product.Height,
		product.Length,
		product.NetWeight,
		product.ProductType,
		product.Seller,
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

func (productRepository *productRepository) ExistInDb(productCode string) (bool, error) {
	var exist bool
	query := "select exists(select 1 from products p where p.product_code = ?)"
	err := productRepository.db.QueryRow(query, productCode).Scan(&exist)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (productRepository *productRepository) Update(id int, product *models.ProductDocResponse) error {
	_, err := productRepository.db.Exec("update products set product_code = ?, description = ?, expiration_rate = ?, recommended_freezing_temperature = ?, freezing_rate = ?, width = ?, height = ?, length = ?, net_weight = ?, product_type_id = ?, seller_id = ? where id = ?",
		product.ProductCode,
		product.Description,
		product.ExpirationRate,
		product.RecommendedFreezingTemperature,
		product.FreezingRate,
		product.Width,
		product.Height,
		product.Length,
		product.NetWeight,
		product.ProductType,
		product.Seller,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (productRepository *productRepository) MatchWithTheSameProductCode(id int, productCode string) (bool, error) {
	var numberOfMatches int
	query := "select count(*) from products where id != ? and product_code = ?"
	err := productRepository.db.QueryRow(query, id, productCode).Scan(&numberOfMatches)
	if err != nil {
		return false, err
	}
	return numberOfMatches > 0, nil
}
