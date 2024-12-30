package builderpattern

import (
	"fmt"
	"testing"
)

func TestBuilderPattern(t *testing.T) {
	carBuilder := &CarBuilder{}
	manufacturingComplex := ManufacturingDirector{}

	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()

	if car.wheels != 4 {
		t.Errorf("Wheels on a car must be 4 and they were %d\n", car.wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'Car' and was %s\n", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they were %d\n", car.Seats)
	}

	// Test Bike Builder
	bikeBuilder := &BikeBuilder{}

	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()

	motorBike := bikeBuilder.GetVehicle()
	motorBike.Seats = 1
	fmt.Println(motorBike.wheels)

	if motorBike.wheels != 2 {
		t.Errorf("Wheels on a motorbike must be 2 and they were %d\n", motorBike.wheels)
	}

	if motorBike.Structure != "MotorBike" {
		t.Errorf("Structure on a motorbike must be 'Motorbike' and was %s\n", motorBike.Structure)
	}
}
