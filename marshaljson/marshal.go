package marshaljson

import (
	"encoding/json"
	"errors"
	"reflect"
)

func verifyField(fieldType reflect.StructField, fieldVal reflect.Value, tabName string) (tabT, bool) {
	tm, ok := tabMap[tabName]
	if !ok {
		return tm, false
	}
	if _, ok := fieldType.Tag.Lookup(tabName); !ok {
		return tm, false
	}
	if tm.restrain != "" && fieldType.Type.String() != tm.restrain {
		return tm, false
	}
	if tabName == tabNameDateTime {
		return tm, true
	}
	if tabName == tabNameDefault && fieldVal.IsZero() {
		return tm, true
	}
	if tabName == tabNameDefaultString && fieldVal.IsZero() {
		return tm, true
	}
	return tm, false
}

// MarshalFormat
// 参数不能是指针类型，上层函数，必须是值接受
// 如果有结构体嵌套的情况，需要把每个结构图都实现 MarshalJSON 方法，
// 否则会把子集的结构体的 MarshalJSON继承到父级里，导致结构图替换时候缺少父级字段
func MarshalFormat(p any) ([]byte, error) {
	ref := reflect.ValueOf(p)
	if ref.Kind() == reflect.Pointer {
		return nil, errors.New("parameter must be a structure")
	}
	typ := ref.Type()
	newField := make([]reflect.StructField, 0, ref.NumField())
	for i := 0; i < ref.NumField(); i++ {
		field := typ.Field(i)
		fieldType := field.Type
		for _, tabName := range tabList {
			tm, ok := verifyField(field, ref.Field(i), tabName)
			if !ok {
				continue
			}
			fieldType = tm.refTypOf
			break
		}
		newField = append(newField, reflect.StructField{
			Name: field.Name,
			Type: fieldType,
			Tag:  field.Tag,
		})
	}

	newStruct := reflect.New(reflect.StructOf(newField)).Elem()
	for i := 0; i < newStruct.NumField(); i++ {
		oldFieldVal := ref.Field(i)
		oldFileType := typ.Field(i)
		var newFieldVal reflect.Value
		newFieldVal = oldFieldVal
		for _, tabName := range tabList {
			tm, ok := verifyField(oldFileType, oldFieldVal, tabName)
			if !ok {
				continue
			}
			if tm.fun == nil {
				continue
			}
			newVal, ok := tm.fun.typeConv(oldFieldVal, oldFileType)
			if ok {
				newFieldVal = newVal
			}
			break
		}
		newStruct.Field(i).Set(newFieldVal)
	}
	return json.Marshal(newStruct.Interface())
}
