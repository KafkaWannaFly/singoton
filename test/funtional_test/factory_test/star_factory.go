package factory

import "math/rand"

type Star struct {
	size        int
	temperature int
}

type IStella interface {
	GetSize() int
	GetTemperature() int
}

func (s Star) GetSize() int {
	return s.size
}

func (s Star) GetTemperature() int {
	return s.temperature
}

type StarFactory struct {
}

func (sf StarFactory) New() IStella {
	return Star{
		size:        rand.Int(),
		temperature: rand.Int(),
	}
}
