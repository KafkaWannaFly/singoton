package factory

import (
	"github.com/KafkaWannaFly/singoton"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_get_from_factory(t *testing.T) {
	star1, err1 := singoton.GetFromFactory[IStella]()
	assert.NotZero(t, star1, "star1 should not be zero")
	assert.Nil(t, err1, "err1 should be nil")

	star2, _ := singoton.GetFromFactory[IStella]()
	assert.NotZero(t, star2, "star2 should not be zero")
	assert.Nil(t, err1, "err2 should be nil")

	assert.NotEqual(t, star1.GetSize(), star2.GetSize(), "star1 and star2 should not be equal")
	assert.NotEqual(t, star1.GetTemperature(), star2.GetTemperature(), "star1 and star2 should not be equal")
}

func Test_get_nothing_from_factory(t *testing.T) {
	_, err := singoton.GetFromFactory[Star]()
	assert.NotNil(t, err, "err should not be nil")
}

func TestMain(m *testing.M) {
	setUp()
	os.Exit(m.Run())
}

func setUp() {
	singoton.RegisterFactory[IStella](StarFactory{})
}
