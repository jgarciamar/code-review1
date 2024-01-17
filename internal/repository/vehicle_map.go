package repository

import (
	"app/internal"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{
		db: defaultDb,
	}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

// Create is a method that creates or inserts a Vehicle into the db
func (r *VehicleMap) Create(v *internal.Vehicle) (err error) {

	for _, vehicle := range r.db {

		//We check if theres a car with that ID already
		if v.Id == vehicle.Id {
			return internal.ErrVehicleIdAlreadyExists
		}
	}

	// If theres no error we update the lastId counter and insert the vehicle in the map

	r.db[v.Id] = *v

	return nil

}

// Delete removes a Vehicle from the map using the delete function
func (r *VehicleMap) Delete(id int) (err error) {
	//We check if such car is indeed in the map
	_, ok := r.db[id]

	if !ok {
		return internal.ErrVehicleNotFound
	}

	delete(r.db, id)
	return
}
