package singoton

import (
	"errors"
)

var dependencyContainer = &_DependencyContainer{
	objectContainerMap:    make(map[Metadata]any),
	interfaceImplementMap: make(map[Metadata]Metadata),
	factoryContainerMap:   make(map[Metadata]any),
}

type IFactory[T any] interface {
	New() T
}

func Register[T any](object T) {
	typeMetadata := createMetadataFromType[T]()
	objectMetadata := createMetadataFromObject[T](object)
	if typeMetadata != objectMetadata {
		// This mean T is an interface, object is an implement of T
		// Need to map interface to implement
		dependencyContainer.addInterfaceImplement(typeMetadata, objectMetadata)
	}

	dependencyContainer.addObject(objectMetadata, object)
}

func Get[T any]() (T, error) {
	typeMetadata := createMetadataFromType[T]()
	objectMetadata, ok := dependencyContainer.getInterfaceImplement(typeMetadata)
	if ok {
		// T is an interface
		object, _ := dependencyContainer.getObject(objectMetadata)
		return object.(T), nil
	}

	// T is not an interface
	object, ok := dependencyContainer.getObject(typeMetadata)
	if ok {
		return object.(T), nil
	} else {
		return zeroValueOf[T](), errors.New(typeMetadata.Name + " was not registered")
	}
}

func UnRegister[T any]() {
	typeMetadata := createMetadataFromType[T]()
	objectMetadata, ok := dependencyContainer.getInterfaceImplement(typeMetadata)
	if ok {
		dependencyContainer.removeInterfaceImplement(typeMetadata)
		dependencyContainer.removeObject(objectMetadata)
	} else {
		dependencyContainer.removeObject(typeMetadata)
	}
}

func IsRegistered[T any]() bool {
	typeMetadata := createMetadataFromType[T]()
	objectMetadata, hasInterface := dependencyContainer.getInterfaceImplement(typeMetadata)
	_, hasInstance := dependencyContainer.getObject(objectMetadata)

	return hasInterface || hasInstance
}

func RegisterFactory[T any](factory IFactory[T]) {
	typeMetadata := createMetadataFromType[T]()
	dependencyContainer.addFactory(typeMetadata, factory)
}

func GetFromFactory[T any]() (T, error) {
	typeMetadata := createMetadataFromType[T]()
	factory, ok := dependencyContainer.getFactory(typeMetadata)
	if ok {
		return factory.(IFactory[T]).New(), nil
	} else {
		return zeroValueOf[T](), errors.New(typeMetadata.Name + " was not registered")
	}
}

func UnRegisterFactory[T any]() {
	typeMetadata := createMetadataFromType[T]()
	dependencyContainer.removeFactory(typeMetadata)
}

func GetDependencyContainer() *map[Metadata]any {
	return dependencyContainer.getObjectMap()
}

func GetInterfaceImplementMap() *map[Metadata]Metadata {
	return dependencyContainer.getInterfaceImplementMap()
}
