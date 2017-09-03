package abstractfactory

import (
	"errors"
	"fmt"
)

type Motorbike interface {
	GetMotorbikeType() int
}

const (
	SportMotorbikeType = 2 * iota
	CruiseMotorbikeType
)

type MotorbikeFactory struct{}

func (m *MotorbikeFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, errors.New(fmt.Sprintf("the vehicle type not found %d", v))
	}
}

type SportMotorbike struct{}
type CruiseMotorbike struct{}

func (m SportMotorbike) GetMotorbikeType() int {
	return SportMotorbikeType
}

func (m SportMotorbike) NumSeats() int {
	return 2
}

func (m SportMotorbike) NumWheels() int {
	return 2
}

func (m *CruiseMotorbike) GetMotorbikeType() int {
	return CruiseMotorbikeType
}

func (m *CruiseMotorbike) NumSeats() int {
	return 2
}

func (m *CruiseMotorbike) NumWheels() int {
	return 2
}
