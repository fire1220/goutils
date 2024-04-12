package common

import (
	"context"
	"fmt"
	"reflect"
)

// IsCancel 判断上下文是否关闭
func IsCancel(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

// ContextDuplicate 复值上下文的key和val到新的上下文
func ContextDuplicate(ctx context.Context) context.Context {
	newCtx := context.Background()
PARENT:
	oldCtx := ctx
	rCtx := reflect.ValueOf(ctx)
	vCtx := rCtx
	if rCtx.Kind() == reflect.Pointer {
		vCtx = rCtx.Elem()
	}
	tCtx := vCtx.Type()
	if vCtx.Kind() != reflect.Struct {
		return newCtx
	}
	var key string
	isParent := false
	for i := 0; i < vCtx.NumField(); i++ {
		v := vCtx.Field(i)
		t := tCtx.Field(i)
		if t.Name == "key" {
			key = fmt.Sprintf("%v", v)
			break
		}
		if t.Name == "Context" {
			ctx, isParent = v.Interface().(context.Context)
		}
	}
	if key != "" {
		val := oldCtx.Value(key)
		newCtx = context.WithValue(newCtx, key, val)
	}
	if isParent {
		goto PARENT
	}
	return newCtx
}

// ContextKeys 获取上下文的所有key
func ContextKeys(ctx context.Context) []string {
	keys := make([]string, 0)
PARENT:
	rCtx := reflect.ValueOf(ctx)
	vCtx := rCtx
	if rCtx.Kind() == reflect.Pointer {
		vCtx = rCtx.Elem()
	}
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
