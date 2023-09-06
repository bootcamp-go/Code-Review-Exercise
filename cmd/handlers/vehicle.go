package handlers

import (
	"app/internal/vehicle/storage"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewControllerVehicle returns a new instance of a vehicle controller.
func NewControllerVehicle(st storage.StorageVehicle) *ControllerVehicle {
	return &ControllerVehicle{st: st}
}

// ControllerVehicle is an struct that represents a vehicle controller.
type ControllerVehicle struct {
	// StorageVehicle is the storage of vehicles.
	st storage.StorageVehicle
}

// GetAll returns all vehicles.
type VehicleHandlerGetAll struct {
	Id				int		`json:"id"`
	Brand			string	`json:"brand"`
	Model			string	`json:"model"`
	Registration	string	`json:"registration"`
	Year			int		`json:"year"`
	Color			string	`json:"color"`
	MaxSpeed		int		`json:"max_speed"`
	FuelType		string	`json:"fuel_type"`
	Transmission	string	`json:"transmission"`
	Passengers 		int		`json:"passengers"`
	Height			int		`json:"height"`
	Width			int		`json:"width"`
	Weight			int		`json:"weight"`
}
type ResponseBodyGetAll struct {
	Message	string			   		`json:"message"`
	Data	[]*VehicleHandlerGetAll `json:"vehicles"`
	Error	bool			   		`json:"error"`
}
func (c *ControllerVehicle) GetAll() (gin.HandlerFunc) {
	return func(ctx *gin.Context) {
		// request
		// ...
		
		// process
		vehicles, err := c.st.GetAll()
		if err != nil {
			var code int; var body *ResponseBodyGetAll
			switch {
			case errors.Is(err, storage.ErrStorageVehicleNotFound):
				code = http.StatusNotFound
				body = &ResponseBodyGetAll{Message: "Not found", Error: true}
			default:
				code = http.StatusInternalServerError
				body = &ResponseBodyGetAll{Message: "Internal server error", Error: true}
			}

			ctx.JSON(code, body)
			return
		}

		// response
		code := http.StatusOK
		body := &ResponseBodyGetAll{Message: "Success", Data: make([]*VehicleHandlerGetAll, len(vehicles)), Error: false}
		for _, vehicle := range vehicles {
			body.Data = append(body.Data, &VehicleHandlerGetAll{
				Id:				vehicle.Id,
				Brand:			vehicle.Attributes.Brand,
				Model:			vehicle.Attributes.Model,
				Registration:	vehicle.Attributes.Registration,
				Year:			vehicle.Attributes.Year,
				Color:			vehicle.Attributes.Color,
				MaxSpeed:		vehicle.Attributes.MaxSpeed,
				FuelType:		vehicle.Attributes.FuelType,
				Transmission:	vehicle.Attributes.Transmission,
				Passengers:		vehicle.Attributes.Passengers,
				Height:			vehicle.Attributes.Height,
				Width:			vehicle.Attributes.Width,
				Weight:			vehicle.Attributes.Weight,
			})
		}

		ctx.JSON(code, body)
	}
}
