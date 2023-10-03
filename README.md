# Singoton

> How do you call an unmarried, 1000 kg man? A singoton.

Centralizing object instances in a single place and covering them with an abstract layer. Also support generics.

## Guide

### Register Objects and Use Them
Step 1: Design struct and interface.
```go
// Path: example/crud.go
package main

// ICrud is an interface for CRUD operations.
type ICrud interface {
	Create(item any)
	Read(id int) any
	Update(item any)
	Delete(id int)
}
```

```go
// Path: example/user.go
package main

// User is a struct for user data.
type User struct {
    ID   int
    Name string
}

// UserCrud is a struct for user CRUD operations.
type UserCrud struct {
    users []User
}

// Create creates a new user.
func (crud *UserCrud) Create(item any) {
    user := item.(User)
    crud.users = append(crud.users, user)
}

// Read reads a user by id.
func (crud *UserCrud) Read(id int) any {
    for _, user := range crud.users {
        if user.ID == id {
            return user
        }
    }
    return nil
}

// Update updates a user.
func (crud *UserCrud) Update(item any) {
    user := item.(User)
    for i, u := range crud.users {
        if u.ID == user.ID {
            crud.users[i] = user
            return
        }
    }
}

// Delete deletes a user by id.
func (crud *UserCrud) Delete(id int) {
    for i, user := range crud.users {
        if user.ID == id {
            crud.users = append(crud.users[:i], crud.users[i+1:]...)
            return
        }
    }
}
```

Step 2: Register it with singoton.
```go
// Path: example/main.go
package main

import (
	"github.com/KafkaWannaFly/singoton"
)

func main() {
	// Notice that the interface is registered with the struct implementing it.
	singoton.Register[ICrud](&UserCrud{})
}

```

Step 3: Use it.

```go
// Path: example/main.go
package main

import (
	"fmt"
	"github.com/KafkaWannaFly/singoton"
)

func useCrud() {
	// If can't find the registered struct, it will return an error.
	icrud, err := singoton.Get[ICrud]()
	if err != nil {
		panic(err)
	}

	user := icrud.Read(1).(User) // Read return any, so we need to cast it to User.
	fmt.Println(user)

	// Similarly, we can use other method of ICrud interface. Create, Update, Delete.
}

```


The example above is a standard way to abstract CRUD operations. Required structs and interfaces. But for people who don't like to write interfaces, you can just register the struct right away.

```go
// Not have to be a pointer. It also can be a value.
singoton.Register[&UserCrud{}]
```
And use it.
```go
// Remember your data type. We register a pointer, so we get a pointer.
userCrud, _ := singoton.Get[*UserCrud]()
```

You can also overwrite or remove the object from dependency container.

```go
package main

import "github.com/KafkaWannaFly/singoton"

func Overwrite() {
	// Overwrite the *UserCrud with a new one. Having 3 users inside.
	singoton.Register(&UserCrud{
		users: []User{
			{ID: 1, Name: "John"},
			{ID: 2, Name: "Jane"},
			{ID: 3, Name: "Jack"},
		},
	})
}

func Remove() {
	// Remove the *UserCrud from dependency container.
	// Notice that data type must be the same as the one you registered.
	singoton.UnRegister[*UserCrud]()
}

```

### Register Object Factory

Sometimes you need to create a new object every time you get it. You can register a factory function to do that.

Step 1: Implement your factory.

```go
// Package factory 
// StarFactory implements IFactory interface.
// It will create a new IStella every time you get it.
// Of course, it doesn't have to be an interface. It can be any data type.
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

```

Step 2: Pretty much the same as above.

```go
package main

import (
	"fmt"
	"github.com/KafkaWannaFly/singoton"
)

func registerFactory() {
	singoton.RegisterFactory[IStella](StarFactory{})
}

func getObjectFromFactory() {
	// Notice the data type. It's IStella.
	stella1, err1 := singoton.GetFromFactory[IStella]()
	if err1 != nil {
		panic(err1)
	}

	fmt.Println(stella1.GetSize())
	fmt.Println(stella1.GetTemperature())
}

```