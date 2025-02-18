package marshaljson

import (
	"errors"
	"reflect"
	"strconv"
)

func stringToInt64(str string) ([]byte, error) {
	_, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return nil, errors.New("tag default value is not int")
	}
	return []byte(str), nil
}

func stringToUint64(str string) ([]byte, error) {
	_, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return nil, errors.New("tag default value is not int")
	}
	return []byte(str), nil
}

func stringToBool(str string) ([]byte, error) {
	if str != "true" && str != "false" {
		return nil, errors.New("tag default value is not bool")
	}
	return []byte(str), nil
}

func stringToFloat64(str string) ([]byte, error) {
	_, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return nil, errors.New("tag default value is not float")
	}
	return []byte(str), nil
}

func stringToArray(str string) ([]byte, error) {
	if str != "[]" && str != "{}" {
		return nil, errors.New("tag default value is not array")
	}
	return []byte(str), nil
}

func stringToObj(str string) ([]byte, error) {
	if str == "{}" {
		return []byte(str), nil
	}
	return []byte(`"` + str + `"`), nil
}

func stringToPrt(str string) ([]byte, error) {
	if str == "{}" || str == "[]" {
		return []byte(str), nil
	}
	return []byte(`"` + str + `"`), nil
}

var kindBoolMap = map[reflect.Kind]func(string) ([]byte, error){
	reflect.Bool: stringToBool,
}

var kindIntMap = map[reflect.Kind]func(string) ([]byte, error){
	reflect.Int:   stringToInt64,
	reflect.Int8:  stringToInt64,
	reflect.Int16: stringToInt64,
	reflect.Int32: stringToInt64,
	reflect.Int64: stringToInt64,
	// reflect.Uint:   stringToInt64,
	// reflect.Uint8:  stringToInt64,
	// reflect.Uint16: stringToInt64,
	// reflect.Uint32: stringToInt64,
	// reflect.Uint64: stringToInt64,
}

var kindUintMap = map[reflect.Kind]func(string) ([]byte, error){
	reflect.Uint:   stringToUint64,
	reflect.Uint8:  stringToUint64,
	reflect.Uint16: stringToUint64,
	reflect.Uint32: stringToUint64,
	reflect.Uint64: stringToUint64,
}

var kindFloatMap = map[reflect.Kind]func(string) ([]byte, error){
	reflect.Float32: stringToFloat64,
	reflect.Float64: stringToFloat64,
}

var kindArrayMap = map[reflect.Kind]func(string) ([]byte, error){
	reflect.Array: stringToArray,
	reflect.Slice: stringToArray,
}

var kindObjMap = map[reflect.Kind]func(string) ([]byte, error){
	reflect.Struct: stringToObj,
	reflect.Map:    stringToObj,
}

var kindPtrMap = map[reflect.Kind]func(string) ([]byte, error){
	reflect.Pointer: stringToPrt,
}

var kindSlice = []map[reflect.Kind]func(string) ([]byte, error){
	kindBoolMap,
	kindIntMap,
	kindUintMap,
	kindFloatMap,
	kindArrayMap,
	kindObjMap,
	kindPtrMap,
}

type tabDefault struct {
	tag reflect.StructField
}

func (d tabDefault) MarshalJSON() ([]byte, error) {
	format := d.tag.Tag.Get(tabNameDefault)
	for _, m := range kindSlice {
		if fn, ok := m[d.tag.Type.Kind()]; ok {
			return fn(format)
		}
	}
	return []byte(`"` + format + `"`), nil
}

func (d tabDefault) typeConv(field reflect.Value, typ reflect.StructField) (reflect.Value, bool) {
	return reflect.ValueOf(tabDefault{tag: typ}), true
}
