package common

import (
	"context"
	"fmt"
	"testing"
)

func TestConvertNum(t *testing.T) {
	y := context.Background()
	a := context.WithValue(y, "name", "jock")
	b := context.WithValue(a, "age", "12")
	x := ContextKeys(b)
	// x = ContextKeys(context.Background())
	fmt.Println(x)
	for _, val := range x {
		fmt.Printf("%v:%v\n", val, b.Value(val))
	}
}
