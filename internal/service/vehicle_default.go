package service

import (
	"app/internal"
	"errors"
	"fmt"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

func (s *VehicleDefault) Create(v *internal.Vehicle) (err error) {

	if err := ValidateVehicle(*v); err != nil {
		return internal.ErrFieldQuality
	}

	err = s.rp.Create(v)

	if err != nil {
		return err
	}

	return

}

func (s *VehicleDefault) GetVehiclesByFuelType(fuelType string) (vehicles []internal.Vehicle, err error) {

	allVehicles, err := s.rp.FindAll()

	if err != nil {
		return
	}

	for _, v := range allVehicles {
		if v.FuelType == fuelType {
			vehicles = append(vehicles, v)
		}
	}

	return
}

func (s *VehicleDefault) Delete(id int) (err error) {

	if err = s.rp.Delete(id); err != nil {
		switch err {
		case internal.ErrVehicleNotFound:
			//Wrapping
			err = fmt.Errorf("%w: id", internal.ErrVehicleNotFound)
		default:
			err = fmt.Errorf("%w id", errors.New("error deleting Vehicle"))
		}
		return
	}
	return
}

func ValidateVehicle(v internal.Vehicle) (err error) {

	if v.Brand == "" {
		return errors.New("brand is required")
	}

	if v.Model == "" {
		return errors.New("model is required")
	}

	return nil
}
