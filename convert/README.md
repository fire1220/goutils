# convert 常用转换工具

---

### describe
> ConvNum 数字转汉字；例："100" 转成 "一百"

## Usage

```
import "github.com/fire1220/goutils/convert"
```

### example:
``` go
package convert

import (
	"fmt"
	"testing"
)

func TestConvertNum(t *testing.T) {
	x := ConvNum(3241004050) // 三十二亿四千一百万零四千零五十
	fmt.Println(x)
}
```
