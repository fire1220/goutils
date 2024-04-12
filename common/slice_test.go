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

	fmt.Printf("%#v\n", SliceColumn(new(int), list, "Id"))          // []int{1, 2}
	fmt.Printf("%#v\n", SliceColumn(new(string), list, "Name"))     // []string{"jock", "fire"}
	fmt.Printf("%#v\n", SliceColumn(new(string), list, "Like"))     // []string{"basketball", "football"}
	fmt.Printf("%#v\n", SliceColumn(new(int), list, "field_EMPTY")) // []int(nil)
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

	fmt.Printf("%#v\n", SliceColumnMap(new(map[int]User), list, "Id"))                  // map[int]common.User{1:common.User{Id:1, Name:"jock", Like:"basketball"}}
	fmt.Printf("%#v\n", SliceColumnMap(new(map[int]string), list, "Id", "Name"))        // map[int]string{1:"jock"}
	fmt.Printf("%#v\n", SliceColumnMap(new(map[int]string), list, "Id", "Like"))        // map[int]string{1:"basketball"}
	fmt.Printf("%#v\n", SliceColumnMap(new(map[int]string), list, "Id", "field_EMPTY")) // map[int]string(nil)
	fmt.Printf("%#v\n", SliceColumnMap(new(map[int]string), list, "field_EMPTY"))       // map[int]string(nil)

}
