package await

import (
	"sort"
	"sync"
)

type AwaitOut[R any] struct {
	Id    int
	Value *R
	Err   error
}

func AwaitAll[T any, R any](list []T, fn func(T) (*R, error)) []AwaitOut[R] {
	var result []AwaitOut[R]

	var wg sync.WaitGroup
	var m sync.Mutex

	wg.Add(len(list))

	for i, elem := range list {
		i := i
		elem := elem
		go func() {
			defer wg.Done()
			r, err := fn(elem)
			if err != nil {
				m.Lock()
				out := AwaitOut[R]{i, nil, err}
				result = append(result, out)
				m.Unlock()
				return
			}

			m.Lock()
			out := AwaitOut[R]{i, r, nil}
			result = append(result, out)
			m.Unlock()
		}()
	}

	wg.Wait()
	return result
}

func FilterFailures[T any](list []AwaitOut[T]) []AwaitOut[T] {
	var result []AwaitOut[T]
	for _, elem := range list {
		if elem.Err != nil {
			result = append(result, elem)
		}
	}
	return result
}

func Sort[T any](list []AwaitOut[T]) {
	sort.Slice(list, func(i, j int) bool {
		return list[i].Id < list[j].Id
	})
}

func FilterSuccess[T any](list []AwaitOut[T]) []T {
	var result []T
	for _, elem := range list {
		if elem.Value != nil && elem.Err == nil {
			result = append(result, *elem.Value)
		}
	}
	return result
}
