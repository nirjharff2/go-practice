package main

import "fmt"

func f1() func(string) string {
	return func(str string) string {
		return "hello" + " " + str
	}
}

func main() {
	greet := f1()

	fmt.Println(greet("Nirjhar"))
	fmt.Println(greet("Emma"))
}
