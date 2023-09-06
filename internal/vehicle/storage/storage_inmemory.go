package storage

func NewStorageVehicleInMemory(db map[int]*VehicleAttributes) *StorageVehicleInMemory {
	return &StorageVehicleInMemory{db: db}
}

// StorageVehicleInMemory is an struct that represents a vehicle storage in memory.
type StorageVehicleInMemory struct {
	// db is the database of vehicles.
	db map[int]*VehicleAttributes
}

// GetAll returns all vehicles
func (s *StorageVehicleInMemory) GetAll() (v []*Vehicle, err error) {
	// check if the database is empty
	if len(s.db) == 0 {
		err = ErrStorageVehicleNotFound
		return
	}

	// get all vehicles from the database
	for key, value := range s.db {
		v = append(v, &Vehicle{Id: key, Attributes: *value})
	}

	return
}