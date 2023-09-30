package singoton

import (
	"errors"
	"reflect"

	metadata "github.com/KafkaWannaFly/singoton/metadata"
)

var interfaceImplementMap = make(map[metadata.Metadata]metadata.Metadata)
var dependencyContainer = make(map[metadata.Metadata]any)

type RegisterOptions[T any] struct {
	InitialValue T
	InitFunction func() T
}

type IFactory[T any] interface {
	New() T
}

func Register[T any](options RegisterOptions[T]) {
	var objectMetadata metadata.Metadata
	if options.InitFunction != nil {
		initValue := options.InitFunction()
		objectMetadata = metadata.New(initValue)
		addToContainer(objectMetadata, initValue)
	} else {
		objectMetadata = metadata.New(options.InitialValue)
		addToContainer(objectMetadata, options.InitialValue)
	}

	// Type T is an interface, we need to save data about which type implements it
	fn := func(T) {}
	interfaceType := reflect.TypeOf(fn).In(0)
	if interfaceType.Kind() == reflect.Interface {
		interfaceMetadata := metadata.Metadata{
			Name:    interfaceType.Name(),
			Package: interfaceType.PkgPath(),
		}
		interfaceImplementMap[interfaceMetadata] = objectMetadata
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
	metadata := metadata.New[T](obj)
	dependencyObj := dependencyContainer[metadata]
	if dependencyObj != nil {
		return dependencyObj.(T), nil
	}

	// Not found object, maybe it is an interface
	implementMetadata, isFound := interfaceImplementMap[metadata]
	if isFound {
		return dependencyContainer[implementMetadata].(T), nil
	}

	return obj, errors.New("not registered")
}

func GetContainer() map[metadata.Metadata]any {
	return dependencyContainer
}

func addToContainer[T any](key metadata.Metadata, value T) {
	dependencyContainer[key] = value
}
