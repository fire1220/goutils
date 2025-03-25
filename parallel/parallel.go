package parallel

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

type Para interface {
	Exec() error
}

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

// ExecObj
//
//	example:
//	type Add struct {
//		Param [2]int
//		Res   int
//	}
//	func (a *Add) Exec() error {
//		a.Res = a.Param[0] + a.Param[1]
//		return nil
//	}
//	y := &Add{Param: [2]int{1, 2}}
//	err := parallel.New().ExecObj(ctx, y)
//	fmt.Printf("%#v\n %v", y, err)
func (p *parallel) ExecObj(ctx context.Context, objList ...Para) error {
	if len(objList) == 0 {
		return nil
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	errCh := make(chan error)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(errCh chan<- error, objList []Para) {
		defer wg.Done()
		wg := sync.WaitGroup{}
		for _, obj := range objList {
			wg.Add(1)
			go func(errCh chan<- error, obj Para) {
				defer func() {
					if e := recover(); e != nil {
						errCh <- fmt.Errorf("panic : err = %v", e)
						cancel()
					}
				}()
				defer wg.Done()
				err := obj.Exec()
				if err != nil {
					errCh <- err
					cancel()
				}
			}(errCh, obj)
		}
		wg.Wait()
		close(errCh)
	}(errCh, objList)
	for err := range errCh {
		return err
	}
	wg.Wait()
	return nil
}
