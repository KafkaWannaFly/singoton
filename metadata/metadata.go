package metadata

import (
	"reflect"
)

type Metadata struct {
	Name    string
	Package string
}

func New[T any](obj T) Metadata {
	objType := reflect.TypeOf(obj)

	if objType != nil {
		return Metadata{
			Name:    objType.Name(),
			Package: objType.PkgPath(),
		}
	}

	// T is interface
	fn := func(T) {}
	interfaceType := reflect.TypeOf(fn).In(0)

	return Metadata{
		Name:    interfaceType.Name(),
		Package: interfaceType.PkgPath(),
	}
}
