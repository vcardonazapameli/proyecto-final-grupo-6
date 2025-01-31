package employee

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

// EmployeeLoader is an interface that represents the loader for vehicles
type EmployeeLoader interface {
	// Load is a method that loads the vehicles
	Load() (v map[int]models.Employee, err error)
}
