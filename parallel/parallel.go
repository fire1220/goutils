package parallel

import (
	"context"
	"errors"
	"sync"
)

type parallel struct {
}

type Handle func(ctx context.Context, param interface{}) (interface{}, error)

func New() *parallel {
	return new(parallel)
}

// Exec goroutines batch execute
func (p *parallel) Exec(ctx context.Context, funcList []Handle, params ...interface{}) ([]interface{}, error) {
	if len(params) != 0 && len(params) != len(funcList) {
		return nil, errors.New("goroutine execute function name Exec params incorrect quantity")
	}
	ctx, cancel := context.WithCancel(ctx)
	ret := make([]interface{}, len(funcList))
	errChan := make(chan interface{})
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(errChan chan<- interface{}, funcList []Handle, ret *[]interface{}, params ...interface{}) {
		defer wg.Done()
		if len(*ret) != len(funcList) {
			errChan <- errors.New("goroutine execute function params error")
			cancel()
			return
		}
		if len(params) != 0 && len(params) != len(funcList) {
			errChan <- errors.New("goroutine execute function params incorrect quantity")
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
			go func(f Handle, p interface{}, index int, errChan chan<- interface{}) {
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
			for range errChan {
			}
			if e, ok := err.(error); ok {
				return nil, e
			} else {
				return nil, errors.New("goroutine execute panic,and is not error type")
			}
		}
	}

	wg.Wait()
	if len(funcList) != len(ret) {
		return nil, errors.New("goroutine execute function list and return list unequal quantity")
	}
	return ret, nil
}
