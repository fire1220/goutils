package marshaljson

import (
	"reflect"
)

type defaultStringT struct {
	tag reflect.StructField
}

func (d defaultStringT) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.tag.Tag.Get(tabDefaultString) + `"`), nil
}

func (d defaultStringT) typeConv(field reflect.Value, typ reflect.StructField) (reflect.Value, bool) {
	return reflect.ValueOf(defaultStringT{tag: typ}), true
}
