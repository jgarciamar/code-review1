package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVehicleMap_Create(t *testing.T) {

	// Arrange

	db := make(map[int]internal.Vehicle)
	rp := repository.NewVehicleMap(db)

	v := internal.Vehicle{
		Id: 1,
		VehicleAttributes: internal.VehicleAttributes{
			Brand:           "Toyota",
			Model:           "Camry",
			Registration:    "ABC123",
			Color:           "Silver",
			FabricationYear: 2022,
			Capacity:        5,
			MaxSpeed:        180.5,
			FuelType:        "Petrol",
			Transmission:    "Automatic",
			Weight:          1500.0,
			Dimensions: internal.Dimensions{
				Length: 4.5,
				Width:  1.8,
				Height: 1.5,
			},
		},
	}

	// Act

	err := rp.Create(&v)

	// Assert

	assert.NoError(t, err)
	assert.NotEmpty(t, db)
	assert.Equal(t, v, db[1])

}
