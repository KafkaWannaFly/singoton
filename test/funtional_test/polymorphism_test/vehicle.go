package test

type VehicleNamer interface {
	GetName() string
}

type Motobike struct {
	Manufacturer string
}

func (v Motobike) GetName() string {
	return "Motobike"
}

type Car struct {
	Speed int
}

func (c Car) GetName() string {
	return "Car"
}

type Tank struct {
	CanonSize int
}

func (t Tank) GetName() string {
	return "Tank"
}
