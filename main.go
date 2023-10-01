package singoton

import (
	"errors"
)

var interfaceImplementMap = make(map[Metadata]Metadata)
var dependencyContainerMap = make(map[Metadata]any)
var factoryContainerMap = make(map[Metadata]any)

type IFactory[T any] interface {
	New() T
}

func Register[T any](object T) {
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
	typeMetadata := createMetadataFromType[T]()
	factoryContainerMap[typeMetadata] = factory
}

func GetFromFactory[T any]() (T, error) {
	typeMetadata := createMetadataFromType[T]()
	factory, ok := factoryContainerMap[typeMetadata]
	if ok {
		return factory.(IFactory[T]).New(), nil
	} else {
		return zeroValueOf[T](), errors.New(typeMetadata.Name + " was not registered")
	}
}

func UnRegister[T any]() {
	typeMetadata := createMetadataFromType[T]()
	objectMetadata, ok := interfaceImplementMap[typeMetadata]
	if ok {
		delete(interfaceImplementMap, typeMetadata)
		delete(dependencyContainerMap, objectMetadata)
	} else {
		delete(dependencyContainerMap, typeMetadata)
	}
}

func IsRegistered[T any]() bool {
	typeMetadata := createMetadataFromType[T]()
	objectMetadata, hasInterface := interfaceImplementMap[typeMetadata]
	_, hasInstance := dependencyContainerMap[objectMetadata]

	return hasInterface || hasInstance
}

func Get[T any]() (T, error) {
	typeMetadata := createMetadataFromType[T]()
	objectMetadata, ok := interfaceImplementMap[typeMetadata]
	if ok {
		// T is an interface
		return dependencyContainerMap[objectMetadata].(T), nil
	}

	// T is not an interface
	object, ok := dependencyContainerMap[typeMetadata]
	if ok {
		return object.(T), nil
	} else {
		return zeroValueOf[T](), errors.New(typeMetadata.Name + " was not registered")
	}
}

func GetDependencyContainer() *map[Metadata]any {
	return &dependencyContainerMap
}

func GetInterfaceImplementMap() *map[Metadata]Metadata {
	return &interfaceImplementMap
}

func addToContainer[T any](key Metadata, value T) {
	dependencyContainerMap[key] = value
}
