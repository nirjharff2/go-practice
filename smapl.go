package main

import (
	"fmt"
	"sync"
	"time"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("working with id", id)
	time.Sleep(time.Second * 1)
}

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}
