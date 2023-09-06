package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

// NewStorageVehicleJSONFile returns a new instance of a vehicle storage using JSON files.
func NewStorageVehicleJSONFile(path string) *StorageVehicleJSONFile {
	return &StorageVehicleJSONFile{Path: path}
}

// StorageVehicleJSONFile is an struct that represents a vehicle storage using JSON files.
type StorageVehicleJSONFile struct {
	// Path is the path of the JSON file.
	Path string
}

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
func (st *StorageVehicleJSONFile) GetAll() (v []*Vehicle, err error) {
	// open file
	f, err := os.Open(st.Path)
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageVehicleInternal, err)
		return
	}
	defer f.Close()
	
	// decode file
	var vj []*VehicleJSON
	err = json.NewDecoder(f).Decode(&vj)
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageVehicleInternal, err)
		return
	}

	// check if there are vehicles
	if len(vj) == 0 {
		err = fmt.Errorf("%w. %v", ErrStorageVehicleNotFound, err)
		return
	}

	// serialization
	v = make([]*Vehicle, len(vj))
	for i, vehicle := range vj {
		v[i] = &Vehicle{
			Id: vehicle.Id,
			Attributes: VehicleAttributes{
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