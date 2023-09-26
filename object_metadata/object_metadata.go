package objectmetadata

import (
	"reflect"
)

type ObjectMetadata struct {
	Name    string
	Package string
}

func New[T any](obj T) ObjectMetadata {
	t := reflect.TypeOf(obj)

	return ObjectMetadata{
		Name:    t.Name(),
		Package: t.PkgPath(),
	}
}
