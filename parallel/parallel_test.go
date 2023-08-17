package parallel

import (
	"context"
	"fmt"
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

	list, err := New().Exec(ctx, []ParallelHandle{
		func1, // function 1
		func2, // function 2
		// func3,
	},
		"func1ParamYes", // function 1 params
		"func2ParamOk",  // function 2 params
		// "func2ParamNo",
	)
	fmt.Printf("%+v\n%+v\n", list, err)
}
