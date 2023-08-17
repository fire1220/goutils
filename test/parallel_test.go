package test

import (
	"context"
	"fmt"
	"go-utils/parallel"
	"testing"
)

func TestParallel(t *testing.T) {
	ctx := context.Background()
	func1 := func(ctx context.Context, param interface{}) (interface{}, error) {
		return param, nil
	}

	func2 := func(ctx context.Context, param interface{}) (interface{}, error) {
		return param, nil
	}

	// func3 := func(ctx context.Context, param interface{}) (interface{}, error) {
	// 	return param, errors.New("测试错误")
	// }

	list, err := parallel.New().Exec(ctx, []parallel.ParallelHandle{
		func1,
		func2,
		// func3,
	},
		"yes",
		"ok",
		// "no",
	)
	fmt.Printf("%+v\n%+v\n", list, err)
}
