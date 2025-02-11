package locality

import (
	"database/sql"

	defaultErrors "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/go-sql-driver/mysql"
)

type LocalityRepositoryDB struct {
	db *sql.DB
}

func NewLocalityRepositoryDB(db *sql.DB) *LocalityRepositoryDB {
	return &LocalityRepositoryDB{db}
}

func (r *LocalityRepositoryDB) Save(loc *models.Locality) error {
	_, err := r.db.Exec("INSERT INTO localities (id, locality_name, province_name, country_name) VALUES (?,?,?,?)",
		&loc.Id,
		&loc.LocalityName,
		&loc.ProvinceName,
		&loc.CountryName,
	)

	if err != nil {
		if valErr, ok := err.(*mysql.MySQLError); ok {
			if valErr.Number == 1062 {
				return defaultErrors.ErrorConflict
			}
			return err
		}
	}
	return nil
}
