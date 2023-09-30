package test

type VehicleNamer interface {
	GetName() string
}

type Motobike struct {
	Manufacturer string `json:"manufacturer"`
}

func (v Motobike) GetName() string {
	return "Motobike"
}

type Car struct {
	Speed int `json:"speed"`
}

func (c Car) GetName() string {
	return "Car"
}

type Tank struct {
	CanonSize int `json:"canon_size"`
}

func (t Tank) GetName() string {
	return "Tank"
}