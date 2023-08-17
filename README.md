# go-utils

---

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

---

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

list, err := parallel.New().Exec(ctx, []parallel.ParallelHandle{func1,func2},"func1ParamYes","func2ParamOk")

fmt.Printf("%+v\n%+v\n", list, err)
```