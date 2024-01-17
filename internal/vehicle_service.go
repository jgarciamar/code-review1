package internal

import "errors"

var (
	ErrFieldRequired = errors.New("field required")
	ErrFieldQuality  = errors.New("field quality")
)

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	Create(v *Vehicle) (err error)
	GetVehiclesByFuelType(fuelType string) (vehicles []Vehicle, err error)
	Delete(id int) (err error)
}
