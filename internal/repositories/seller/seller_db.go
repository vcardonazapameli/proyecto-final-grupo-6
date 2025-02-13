package seller

import (
	"database/sql"
	"errors"

	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/go-sql-driver/mysql"
)

type SellerRepositoryDB struct {
	db *sql.DB
}

func NewSellerRepositoryDB(db *sql.DB) *SellerRepositoryDB {
	return &SellerRepositoryDB{db}
}

// Delete implements SellerRepository.
func (r SellerRepositoryDB) Delete(id int) error {
	res, err := r.db.Exec("DELETE FROM sellers WHERE id = ?", id)
	if err != nil {
		return err // Error en la ejecución del SQL
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err // Error inesperado al obtener filas afectadas
	}

	if rowsAffected == 0 {
		return customErrors.ErrorNotFound // Devolver 404 si no se eliminó nada
	}

	return nil // Eliminación exitosa
}

// GetAll implements SellerRepository.
func (r SellerRepositoryDB) GetAll() (map[int]models.Seller, error) {
	rows, err := r.db.Query("SELECT id, cid, company_name, address, telephone, locality_id FROM sellers")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sellers := make(map[int]models.Seller)

	for rows.Next() {
		var seller models.Seller
		if err := rows.Scan(&seller.Id, &seller.Cid, &seller.CompanyName, &seller.Address, &seller.Telephone, &seller.LocalityID); err != nil {
			return nil, err
		}

		sellers[seller.Id] = seller
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return sellers, nil
}

// GetByID implements SellerRepository.
func (r SellerRepositoryDB) GetByID(id int) (models.Seller, error) {
	row := r.db.QueryRow("SELECT id, cid, company_name, address, telephone, locality_id FROM sellers WHERE id = ?", id)
	var seller models.Seller
	err := row.Scan(&seller.Id, &seller.Cid, &seller.CompanyName, &seller.Address, &seller.Telephone, &seller.LocalityID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Seller{}, customErrors.ErrorNotFound
		}
		return models.Seller{}, err
	}
	return seller, nil
}

// Save implements SellerRepository.
func (r SellerRepositoryDB) Save(s *models.Seller) error {
	res, err := r.db.Exec("INSERT INTO sellers (cid, company_name, address, telephone, locality_id) VALUES (?, ?, ?, ?, ?)",
		s.Cid,
		s.CompanyName,
		s.Address,
		s.Telephone,
		s.LocalityID,
	)

	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1452: // LocalityID not found
				return customErrors.ErrorNotFound
			case 1062: // Duplicated CID
				return customErrors.ErrorConflict
			default:
				return err
			}
		}
		return err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return err
	}

	s.Id = int(id)

	return nil
}

// SearchByCID implements SellerRepository.
func (r SellerRepositoryDB) SearchByCID(cid int) (models.Seller, bool) {
	row := r.db.QueryRow("SELECT id, cid, company_name, address, telephone, locality_id FROM sellers WHERE cid = ?", cid)
	var seller models.Seller
	err := row.Scan(&seller.Id, &seller.Cid, &seller.CompanyName, &seller.Address, &seller.Telephone, &seller.LocalityID)

	if err != nil {
		return models.Seller{}, false
	}
	return seller, true
}

// Update implements SellerRepository.
func (r SellerRepositoryDB) Update(models.Seller) {
	panic("unimplemented")
}
