package employee

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type EmployeeService interface {
	GetAll() (map[int]models.Employee, error)
	GetById(id int) (*models.Employee, error)
	Create(request models.RequestEmployee) (*models.Employee, error)
}
