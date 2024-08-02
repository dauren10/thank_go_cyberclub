// Инкапусляция: функция-обертка для WaitGroup
package main

import (
	"fmt"
	"sync"
	"time"
)

// RunConc выполняет функции одновременно и дожидается их окончания.
func RunConc(funcs ...func()) {
	var wg sync.WaitGroup
	wg.Add(len(funcs))
	for _, fn := range funcs {
		go func() {
			defer wg.Done()
			fn()
		}()
	}
	wg.Wait()
}

func main() {
	work := func() {
		time.Sleep(50 * time.Millisecond)
		fmt.Print(".")
	}

	start := time.Now()
	RunConc(work, work, work)
	elapsed := time.Now().Sub(start).Milliseconds()
	fmt.Printf("3 functions took %d ms\n", elapsed)

	// ...3 functions took 50 ms
}
