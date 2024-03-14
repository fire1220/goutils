package marshal

import (
	"encoding/json"
	"reflect"
	"time"
)

type DateTime time.Time

func (d DateTime) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	if t.IsZero() {
		return []byte(`"0000-00-00 00:00:00"`), nil
	}
	return []byte(`"` + t.Format(`2006-01-02 15:04:05`) + `"`), nil
}

func Marshal(p any) ([]byte, error) {
	ref := reflect.ValueOf(p)
	typ := ref.Type()
	newField := make([]reflect.StructField, 0, ref.NumField())
	dateTimeReflectType := reflect.TypeOf(DateTime{})
	for i := 0; i < ref.NumField(); i++ {
		field := typ.Field(i)
		fieldType := field.Type
		if field.Type.String() == "time.Time" {
			fieldType = dateTimeReflectType
		}
		newField = append(newField, reflect.StructField{
			Name: field.Name,
			Type: fieldType,
			Tag:  field.Tag,
		})
	}
	newStruct := reflect.New(reflect.StructOf(newField)).Elem()
	for i := 0; i < newStruct.NumField(); i++ {
		oldField := ref.Field(i)
		if oldField.Type().String() != "time.Time" {
			newStruct.Field(i).Set(oldField)
			continue
		}
		if v, ok := oldField.Interface().(time.Time); ok {
			newStruct.Field(i).Set(reflect.ValueOf(DateTime(v)))
		}
	}
	return json.Marshal(newStruct.Interface())
}
