package marshal

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"
)

type dateTime struct {
	t   time.Time
	tag reflect.StructTag
}

func (d dateTime) MarshalJSON() ([]byte, error) {
	t := d.t
	formatData := d.tag.Get("datetime")
	format, ok := strings.CutSuffix(formatData, "omitempty")
	format = strings.Trim(format, ",")
	if format == "" {
		format = time.DateTime
	}
	mapTime := map[string]string{
		time.DateTime: "0000-00-00 00:00:00",
		time.DateOnly: "0000-00-00",
		time.TimeOnly: "00:00:00",
	}
	if t.IsZero() {
		if ok {
			return []byte(`""`), nil
		}
		if v, ok := mapTime[format]; ok {
			return []byte(`"` + v + `"`), nil
		} else {
			return []byte(`""`), nil
		}
	}
	return []byte(`"` + t.Format(format) + `"`), nil
}

func Marshal(p any) ([]byte, error) {
	ref := reflect.ValueOf(p)
	typ := ref.Type()
	newField := make([]reflect.StructField, 0, ref.NumField())
	dateTimeReflectType := reflect.TypeOf(dateTime{})
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
		oldFieldType := typ.Field(i)
		if oldField.Type().String() != "time.Time" {
			newStruct.Field(i).Set(oldField)
			continue
		}
		if v, ok := oldField.Interface().(time.Time); ok {
			newStruct.Field(i).Set(reflect.ValueOf(dateTime{t: v, tag: oldFieldType.Tag}))
		}
	}
	return json.Marshal(newStruct.Interface())
}
