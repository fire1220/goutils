package wmutil

import (
	"context"
	"errors"
	"sync"
)

type Parallel struct {
}

func NewParallel() *Parallel {
	return new(Parallel)
}

type ParallelHandle func(ctx context.Context, param interface{}) (interface{}, error)

// Exec 协成批量执行方法
func (t *Parallel) Exec(ctx context.Context, funcList []ParallelHandle, params ...interface{}) ([]interface{}, error) {
	if len(params) != 0 && len(params) != len(funcList) {
		return nil, errors.New("协程执行的方法和参数数量不对应")
	}
	ctx, cancel := context.WithCancel(ctx)
	ret := make([]interface{}, len(funcList))
	errChan := make(chan interface{})
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(errChan chan<- interface{}, funcList []ParallelHandle, ret *[]interface{}, params ...interface{}) {
		defer wg.Done()
		if len(*ret) != len(funcList) {
			errChan <- errors.New("协程执行函数参数错误！")
			cancel()
			return
		}
		if len(params) != 0 && len(params) != len(funcList) {
			errChan <- errors.New("协程执行的方法和参数数量不对应！")
			cancel()
			return
		}
		wg := sync.WaitGroup{}
		for k, f := range funcList {
			var p interface{}
			if len(params) != 0 {
				p = params[k]
			}
			wg.Add(1)
			go func(f ParallelHandle, p interface{}, index int, errChan chan<- interface{}) {
				defer func() {
					if err := recover(); err != nil {
						errChan <- err
						cancel()
					}
				}()
				defer wg.Done()
				res, err := f(ctx, p)
				if err != nil {
					errChan <- err
					cancel()
				}
				(*ret)[index] = res
			}(f, p, k, errChan)
		}
		wg.Wait()
		close(errChan)
	}(errChan, funcList, &ret, params...)

	for {
		if err, ok := <-errChan; !ok {
			break
		} else if err != nil {
			if e, ok := err.(error); ok {
				return nil, e
			} else {
				return nil, errors.New("协程执行panic，并且类型不是error")
			}
		}
	}

	wg.Wait()
	if len(funcList) != len(ret) {
		return nil, errors.New("协程执行,方法列表和返回值列表不相等")
	}
	return ret, nil
}
