# parallel

---

### describe
> goroutines centralized access

## Usage

```
import "github.com/fire1220/goutils/parallel"
```

### example:
```
package main

import (
	"context"
	"fmt"
	"github.com/fire1220/goutils/parallel"
)

func main() {
	ctx := context.Background()
	func1 := func(ctx context.Context, param interface{}) (interface{}, error) {
		return param, nil
	}
	func2 := func(ctx context.Context, param interface{}) (interface{}, error) {
		return param, nil
	}
	list, err := parallel.New().Exec(ctx, []parallel.ParallelHandle{
		func1, // function 1
		func2, // function 2
	},
		"func1ParamYes", // function 1 params
		"func2ParamOk",  // function 1 params
	)
	fmt.Printf("%+v\n%+v\n", list, err)
}
```
