package internal

import "errors"

var (
	ErrVehicleIdAlreadyExists = errors.New("vehicle id already exists")
	ErrVehicleNotFound        = errors.New("vehicle not found")
)

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	Create(v *Vehicle) (err error)
	Delete(id int) (err error)
}
