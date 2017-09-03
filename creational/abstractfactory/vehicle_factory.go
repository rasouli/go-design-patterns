package abstractfactory

import (
	"errors"
	"fmt"
)

type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}

const (
	CarFactoryType = 5 * iota
	MotorbikeFactoryType
)

func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil

	default:
		return nil, errors.New(fmt.Sprintf("Factory with id %d not found", f))
	}
}
