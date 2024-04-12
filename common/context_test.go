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

func TestContextDuplicate(t *testing.T) {
	y := context.Background()
	a := context.WithValue(y, "name", "jock")
	b := context.WithValue(a, "age", "12")
	x := ContextDuplicate(b)
	fmt.Println(b)                        // context.Background.WithValue(type string, val jock).WithValue(type string, val 12)
	fmt.Println(x)                        // context.Background.WithValue(type string, val 12).WithValue(type string, val jock)
	fmt.Println("age:", x.Value("age"))   // age: 12
	fmt.Println("name:", x.Value("name")) // name: jock
}
