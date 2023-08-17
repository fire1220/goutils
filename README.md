# go-utils

## Installation
```shell
go get github.com/fire1220/go-utils
```
 ## Usage

```
import "github.com/fire1220/go-utils"
```

---


# parallel

### describe
> goroutines centralized access

## Usage

```
import "github.com/fire1220/go-utils/parallel"
```
## 
```
list, err := parallel.New().Exec(ctx, []parallel.ParallelHandle{function list}, params...)
```

### example:
```
ctx := context.Background()

func1 := func(ctx context.Context, param interface{}) (interface{}, error) {
    return param, nil
}

func2 := func(ctx context.Context, param interface{}) (interface{}, error) {
    return param, nil
}

list, err := New().Exec(ctx, []ParallelHandle{
	func1, // function 1
	func2, // function 2
},
	"func1ParamYes", // function 1 params
	"func2ParamOk",  // function 2 params
)

fmt.Printf("%+v\n%+v\n", list, err)
```