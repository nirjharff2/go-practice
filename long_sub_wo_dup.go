// Given a string, find the length of the longest substring without repeating characters.

package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)

	max_len := 0

	for i := 0; i < len(str); i++ {
		mp := make(map[byte]bool)
		curr_len := 0

		for j := i; j < len(str); j++ {
			if mp[str[j]] {
				break
			}
			mp[str[j]] = true
			curr_len += 1
		}

		if curr_len > max_len {
			max_len = curr_len
		}
	}

	fmt.Printf("%d", max_len)
}
