package test

import (
	"context"
	"fmt"
	"github.com/fire1220/goutils/parallel"
	"strconv"
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

type Like struct {
	Param int
	Res   string
}

func (u *Like) Exec() error {
	u.Res = strconv.Itoa(u.Param*100) + " 篮球"
	return nil
}

type Age struct {
	Param int
	Res   int
}

func (u *Age) Exec() error {
	u.Res = u.Param + 10
	return nil
}

func TestExecWithObj(t *testing.T) {
	ctx := context.Background()
	like := &Like{Param: 1}
	age := &Age{Param: 2}
	err := parallel.New().ExecWithObj(ctx, []parallel.Para{like, age})
	fmt.Printf("%#v\n", like)
	fmt.Printf("%#v\n", age)
	fmt.Printf("%+v\n", err)
}
