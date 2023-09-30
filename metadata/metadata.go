package metadata

import (
	"reflect"
)

type Metadata struct {
	Name    string
	Package string
}

func New[T any]() Metadata {
	fn := func(T) {}
	interfaceType := reflect.TypeOf(fn).In(0)

	return Metadata{
		Name:    interfaceType.Name(),
		Package: interfaceType.PkgPath(),
	}
}
