package singoton

import (
	"encoding/json"
	"reflect"
)

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

/**
 * Return JSON string of Metadata
 */
func (metadata Metadata) ToString() string {
	bytes, _ := json.MarshalIndent(metadata, "", "  ")
	return string(bytes)
}
