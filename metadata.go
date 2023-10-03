package singoton

import (
	"encoding/json"
	"reflect"
)

// Metadata is a struct to store type name and package name of a type
type Metadata struct {
	Name    string `json:"name"`
	Package string `json:"package"`
}

func createMetadataFromType[T any]() Metadata {
	fn := func(T) {}
	interfaceType := reflect.TypeOf(fn).In(0)

	return Metadata{
		Name:    interfaceType.Name(),
		Package: interfaceType.PkgPath(),
	}
}

func createMetadataFromObject[T any](obj T) Metadata {
	typeOfObj := reflect.TypeOf(obj)
	return Metadata{
		Name:    typeOfObj.Name(),
		Package: typeOfObj.PkgPath(),
	}
}

// ToString Return JSON string of Metadata
func (metadata Metadata) ToString() string {
	bytes, _ := json.MarshalIndent(metadata, "", "  ")
	return string(bytes)
}
