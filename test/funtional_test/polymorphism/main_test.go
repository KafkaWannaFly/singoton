package polymorphism

import (
	"os"
	"testing"

	"github.com/KafkaWannaFly/singoton"
	"github.com/stretchr/testify/assert"
)

func Test_get_different_implement_of_an_interface(t *testing.T) {
	var vehicles []VehicleNamer

	motobike, _ := singoton.Get[Motobike]()
	vehicles = append(vehicles, motobike)

	car, _ := singoton.Get[Car]()
	vehicles = append(vehicles, car)

	tank, _ := singoton.Get[Tank]()
	vehicles = append(vehicles, tank)

	assert.Equal(t, vehicles[0].GetName(), "Motobike", "Motobike name should be Motobike")
	assert.Equal(t, vehicles[0].(Motobike).Manufacturer, "Harley Davidson", "Motobike manufacturer should be Harley Davidson")

	assert.Equal(t, vehicles[1].GetName(), "Car", "Car name should be Car")
	assert.Equal(t, vehicles[1].(Car).Speed, 100, "Car speed should be 100")

	assert.Equal(t, vehicles[2].GetName(), "Tank", "Tank name should be Tank")
	assert.Equal(t, vehicles[2].(Tank).CanonSize, 100, "Tank canon size should be 100")
}

func Test_get_interface(t *testing.T) {
	item, _ := singoton.Get[ILeviating]()

	assert.Equal(t, item.Leviate(), 10, "AirCraftCarier should leviate 10")

	interfaceImplementMap := singoton.GetInterfaceImplementMap()
	assert.NotEmpty(t, interfaceImplementMap, "interfaceImplementMap should not be empty")
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
	singoton.Register(Motobike{
		Manufacturer: "Harley Davidson",
	})

	singoton.Register(Car{
		Speed: 100,
	})

	singoton.Register(Tank{
		CanonSize: 100,
	})
}

func registerShip() {
	singoton.Register[ILeviating](AirCraftCarier{
		NumberOfPlanes: 10,
	})
}
