package employee

import (
	"database/sql"
	"errors"

	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func NewEmployeeMap(db *sql.DB) EmployeeRepository {
	return &EmployeeMap{db: db}
}

type EmployeeMap struct {
	db *sql.DB
}

// FindByCardNumberID implements EmployeeRepository.
func (r *EmployeeMap) FindByCardNumberID(cardNumberID string) (*models.Employee, error) {
	row := r.db.QueryRow("SELECT id, first_name, last_name, id_card_number, warehouse_id FROM employees WHERE id_card_number = ?", cardNumberID)
	var employee models.Employee
	if err := row.Scan(&employee.Id, &employee.FirstName, &employee.LastName, &employee.CardNumberID, &employee.WarehouseID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, customErrors.ErrorNotFound
		}
		return nil, err
	}
	return &employee, nil
}

// FindAll is a method that returns a map of all vehicles
func (r *EmployeeMap) GetAll() (map[int]models.Employee, error) {
	rows, err := r.db.Query("SELECT id, first_name, last_name, id_card_number, warehouse_id FROM  employees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	e := make(map[int]models.Employee)
	for rows.Next() {
		var employee models.Employee
		if err := rows.Scan(&employee.Id, &employee.FirstName, &employee.LastName, &employee.CardNumberID, &employee.WarehouseID); err != nil {
			return nil, err
		}
		e[employee.Id] = employee
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return e, nil
}

// FindAll is a method that returns a map of all vehicles
func (r *EmployeeMap) GetById(id int) (*models.Employee, error) {

	row := r.db.QueryRow("SELECT id, first_name, last_name, id_card_number, warehouse_id FROM employees WHERE id = ?", id)
	var employee models.Employee
	if err := row.Scan(&employee.Id, &employee.FirstName, &employee.LastName, &employee.CardNumberID, &employee.WarehouseID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, customErrors.ErrorNotFound
		}
		return nil, err
	}
	return &employee, nil
}

// Create implements EmployeeRepository.
func (r *EmployeeMap) Create(newEmployee models.Employee) (*models.Employee, error) {

	result, err := r.db.Exec("INSERT INTO employees (first_name, last_name, id_card_number, warehouse_id) VALUES (?, ?, ?, ?)",
		newEmployee.FirstName, newEmployee.LastName, newEmployee.CardNumberID, newEmployee.WarehouseID)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	newEmployee.Id = int(id)
	return &newEmployee, nil
}

// Update implements EmployeeRepository.
func (r *EmployeeMap) Update(id int, request *models.Employee) error {
	result, err := r.db.Exec("UPDATE employees SET first_name = ?, last_name = ?, id_card_number = ?, wareHouse_id = ? WHERE id = ?",
		request.FirstName, request.LastName, request.CardNumberID, request.WarehouseID, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return customErrors.ErrorNotFound
	}

	return nil
}

// Create implements EmployeeRepository.
func (r *EmployeeMap) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM employees WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {

		return customErrors.ErrorNotFound
	}

	return nil
}
