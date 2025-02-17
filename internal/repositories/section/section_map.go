package section

import (
	"database/sql"

	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/validators"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

type SectionMap struct {
	db *sql.DB
}

func NewSectionMap(db *sql.DB) *SectionMap {
	return &SectionMap{
		db: db,
	}
}

func (r *SectionMap) SectionExists(id int) bool {
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM sections WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func (r *SectionMap) SectionNumberExists(sn string) bool {
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM sections WHERE section_number = ?)", sn).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func (r *SectionMap) GetAll() (map[int]models.Section, error) {
	rows, err := r.db.Query("SELECT id, section_number, current_capacity, current_temperature, maximum_capacity, minimum_capacity, minimum_temperature, product_type_id, warehouse_id FROM sections WHERE is_deleted = FALSE")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	sections := make(map[int]models.Section)
	for rows.Next() {
		var section models.Section
		err := rows.Scan(
			&section.Id,
			&section.SectionNumber,
			&section.CurrentCapacity,
			&section.CurrentTemperature,
			&section.MaximumCapacity,
			&section.MinimumCapacity,
			&section.MinimumTemperature,
			&section.ProductTypeId,
			&section.WarehouseId)
		if err != nil {
			return nil, err
		}
		sections[section.Id] = section
	}
	return sections, nil
}

func (r *SectionMap) GetByID(id int) (models.Section, error) {
	row := r.db.QueryRow("SELECT id, section_number, current_capacity, current_temperature, maximum_capacity, minimum_capacity, minimum_temperature, product_type_id, warehouse_id FROM sections WHERE id = ? AND is_deleted = FALSE", id)
	var section models.Section
	err := row.Scan(
		&section.Id,
		&section.SectionNumber,
		&section.CurrentCapacity,
		&section.CurrentTemperature,
		&section.MaximumCapacity,
		&section.MinimumCapacity,
		&section.MinimumTemperature,
		&section.ProductTypeId,
		&section.WarehouseId)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Section{}, customErrors.ErrorNotFound
		}
		return models.Section{}, err
	}
	return section, nil
}

func (r *SectionMap) Recover(id int) error {
	if !r.SectionExists(id) {
		return customErrors.ErrorNotFound
	}
	_, err := r.db.Exec("UPDATE sections SET is_deleted = FALSE WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *SectionMap) Create(section models.SectionAttributes) (models.Section, error) {
	if err := validators.ValidateNoEmptyFields(section); err != nil {
		return models.Section{}, customErrors.ErrorUnprocessableContent
	}

	if r.SectionNumberExists(section.SectionNumber) {
		return models.Section{}, customErrors.ErrorConflict
	}

	row, err := r.db.Exec("INSERT INTO sections (section_number, current_capacity, current_temperature, maximum_capacity, minimum_capacity, minimum_temperature, product_type_id, warehouse_id, is_deleted) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		section.SectionNumber, section.CurrentCapacity, section.CurrentTemperature, section.MaximumCapacity, section.MinimumCapacity, section.MinimumTemperature, section.ProductTypeId, section.WarehouseId, false)
	if err != nil {
		return models.Section{}, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		return models.Section{}, err
	}
	return mappers.SectionAttributesToSection(section, int(id)), nil
}

func (r *SectionMap) Update(id int, section models.Section) (models.Section, error) {
	if err := validators.ValidateNoEmptyFields(section.SectionAttributes); err != nil {
		return models.Section{}, customErrors.ErrorUnprocessableContent
	}

	if !r.SectionExists(id) {
		return models.Section{}, customErrors.ErrorNotFound
	}

	if section.SectionNumber != "" && r.SectionNumberExists(section.SectionNumber) {
		return models.Section{}, customErrors.ErrorConflict
	}
	_, err := r.db.Exec("UPDATE sections SET section_number = ?, current_capacity = ?, current_temperature = ?, maximum_capacity = ?, minimum_capacity = ?, minimum_temperature = ?, product_type_id = ?, warehouse_id = ? WHERE id = ?",
		section.SectionNumber, section.CurrentCapacity, section.CurrentTemperature, section.MaximumCapacity, section.MinimumCapacity, section.MinimumTemperature, section.ProductTypeId, section.WarehouseId, id)
	if err != nil {
		return models.Section{}, err
	}
	return section, nil
}

func (r *SectionMap) Delete(id int) error {
	if !r.SectionExists(id) {
		return customErrors.ErrorNotFound
	}
	res, err := r.db.Exec("UPDATE sections SET is_deleted = TRUE WHERE id = ?", id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return customErrors.ErrorNotFound
	}
	return nil
}

func (r *SectionMap) GetSectionReports(sectionId int) ([]models.SectionReport, error) {
	if sectionId != 0 {
		exists := r.SectionExists(sectionId)
		if !exists {
			return nil, customErrors.ErrorNotFound
		}
	}
	query, args := sectionReportQuery(sectionId)
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, customErrors.HandleSqlError(err)
	}
	defer rows.Close()
	var section_reports []models.SectionReport
	for rows.Next() {
		var report models.SectionReport
		err := rows.Scan(
			&report.SectionId,
			&report.SectionNumber,
			&report.ProductsCount,
		)
		if err != nil {
			continue
		}
		section_reports = append(section_reports, report)
	}
	if err = rows.Err(); err != nil {
		return nil, customErrors.HandleSqlError(err)
	}

	if sectionId != 0 && len(section_reports) == 0 {
		return nil, customErrors.ErrorNotFound
	}
	return section_reports, nil
}
func sectionReportQuery(sectionId int) (string, []any) {
	var query string
	var args []interface{}

	if sectionId != 0 {
		query = `
            SELECT s.id, s.section_number, COUNT(pb.id) as products_count
            FROM sections s
            LEFT JOIN product_batches pb ON s.id = pb.section_id
            WHERE s.id = ? AND s.is_deleted = FALSE
            GROUP BY s.id, s.section_number`
		args = append(args, sectionId)
	} else {
		query = `
            SELECT s.id, s.section_number, COUNT(pb.id) as products_count
            FROM sections s
            LEFT JOIN product_batches pb ON s.id = pb.section_id
            WHERE s.is_deleted = FALSE
            GROUP BY s.id, s.section_number`
	}
	return query, args
}
