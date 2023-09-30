package test

type ILeviating interface {
	Leviate() int
}

type AirCraftCarier struct {
	NumberOfPlanes int
}

func (a AirCraftCarier) Leviate() int {
	return a.NumberOfPlanes
}
