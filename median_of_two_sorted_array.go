package main

import "fmt"

func main() {
	var n int
	var m int

	fmt.Scan(&n)
	fmt.Scan(&m)

	arr1 := make([]int, n)
	arr2 := make([]int, m)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr1[i])
	}

	for i := 0; i < m; i++ {
		fmt.Scan(&arr2[i])
	}

}
