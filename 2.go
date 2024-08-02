// Наивная группа ожидания
package main

import (
	"fmt"
	"sync"
	"time"
)

type WaitGroup struct {
	n int
}

func (wg *WaitGroup) Add(delta int) {
	wg.n += delta
	if wg.n < 0 {
		panic("negative counter")
	}
}

func (wg *WaitGroup) Done() {
	wg.Add(-1)
}

func (wg *WaitGroup) Wait() {
	for wg.n > 0 {
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		time.Sleep(50 * time.Millisecond)
		fmt.Print(".")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("work done")
}
