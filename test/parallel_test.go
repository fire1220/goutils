package test

import (
	"context"
	"fmt"
	utils "go-utils"
	"testing"
)

func TestParallel(t *testing.T) {
	// ctxBase := context.Background()
	// ctx := context.WithValue(ctxBase, wmutil.ParallelDebugName, true)
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

	list, err := utils.NewParallel().Exec(ctx, []utils.ParallelHandle{
		func1,
		func2,
		// func3,
	},
		"yes1",
		"ok2",
		// "no3",
	)
	fmt.Printf("%+v\n%+v\n", list, err)
}
