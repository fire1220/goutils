# parallel

---

### describe
> 启动多协成，统一处理并返回各个协成结果,
> 方法名和返回值都是slice，会额外返回一个error

## Usage

```
import "github.com/fire1220/goutils/parallel"
```

### example:
```go
package main

import (
	"context"
	"fmt"
	"github.com/fire1220/goutils/parallel"
)

func fun1(ctx context.Context, param interface{}) (interface{}, error) {
	return param, nil
}

func fun2(ctx context.Context, param interface{}) (interface{}, error) {
	return param, nil
}

func main() {
	ctx := context.Background()
	list, err := parallel.New().Exec(ctx, []parallel.Handle{
		fun1, // function 1
		fun2, // function 2
	},
		"张三", // function 1 params
		"李四", // function 1 params
	)
	fmt.Printf("%+v%+v\n", list, err) // [张三 李四]<nil>
}
```
