package singoton

import (
	"unsafe"

	metadata "github.com/KafkaWannaFly/singoton/metadata"
)

var dependencyContainer = make(map[metadata.Metadata]*any)

type RegisterOptions[T any] struct {
	InitialValue *T
	InitFunction func() *T
}

type IFactory[T any] interface {
	InitObject() T
}

func Register[T any](options RegisterOptions[T]) {
	if options.InitialValue != nil {
		addToContainer(metadata.New(*options.InitialValue), options.InitialValue)
	} else if options.InitFunction != nil {
		initValue := options.InitFunction()
		addToContainer(metadata.New(*initValue), initValue)
	} else {
		var obj T
		addToContainer(metadata.New(obj), &obj)
	}
}

func RegisterFactory[T any](factory IFactory[T]) {
	panic("Not implemented")
}

func UnRegister[T any]() *T {
	panic("Not implemented")
}

func Get[T any]() *T {
	var obj T
	key := metadata.New[T](obj)
	unsafePointer := unsafe.Pointer(dependencyContainer[key])
	return (*T)(unsafePointer)
}

func GetContainer() map[metadata.Metadata]*any {
	return dependencyContainer
}

func addToContainer[T any](key metadata.Metadata, value *T) {
	unsafePointer := unsafe.Pointer(value)
	dependencyContainer[key] = (*any)(unsafePointer)
}
