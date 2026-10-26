package main

import (
	"errors"
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
	str, err := OpenFile("test.go")
	var noTxtError *NoTxtError

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File doesn't exist!")
	} else if errors.As(err, &noTxtError) {
		fmt.Println(noTxtError.Message)
	}

	fmt.Println("Filename:", str)
	fmt.Println("Error:", err)
}
