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
	panic("unimplemented")
}

// GetAll implements SellerRepository.
func (r SellerRepositoryDB) GetAll() (map[int]models.Seller, error) {
	panic("unimplemented")
}

// GetByID implements SellerRepository.
func (r SellerRepositoryDB) GetByID(id int) (models.Seller, error) {
	panic("unimplemented")
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
func (r SellerRepositoryDB) SearchByCID(int) (models.Seller, bool) {
	panic("unimplemented")
}

// Update implements SellerRepository.
func (r SellerRepositoryDB) Update(models.Seller) {
	panic("unimplemented")
}
