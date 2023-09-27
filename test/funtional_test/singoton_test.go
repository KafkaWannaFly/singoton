package test

import (
	"testing"

	"github.com/KafkaWannaFly/singoton"
	"github.com/stretchr/testify/assert"
)

func TestRegisterAndGet(t *testing.T) {
	singoton.Register(singoton.RegisterOptions[DummyStruct0]{
		InitialValue: &DummyStruct0{
			DummyInt: 000,
			DummyStr: "000",
		},
	})
	singoton.Register(singoton.RegisterOptions[DummyStruct1]{
		InitFunction: func() *DummyStruct1 {
			return &DummyStruct1{
				DummyInt: 111,
				DummyStr: "111",
			}
		},
	})
	singoton.Register(singoton.RegisterOptions[DummyStruct2]{})
	singoton.Register(singoton.RegisterOptions[DummyStruct3]{})
	singoton.Register(singoton.RegisterOptions[DummyStruct4]{})
	singoton.Register(singoton.RegisterOptions[DummyStruct5]{})
	singoton.Register(singoton.RegisterOptions[DummyStruct6]{})

	dependencyContainer := singoton.GetContainer()
	size := len(dependencyContainer)
	assert.Equal(t, 7, size)

	dummy0 := singoton.Get[DummyStruct0]()
	assert.Equal(t, dummy0.DummyInt, 000)
	assert.Equal(t, dummy0.DummyStr, "000")

	dummy1 := singoton.Get[DummyStruct1]()
	assert.Equal(t, dummy1.DummyInt, 111)
	assert.Equal(t, dummy1.DummyStr, "111")
}
