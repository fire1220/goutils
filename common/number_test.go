package common

import (
	"fmt"
	"testing"
)

func TestNumberNum(t *testing.T) {
	x := NumberConvChinese(3241004050) // 三十二亿四千一百万零四千零五十
	fmt.Println(x)
	// for i := 0; i <= 30000; i++ {
	// 	x := ConvertNum(uint(i))
	// 	fmt.Println(i, x)
	// }
}
