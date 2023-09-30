package test

import (
	"os"
	"testing"

	"github.com/KafkaWannaFly/singoton"
	"github.com/stretchr/testify/assert"
)

func TestGetDifferentImplementOfAnInterface(t *testing.T) {
	var vehicles []VehicleNamer

	motobike, _ := singoton.Get[Motobike]()
	vehicles = append(vehicles, motobike)

	car, _ := singoton.Get[Car]()
	vehicles = append(vehicles, car)

	tank, _ := singoton.Get[Tank]()
	vehicles = append(vehicles, tank)

	assert.Equal(t, vehicles[0].GetName(), "Motobike")
	assert.Equal(t, vehicles[1].GetName(), "Car")
	assert.Equal(t, vehicles[2].GetName(), "Tank")
}

func TestGetInterface(t *testing.T) {
	leviate, _ := singoton.Get[ILeviating]()

	assert.Equal(t, leviate.Leviate(), 10)
}

func TestMain(m *testing.M) {
	setUp()
	os.Exit(m.Run())
}

func setUp() {
	registerVehicle()
	registerShip()
}

func registerVehicle() {
	singoton.Register(singoton.RegisterOptions[VehicleNamer]{
		InitialValue: Motobike{
			Manufacturer: "Harley Davidson",
		},
	})

	singoton.Register(singoton.RegisterOptions[VehicleNamer]{
		InitFunction: func() VehicleNamer {
			return Car{
				Speed: 100,
			}
		},
	})

	singoton.Register(singoton.RegisterOptions[VehicleNamer]{
		InitialValue: Tank{
			CanonSize: 100,
		},
	})
}

func registerShip() {
	singoton.Register(singoton.RegisterOptions[ILeviating]{
		InitialValue: AirCraftCarier{
			NumberOfPlanes: 10,
		},
	})
}
