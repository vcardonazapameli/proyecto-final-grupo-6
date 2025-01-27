package employee

import (
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/employee"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func NewEmployeeDefault(rp repository.EmployeeRepository) EmployeeService {
	return &EmployeeDefault{rp: rp}
}

// EmployeeDefault is a struct that represents the default service for vehicles
type EmployeeDefault struct {
	// rp is the repository that will be used by the service
	rp repository.EmployeeRepository
}

// GetAll implements EmployeeService.
func (e *EmployeeDefault) GetAll() (map[int]models.Employee, error) {
	data, err := e.rp.GetAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}
