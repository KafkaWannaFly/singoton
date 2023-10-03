package polymorphism

type VehicleNamer interface {
	GetName() string
}

type Motorbike struct {
	Manufacturer string `json:"manufacturer"`
}

func (v Motorbike) GetName() string {
	return "Motorbike"
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
