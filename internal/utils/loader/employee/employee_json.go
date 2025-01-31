package employee

import (
	"encoding/json"
	"os"

	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

// NewVehicleJSONFile is a function that returns a new instance of VehicleJSONFile
func NewEmployeeJSONFile(path string) EmployeeLoader {
	return &EmployeeJSONFile{
		path: path,
	}
}

// VehicleJSONFile is a struct that implements the LoaderVehicle interface
type EmployeeJSONFile struct {
	// path is the path to the file that contains the vehicles in JSON format
	path string
}

// Load is a method that loads the vehicles
func (l *EmployeeJSONFile) Load() (v map[int]models.Employee, err error) {
	// open file
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	//decode file
	var employeeJSON []models.EmployeeDoc
	err = json.NewDecoder(file).Decode(&employeeJSON)
	if err != nil {
		return
	}

	// serialize vehicles
	v = make(map[int]models.Employee)
	for _, em := range employeeJSON {
		v[em.Id] = models.Employee{
			Id: em.Id,
			EmployeeAttributes: models.EmployeeAttributes{
				CardNumberID: em.CardNumberID,
				FirstName:    em.FirstName,
				LastName:     em.LastName,
				WarehouseID:  em.WarehouseID,
			},
		}
	}

	return
}
