package builder

import (
	"testing"
)

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4, but got %d", car.Wheels)
	}

	if car.Structure != "car" {
		t.Errorf("Structure on a car must be 'car' but got %s", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 but got %d ", car.Seats)
	}

	bikeBuilder := &MotorBike{}

	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()

	motorBike := bikeBuilder.GetVehicle()

	if motorBike.Wheels != 2 {
		t.Errorf("Wheels on a bike must be 2 but got : %d ", motorBike.Wheels)
	}

	if motorBike.Seats != 1 {
		t.Errorf("Seats on motor bike must be 1 but got :%d", motorBike.Seats)
	}

	if motorBike.Structure != "motorbike" {
		t.Errorf("Expected the structure on the motorbike be 'motorbike' but got %s", motorBike.Structure)
	}

}
