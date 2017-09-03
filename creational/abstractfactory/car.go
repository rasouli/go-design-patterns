package abstractfactory

import (
	"errors"
	"fmt"
)

type Car interface {
	NumDoors() int
	NumWheels() int
	NumSeats() int
}

const (
	LuxuryCarType = iota
	FamilyCarType
)

type CarFactory struct{}

func (c *CarFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New(fmt.Sprint("Did Not Understand Car Type %d", v))
	}
}

type LuxuryCar struct{}

func (c *LuxuryCar) NumDoors() int {
	return 4
}

func (c *LuxuryCar) NumWheels() int {
	return 4
}

func (c *LuxuryCar) NumSeats() int {
	return 5
}

type FamilyCar struct{}

func (c *FamilyCar) NumDoors() int {
	return 5
}

func (c *FamilyCar) NumWheels() int {
	return 4
}

func (c *FamilyCar) NumSeats() int {
	return 5
}
