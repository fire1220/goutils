package common

import (
	"reflect"
)

// SliceColumn 取出slice的key，返回key的t类型类别
func SliceColumn[T1, T2 any](arr []T1, key string, keyType *T2) []T2 {
	if len(arr) == 0 {
		return nil
	}
	list := make([]T2, 0, len(arr))
	for k, val := range arr {
		r := reflect.ValueOf(val)
		if r.Kind() == reflect.Ptr {
			r = r.Elem()
		}
		if k == 0 {
			if r.Kind() != reflect.Struct {
				return nil
			}
			existsKey := false
			ty := r.Type()
			for i := 0; i < r.NumField(); i++ {
				if ty.Field(i).Name == key {
					existsKey = true
				}
			}
			if !existsKey {
				return nil
			}
		}
		if v, ok := r.FieldByName(key).Interface().(T2); ok {
			list = append(list, v)
		}
	}
	return list
}
