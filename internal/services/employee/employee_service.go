package employee

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type EmployeeService interface {
	GetAll() ([]models.EmployeeDoc, error)
	GetById(id int) (*models.EmployeeDoc, error)
	Create(request models.RequestEmployee) (*models.EmployeeDoc, error)
	Delete(id int) error
}
