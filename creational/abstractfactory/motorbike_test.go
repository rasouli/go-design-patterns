package abstractfactory

import (
	"fmt"
	"testing"
)

func TestMotorbikeFactory(t *testing.T) {
	motorbikeF, err := BuildFactory(MotorbikeFactoryType)

	if err != nil {
		t.Fatal(fmt.Sprintf("could not build motorbike factory, error: %#v", err))
	}

	motorbikeVehicle, err := motorbikeF.NewVehicle(SportMotorbikeType)

	if err != nil {
		t.Fatal(fmt.Sprintf("could not make motorbike type of SportMotorbikeType , error: %#v", err))
	}

	t.Logf(fmt.Sprintf("the motorbike data: %#v", motorbikeVehicle))

	sportBike, _ := motorbikeVehicle.(SportMotorbike)

	// if !ok {
	// 	t.Errorf("could not cast sportbike")
	// }

	if sportBike.NumSeats() != 1 {
		t.Errorf("expected the sport bike to have one seat got %d", sportBike.NumSeats())
	}

}
