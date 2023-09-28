package singoton

import (
	"errors"

	metadata "github.com/KafkaWannaFly/singoton/metadata"
)

var dependencyContainer = make(map[metadata.Metadata]any)

type RegisterOptions[T any] struct {
	InitialValue T
	InitFunction func() T
}

type IFactory[T any] interface {
	InitObject() T
}

func Register[T any](options RegisterOptions[T]) {
	if options.InitFunction != nil {
		initValue := options.InitFunction()
		addToContainer(metadata.New(initValue), initValue)
	} else {
		addToContainer(metadata.New(options.InitialValue), options.InitialValue)
	}
}

func RegisterFactory[T any](factory IFactory[T]) {
	panic("Not implemented")
}

func UnRegister[T any]() T {
	panic("Not implemented")
}

func IsRegistered[T any]() bool {
	panic("Not implemented")
}

func Get[T any]() (T, error) {
	var obj T
	key := metadata.New[T](obj)
	value := dependencyContainer[key]
	if value == nil {
		return obj, errors.New("not registered")
	} else {
		return value.(T), nil
	}
}

func GetContainer() map[metadata.Metadata]any {
	return dependencyContainer
}

func addToContainer[T any](key metadata.Metadata, value T) {
	dependencyContainer[key] = value
}
