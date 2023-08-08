package utils

import "fmt"

func PrintList[T any](list []T, fn func(T) string) {
	for _, elem := range list {
		fmt.Println(fn(elem))
	}
}
