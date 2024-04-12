package common

import (
	"reflect"
)

// SliceColumn 取出slice的key，返回key的t类型类别
func SliceColumn[T1, T2 any](keyType *T2, s []T1, key string) []T2 {
	if len(s) == 0 {
		return nil
	}
	list := make([]T2, 0, len(s))
	for k, val := range s {
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
					break
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

func SliceColumnMap[TSlice, TKey any, TVal comparable](mapType *map[TVal]TKey, s []TSlice, key string, valS ...string) map[TVal]TKey {
	if len(s) == 0 {
		return nil
	}
	val := ""
	if len(valS) > 0 {
		val = valS[0]
	}
	retMap := make(map[TVal]TKey, len(s))
	for k, v := range s {
		r := reflect.ValueOf(v)
		if r.Kind() == reflect.Ptr {
			r = r.Elem()
		}
		t := r.Type()
		if k == 0 {
			existsKey := false
			existsVal := false
			if val == "" {
				existsVal = true
			}
			for i := 0; i < r.NumField(); i++ {
				if key == t.Field(i).Name {
					existsKey = true
				}
				if val != "" && val == t.Field(i).Name {
					existsVal = true
				}
			}
			if !existsKey || !existsVal {
				return nil
			}
			mKey, kOk := r.FieldByName(key).Interface().(TVal)
			var mVal TKey
			vOk := false
			if val == "" {
				mVal, vOk = r.Interface().(TKey)
			} else {
				mVal, vOk = r.FieldByName(val).Interface().(TKey)
			}
			if kOk && vOk {
				retMap[mKey] = mVal
			}
		}
	}
	return retMap
}
