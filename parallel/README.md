# parallel

---

### describe
> goroutines centralized access


```go
list, err := parallel.New().Exec(ctx, []parallel.ParallelHandle{function list}, params...)
```

### example:
```go
ctx := context.Background()

func1 := func(ctx context.Context, param interface{}) (interface{}, error) {
    return param, nil
}

func2 := func(ctx context.Context, param interface{}) (interface{}, error) {
    return param, nil
}

list, err := parallel.New().Exec(ctx, []parallel.ParallelHandle{func1,func2},"yes","ok")

fmt.Printf("%+v\n%+v\n", list, err)
```