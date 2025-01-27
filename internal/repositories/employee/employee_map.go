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
func (r *EmployeeMap) GetAll() (v map[int]models.Employee, err error) {
	return
}
