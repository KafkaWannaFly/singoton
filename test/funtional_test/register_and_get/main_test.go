package test

import (
	"log"
	"os"
	"testing"

	"github.com/KafkaWannaFly/singoton"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	dummy0, _ := singoton.Get[DummyStruct0]()
	assert.Equal(t, dummy0.DummyInt, 000)
	assert.Equal(t, dummy0.DummyStr, "000")

	dummy1, _ := singoton.Get[DummyStruct1]()
	assert.Equal(t, dummy1.DummyInt, 111)
	assert.Equal(t, dummy1.DummyStr, "111")
}

func TestGetNotExisted(t *testing.T) {
	_, err := singoton.Get[string]()
	assert.NotNil(t, err)
}

func TestOverwriteAndGet(t *testing.T) {
	dummy2, _ := singoton.Get[DummyStruct2]()
	assert.Equal(t, dummy2.DummyInt, 0)
	assert.Equal(t, dummy2.DummyStr, "")

	singoton.Register(singoton.RegisterOptions[DummyStruct2]{
		InitialValue: DummyStruct2{
			DummyInt: 999,
			DummyStr: "999",
		},
	})

	newDummy2, _ := singoton.Get[DummyStruct2]()
	assert.Equal(t, newDummy2.DummyInt, 999)
	assert.Equal(t, newDummy2.DummyStr, "999")
}

func TestMain(m *testing.M) {
	setUp()
	os.Exit(m.Run())
}

func setUp() {
	registerDummy()
}

func registerDummy() {
	singoton.Register(singoton.RegisterOptions[DummyStruct0]{
		InitialValue: DummyStruct0{
			DummyInt: 000,
			DummyStr: "000",
		},
	})
	singoton.Register(singoton.RegisterOptions[DummyStruct1]{
		InitFunction: func() DummyStruct1 {
			return DummyStruct1{
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

	container := singoton.GetContainer()
	log.Println("Registering", len(container), "objects")
}