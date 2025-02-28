package employee_test

import (
	"testing"

	rpEmp "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/employee"
	rpIor "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/inbound_order"
	rpWh "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/warehouse"
	sv "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/employee"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEmployeeService(t *testing.T) {
	t.Run("create_ok", func(t *testing.T) {
		// Arrange
		repoEmp := new(rpEmp.EmployeeRepositoryMock)
		repoIor := new(rpIor.InboundOrderRepositoryMock)
		repoWh := new(rpWh.WarehouseRepositoryMock)
		service := sv.NewEmployeeDefault(repoEmp, repoIor, repoWh)

		newEmployee := models.RequestEmployee{
			CardNumberID: "123456",
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  1,
		}

		newEmployeeMap := models.Employee{
			Id: 0,
			EmployeeAttributes: models.EmployeeAttributes{
				CardNumberID: "123456",
				FirstName:    "John",
				LastName:     "Doe",
				WarehouseID:  1,
			},
		}

		createdEmployeeRepo := models.Employee{
			Id: 1,
			EmployeeAttributes: models.EmployeeAttributes{
				CardNumberID: "123456",
				FirstName:    "John",
				LastName:     "Doe",
				WarehouseID:  1,
			},
		}

		expectedEmployeeMap := models.EmployeeDoc{
			Id:           1,
			CardNumberID: "123456",
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  1,
		}

		repoEmp.On("FindByCardNumberID", newEmployee.CardNumberID).Return(nil, customErrors.ErrorNotFound)
		repoWh.On("GetById", 1).Return(&models.WarehouseDocResponse{ID: 1}, nil)
		repoEmp.On("Create", newEmployeeMap).Return(&createdEmployeeRepo, nil)

		// Act
		createdEmployee, err := service.Create(newEmployee)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, &expectedEmployeeMap, createdEmployee)
		repoEmp.AssertExpectations(t)
		repoWh.AssertExpectations(t)
	})

	t.Run("create_fail_validation", func(t *testing.T) {
		// Arrange
		repoEmp := new(rpEmp.EmployeeRepositoryMock)
		repoIor := new(rpIor.InboundOrderRepositoryMock)
		repoWh := new(rpWh.WarehouseRepositoryMock)
		service := sv.NewEmployeeDefault(repoEmp, repoIor, repoWh)

		newEmployee := models.RequestEmployee{
			CardNumberID: "",
			FirstName:    "",
			LastName:     "",
			WarehouseID:  0,
		}

		// Act
		createdEmployee, err := service.Create(newEmployee)

		// Assert
		require.Error(t, err)
		assert.Nil(t, createdEmployee)
		assert.Equal(t, customErrors.ErrorUnprocessableContent, err)
	})

	t.Run("create_conflict_invalid_warehouse", func(t *testing.T) {
		// Arrange
		repoEmp := new(rpEmp.EmployeeRepositoryMock)
		repoIor := new(rpIor.InboundOrderRepositoryMock)
		repoWh := new(rpWh.WarehouseRepositoryMock)
		service := sv.NewEmployeeDefault(repoEmp, repoIor, repoWh)

		newEmployee := models.RequestEmployee{
			CardNumberID: "123456",
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  999, // ID de almacén no válido
		}

		// si no existe el cardNumberID en la base de datos, se puede crear el empleado
		repoEmp.On("FindByCardNumberID", newEmployee.CardNumberID).Return(nil, customErrors.ErrorNotFound)
		repoWh.On("GetById", newEmployee.WarehouseID).Return(nil, customErrors.ErrorNotFound)

		// Act
		createdEmployee, err := service.Create(newEmployee)

		// Assert
		require.Error(t, err)
		assert.Nil(t, createdEmployee)
		assert.Equal(t, customErrors.ErrorConflict, err)
		repoWh.AssertExpectations(t)
	})

	t.Run("create_fail", func(t *testing.T) {
		// Arrange
		repoEmp := new(rpEmp.EmployeeRepositoryMock)
		repoIor := new(rpIor.InboundOrderRepositoryMock)
		repoWh := new(rpWh.WarehouseRepositoryMock)
		service := sv.NewEmployeeDefault(repoEmp, repoIor, repoWh)

		newEmployee := models.RequestEmployee{
			CardNumberID: "123456",
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  1,
		}

		expectedEmployeeRepo := models.Employee{
			Id: 0,
			EmployeeAttributes: models.EmployeeAttributes{
				CardNumberID: "123456",
				FirstName:    "John",
				LastName:     "Doe",
				WarehouseID:  1,
			},
		}

		// si no existe el cardNumberID en la base de datos, se puede crear el empleado
		repoEmp.On("FindByCardNumberID", newEmployee.CardNumberID).Return(nil, customErrors.ErrorNotFound)
		repoWh.On("GetById", newEmployee.WarehouseID).Return(&models.WarehouseDocResponse{ID: 1}, nil)
		repoEmp.On("Create", expectedEmployeeRepo).Return(nil, customErrors.ErrorInternalServerError)

		// Act
		createdEmployee, err := service.Create(newEmployee)

		// Assert
		require.Error(t, err)
		assert.Nil(t, createdEmployee)
		assert.Equal(t, customErrors.ErrorInternalServerError, err)
		repoEmp.AssertExpectations(t)
		repoWh.AssertExpectations(t)
	})

	t.Run("create_conflict_cardNumber", func(t *testing.T) {
		// Arrange
		repoEmp := new(rpEmp.EmployeeRepositoryMock)
		repoIor := new(rpIor.InboundOrderRepositoryMock)
		repoWh := new(rpWh.WarehouseRepositoryMock)
		service := sv.NewEmployeeDefault(repoEmp, repoIor, repoWh)

		newEmployee := models.RequestEmployee{
			CardNumberID: "123456",
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  1,
		}

		expectedEmployee := models.Employee{
			Id: 2,
			EmployeeAttributes: models.EmployeeAttributes{
				CardNumberID: "123456",
				FirstName:    "Carlos",
				LastName:     "Garcia",
				WarehouseID:  2,
			},
		}

		repoEmp.On("FindByCardNumberID", newEmployee.CardNumberID).Return(&expectedEmployee, nil)

		// Act
		createdEmployee, err := service.Create(newEmployee)

		// Assert
		require.Error(t, err)
		assert.Nil(t, createdEmployee)
		assert.Equal(t, customErrors.ErrorConflict, err)
		repoEmp.AssertExpectations(t)
	})

	t.Run("find_all", func(t *testing.T) {
		// Arrange
		repoEmp := new(rpEmp.EmployeeRepositoryMock)
		repoIor := new(rpIor.InboundOrderRepositoryMock)
		repoWh := new(rpWh.WarehouseRepositoryMock)
		service := sv.NewEmployeeDefault(repoEmp, repoIor, repoWh)

		expectedEmployees := map[int]models.Employee{
			1: {
				Id: 1,
				EmployeeAttributes: models.EmployeeAttributes{
					CardNumberID: "123456",
					FirstName:    "John",
					LastName:     "Doe",
					WarehouseID:  1,
				},
			},
			2: {
				Id: 2,
				EmployeeAttributes: models.EmployeeAttributes{
					CardNumberID: "654321",
					FirstName:    "Sara",
					LastName:     "Ramirez",
					WarehouseID:  2,
				},
			},
		}

		expectedEmployeeDocs := []models.EmployeeDoc{
			{
				Id:           1,
				CardNumberID: "123456",
				FirstName:    "John",
				LastName:     "Doe",
				WarehouseID:  1,
			},
			{
				Id:           2,
				CardNumberID: "654321",
				FirstName:    "Sara",
				LastName:     "Ramirez",
				WarehouseID:  2,
			},
		}

		repoEmp.On("GetAll").Return(expectedEmployees, nil)

		// Act
		employees, err := service.GetAll()

		// Assert
		require.NoError(t, err)
		assert.ElementsMatch(t, expectedEmployeeDocs, employees)
		repoEmp.AssertExpectations(t)
	})

	t.Run("find_all_fail", func(t *testing.T) {
		// Arrange
		repoEmp := new(rpEmp.EmployeeRepositoryMock)
		repoIor := new(rpIor.InboundOrderRepositoryMock)
		repoWh := new(rpWh.WarehouseRepositoryMock)
		service := sv.NewEmployeeDefault(repoEmp, repoIor, repoWh)

		repoEmp.On("GetAll").Return(nil, customErrors.ErrorInternalServerError)
		// Act
		employees, err := service.GetAll()

		// Assert
		require.Error(t, err)
		assert.Nil(t, employees)
		repoEmp.AssertExpectations(t)
	})

	t.Run("find_by_id_non_existent", func(t *testing.T) {
		// Arrange
		repoEmp := new(rpEmp.EmployeeRepositoryMock)
		repoIor := new(rpIor.InboundOrderRepositoryMock)
		repoWh := new(rpWh.WarehouseRepositoryMock)
		service := sv.NewEmployeeDefault(repoEmp, repoIor, repoWh)

		invalidId := 134
		repoEmp.On("GetById", invalidId).Return(nil, customErrors.ErrorNotFound)

		// Act
		employee, err := service.GetById(invalidId)

		// Assert
		require.Error(t, err)
		assert.Nil(t, employee)
		assert.Equal(t, customErrors.ErrorNotFound, err)
		repoEmp.AssertExpectations(t)
	})

	t.Run("find_by_id_existent", func(t *testing.T) {
		/// Arrange
		repoEmp := new(rpEmp.EmployeeRepositoryMock)
		repoIor := new(rpIor.InboundOrderRepositoryMock)
		repoWh := new(rpWh.WarehouseRepositoryMock)
		service := sv.NewEmployeeDefault(repoEmp, repoIor, repoWh)

		expectedEmployee := models.Employee{
			Id: 1,
			EmployeeAttributes: models.EmployeeAttributes{
				CardNumberID: "123456",
				FirstName:    "John",
				LastName:     "Doe",
				WarehouseID:  1,
			},
		}
		expectedEmployeeDoc := models.EmployeeDoc{
			Id:           1,
			CardNumberID: "123456",
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  1,
		}
		repoEmp.On("GetById", expectedEmployee.Id).Return(&expectedEmployee, nil)

		// Act
		employee, err := service.GetById(expectedEmployeeDoc.Id)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, &expectedEmployeeDoc, employee)
		repoEmp.AssertExpectations(t)
	})

	t.Run("update_existent", func(t *testing.T) {
		// Arrange
		repoEmp := new(rpEmp.EmployeeRepositoryMock)
		repoIor := new(rpIor.InboundOrderRepositoryMock)
		repoWh := new(rpWh.WarehouseRepositoryMock)
		service := sv.NewEmployeeDefault(repoEmp, repoIor, repoWh)

		idEmployee := 1
		cardNumberID := "41242"
		firstName := "Carlos"
		lastName := "Garcia"
		warehouseID := 2

		requestEmployee := models.UpdateEmployee{
			CardNumberID: &cardNumberID,
			FirstName:    &firstName,
			LastName:     &lastName,
			WarehouseID:  &warehouseID,
		}

		expectedEmployeeRepo := models.Employee{
			Id: 1,
			EmployeeAttributes: models.EmployeeAttributes{
				CardNumberID: "123456",
				FirstName:    "John",
				LastName:     "Doe",
				WarehouseID:  1,
			},
		}

		expectedUpdatedEmployee := models.Employee{
			Id: idEmployee,
			EmployeeAttributes: models.EmployeeAttributes{
				CardNumberID: cardNumberID,
				FirstName:    firstName,
				LastName:     lastName,
				WarehouseID:  warehouseID,
			},
		}

		expectedEmployeeMap := models.EmployeeDoc{
			Id:           idEmployee,
			CardNumberID: cardNumberID,
			FirstName:    firstName,
			LastName:     lastName,
			WarehouseID:  warehouseID,
		}

		repoEmp.On("GetById", idEmployee).Return(&expectedEmployeeRepo, nil)
		repoEmp.On("FindByCardNumberID", cardNumberID).Return(nil, customErrors.ErrorNotFound)
		repoWh.On("GetById", warehouseID).Return(&models.WarehouseDocResponse{ID: 2}, nil)
		repoEmp.On("Update", idEmployee, &expectedUpdatedEmployee).Return(nil)

		// Act
		updatedEmployee, err := service.Update(idEmployee, requestEmployee)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, &expectedEmployeeMap, updatedEmployee)
		repoEmp.AssertExpectations(t)
		repoWh.AssertExpectations(t)
	})

	t.Run("update_non_existent", func(t *testing.T) {
		/// Arrange
		repoEmp := new(rpEmp.EmployeeRepositoryMock)
		repoIor := new(rpIor.InboundOrderRepositoryMock)
		repoWh := new(rpWh.WarehouseRepositoryMock)
		service := sv.NewEmployeeDefault(repoEmp, repoIor, repoWh)

		idEmployee := 1
		cardNumberID := "41242"
		firstName := "Carlos"
		lastName := "Garcia"
		warehouseID := 2

		requestEmployee := models.UpdateEmployee{
			CardNumberID: &cardNumberID,
			FirstName:    &firstName,
			LastName:     &lastName,
			WarehouseID:  &warehouseID,
		}

		repoEmp.On("GetById", idEmployee).Return(nil, customErrors.ErrorNotFound)

		// Act
		updatedEmployee, err := service.Update(idEmployee, requestEmployee)

		// Assert
		require.Error(t, err)
		assert.Nil(t, updatedEmployee)
		assert.Equal(t, customErrors.ErrorNotFound, err)
		repoEmp.AssertExpectations(t)
	})

	t.Run("update_fail_unprocessable_content", func(t *testing.T) {
		// Arrange
		repoEmp := new(rpEmp.EmployeeRepositoryMock)
		repoIor := new(rpIor.InboundOrderRepositoryMock)
		repoWh := new(rpWh.WarehouseRepositoryMock)
		service := sv.NewEmployeeDefault(repoEmp, repoIor, repoWh)

		idEmployee := 1
		cardNumberID := ""
		firstName := "   "
		lastName := "Garcia"
		warehouseID := 0

		requestEmployee := models.UpdateEmployee{
			CardNumberID: &cardNumberID,
			FirstName:    &firstName,
			LastName:     &lastName,
			WarehouseID:  &warehouseID,
		}

		expectedEmployeeRepo := models.Employee{
			Id: 1,
			EmployeeAttributes: models.EmployeeAttributes{
				CardNumberID: "123456",
				FirstName:    "John",
				LastName:     "Doe",
				WarehouseID:  1,
			},
		}

		repoEmp.On("GetById", idEmployee).Return(&expectedEmployeeRepo, nil)

		// Act
		updatedEmployee, err := service.Update(idEmployee, requestEmployee)

		// Assert
		require.Error(t, err)
		assert.Nil(t, updatedEmployee)
		assert.Equal(t, customErrors.ErrorUnprocessableContent, err)
		repoEmp.AssertExpectations(t)
	})

	t.Run("update_conflict_cardnumber", func(t *testing.T) {
		// Arrange
		repoEmp := new(rpEmp.EmployeeRepositoryMock)
		repoIor := new(rpIor.InboundOrderRepositoryMock)
		repoWh := new(rpWh.WarehouseRepositoryMock)
		service := sv.NewEmployeeDefault(repoEmp, repoIor, repoWh)

		idEmployee := 1
		cardNumberID := "41242"
		firstName := "Carlos"
		lastName := "Garcia"
		warehouseID := 2

		requestEmployee := models.UpdateEmployee{
			CardNumberID: &cardNumberID,
			FirstName:    &firstName,
			LastName:     &lastName,
			WarehouseID:  &warehouseID,
		}

		expectedEmployeeRepo := models.Employee{
			Id: 1,
			EmployeeAttributes: models.EmployeeAttributes{
				CardNumberID: "123456",
				FirstName:    "John",
				LastName:     "Doe",
				WarehouseID:  1,
			},
		}

		employeeWithSameCardNumber := models.Employee{
			Id: 2,
			EmployeeAttributes: models.EmployeeAttributes{
				CardNumberID: "41242",
				FirstName:    "Jose",
				LastName:     "Garcia",
				WarehouseID:  1,
			},
		}

		repoEmp.On("GetById", idEmployee).Return(&expectedEmployeeRepo, nil)
		repoEmp.On("FindByCardNumberID", cardNumberID).Return(&employeeWithSameCardNumber, nil)

		// Act
		updatedEmployee, err := service.Update(idEmployee, requestEmployee)

		// Assert
		require.Error(t, err)
		assert.Nil(t, updatedEmployee)
		assert.Equal(t, customErrors.ErrorConflict, err)
		repoEmp.AssertExpectations(t)
	})

	t.Run("update_conflict_warehouse", func(t *testing.T) {
		// Arrange
		repoEmp := new(rpEmp.EmployeeRepositoryMock)
		repoIor := new(rpIor.InboundOrderRepositoryMock)
		repoWh := new(rpWh.WarehouseRepositoryMock)
		service := sv.NewEmployeeDefault(repoEmp, repoIor, repoWh)

		idEmployee := 1
		cardNumberID := "41242"
		firstName := "Carlos"
		lastName := "Garcia"
		warehouseID := 2132

		requestEmployee := models.UpdateEmployee{
			CardNumberID: &cardNumberID,
			FirstName:    &firstName,
			LastName:     &lastName,
			WarehouseID:  &warehouseID,
		}

		expectedEmployeeRepo := models.Employee{
			Id: 1,
			EmployeeAttributes: models.EmployeeAttributes{
				CardNumberID: "123456",
				FirstName:    "John",
				LastName:     "Doe",
				WarehouseID:  1,
			},
		}

		repoEmp.On("GetById", idEmployee).Return(&expectedEmployeeRepo, nil)
		repoEmp.On("FindByCardNumberID", cardNumberID).Return(nil, customErrors.ErrorNotFound)
		repoWh.On("GetById", warehouseID).Return(nil, customErrors.ErrorNotFound)

		// Act
		updatedEmployee, err := service.Update(idEmployee, requestEmployee)

		// Assert
		require.Error(t, err)
		assert.Nil(t, updatedEmployee)
		assert.Equal(t, customErrors.ErrorConflict, err)
		repoEmp.AssertExpectations(t)
	})

	t.Run("update_fail", func(t *testing.T) {
		// Arrange
		repoEmp := new(rpEmp.EmployeeRepositoryMock)
		repoIor := new(rpIor.InboundOrderRepositoryMock)
		repoWh := new(rpWh.WarehouseRepositoryMock)
		service := sv.NewEmployeeDefault(repoEmp, repoIor, repoWh)

		idEmployee := 1
		cardNumberID := "41242"
		firstName := "Carlos"
		lastName := "Garcia"
		warehouseID := 2

		requestEmployee := models.UpdateEmployee{
			CardNumberID: &cardNumberID,
			FirstName:    &firstName,
			LastName:     &lastName,
			WarehouseID:  &warehouseID,
		}

		expectedEmployeeRepo := models.Employee{
			Id: 1,
			EmployeeAttributes: models.EmployeeAttributes{
				CardNumberID: "123456",
				FirstName:    "John",
				LastName:     "Doe",
				WarehouseID:  1,
			},
		}

		expectedUpdatedEmployee := models.Employee{
			Id: idEmployee,
			EmployeeAttributes: models.EmployeeAttributes{
				CardNumberID: cardNumberID,
				FirstName:    firstName,
				LastName:     lastName,
				WarehouseID:  warehouseID,
			},
		}

		repoEmp.On("GetById", idEmployee).Return(&expectedEmployeeRepo, nil)
		repoEmp.On("FindByCardNumberID", cardNumberID).Return(nil, customErrors.ErrorNotFound)
		repoWh.On("GetById", warehouseID).Return(&models.WarehouseDocResponse{ID: 2}, nil)
		repoEmp.On("Update", idEmployee, &expectedUpdatedEmployee).Return(customErrors.ErrorInternalServerError)

		// Act
		updatedEmployee, err := service.Update(idEmployee, requestEmployee)

		// Assert
		require.Error(t, err)
		assert.Nil(t, updatedEmployee)
		assert.Equal(t, customErrors.ErrorInternalServerError, err)

		repoEmp.AssertExpectations(t)
		repoWh.AssertExpectations(t)
	})

	t.Run("delete_non_existent", func(t *testing.T) {
		// Arrange
		repoEmp := new(rpEmp.EmployeeRepositoryMock)
		repoIor := new(rpIor.InboundOrderRepositoryMock)
		repoWh := new(rpWh.WarehouseRepositoryMock)
		service := sv.NewEmployeeDefault(repoEmp, repoIor, repoWh)

		idEmployee := 1

		repoEmp.On("Delete", idEmployee).Return(customErrors.ErrorNotFound)

		// Act
		err := service.Delete(idEmployee)

		// Assert
		require.Error(t, err)
		assert.Equal(t, customErrors.ErrorNotFound, err)
		repoEmp.AssertExpectations(t)
	})

	t.Run("delete_ok", func(t *testing.T) {
		// Arrange
		repoEmp := new(rpEmp.EmployeeRepositoryMock)
		repoIor := new(rpIor.InboundOrderRepositoryMock)
		repoWh := new(rpWh.WarehouseRepositoryMock)
		service := sv.NewEmployeeDefault(repoEmp, repoIor, repoWh)

		idEmployee := 1

		repoEmp.On("Delete", idEmployee).Return(nil)
		repoEmp.On("GetById", idEmployee).Return(nil, customErrors.ErrorNotFound)

		// Act
		err := service.Delete(idEmployee)

		// Assert
		require.NoError(t, err)

		// Verificar que el empleado ya no está en la lista
		employee, err := service.GetById(idEmployee)
		require.Error(t, err)
		assert.Nil(t, employee)
		assert.Equal(t, customErrors.ErrorNotFound, err)
		repoEmp.AssertExpectations(t)
	})

	t.Run("get_report_inbound_order_by_id_success", func(t *testing.T) {
		// Arrange
		repoEmp := new(rpEmp.EmployeeRepositoryMock)
		repoIor := new(rpIor.InboundOrderRepositoryMock)
		repoWh := new(rpWh.WarehouseRepositoryMock)
		service := sv.NewEmployeeDefault(repoEmp, repoIor, repoWh)

		employeeID := 1
		expectedEmployee := models.Employee{
			Id: 1,
			EmployeeAttributes: models.EmployeeAttributes{
				CardNumberID: "123456",
				FirstName:    "John",
				LastName:     "Doe",
				WarehouseID:  1,
			},
		}
		expectedReport := models.EmployeeWithOrders{
			Id:                 1,
			CardNumberID:       "123456",
			FirstName:          "John",
			LastName:           "Doe",
			WarehouseID:        1,
			InboundOrdersCount: 2,
		}

		repoEmp.On("GetById", employeeID).Return(&expectedEmployee, nil)
		repoIor.On("GetReportByEmployeeID", employeeID).Return(&expectedReport, nil)

		// Act
		report, err := service.GetReportInboundOrders(&employeeID)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, &expectedReport, report)
		repoEmp.AssertExpectations(t)
		repoIor.AssertExpectations(t)
	})

	t.Run("get_report_inbound_orders_with_id_employee_not_found", func(t *testing.T) {
		// Arrange
		repoEmp := new(rpEmp.EmployeeRepositoryMock)
		repoIor := new(rpIor.InboundOrderRepositoryMock)
		repoWh := new(rpWh.WarehouseRepositoryMock)
		service := sv.NewEmployeeDefault(repoEmp, repoIor, repoWh)

		employeeID := 1

		repoEmp.On("GetById", employeeID).Return(nil, customErrors.ErrorNotFound)

		// Act
		report, err := service.GetReportInboundOrders(&employeeID)

		// Assert
		require.Error(t, err)
		assert.Nil(t, report)
		assert.Equal(t, customErrors.ErrorNotFound, err)
		repoEmp.AssertExpectations(t)
	})

	t.Run("get_report_inbound_orders_by_id_fail", func(t *testing.T) {
		// Arrange
		repoEmp := new(rpEmp.EmployeeRepositoryMock)
		repoIor := new(rpIor.InboundOrderRepositoryMock)
		repoWh := new(rpWh.WarehouseRepositoryMock)
		service := sv.NewEmployeeDefault(repoEmp, repoIor, repoWh)

		employeeID := 1
		expectedEmployee := models.Employee{
			Id: 1,
			EmployeeAttributes: models.EmployeeAttributes{
				CardNumberID: "123456",
				FirstName:    "John",
				LastName:     "Doe",
				WarehouseID:  1,
			},
		}

		repoEmp.On("GetById", employeeID).Return(&expectedEmployee, nil)
		repoIor.On("GetReportByEmployeeID", employeeID).Return(nil, customErrors.ErrorInternalServerError)

		// Act
		report, err := service.GetReportInboundOrders(&employeeID)

		// Assert
		require.Error(t, err)
		assert.Nil(t, report)
		assert.Equal(t, customErrors.ErrorInternalServerError, err)
		repoEmp.AssertExpectations(t)
	})

	t.Run("get_all_report_inbound_orders_success", func(t *testing.T) {
		// Arrange
		repoEmp := new(rpEmp.EmployeeRepositoryMock)
		repoIor := new(rpIor.InboundOrderRepositoryMock)
		repoWh := new(rpWh.WarehouseRepositoryMock)
		service := sv.NewEmployeeDefault(repoEmp, repoIor, repoWh)

		expectedReport := []models.EmployeeWithOrders{
			{
				Id:                 1,
				CardNumberID:       "123456",
				FirstName:          "John",
				LastName:           "Doe",
				WarehouseID:        1,
				InboundOrdersCount: 2,
			},
			{
				Id:                 2,
				CardNumberID:       "654321",
				FirstName:          "Sara",
				LastName:           "Ramirez",
				WarehouseID:        2,
				InboundOrdersCount: 4,
			},
		}

		repoIor.On("GetAllReport").Return(expectedReport, nil)

		// Act
		report, err := service.GetReportInboundOrders(nil)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, expectedReport, report)
		repoIor.AssertExpectations(t)
	})

	t.Run("get_all_fail_report_inbound_orders", func(t *testing.T) {
		// Arrange
		repoEmp := new(rpEmp.EmployeeRepositoryMock)
		repoIor := new(rpIor.InboundOrderRepositoryMock)
		repoWh := new(rpWh.WarehouseRepositoryMock)
		service := sv.NewEmployeeDefault(repoEmp, repoIor, repoWh)

		repoIor.On("GetAllReport").Return(nil, customErrors.ErrorInternalServerError)

		// Act
		report, err := service.GetReportInboundOrders(nil)

		// Assert
		require.Error(t, err)
		assert.Nil(t, report)
		assert.Equal(t, customErrors.ErrorInternalServerError, err)
		repoIor.AssertExpectations(t)
	})

}
