package repository

import (
	"app/internal/domain"
	"encoding/json"
	"fmt"
	"os"
)

// NewRepositoryVehicleJSONFile returns a new instance of a vehicle repository using JSON files.
func NewRepositoryVehicleJSONFile(path string) *RepositoryVehicleJSONFile {
	return &RepositoryVehicleJSONFile{Path: path}
}

// RepositoryVehicleJSONFile is an struct that represents a vehicle repository using JSON files.
type RepositoryVehicleJSONFile struct {
	// Path is the path of the JSON file.
	Path string
}

// VehicleJSON is an struct that represents a vehicle in JSON format.
type VehicleJSON struct {
	Id 			 int		`json:"id"`
	Brand 		 string		`json:"brand"`
	Model 		 string		`json:"model"`
	Registration string		`json:"registration"`
	Year 		 int		`json:"year"`
	Color 		 string		`json:"color"`
	MaxSpeed 	 int		`json:"max_speed"`
	FuelType 	 string		`json:"fuel_type"`
	Transmission string		`json:"transmission"`
	Passengers 	 int		`json:"passengers"`
	Height 		 float64	`json:"height"`
	Width 		 float64	`json:"width"`
	Weight 		 float64	`json:"weight"`
}

// GetAll returns all vehicles.
func (rp *RepositoryVehicleJSONFile) GetAll() (v []*domain.Vehicle, err error) {
	// open file
	f, err := os.Open(rp.Path)
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrRepositoryVehicleInternal, err)
		return
	}
	defer f.Close()
	
	// decode file
	var vj []*VehicleJSON
	err = json.NewDecoder(f).Decode(&vj)
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrRepositoryVehicleInternal, err)
		return
	}

	// check if there are vehicles
	if len(vj) == 0 {
		err = fmt.Errorf("%w. %v", ErrRepositoryVehicleNotFound, err)
		return
	}

	// serialization
	v = make([]*domain.Vehicle, len(vj))
	for i, vehicle := range vj {
		v[i] = &domain.Vehicle{
			Id: vehicle.Id,
			Attributes: domain.VehicleAttributes{
				Brand: vehicle.Brand,
				Model: vehicle.Model,
				Registration: vehicle.Registration,
				Year: vehicle.Year,
				Color: vehicle.Color,
				MaxSpeed: vehicle.MaxSpeed,
				FuelType: vehicle.FuelType,
				Transmission: vehicle.Transmission,
				Passengers: vehicle.Passengers,
				Height: vehicle.Height,
				Width: vehicle.Width,
				Weight: vehicle.Weight,
			},
		}
	}

	return
}