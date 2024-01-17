package service_test

import (
	"app/internal"
	"app/internal/repository"
	"app/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVehicleDefault_Create(t *testing.T) {

	t.Run("Sucess: Created a Vehicle", func(t *testing.T) {

		db := make(map[int]internal.Vehicle)
		rp := repository.NewVehicleMap(db)
		sv := service.NewVehicleDefault(rp)

		v := internal.Vehicle{
			Id: 1,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           "Ford",
				Model:           "Fiesta",
				FuelType:        "Gasoline",
				Registration:    "ABC-1234",
				Color:           "Red",
				FabricationYear: 2020,
				Capacity:        4,
				MaxSpeed:        200,
				Transmission:    "Manual",
				Weight:          1000,
				Dimensions: internal.Dimensions{
					Height: 1.5,
					Length: 4.0,
					Width:  1.8,
				},
			},
		}

		err := sv.Create(&v)

		assert.NoError(t, err, "Should not return an error")
		assert.NotEqual(t, len(db), 0, "The db should not be empty")

	})
	t.Run("Error 01: Should fail after not passing a valid Brand to the Create method", func(t *testing.T) {

		db := make(map[int]internal.Vehicle)
		rp := repository.NewVehicleMap(db)
		sv := service.NewVehicleDefault(rp)

		v := internal.Vehicle{
			Id: 1,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           "",
				Model:           "Fiesta",
				FuelType:        "Gasoline",
				Registration:    "ABC-1234",
				Color:           "Red",
				FabricationYear: 2020,
				Capacity:        4,
				MaxSpeed:        200,
				Transmission:    "Manual",
				Weight:          1000,
				Dimensions: internal.Dimensions{
					Height: 1.5,
					Length: 4.0,
					Width:  1.8,
				},
			},
		}

		err := sv.Create(&v)

		assert.ErrorIs(t, err, internal.ErrFieldQuality, "Should return an error since Brand is empty")
	})
}

func TestDefaultVehicle_Delete(t *testing.T) {
	t.Run("Success: Delete a Vehicle", func(t *testing.T) {

		db := map[int]internal.Vehicle{
			1: internal.Vehicle{
				Id: 1,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Ford",
					Model:           "Fiesta",
					FuelType:        "Gasoline",
					Registration:    "ABC-1234",
					Color:           "Red",
					FabricationYear: 2020,
					Capacity:        4,
					MaxSpeed:        200,
					Transmission:    "Manual",
					Weight:          1000,
					Dimensions: internal.Dimensions{
						Height: 1.5,
						Length: 4.0,
						Width:  1.8,
					},
				},
			},
		}

		rp := repository.NewVehicleMap(db)
		sv := service.NewVehicleDefault(rp)

		err := sv.Delete(1)

		assert.NoError(t, err, "Should not return an error")
		assert.Equal(t, len(db), 0, "The db should be empty")
	})
}
