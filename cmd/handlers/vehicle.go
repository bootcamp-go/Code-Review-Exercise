package handlers

import (
	"app/internal/vehicle/storage"

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
		// ...

		// response
		// ...
	}
}
