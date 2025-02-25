package employee

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type EmployeeRepository interface {
	GetAll() (map[int]models.Employee, error)
	GetById(id int) (*models.Employee, error)
	Create(newEmployee models.Employee) (*models.Employee, error)
	Update(id int, request *models.Employee) error
	Delete(id int) error
	FindByCardNumberID(cardNumberID string) (*models.Employee, error)
}
 