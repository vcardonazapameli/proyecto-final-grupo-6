package carrier

import (
	"database/sql"

	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

type carrierRepository struct {
	db *sql.DB
}

func NewCarrierRepository(db *sql.DB) CarrierRepository {
	repository := &carrierRepository{
		db: db,
	}
	return repository
}

func (carrierRepository *carrierRepository) Create(carrier *models.CarrierDocResponse) error {
	result, err := carrierRepository.db.Exec("insert into carriers (cid, company_name, address, telephone, locality_id) values (?, ?, ?, ?, ?)",
		carrier.Cid,
		carrier.Company_name,
		carrier.Address,
		carrier.Telephone,
		carrier.Locality_id,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	carrier.Id = int(id)
	return nil
}

func (carrierRepository *carrierRepository) ExistInDb(cid string) (bool, error) {
	var exist bool
	query := "select exists(select 1 from carriers c where c.cid = ?)"
	err := carrierRepository.db.QueryRow(query, cid).Scan(&exist)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (carrierRepository *carrierRepository) ExistLocalityInDb(id int) (bool, error) {
	var exist bool
	query := "select exists(select 1 from localities l where l.id = ?)"
	err := carrierRepository.db.QueryRow(query, id).Scan(&exist)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (carrierRepository *carrierRepository) ExistCarrierInDb(id int) (bool, error) {
	var exist bool
	query := "select exists(select 1 from carriers c where c.id = ?)"
	err := carrierRepository.db.QueryRow(query, id).Scan(&exist)
	if err != nil {
		return false, err
	}
	return exist, nil
}
