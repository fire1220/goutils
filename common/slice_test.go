package common

import (
	"fmt"
	"testing"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Like string `json:"like"`
}

func TestSliceColumn(t *testing.T) {
	list := make([]User, 0, 10)
	list = append(list, User{
		Id:   1,
		Name: "jock",
		Like: "basketball",
	})
	list = append(list, User{
		Id:   2,
		Name: "fire",
		Like: "football",
	})

	fmt.Printf("%#v\n", SliceColumn(new(int), list, "Id"))
	fmt.Printf("%#v\n", SliceColumn(new(string), list, "Name"))
	fmt.Printf("%#v\n", SliceColumn(new(string), list, "Like"))
	fmt.Printf("%#v\n", SliceColumn(new(int), list, "Id_ABC"))
}

func TestSliceColumnMap(t *testing.T) {
	list := make([]User, 0, 10)
	list = append(list, User{
		Id:   1,
		Name: "jock",
		Like: "basketball",
	})
	list = append(list, User{
		Id:   2,
		Name: "fire",
		Like: "football",
	})

	fmt.Printf("%#v\n", SliceColumnMap(new(map[int]User), list, "Id"))
	fmt.Printf("%#v\n", SliceColumnMap(new(map[int]string), list, "Id", "Name"))
	fmt.Printf("%#v\n", SliceColumnMap(new(map[int]string), list, "Id", "Like"))
}
