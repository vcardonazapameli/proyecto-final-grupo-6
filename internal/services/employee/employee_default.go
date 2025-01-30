package employee

import (
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/employee"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
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
func (e *EmployeeDefault) GetAll() ([]models.EmployeeDoc, error) {

	data, err := e.rp.GetAll()
	if err != nil {
		return nil, err
	}

	//mapear
	var dataMap []models.EmployeeDoc
	for _, value := range data {
		employeeMap := mappers.EmployeeToEmployeeDoc(value)
		dataMap = append(dataMap, employeeMap)
	}

	return dataMap, nil
}

// GetById implements EmployeeService.
func (e *EmployeeDefault) GetById(id int) (*models.EmployeeDoc, error) {
	data, err := e.rp.GetById(id)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, customErrors.ErrorNotFound
	}

	dataMap := mappers.EmployeeToEmployeeDoc(*data)

	return &dataMap, nil
}

// Create implements EmployeeService.
func (e *EmployeeDefault) Create(request models.RequestEmployee) (*models.EmployeeDoc, error) {

	//1. validate model

	//2. validate if employee already exist
	existingEmployee, err := e.rp.FindByCardNumberID(request.CardNumberID)
	if err == nil && existingEmployee != nil {
		return nil, customErrors.ErrorConflict
	}

	//3. map request to model
	newEmployeeMap := mappers.RequestEmployeeToEmployee(request)

	//4. create new Employee
	data, err := e.rp.Create(newEmployeeMap)
	if err != nil {
		return nil, customErrors.ErrorInternalServerError
	}

	//5. map model to doc
	dataMap := mappers.EmployeeToEmployeeDoc(data)
	return &dataMap, nil
}

// Delete implements EmployeeService.
func (e *EmployeeDefault) Delete(id int) error {
	err := e.rp.Delete(id)

	if err != nil {
		return customErrors.ErrorNotFound
	}

	return nil
}
