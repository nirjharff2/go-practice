package main

import "fmt"

func main() {
	var n int
	var target int

	fmt.Scan(&n)
	fmt.Scan(&target)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	st := make(map[int]int)

	for i := 0; i < n; i++ {
		rem := target - arr[i]
		if _, exist := st[rem]; exist {
			fmt.Printf("%d %d", st[arr[i]], i)
			return
		}
		st[arr[i]] = i
	}
}
