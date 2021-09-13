package main

import (
	"fmt"
	"sync"
)

var n = 0

func increase(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	n = n + 1
	m.Unlock()
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increase(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of n", n)
}
