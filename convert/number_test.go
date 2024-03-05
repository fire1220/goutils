package convert

import (
	"fmt"
	"testing"
)

func TestConvertNum(t *testing.T) {
	x := ConvertNum(3241004050)
	fmt.Println(x)
	// for i := 0; i <= 30000; i++ {
	// 	x := ConvertNum(uint(i))
	// 	fmt.Println(i, x)
	// }
}
