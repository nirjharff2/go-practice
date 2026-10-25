package main

import "fmt"

func PrintSlice[T any](sl []T) { // any, interface{}, int | string | bool etc..
	for _, item := range sl {
		fmt.Println(item)
	}
}

func main() {
	arrInt := []int{1, 4, 6, 2}
	arrStr := []string{"hello", "world"}

	PrintSlice(arrInt)
	PrintSlice(arrStr)
}
