package metadata

import (
	"reflect"
)

type Metadata struct {
	Name    string
	Package string
}

func New[T any](obj T) Metadata {
	t := reflect.TypeOf(obj)

	return Metadata{
		Name:    t.Name(),
		Package: t.PkgPath(),
	}
}
