package pointer

import (
	"github.com/KafkaWannaFly/singoton"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_get(t *testing.T) {
	userCrud, err := singoton.Get[*UserCrud]()
	assert.Nil(t, err)
	assert.NotNil(t, userCrud)

	users := userCrud.GetUsers()
	assert.NotEqual(t, 0, len(*users))
}

func Test_create(t *testing.T) {
	icrud, err := singoton.Get[ICrud]()

	assert.Nil(t, err)
	assert.NotNil(t, icrud)

	users := icrud.(*UserCrud).GetUsers()
	assert.Equal(t, 3, len(*users))

	icrud.Create(User{ID: 4, Name: "Jill"})
	assert.Equal(t, 4, len(*users))
}

func Test_update(t *testing.T) {
	icrud, err := singoton.Get[ICrud]()
	assert.Nil(t, err)
	assert.NotNil(t, icrud)

	icrud.Update(User{ID: 1, Name: "John Wick"})
	user := icrud.Read(1).(User)
	assert.Equal(t, "John Wick", user.Name)
	assert.NotEqual(t, "John", user.Name)
}

func Test_delete(t *testing.T) {
	icrud, err := singoton.Get[ICrud]()
	assert.Nil(t, err)
	assert.NotNil(t, icrud)

	users := icrud.(*UserCrud).GetUsers()
	currentLen := len(*users)

	icrud.Delete(1)
	assert.Equal(t, currentLen-1, len(*users))

	user1 := icrud.Read(1)
	assert.Nil(t, user1)
}

func Test_unregister(t *testing.T) {
	singoton.UnRegister[ICrud]()
	icrud, err := singoton.Get[ICrud]()

	assert.NotNil(t, err)
	assert.Nil(t, icrud)
}

func TestMain(m *testing.M) {
	setUp()
	os.Exit(m.Run())
}

func setUp() {
	singoton.Register[ICrud](&UserCrud{})

	icrud, _ := singoton.Get[ICrud]()
	icrud.Create(User{ID: 1, Name: "John"})
	icrud.Create(User{ID: 2, Name: "Jane"})
	icrud.Create(User{ID: 3, Name: "Jack"})
}
