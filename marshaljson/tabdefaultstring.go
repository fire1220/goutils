package marshaljson

import (
	"reflect"
)

type tabDefaultString struct {
	tag reflect.StructField
}

func (d tabDefaultString) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.tag.Tag.Get(tabNameDefaultString) + `"`), nil
}

func (d tabDefaultString) typeConv(field reflect.Value, typ reflect.StructField) (reflect.Value, bool) {
	return reflect.ValueOf(tabDefaultString{tag: typ}), true
}
