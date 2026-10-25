package main

import "fmt"

type poster interface {
	post_your_status(string)
}

type post struct {
	platform poster
}

func (p post) post_status(post_str string) {
	p.platform.post_your_status(post_str)
}

type facebook struct{}

func (f facebook) post_your_status(str string) {
	fmt.Println("Posting your status on facebook.\n\tStatus:", str)
}

type twitter struct{}

func (t twitter) post_your_status(str string) {
	fmt.Println("Posting your status on twitter.\n\tStatus:", str)
}

func main() {
	fc := facebook{}
	pst1 := post{
		platform: fc,
	}

	pst1.post_status("Hello there, this is my first post")

	tw := twitter{}
	pst2 := post{
		platform: tw,
	}

	pst2.post_status("Hello there, this is my second post")
}
