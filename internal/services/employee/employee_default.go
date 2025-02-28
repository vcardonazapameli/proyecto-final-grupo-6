package employee

import (
	"log"

	rpEmployee "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/employee"
	rpInboundOrders "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/inbound_order"
	rpWarehouse "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/warehouse"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/validators"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func NewEmployeeDefault(rp rpEmployee.EmployeeRepository, ior rpInboundOrders.InboundOrderRepository, wh rpWarehouse.WarehouseRepository) EmployeeService {
	return &EmployeeDefault{rp: rp, ior: ior, wh: wh}
}

type EmployeeDefault struct {
	rp  rpEmployee.EmployeeRepository
	ior rpInboundOrders.InboundOrderRepository
	wh  rpWarehouse.WarehouseRepository
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

	dataMap := mappers.EmployeeToEmployeeDoc(*data)

	return &dataMap, nil
}

// Create implements EmployeeService.
func (e *EmployeeDefault) Create(request models.RequestEmployee) (*models.EmployeeDoc, error) {

	//validate model (empty fields)
	err := validators.ValidateCreateEmployee(request)
	if err != nil {
		log.Print(err.Error())
		return nil, customErrors.ErrorUnprocessableContent
	}

	//validate if cardNumberID already exist
	exists, err := e.rp.FindByCardNumberID(request.CardNumberID)
	if err == nil && exists != nil {
		log.Print("CardNumberID already exist")
		return nil, customErrors.ErrorConflict
	}

	//Validate if warehouseID exist
	warehouse, _ := e.wh.GetById(request.WarehouseID)
	if warehouse == nil {
		log.Printf("Warehouse not found: %+v", warehouse)
		return nil, customErrors.ErrorConflict
	}

	//map request to model
	newEmployeeMap := mappers.RequestEmployeeToEmployee(request)

	//create new Employee
	data, err := e.rp.Create(newEmployeeMap)
	if err != nil {
		return nil, customErrors.ErrorInternalServerError
	}

	//map model to doc
	dataMap := mappers.EmployeeToEmployeeDoc(*data)
	return &dataMap, nil
}

// Update implements EmployeeService.
func (e *EmployeeDefault) Update(id int, request models.UpdateEmployee) (*models.EmployeeDoc, error) {

	///Get the existing employee
	existingEmployee, err := e.rp.GetById(id)
	if err != nil {
		return nil, customErrors.ErrorNotFound
	}

	//Validate request fields
	err = validators.ValidateUpdateEmployee(request)
	if err != nil {
		return nil, customErrors.ErrorUnprocessableContent
	}

	// Validate if cardNumberid already exist y Actualizar los campos que no son nil
	if request.CardNumberID != nil {
		//validate if cardNumberID already exist
		exists, err := e.rp.FindByCardNumberID(*request.CardNumberID)
		if err == nil && exists != nil && exists.Id != id {
			return nil, customErrors.ErrorConflict
		}

		existingEmployee.EmployeeAttributes.CardNumberID = *request.CardNumberID
	}
	if request.FirstName != nil {
		existingEmployee.EmployeeAttributes.FirstName = *request.FirstName
	}
	if request.LastName != nil {
		existingEmployee.EmployeeAttributes.LastName = *request.LastName
	}
	if request.WarehouseID != nil {
		//Validate if warehouseID exist
		warehouse, _ := e.wh.GetById(*request.WarehouseID)
		if warehouse == nil {
			log.Printf("Warehouse not found: %+v", warehouse)
			return nil, customErrors.ErrorConflict
		}
		existingEmployee.EmployeeAttributes.WarehouseID = *request.WarehouseID
	}

	// Save the updated employee
	err = e.rp.Update(id, existingEmployee)
	if err != nil {
		return nil, err
	}

	// Map model to doc
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

// GetReportInboundOrders implements EmployeeService.
func (e *EmployeeDefault) GetReportInboundOrders(employeeID *int) (any, error) {
	// GetReportInboundOrders implements EmployeeService.
	if employeeID != nil {
		// Validate if employee exists
		employeeExists, err := e.rp.GetById(*employeeID)
		if err != nil {
			return nil, err
		}

		employee, err := e.ior.GetReportByEmployeeID(employeeExists.Id)
		if err != nil {
			return nil, err
		}

		return employee, nil
	}

	report, err := e.ior.GetAllReport()
	if err != nil {
		log.Print("Error getting ALL report")
		return nil, customErrors.ErrorInternalServerError
	}

	return report, nil

}
