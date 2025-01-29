package employee

import (
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func NewEmployeeMap(db map[int]models.Employee) EmployeeRepository {
	// default db
	defaultDb := make(map[int]models.Employee)
	if db != nil {
		defaultDb = db
	}
	return &EmployeeMap{db: defaultDb}
}

// EmployeeMap is a struct that represents a vehicle repository
type EmployeeMap struct {
	// db is a map of vehicles
	db map[int]models.Employee
}

// FindAll is a method that returns a map of all vehicles
func (r *EmployeeMap) GetAll() (map[int]models.Employee, error) {
	e := make(map[int]models.Employee)

	// copy db
	for key, value := range r.db {
		e[key] = value
	}

	return e, nil
}

// FindAll is a method that returns a map of all vehicles
func (r *EmployeeMap) GetById(id int) (*models.Employee, error) {
	var e models.Employee

	for _, value := range r.db {
		if value.Id == id {
			e = value
			return &e, nil
		}
	}

	return nil, nil
}
