package polymorphism

type ILeviating interface {
	Leviate() int
}

type AirCraftCarier struct {
	NumberOfPlanes int
}

func (a AirCraftCarier) Leviate() int {
	return a.NumberOfPlanes
}

type SpaceShip struct {
	EnginePower int
}

func (s SpaceShip) Leviate() int {
	return s.EnginePower
}
