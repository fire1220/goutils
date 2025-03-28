package test

import (
	"context"
	"fmt"
	"github.com/fire1220/goutils/parallel"
	"strings"
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
	// 	return param, errors.New("testing error")
	// }

	list, err := parallel.New().Exec(ctx, []parallel.Handle{
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

func TestParallel2(t *testing.T) {
	ctx := context.Background()
	var res1 string
	var res2 int
	_, err := parallel.New().Exec(ctx, []parallel.Handle{
		func(ctx context.Context, param interface{}) (interface{}, error) {
			p, _ := param.(string)
			res1 = "hello" + p
			return nil, nil
		},
		func(ctx context.Context, param interface{}) (interface{}, error) {
			p, _ := param.(int)
			res2 = 100 + p
			return nil, nil
		},
	}, "func1ParamYes", 12)
	fmt.Printf("%+v %+v\n%+v\n", res1, res2, err)
}

type Concat struct {
	Param [2]string
	Res   string
}

func (c *Concat) Exec() error {
	c.Res = strings.Join(c.Param[:], " ")
	return nil
}

type Add struct {
	Param [2]int
	Res   int
}

func (a *Add) Exec() error {
	a.Res = a.Param[0] + a.Param[1]
	return nil
}

func TestExecWithObj(t *testing.T) {
	ctx := context.Background()
	x := &Concat{Param: [2]string{"hello", "world"}}
	y := &Add{Param: [2]int{1, 2}}
	err := parallel.New().ExecObj(ctx, x, y)
	fmt.Printf("%#v\n", x)
	fmt.Printf("%#v\n", y)
	fmt.Printf("%+v\n", err)
}
