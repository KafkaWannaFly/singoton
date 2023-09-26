package test

import (
	"testing"

	"github.com/KafkaWannaFly/singoton"
	"github.com/stretchr/testify/assert"
)

func TestRegisterAndGet(t *testing.T) {
	singoton.Register(singoton.RegisterOptions[DummyStruct0]{
		InitialValue: &DummyStruct0{
			DummyInt: 111,
			DummyStr: "222",
		},
	})
	singoton.Register(singoton.RegisterOptions[DummyStruct1]{})
	singoton.Register(singoton.RegisterOptions[DummyStruct2]{})
	singoton.Register(singoton.RegisterOptions[DummyStruct3]{})
	singoton.Register(singoton.RegisterOptions[DummyStruct4]{})
	singoton.Register(singoton.RegisterOptions[DummyStruct5]{})
	singoton.Register(singoton.RegisterOptions[DummyStruct6]{})

	dependencyContainer := singoton.GetContainer()
	size := len(dependencyContainer)
	assert.Equal(t, 7, size)

	dummy := singoton.Get[DummyStruct0]()
	assert.Equal(t, dummy.DummyInt, 111)
	assert.Equal(t, dummy.DummyStr, "222")
}
