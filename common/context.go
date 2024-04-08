package common

import (
	"context"
	"fmt"
	"reflect"
)

func ContextKeys(ctx context.Context) []string {
	keys := make([]string, 0)
PARENT:
	rCtx := reflect.ValueOf(ctx)
	vCtx := rCtx.Elem()
	tCtx := vCtx.Type()
	isParent := false
	if vCtx.Kind() != reflect.Struct {
		return keys
	}

	for i := 0; i < vCtx.NumField(); i++ {
		v := vCtx.Field(i)
		t := tCtx.Field(i)
		if t.Name == "key" {
			keys = append(keys, fmt.Sprintf("%v", v))
		}
		if t.Name == "Context" {
			ctx, isParent = v.Interface().(context.Context)
		}
	}
	if isParent {
		goto PARENT
	}
	return keys
}
