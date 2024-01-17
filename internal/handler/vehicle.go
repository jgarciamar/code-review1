package handler

import (
	"app/internal"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{
		sv: sv,
	}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// Create is a method that returns a handler for the route POST /vehicles
func (h *VehicleDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var body VehicleJSON

		// We try to parse the body of the request to a VehicleJSON struct
		if err := request.JSON(r, &body); err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid request")
			return
		}

		// Process
		// We serialize a Vehicle struct without any Id

		vehicle := internal.Vehicle{
			Id: body.ID,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           body.Brand,
				Model:           body.Model,
				Registration:    body.Registration,
				Color:           body.Color,
				FabricationYear: body.FabricationYear,
				Capacity:        body.Capacity,
				MaxSpeed:        body.MaxSpeed,
				FuelType:        body.FuelType,
				Transmission:    body.Transmission,
				Weight:          body.Weight,
				Dimensions: internal.Dimensions{
					Height: body.Height,
					Length: body.Length,
					Width:  body.Width,
				},
			},
		}

		if err := h.sv.Create(&vehicle); err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleIdAlreadyExists):
				// We return a 409
				response.Error(w, http.StatusConflict, "Id already exists")
			case errors.Is(err, internal.ErrFieldQuality):
				response.Error(w, http.StatusConflict, "invalid request, field required")
			default:
				response.Error(w, http.StatusInternalServerError, "Internal server error")
			}
			return
		}

		// Finally we return the response with its id

		responseJSON := VehicleJSON{
			ID: vehicle.Id,
			// Vehicle Attributes
			Brand:           vehicle.Brand,
			Model:           vehicle.Model,
			Registration:    vehicle.Registration,
			Color:           vehicle.Color,
			FabricationYear: vehicle.FabricationYear,
			Capacity:        vehicle.Capacity,
			MaxSpeed:        vehicle.MaxSpeed,
			FuelType:        vehicle.FuelType,
			Transmission:    vehicle.Transmission,
			Weight:          vehicle.Weight,
			// Dimensions
			Height: vehicle.Height,
			Length: vehicle.Length,
			Width:  vehicle.Width,
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "Sucess",
			"data":    responseJSON,
		})
	}
}

func (h *VehicleDefault) GetVehiclesByFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fuelType := chi.URLParam(r, "type")

		vehiclesByFuelType, err := h.sv.GetVehiclesByFuelType(fuelType)

		if err != nil {
			fmt.Println(err)
			return
		}

		if len(vehiclesByFuelType) == 0 {
			response.Error(w, http.StatusNotFound, "no vehicles found with this fuel type")
			return
		}

		response.JSON(w, http.StatusOK, vehiclesByFuelType)
	}
}

func (h *VehicleDefault) Delete() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			response.Error(w, http.StatusNotFound, "invalid id")
			return
		}

		if err = h.sv.Delete(id); err != nil {
			response.Error(w, http.StatusNotFound, "Did not found vehicle")
			return
		}

		response.JSON(w, http.StatusNoContent, nil)
	}
}
