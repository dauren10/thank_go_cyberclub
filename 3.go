// Передача группы ожидания по указателю
package main

import (
	"fmt"
	"sync"
	"time"
)

func runWork(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		time.Sleep(50 * time.Millisecond)
		fmt.Println("work done")
		wg.Done()
	}()
}

func main() {
	var wg sync.WaitGroup
	runWork(&wg)
	wg.Wait()
	fmt.Println("all done")
}
