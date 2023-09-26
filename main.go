package singoton

import (
	"unsafe"

	objectmetadata "github.com/KafkaWannaFly/singoton/object_metadata"
)

var dependencyContainer = make(map[objectmetadata.ObjectMetadata]*any)

type RegisterOptions[T any] struct {
	InitialValue *T
	InitFunction func() *T
}

func Register[T any](options RegisterOptions[T]) {
	if options.InitialValue != nil {
		addToContainer(objectmetadata.New(*options.InitialValue), options.InitialValue)
	} else if options.InitFunction != nil {
		initValue := options.InitFunction()
		addToContainer(objectmetadata.New(*initValue), initValue)
	} else {
		var obj T
		addToContainer(objectmetadata.New(obj), &obj)
	}
}

func Get[T any]() *T {
	var obj T
	key := objectmetadata.New[T](obj)
	unsafePointer := unsafe.Pointer(dependencyContainer[key])
	return (*T)(unsafePointer)
}

func addToContainer[T any](key objectmetadata.ObjectMetadata, value *T) {
	unsafePointer := unsafe.Pointer(value)
	dependencyContainer[key] = (*any)(unsafePointer)
}

func GetContainer() map[objectmetadata.ObjectMetadata]*any {
	return dependencyContainer
}
