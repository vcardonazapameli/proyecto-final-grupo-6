package employee

// VehicleLoader is an interface that represents the loader for vehicles
type VehicleLoader interface {
	// Load is a method that loads the vehicles
	Load() (v any, err error)
}
