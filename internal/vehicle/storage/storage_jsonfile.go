package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

// StorageVehicleJSONFile is an struct that represents a vehicle storage using JSON files.
type StorageVehicleJSONFile struct {
	// Path is the path of the JSON file.
	Path string
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
	err = json.NewDecoder(f).Decode(&v)
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageVehicleInternal, err)
		return
	}

	// check if there are vehicles
	if len(v) == 0 {
		err = fmt.Errorf("%w. %v", ErrStorageVehicleNotFound, err)
		return
	}

	return
}