package main

import "fmt"

func f1() (func() int, func() int) {

	i := 0

	inc := func() int {
		i += 1
		return i
	}

	dec := func() int {
		i -= 1
		return i
	}

	return inc, dec
}

func main() {
	inc, dec := f1()

	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(dec())
}
