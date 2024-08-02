// Инкапусляция: тип-обертка для WaitGroup
package main

import (
	"fmt"
	"sync"
	"time"
)

// ConcRunner выполняет заданные функции одновременно.
type ConcRunner struct {
	wg    sync.WaitGroup
	funcs []func()
}

// NewConcRunner создает новый экземпляр ConcRunner.
func NewConcRunner() *ConcRunner {
	return &ConcRunner{wg: sync.WaitGroup{}}
}

// Add добавляет функцию, не выполняя ее.
func (cg *ConcRunner) Add(fn func()) {
	cg.funcs = append(cg.funcs, fn)
}

// Run выполняет функции одновременно и дожидается их окончания.
func (cg *ConcRunner) Run() {
	cg.wg.Add(len(cg.funcs))
	for _, fn := range cg.funcs {
		go func() {
			defer cg.wg.Done()
			fn()
		}()
	}
	cg.wg.Wait()
}

func timeit(cg *ConcRunner) {
	start := time.Now()
	cg.Run()
	elapsed := time.Now().Sub(start).Milliseconds()
	fmt.Printf("3 functions took %d ms\n", elapsed)
}

func main() {
	work := func() {
		time.Sleep(50 * time.Millisecond)
		fmt.Print(".")
	}

	cr := NewConcRunner()

	// формируем набор функций
	cr.Add(work)
	cr.Add(work)
	cr.Add(work)

	// выполняем настроенные функции
	timeit(cr)
	// ...3 functions took 50 ms

	// и еще разок
	timeit(cr)
	// ...3 functions took 50 ms
}
