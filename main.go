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

func Get[T any]() *T {
	var obj T
	key := metadata.New[T](obj)
	unsafePointer := unsafe.Pointer(dependencyContainer[key])
	return (*T)(unsafePointer)
}

func addToContainer[T any](key metadata.Metadata, value *T) {
	unsafePointer := unsafe.Pointer(value)
	dependencyContainer[key] = (*any)(unsafePointer)
}

func GetContainer() map[metadata.Metadata]*any {
	return dependencyContainer
}
