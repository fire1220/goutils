package marshaljson

import (
	"reflect"
	"time"
)

type tabDateTime struct {
	t   time.Time
	tag reflect.StructTag
}

func (d tabDateTime) MarshalJSON() ([]byte, error) {
	t := d.t
	format := d.tag.Get(tabNameDateTime)
	mapTime := map[string]string{
		time.DateTime: "0000-00-00 00:00:00",
		time.DateOnly: "0000-00-00",
		time.TimeOnly: "00:00:00",
	}
	if t.IsZero() {
		if v, ok := mapTime[format]; ok {
			return []byte(`"` + v + `"`), nil
		} else {
			return []byte(`""`), nil
		}
	}
	return []byte(`"` + t.Format(format) + `"`), nil
}

func (d tabDateTime) typeConv(field reflect.Value, typ reflect.StructField) (reflect.Value, bool) {
	if v, ok := field.Interface().(time.Time); ok {
		return reflect.ValueOf(tabDateTime{t: v, tag: typ.Tag}), true
	}
	return reflect.Value{}, false
}
