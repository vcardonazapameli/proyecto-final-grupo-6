package employee

import (
	"fmt"

	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/employee"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/validators"
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

	//1. validate model (empty fields)
	err := validators.ValidateNoEmptyFields(request)
	if err != nil {
		return nil, fmt.Errorf("%w : %s", customErrors.ErrorUnprocessableContent, err)
	}

	//2. validate if cardNumberID already exist
	existCardNumberID, err := e.rp.FindByCardNumberID(request.CardNumberID)
	if err == nil && existCardNumberID != nil {
		return nil, customErrors.ErrorConflict
	}
	if err != nil {
		return nil, customErrors.ErrorInternalServerError
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

// Update implements EmployeeService.
func (e *EmployeeDefault) Update(id int, request models.UpdateEmployee) (*models.EmployeeDoc, error) {

	/// 1. Get the existing employee
	existingEmployee, err := e.rp.GetById(id)
	if err != nil {
		return nil, customErrors.ErrorNotFound
	}

	// 2. Validar y actualizar los campos que no son nil
	if request.CardNumberID != nil && *request.CardNumberID != "" {

		//3. validate if cardNumberID already exist
		exists, err := e.rp.FindByCardNumberID(*request.CardNumberID)
		if exists != nil && exists.Id != id {
			return nil, customErrors.ErrorConflict
		}
		if err != nil {
			return nil, customErrors.ErrorInternalServerError
		}

		existingEmployee.EmployeeAttributes.CardNumberID = *request.CardNumberID
	}
	if request.FirstName != nil && *request.FirstName != "" {
		existingEmployee.EmployeeAttributes.FirstName = *request.FirstName
	}
	if request.LastName != nil && *request.LastName != "" {
		existingEmployee.EmployeeAttributes.LastName = *request.LastName
	}
	if request.WarehouseID != nil && *request.WarehouseID > 0 {
		existingEmployee.EmployeeAttributes.WarehouseID = *request.WarehouseID
	}

	// 4. Save the updated employee
	err = e.rp.Update(id, existingEmployee)
	if err != nil {
		return nil, customErrors.ErrorInternalServerError
	}

	// 4. Map model to doc
	dataMap := mappers.EmployeeToEmployeeDoc(*existingEmployee)
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
