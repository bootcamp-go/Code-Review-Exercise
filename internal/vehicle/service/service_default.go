package service

import (
	"app/internal/domain"
	"app/internal/vehicle/repository"
	"errors"
	"fmt"
)

// ServiceVehicleDefault is an struct that represents a vehicle service.
type ServiceVehicleDefault struct {
	rp repository.RepositoryVehicle
}

// NewServiceVehicleDefault returns a new instance of a vehicle service.
func NewServiceVehicleDefault(rp repository.RepositoryVehicle) *ServiceVehicleDefault {
	return &ServiceVehicleDefault{rp: rp}
}

// GetAll returns all vehicles.
func (s *ServiceVehicleDefault) GetAll() (v []*domain.Vehicle, err error) {
	v, err = s.rp.GetAll()
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrRepositoryVehicleNotFound):
			err = fmt.Errorf("%w. %v", ErrServiceVehicleNotFound, err)
		default:
			err = fmt.Errorf("%w. %v", ErrServiceVehicleInternal, err)
		}
	
		return
	}

	return
}