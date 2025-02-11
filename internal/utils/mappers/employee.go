package mappers

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

func EmployeeDocToEmployee(EmployeeDoc models.EmployeeDoc) models.Employee {
	return models.Employee{
		Id: EmployeeDoc.Id,
		EmployeeAttributes: models.EmployeeAttributes{
			CardNumberID: EmployeeDoc.CardNumberID,
			FirstName:    EmployeeDoc.FirstName,
			LastName:     EmployeeDoc.LastName,
			WarehouseID:  EmployeeDoc.WarehouseID,
		},
	}
}

func RequestEmployeeToEmployee(request models.RequestEmployee) models.Employee {
	return models.Employee{
		Id:                 0,
		EmployeeAttributes: models.EmployeeAttributes(request),
	}
}

func EmployeeToEmployeeDoc(Employee models.Employee) models.EmployeeDoc {
	return models.EmployeeDoc{
		Id:           Employee.Id,
		CardNumberID: Employee.EmployeeAttributes.CardNumberID,
		FirstName:    Employee.EmployeeAttributes.FirstName,
		LastName:     Employee.EmployeeAttributes.LastName,
		WarehouseID:  Employee.EmployeeAttributes.WarehouseID,
	}
}
