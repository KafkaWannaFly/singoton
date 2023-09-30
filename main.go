package singoton

import (
	"errors"
)

var interfaceImplementMap = make(map[Metadata]Metadata)
var dependencyContainer = make(map[Metadata]any)

type RegisterOptions[T any] struct {
	InitialValue T
	InitFunction func() T
}

type IFactory[T any] interface {
	New() *T
}

func Register[T any](options RegisterOptions[T]) {
	var object T
	if options.InitFunction != nil {
		object = options.InitFunction()
	} else {
		object = options.InitialValue
	}

	typeMetadata := createMetadataFromType[T]()
	objectMetadata := createMetadataFromObject[T](object)
	if typeMetadata != objectMetadata {
		// This mean T is an interface, object is an implement of T
		// Need to map interface to implement
		interfaceImplementMap[typeMetadata] = objectMetadata
	}

	// Add object to container
	addToContainer(objectMetadata, object)
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
	typeMetadata := createMetadataFromType[T]()
	objectMetadata, ok := interfaceImplementMap[typeMetadata]
	if ok {
		// T is an interface
		return dependencyContainer[objectMetadata].(T), nil
	}

	// T is not an interface
	object, ok := dependencyContainer[typeMetadata]
	if ok {
		return object.(T), nil
	} else {
		return zeroValueOf[T](), errors.New(typeMetadata.Name + " was not registered")
	}
}

func GetDependencyContainer() *map[Metadata]any {
	return &dependencyContainer
}

func GetInterfaceImplementMap() *map[Metadata]Metadata {
	return &interfaceImplementMap
}

func addToContainer[T any](key Metadata, value T) {
	dependencyContainer[key] = value
}
