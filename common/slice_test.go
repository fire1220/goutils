package common

import (
	"fmt"
	"testing"
)

func TestSliceColumn(t *testing.T) {
	type User struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Like string `json:"like"`
	}
	a := make([]User, 0, 10)
	a = append(a, User{
		Id:   1,
		Name: "jock",
		Like: "basketball",
	})
	a = append(a, User{
		Id:   2,
		Name: "fire",
		Like: "football",
	})

	fmt.Printf("%#v\n", SliceColumn(a, "Id", new(int)))
	fmt.Printf("%#v\n", SliceColumn(a, "Name", new(string)))
	fmt.Printf("%#v\n", SliceColumn(a, "Like", new(string)))
	fmt.Printf("%#v\n", SliceColumn(a, "Id_ABC", new(int)))
}
