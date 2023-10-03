package singoton

import (
	"errors"
)

var dependencyContainer = &_DependencyContainer{
	objectContainerMap:    make(map[Metadata]any),
	interfaceImplementMap: make(map[Metadata]Metadata),
	factoryContainerMap:   make(map[Metadata]any),
}

// IFactory is an interface for factory. You must implement this interface to use RegisterFactory.
//
// Type T can be struct or interface.
type IFactory[T any] interface {
	// New return an object of type T
	New() T
}

// Register an object of type T to dependency container
// Type T can be struct or interface
// Example:
//
//		type IVehicle interface {
//			Run()
//		}
//
//		type Car struct {}
//
//		func (c *Car) Run() {}
//
//		func main() {
//	    	// Register IVehicle interface, with Car implement, pointer of Car
//			singoton.Register[IVehicle](new(Car))
//			// Or you can just register Car, without interface
//			singoton.Register(Car{...})
//		}
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

// Get an object of type T from dependency container. Return error if T was not registered.
//
// Type T can be struct or interface. If T is an interface, it will return an implement of T. You must use the same type when register.
//
// Example:
//
//	func main() {
//		singoton.Register[IVehicle](new(Car))
//		vehicle, _ := singoton.Get[IVehicle]()
//		vehicle.Run()
//	}
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

// UnRegister an object of type T from dependency container.
//
// Type T can be struct or interface. Must be the same type when register.
//
// Example:
//
//	func main() {
//		singoton.Register[IVehicle](new(Car))
//		singoton.UnRegister[IVehicle]()
//	}
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

// IsRegistered check if an object of type T was registered to dependency container.
//
// Type T can be struct or interface. Must be the same type when register.
//
// Example:
//
//	func main() {
//		singoton.Register[IVehicle](new(Car))
//		singoton.IsRegistered[IVehicle]() // return true
//		singoton.IsRegistered[Car]() // return true
//		singoton.IsRegistered[Bus]() // return false
//	}
func IsRegistered[T any]() bool {
	typeMetadata := createMetadataFromType[T]()
	objectMetadata, hasInterface := dependencyContainer.getInterfaceImplement(typeMetadata)
	_, hasInstance := dependencyContainer.getObject(objectMetadata)

	return hasInterface || hasInstance
}

// RegisterFactory register a factory of type T to dependency container.
//
// Factory must implement IFactory[T] interface. Type T can be struct or interface.
//
// Example:
//
//	type IStella interface {
//		GetSize() int
//		GetTemperature() int
//	}
//
//	type Star struct {
//		size        int
//		temperature int
//	}
//
//	func (s Star) GetSize() int {
//		return s.size
//	}
//
//	func (s Star) GetTemperature() int {
//		return s.temperature
//	}
//
//	type StarFactory struct {
//	}
//
//	func (sf StarFactory) New() IStella {
//		return Star{
//			size:        rand.Int(),
//			temperature: rand.Int(),
//		}
//	}
//
//	func main() {
//		singoton.RegisterFactory[IStella](StarFactory{})
//	}
func RegisterFactory[T any](factory IFactory[T]) {
	typeMetadata := createMetadataFromType[T]()
	dependencyContainer.addFactory(typeMetadata, factory)
}

// GetFromFactory get an object of type T from factory. Return error if T was not registered.
//
// Type T can be struct or interface. You must use the same type when register.
//
// Example:
//
//	func main() {
//		singoton.RegisterFactory[IStella](StarFactory{})
//		star, _ := singoton.GetFromFactory[IStella]()
//		star.GetSize()
//		star.GetTemperature()
//	}
func GetFromFactory[T any]() (T, error) {
	typeMetadata := createMetadataFromType[T]()
	factory, ok := dependencyContainer.getFactory(typeMetadata)
	if ok {
		return factory.(IFactory[T]).New(), nil
	} else {
		return zeroValueOf[T](), errors.New(typeMetadata.Name + " was not registered")
	}
}

// UnRegisterFactory unregister a factory of type T from dependency container.
//
// Type T can be struct or interface. Must be the same type when register.
func UnRegisterFactory[T any]() {
	typeMetadata := createMetadataFromType[T]()
	dependencyContainer.removeFactory(typeMetadata)
}

// GetDependencyContainer return a map of all registered objects.
func GetDependencyContainer() *map[Metadata]any {
	return dependencyContainer.getObjectMap()
}

// GetInterfaceImplement return a map of all registered interface and it's implementation.
func GetInterfaceImplement() *map[Metadata]Metadata {
	return dependencyContainer.getInterfaceImplementMap()
}
