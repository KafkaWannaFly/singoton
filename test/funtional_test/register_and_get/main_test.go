package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/KafkaWannaFly/singoton"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func (s *TestSuite) TestGet() {
	dummy0, _ := singoton.Get[DummyStruct0]()
	s.Equal(dummy0.DummyInt, 000)
	s.Equal(dummy0.DummyStr, "000")

	dummy1, _ := singoton.Get[DummyStruct1]()
	s.Equal(dummy1.DummyInt, 111)
	s.Equal(dummy1.DummyStr, "111")
}

func (s *TestSuite) TestGetNotExisted() {
	_, err := singoton.Get[string]()
	s.NotNil(err)
}

func (s *TestSuite) TestOverwriteAndGet() {
	dummy2, _ := singoton.Get[DummyStruct2]()
	s.Assert().Equal(dummy2.DummyInt, 0)
	s.Assert().Equal(dummy2.DummyStr, "")

	singoton.Register(DummyStruct2{
		DummyInt: 999,
		DummyStr: "999",
	})

	newDummy2, _ := singoton.Get[DummyStruct2]()
	s.Assert().Equal(newDummy2.DummyInt, 999)
	s.Assert().Equal(newDummy2.DummyStr, "999")
}

func TestMain(t *testing.T) {
	setUp()
	suite.Run(t, new(TestSuite))
}

func setUp() {
	registerDummy()
}

func registerDummy() {
	singoton.Register(DummyStruct0{
		DummyInt: 000,
		DummyStr: "000",
	})
	singoton.Register(DummyStruct1{
		DummyInt: 111,
		DummyStr: "111",
	})
	singoton.Register(DummyStruct2{})
	singoton.Register(DummyStruct3{})
	singoton.Register(DummyStruct4{})
	singoton.Register(DummyStruct5{})
	singoton.Register(DummyStruct6{})

	container := singoton.GetDependencyContainer()
	InfoColor := "\033[1;34m%s\033[0m"
	log.Printf(InfoColor, fmt.Sprintf("Registered %d objects", len(*container)))
}
