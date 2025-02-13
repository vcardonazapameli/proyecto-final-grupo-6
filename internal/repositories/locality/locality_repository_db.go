package locality

import (
	"database/sql"
	"errors"

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

func (r *LocalityRepositoryDB) Save(locId int, locName string, provId int) error {
	_, err := r.db.Exec("INSERT INTO localities (id, locality_name, province_id) VALUES (?,?,?)",
		locId,
		locName,
		provId,
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

func (r *LocalityRepositoryDB) GetProvinceWithCountryNames(provinceName string, countryName string) (models.Province, error) {
	query := `
	SELECT 
		p.id, p.province_name
	FROM provinces p
	JOIN countries c ON p.id_country_fk = c.id
	WHERE p.province_name COLLATE utf8mb4_general_ci = ? 
	AND c.country_name COLLATE utf8mb4_general_ci = ?
`

	var province models.Province

	err := r.db.QueryRow(query, provinceName, countryName).Scan(
		&province.Id, &province.Name,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Province{}, defaultErrors.ErrorNotFound
		}
		return models.Province{}, err
	}

	return province, nil
}
