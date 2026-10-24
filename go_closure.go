package main

import "fmt"

func f1() func() int {
	i := 0

	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextInt := f1()

	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3
}
