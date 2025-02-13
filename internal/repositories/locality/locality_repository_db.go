package locality

import (
	"database/sql"
	"errors"

	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
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

func (r *LocalityRepositoryDB) GetSellersByLocalityIDCount(locId int) (models.LocalitySellerCountDoc, error) {
	var locSeller models.LocalitySellerCountDoc
	row := r.db.QueryRow(`
	SELECT COUNT(s.id), l.id, l.locality_name 
	FROM localities l 
	LEFT JOIN sellers s ON s.locality_id = l.id 
	WHERE l.id = ?
	GROUP BY l.id, l.locality_name
`, locId)

	if err := row.Scan(&locSeller.SellerCount, &locSeller.LocalityID, &locSeller.LocalityName); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.LocalitySellerCountDoc{}, customErrors.ErrorNotFound
		}
		return models.LocalitySellerCountDoc{}, err
	}
	return locSeller, nil
}

func (r *LocalityRepositoryDB) GetAllSellersByLocalityIDCount() ([]models.LocalitySellerCountDoc, error) {
	localities := make([]models.LocalitySellerCountDoc, 0)

	row, err := r.db.Query(`
	SELECT COUNT(s.id), l.id, l.locality_name 
	FROM localities l 
	LEFT JOIN sellers s ON s.locality_id = l.id
	GROUP BY l.id, l.locality_name
`)
	if err != nil {
		return []models.LocalitySellerCountDoc{}, err
	}

	defer row.Close()
	for row.Next() {
		var locSeller models.LocalitySellerCountDoc

		if err := row.Scan(&locSeller.SellerCount, &locSeller.LocalityID, &locSeller.LocalityName); err != nil {
			return []models.LocalitySellerCountDoc{}, err
		}
		localities = append(localities, locSeller)
	}

	if row.Err() != nil {
		return []models.LocalitySellerCountDoc{}, nil
	}

	return localities, nil
}
