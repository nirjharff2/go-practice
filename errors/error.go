package main

import (
	"fmt"
	"os"
	"strings"
)

type NoTxtError struct {
	Message string
}

func (n NoTxtError) Error() string {
	return n.Message
}

func OpenFile(filename string) (string, error) {
	if !strings.HasSuffix(filename, ".txt") {
		return "", &NoTxtError{
			"Non txt files are not allowed.",
		}
	}

	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	return filename, nil
}

func main() {
	str, err := OpenFile("error.txt")

	fmt.Println(str)
	fmt.Println(err)
}
